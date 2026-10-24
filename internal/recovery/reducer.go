package recovery

import (
	"github.com/rarimo/noaa-indexer/internal/data"
)

// Apply returns the state after e. Events must be applied in chain order.
//
// TODO(real ABI): PROVISIONAL - a lock keeps the pending request and blocks its finalization.
func Apply(s data.RecoveryState, e data.RecoveryEvent) data.RecoveryState {
	locked := s.Status == data.StatusLocked

	switch e.Type {
	case data.EventRequested:
		s.NewIdentityKey, s.ReadyAt, s.LockedUntil = e.NewIdentityKey, e.ReadyAt, nil
	case data.EventCancelled, data.EventFinalized:
		s.NewIdentityKey, s.ReadyAt = data.Numeric{}, nil
	case data.EventContested:
		s.NewIdentityKey, s.ReadyAt, s.LockedUntil = data.Numeric{}, nil, e.LockedUntil
	case data.EventLocked:
		locked = true
	case data.EventUnlocked:
		locked = false
	}

	switch {
	case locked:
		s.Status = data.StatusLocked
	case !s.NewIdentityKey.IsZero():
		s.Status = data.StatusPending
	case s.LockedUntil != nil:
		s.Status = data.StatusContested
	default:
		s.Status = data.StatusActive
	}

	s.PassportHash = e.PassportHash
	s.RecoveryEventID = e.ID
	return s
}
