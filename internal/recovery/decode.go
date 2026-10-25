package recovery

import (
	"github.com/deantech-io/noaa-indexer/internal/contracts"
	"github.com/deantech-io/noaa-indexer/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

// Decode converts a log into an event and the values read from the contract.
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

	switch e.Type {
	case data.EventRequested:
		ev, err := filterer.ParseRecoveryRequested(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryRequested")
		}
		if !ev.ReadyAt.IsUint64() {
			return e, errors.From(errors.New("readyAt out of range"), logan.F{"ready_at": ev.ReadyAt.String()})
		}
		e.PassportHash = data.Hash(ev.PassportHash)
		e.NewIdentityKey = data.NumericFromBytes(ev.NewIdentityKey)
		e.ReadyAt = data.Time(ev.ReadyAt.Uint64())
	case data.EventCancelled:
		ev, err := filterer.ParseRecoveryCancelled(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryCancelled")
		}
		e.PassportHash = data.Hash(ev.PassportHash)
	case data.EventContested:
		ev, err := filterer.ParseRecoveryContested(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryContested")
		}
		e.PassportHash = data.Hash(ev.PassportHash)
	case data.EventFinalized:
		ev, err := filterer.ParseRecoveryFinalized(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryFinalized")
		}
		e.PassportHash = data.Hash(ev.PassportHash)
	case data.EventLocked:
		ev, err := filterer.ParseRecoveryLocked(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryLocked")
		}
		e.PassportHash = data.Hash(ev.PassportHash)
	case data.EventUnlocked:
		ev, err := filterer.ParseRecoveryUnlocked(l)
		if err != nil {
			return e, errors.Wrap(err, "failed to parse RecoveryUnlocked")
		}
		e.PassportHash = data.Hash(ev.PassportHash)
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
