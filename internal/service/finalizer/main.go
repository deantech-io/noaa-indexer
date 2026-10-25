package finalizer

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/noaa-indexer/internal/chain"
	"github.com/rarimo/noaa-indexer/internal/config"
	"github.com/rarimo/noaa-indexer/internal/contracts"
	"github.com/rarimo/noaa-indexer/internal/data"
	"github.com/rarimo/noaa-indexer/internal/data/pg"
	"github.com/rarimo/noaa-indexer/internal/metrics"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const serviceName = "finalizer"

// errRejected means nothing was sent, so the request may be retried.
var errRejected = errors.New("rejected before sending")

type submitter interface {
	submit(ctx context.Context, calldata []byte) (common.Hash, error)
}

type Finalizer struct {
	log       *logan.Entry
	client    *chain.Client
	db        *pgdb.DB
	chainID   uint64
	contract  common.Address
	cfg       config.Finalizer
	abi       *abi.ABI
	submitter submitter

	state       pg.StateQ
	events      pg.EventsQ
	submissions pg.SubmissionsQ
}

func New(cfg config.Config) *Finalizer {
	network, contract, fcfg := cfg.Network(), cfg.Contract().Address, cfg.Finalizer()

	parsed, err := contracts.NoAARecoveryMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to parse NoAARecovery ABI"))
	}

	var s submitter
	switch fcfg.Mode {
	case config.ModeRelayer:
		s = newRelayer(fcfg.RelayerURL, fcfg.RelayerTimeout, contract)
	case config.ModeWallet:
		s = newWallet(network.Client, fcfg.WalletPrivateKey, network.ChainID, contract)
	}

	return &Finalizer{
		log: cfg.Log().WithFields(logan.F{
			"service": serviceName,
			"address": contract.Hex(),
			"mode":    fcfg.Mode,
		}),
		client:    network.Client,
		db:        cfg.DB().Clone(),
		chainID:   network.ChainID,
		contract:  contract,
		cfg:       fcfg,
		abi:       parsed,
		submitter: s,
	}
}

func (f *Finalizer) Run(ctx context.Context) error {
	contractID, err := pg.NewContractQ(f.db).Ensure(f.chainID, f.contract)
	if err != nil {
		return errors.Wrap(err, "failed to ensure contract")
	}
	f.state = pg.NewStateQ(f.db, contractID)
	f.events = pg.NewEventsQ(f.db, f.chainID, contractID)
	f.submissions = pg.NewSubmissionsQ(f.db, contractID)

	f.log.Info("finalizer started")

	// A 'submitting' row at start-up was interrupted mid-flight.
	stale, err := f.submissions.SelectByStatus(data.SubmissionSubmitting)
	if err != nil {
		return errors.Wrap(err, "failed to select interrupted submissions")
	}
	for _, s := range stale {
		f.markUnknown(s.ID, s.PassportHash, errors.New("interrupted while submitting"))
	}

	for {
		if err := f.tick(ctx); err != nil && ctx.Err() == nil {
			f.log.WithError(err).Error("finalizer tick failed")
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(f.cfg.Period):
		}
	}
}

func (f *Finalizer) tick(ctx context.Context) error {
	if err := f.track(ctx); err != nil {
		return errors.Wrap(err, "failed to track submissions")
	}

	// The contract compares readyAt with block.timestamp.
	var head *types.Header
	err := f.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) (err error) {
		head, err = ec.HeaderByNumber(ctx, nil)
		return err
	})
	if err != nil {
		return errors.Wrap(err, "failed to get latest header")
	}
	now := time.Unix(int64(head.Time), 0).UTC()

	ready, err := f.state.SelectReady(now)
	if err != nil {
		return errors.Wrap(err, "failed to select ready requests")
	}

	for _, s := range ready {
		if err = f.finalize(ctx, s, now); err != nil {
			f.log.WithError(err).WithField("passport_hash", s.PassportHash.Hex()).Error("failed to finalize")
		}
	}
	return nil
}

func (f *Finalizer) finalize(ctx context.Context, s data.RecoveryState, now time.Time) error {
	log := f.log.WithField("passport_hash", s.PassportHash.Hex())

	request, err := f.events.LatestRequest(s.PassportHash)
	if err != nil {
		return errors.Wrap(err, "failed to get request event")
	}
	if request == nil {
		return errors.New("pending state without request event")
	}

	blocked, err := f.submissions.Blocked(request.ID)
	if err != nil {
		return errors.Wrap(err, "failed to check submissions")
	}
	if blocked {
		return nil
	}

	// The indexed state lags behind the chain.
	current, err := f.isCurrent(ctx, s, now)
	if err != nil {
		return errors.Wrap(err, "failed to check on-chain state")
	}
	if !current {
		log.Info("on-chain state changed, skipping")
		metrics.Submissions.WithLabelValues("skipped").Inc()
		return nil
	}

	calldata, err := f.abi.Pack("finalizeRecovery", common.Hash(s.PassportHash))
	if err != nil {
		return errors.Wrap(err, "failed to pack calldata")
	}

	id, err := f.submissions.Claim(request.ID)
	if err != nil {
		return errors.Wrap(err, "failed to claim submission")
	}
	if id == 0 {
		return nil
	}

	hash, err := f.submitter.submit(ctx, calldata)
	switch {
	case err == nil:
		txHash := data.Hash(hash)
		if err = f.update(id, data.SubmissionSubmitted, &txHash, "", data.SubmissionSubmitting, data.SubmissionUnknown); err != nil {
			return errors.Wrap(err, "failed to save submitted", logan.F{"tx_hash": hash.Hex()})
		}
		log.WithField("tx_hash", hash.Hex()).Info("finalization submitted")
		metrics.Submissions.WithLabelValues(string(data.SubmissionSubmitted)).Inc()
		return nil
	case errors.Cause(err) == errRejected:
		if uerr := f.update(id, data.SubmissionRejected, nil, err.Error(), data.SubmissionSubmitting, data.SubmissionUnknown); uerr != nil {
			return uerr
		}
		log.WithError(err).Warn("finalization rejected")
		metrics.Submissions.WithLabelValues(string(data.SubmissionRejected)).Inc()
		return nil
	default:
		f.markUnknown(id, s.PassportHash, err)
		return nil
	}
}

