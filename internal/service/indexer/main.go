package indexer

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rarimo/noaa-indexer/internal/chain"
	"github.com/rarimo/noaa-indexer/internal/config"
	"github.com/rarimo/noaa-indexer/internal/data"
	"github.com/rarimo/noaa-indexer/internal/data/pg"
	"github.com/rarimo/noaa-indexer/internal/metrics"
	"github.com/rarimo/noaa-indexer/internal/recovery"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const serviceName = "indexer"

// ErrParentHashMismatch means the checkpoint block is no longer on chain. The indexer stops on it.
var ErrParentHashMismatch = errors.New("parent hash mismatch")

type Indexer struct {
	log      *logan.Entry
	client   *chain.Client
	db       *pgdb.DB
	chainID  uint64
	contract config.Contract
	cfg      config.Indexer
	topics   []common.Hash

	contractID int64
	chunk      uint64
	halted     atomic.Bool
}

func New(cfg config.Config) *Indexer {
	contract := cfg.Contract()
	return &Indexer{
		log: cfg.Log().WithFields(logan.F{
			"service": serviceName,
			"address": contract.Address.Hex(),
		}),
		client:   cfg.Network().Client,
		db:       cfg.DB().Clone(),
		chainID:  cfg.Network().ChainID,
		contract: contract,
		cfg:      cfg.Indexer(),
		topics:   recovery.Topics(),
		chunk:    cfg.Indexer().ChunkSize,
	}
}

func (ix *Indexer) Healthy() bool {
	return !ix.halted.Load()
}

// Run returns ErrParentHashMismatch on a halt. Other errors are logged and retried.
func (ix *Indexer) Run(ctx context.Context) error {
	var err error
	if ix.contractID, err = pg.NewContractQ(ix.db).Ensure(ix.chainID, ix.contract.Address); err != nil {
		return errors.Wrap(err, "failed to ensure contract")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var halt error
	ix.log.Info("indexer started")
	running.WithBackOff(ctx, ix.log, serviceName, func(ctx context.Context) error {
		err := ix.catchUp(ctx)
		if errors.Cause(err) != ErrParentHashMismatch {
			return err
		}
		halt = err
		ix.halted.Store(true)
		metrics.Halted.Set(1)
		ix.log.WithError(err).Error("indexer stopped: checkpoint block is not on chain")
		cancel()
		return nil
	}, ix.cfg.PollPeriod, ix.cfg.PollPeriod, 10*ix.cfg.PollPeriod)
	return halt
}

// catchUp indexes ranges until it reaches the finalized block.
func (ix *Indexer) catchUp(ctx context.Context) error {
	for ctx.Err() == nil {
		caughtUp, err := ix.step(ctx)
		if err != nil || caughtUp {
			return err
		}
	}
	return nil
}

// step indexes one range and reports whether it reached the finalized block.
func (ix *Indexer) step(ctx context.Context) (bool, error) {
	cp, err := pg.NewContractQ(ix.db).Checkpoint(ix.contractID)
	if err != nil {
		return false, errors.Wrap(err, "failed to get checkpoint")
	}

	from, indexed := ix.contract.Block, uint64(0)
	if cp != nil {
		from, indexed = cp.BlockNumber+1, cp.BlockNumber
	}

	var r chunk
	err = ix.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) (err error) {
		r, err = ix.read(ctx, ec, from)
		return err
	})
	if err != nil {
		// Providers reject ranges that are too wide or return too many logs.
		if _, ok := err.(rpc.Error); ok {
			ix.chunk = max(1, ix.chunk/2)
		}
		return false, errors.Wrap(err, "failed to read range", logan.F{"from": from})
	}
	ix.chunk = min(ix.cfg.ChunkSize, ix.chunk*2)

	head := ix.head(ctx)
	// A fallback node may lag behind the checkpoint.
	if r.first == nil {
		report(head, indexed)
		return true, nil
	}

	if cp != nil && data.Hash(r.first.ParentHash) != cp.BlockHash {
		return false, errors.Wrap(ErrParentHashMismatch, "checkpoint mismatch", logan.F{
			"checkpoint_block": cp.BlockNumber,
			"checkpoint_hash":  cp.BlockHash.Hex(),
			"parent_hash":      r.first.ParentHash.Hex(),
		})
	}

	events := make([]data.RecoveryEvent, 0, len(r.logs))
	for _, l := range r.logs {
		e, err := recovery.Decode(l)
		if err != nil {
			return false, errors.Wrap(err, "failed to decode log", logan.F{"tx_hash": l.TxHash.Hex(), "log_index": l.Index})
		}
		e.MinedAt = r.minedAt[l.BlockHash]
		events = append(events, e)
	}

	to := r.last.Number.Uint64()
	unexpected, err := ix.save(cp, events, r.last)
	if err != nil {
		return false, errors.Wrap(err, "failed to save range", logan.F{"from": from, "to": to})
	}

	for _, e := range unexpected {
		ix.log.WithFields(logan.F{
			"passport_hash": e.PassportHash.Hex(),
			"event_type":    e.Type,
			"tx_hash":       e.TxHash.Hex(),
		}).Warn("unexpected transition")
		metrics.UnexpectedTransitions.WithLabelValues(string(e.Type)).Inc()
	}

	for _, e := range events {
		metrics.Events.WithLabelValues(string(e.Type)).Inc()
	}
	report(head, to)
	ix.log.WithFields(logan.F{"from": from, "to": to, "events": len(events)}).Debug("indexed range")

	return to == r.finalized, nil
}

