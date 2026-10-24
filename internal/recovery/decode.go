package recovery

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/noaa-indexer/internal/contracts"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	filterer = mustFilterer()

	eventTypes = map[string]data.EventType{
		"RecoveryRequested": data.EventRequested,
		"RecoveryCancelled": data.EventCancelled,
		"RecoveryContested": data.EventContested,
		"RecoveryFinalized": data.EventFinalized,
		"RecoveryLocked":    data.EventLocked,
		"RecoveryUnlocked":  data.EventUnlocked,
	}

	topics = mustTopics()
)

func Topics() []common.Hash {
	res := make([]common.Hash, 0, len(topics))
	for t := range topics {
		res = append(res, t)
	}
	return res
}

// Decode converts a log into an event without MinedAt.
func Decode(l types.Log) (data.RecoveryEvent, error) {
	e := data.RecoveryEvent{
		TxHash:      data.Hash(l.TxHash),
		LogIndex:    uint64(l.Index),
		BlockNumber: l.BlockNumber,
		BlockHash:   data.Hash(l.BlockHash),
	}
	if len(l.Topics) == 0 {
		return e, errors.New("log has no topics")
	}

	var ok bool
	if e.Type, ok = topics[l.Topics[0]]; !ok {
		return e, errors.From(errors.New("unknown event"), logan.F{"topic": l.Topics[0].Hex()})
	}

	var err error
	switch e.Type {
	case data.EventRequested:
		var ev *contracts.NoAARecoveryRecoveryRequested
		if ev, err = filterer.ParseRecoveryRequested(l); err == nil {
			e.PassportHash, e.NewIdentityKey, e.ReadyAt = data.Hash(ev.PassportHash), data.NewNumeric(ev.NewIdentityKey), data.Time(ev.ReadyAt)
		}
	case data.EventCancelled:
		var ev *contracts.NoAARecoveryRecoveryCancelled
		if ev, err = filterer.ParseRecoveryCancelled(l); err == nil {
			e.PassportHash = data.Hash(ev.PassportHash)
		}
	case data.EventContested:
		var ev *contracts.NoAARecoveryRecoveryContested
		if ev, err = filterer.ParseRecoveryContested(l); err == nil {
			e.PassportHash, e.LockedUntil = data.Hash(ev.PassportHash), data.Time(ev.LockedUntil)
		}
	case data.EventFinalized:
		var ev *contracts.NoAARecoveryRecoveryFinalized
		if ev, err = filterer.ParseRecoveryFinalized(l); err == nil {
			e.PassportHash, e.NewIdentityKey = data.Hash(ev.PassportHash), data.NewNumeric(ev.NewIdentityKey)
		}
	case data.EventLocked:
		var ev *contracts.NoAARecoveryRecoveryLocked
		if ev, err = filterer.ParseRecoveryLocked(l); err == nil {
			e.PassportHash = data.Hash(ev.PassportHash)
		}
	case data.EventUnlocked:
		var ev *contracts.NoAARecoveryRecoveryUnlocked
		if ev, err = filterer.ParseRecoveryUnlocked(l); err == nil {
			e.PassportHash = data.Hash(ev.PassportHash)
		}
	}
	if err != nil {
		return e, errors.Wrap(err, "failed to parse log", logan.F{"event": e.Type})
	}
	return e, nil
}

func mustFilterer() *contracts.NoAARecoveryFilterer {
	f, err := contracts.NewNoAARecoveryFilterer(common.Address{}, nil)
	if err != nil {
		panic(errors.Wrap(err, "failed to create filterer"))
	}
	return f
}

func mustTopics() map[common.Hash]data.EventType {
	abi, err := contracts.NoAARecoveryMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to parse NoAARecovery ABI"))
	}

	res := make(map[common.Hash]data.EventType, len(eventTypes))
	for name, typ := range eventTypes {
		ev, ok := abi.Events[name]
		if !ok {
			panic(errors.From(errors.New("event missing from ABI"), logan.F{"event": name}))
		}
		res[ev.ID] = typ
	}
	return res
}