func (f *Finalizer) isCurrent(ctx context.Context, s data.RecoveryState, now time.Time) (bool, error) {
	var (
		request struct {
			NewIdentityKey *big.Int
			ReadyAt        uint64
		}
		lockedUntil uint64
	)
	err := f.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) error {
		caller, err := contracts.NewNoAARecoveryCaller(f.contract, ec)
		if err != nil {
			return err
		}
		opts := &bind.CallOpts{Context: ctx}
		if request, err = caller.Requests(opts, s.PassportHash); err != nil {
			return err
		}
		lockedUntil, err = caller.LockedUntil(opts, s.PassportHash)
		return err
	})
	if err != nil {
		return false, err
	}

	readyAt := data.Time(request.ReadyAt)
	return !s.NewIdentityKey.IsZero() && data.NewNumeric(request.NewIdentityKey).Equal(s.NewIdentityKey) &&
		readyAt != nil && s.ReadyAt != nil && readyAt.Equal(*s.ReadyAt) && !readyAt.After(now) &&
		lockedUntil <= uint64(now.Unix()), nil
}

func (f *Finalizer) track(ctx context.Context) error {
	open, err := f.submissions.SelectByStatus(data.SubmissionSubmitted, data.SubmissionUnknown)
	if err != nil {
		return err
	}

	unknown := 0
	for _, s := range open {
		log := f.log.WithFields(logan.F{"passport_hash": s.PassportHash.Hex(), "submission_id": s.ID})

		if s.TxHash != nil {
			var receipt *types.Receipt
			err = f.client.Do(ctx, func(ctx context.Context, ec *ethclient.Client) (err error) {
				receipt, err = ec.TransactionReceipt(ctx, common.Hash(*s.TxHash))
				return err
			})
			switch {
			case err == nil:
				status := data.SubmissionConfirmed
				if receipt.Status != types.ReceiptStatusSuccessful {
					status = data.SubmissionFailed
				}
				if err = f.update(s.ID, status, nil, s.Error, s.Status); err != nil {
					return err
				}
				log.WithField("tx_hash", s.TxHash.Hex()).Infof("finalization %s", status)
				metrics.Submissions.WithLabelValues(string(status)).Inc()
				continue
			case errors.Cause(err) != ethereum.NotFound:
				return errors.Wrap(err, "failed to get receipt", logan.F{"tx_hash": s.TxHash.Hex()})
			}
		}

		// No receipt: close once the request is gone.
		pending, err := f.requestPending(s)
		if err != nil {
			return err
		}
		if !pending {
			if err = f.update(s.ID, data.SubmissionResolved, nil, s.Error, s.Status); err != nil {
				return err
			}
			log.Info("request resolved, submission closed")
			continue
		}

		// Covers unknown outcomes and dropped transactions.
		if f.cfg.UnknownRetryAfter > 0 && time.Since(s.UpdatedAt) > f.cfg.UnknownRetryAfter {
			if err = f.update(s.ID, data.SubmissionExpired, nil, s.Error, s.Status); err != nil {
				return err
			}
			log.Warn("submission expired, allowing a new attempt")
			continue
		}
		if s.Status == data.SubmissionUnknown {
			unknown++
		}
	}

	metrics.UnknownSubmissions.Set(float64(unknown))
	return nil
}

func (f *Finalizer) requestPending(s data.Submission) (bool, error) {
	state, err := f.state.Get(s.PassportHash)
	if err != nil {
		return false, errors.Wrap(err, "failed to get state")
	}
	if state == nil || state.NewIdentityKey.IsZero() {
		return false, nil
	}

	latest, err := f.events.LatestRequest(s.PassportHash)
	if err != nil {
		return false, errors.Wrap(err, "failed to get request event")
	}
	return latest != nil && latest.ID == s.RecoveryEventID, nil
}

func (f *Finalizer) markUnknown(id int64, passportHash data.Hash, cause error) {
	log := f.log.WithFields(logan.F{"passport_hash": passportHash.Hex(), "submission_id": id})
	log.WithError(cause).Error("finalization outcome unknown, request blocked until resolved")
	metrics.Submissions.WithLabelValues(string(data.SubmissionUnknown)).Inc()

	if err := f.update(id, data.SubmissionUnknown, nil, cause.Error(), data.SubmissionSubmitting); err != nil {
		log.WithError(err).Error("failed to mark submission unknown")
	}
}

func (f *Finalizer) update(id int64, to data.SubmissionStatus, txHash *data.Hash, errMsg string, from ...data.SubmissionStatus) error {
	fields := logan.F{"submission_id": id, "status": to}
	ok, err := f.submissions.Update(id, to, txHash, errMsg, from...)
	if err != nil {
		return errors.Wrap(err, "failed to update submission", fields)
	}
	if !ok {
		return errors.From(errors.New("submission changed concurrently"), fields)
	}
	return nil
}