type chunk struct {
	finalized   uint64
	first, last *types.Header // nil if from is above the finalized block
	logs        []types.Log
	minedAt     map[common.Hash]time.Time
}

// read fetches one range from a single endpoint, so a lagging fallback cannot mix with another
// node's view of the chain.
func (ix *Indexer) read(ctx context.Context, ec *ethclient.Client, from uint64) (chunk, error) {
	var r chunk
	finalized, err := ec.HeaderByNumber(ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		return r, err
	}
	r.finalized = finalized.Number.Uint64()
	if from > r.finalized {
		return r, nil
	}
	to := min(r.finalized, from+ix.chunk-1)

	if r.first, err = ec.HeaderByNumber(ctx, new(big.Int).SetUint64(from)); err != nil {
		return r, err
	}
	r.logs, err = ec.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(from),
		ToBlock:   new(big.Int).SetUint64(to),
		Addresses: []common.Address{ix.contract.Address},
		Topics:    [][]common.Hash{ix.topics},
	})
	if err != nil {
		return r, err
	}
	if r.last, err = ec.HeaderByNumber(ctx, new(big.Int).SetUint64(to)); err != nil {
		return r, err
	}

	r.minedAt = make(map[common.Hash]time.Time)
	for _, l := range r.logs {
		if _, ok := r.minedAt[l.BlockHash]; ok {
			continue
		}
		h, err := ec.HeaderByHash(ctx, l.BlockHash)
		if err != nil {
			return r, err
		}
		r.minedAt[l.BlockHash] = time.Unix(int64(h.Time), 0).UTC()
	}
	return r, nil
}

// save stores events, updates state and moves the checkpoint in one transaction.
// It returns the new events the contract should not have emitted in the stored state.
func (ix *Indexer) save(cp *data.Checkpoint, events []data.RecoveryEvent, last *types.Header) ([]data.RecoveryEvent, error) {
	contracts := pg.NewContractQ(ix.db)
	blocks := pg.NewBlockQ(ix.db, ix.chainID)
	eventsQ := pg.NewEventsQ(ix.db, ix.chainID, ix.contractID)
	stateQ := pg.NewStateQ(ix.db, ix.contractID)

	var unexpected []data.RecoveryEvent
	err := ix.db.Transaction(func() error {
		if err := contracts.Lock(ix.chainID); err != nil {
			return errors.Wrap(err, "failed to lock")
		}

		for _, e := range events {
			inserted, err := eventsQ.Insert(&e)
			if err != nil {
				return errors.Wrap(err, "failed to insert event")
			}
			// Stored events are already applied.
			if !inserted {
				continue
			}

			state, err := stateQ.Get(e.PassportHash)
			if err != nil {
				return errors.Wrap(err, "failed to get state")
			}
			if state == nil {
				state = &data.RecoveryState{}
			}
			if !recovery.Expected(*state, e) {
				unexpected = append(unexpected, e)
			}
			if err = stateQ.Save(recovery.Apply(*state, e)); err != nil {
				return errors.Wrap(err, "failed to save state")
			}
		}

		blockID, err := blocks.Ensure(last.Number.Uint64(), data.Hash(last.Hash()), time.Unix(int64(last.Time), 0).UTC())
		if err != nil {
			return errors.Wrap(err, "failed to save checkpoint block")
		}
		var prev *int64
		if cp != nil {
			prev = &cp.BlockID
		}
		moved, err := contracts.SetCheckpoint(ix.contractID, prev, blockID)
		if err != nil {
			return errors.Wrap(err, "failed to set checkpoint")
		}
		if !moved {
			return errors.New("checkpoint was moved concurrently")
		}
		return nil
	})
	return unexpected, err
}

// head returns 0 on error: it only feeds the lag metric.
func (ix *Indexer) head(ctx context.Context) uint64 {
	var head uint64
	err := ix.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) (err error) {
		head, err = ec.BlockNumber(ctx)
		return err
	})
	if err != nil {
		ix.log.WithError(err).Warn("failed to get head block")
	}
	return head
}

func report(head, indexed uint64) {
	metrics.IndexedBlock.Set(float64(indexed))
	if head > 0 {
		metrics.HeadBlock.Set(float64(head))
		metrics.LagBlocks.Set(float64(head) - float64(indexed))
	}
}

// Rebuild recomputes recovery_state from recovery_event.
func Rebuild(cfg config.Config) error {
	db := cfg.DB().Clone()
	chainID := cfg.Network().ChainID
	contracts := pg.NewContractQ(db)

	contractID, err := contracts.Ensure(chainID, cfg.Contract().Address)
	if err != nil {
		return errors.Wrap(err, "failed to ensure contract")
	}
	eventsQ := pg.NewEventsQ(db, chainID, contractID)
	stateQ := pg.NewStateQ(db, contractID)

	var passports int
	err = db.Transaction(func() error {
		if err := contracts.Lock(chainID); err != nil {
			return errors.Wrap(err, "failed to lock")
		}

		events, err := eventsQ.Select()
		if err != nil {
			return errors.Wrap(err, "failed to select events")
		}

		states := make(map[data.Hash]data.RecoveryState)
		for _, e := range events {
			states[e.PassportHash] = recovery.Apply(states[e.PassportHash], e)
		}
		passports = len(states)

		for _, s := range states {
			stored, err := stateQ.Get(s.PassportHash)
			if err != nil {
				return errors.Wrap(err, "failed to get state")
			}
			if stored != nil {
				s.ID = stored.ID
			}
			if err = stateQ.Save(s); err != nil {
				return errors.Wrap(err, "failed to save state")
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	cfg.Log().WithField("passports", passports).Info("recovery state rebuilt")
	return nil
}
