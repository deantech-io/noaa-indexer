package recovery

import (
	"github.com/rarimo/noaa-indexer/internal/data"
)

// Expected reports whether the contract allows e in state s.
func Expected(s data.RecoveryState, e data.RecoveryEvent) bool {
	pending := !s.NewIdentityKey.IsZero()

	switch e.Type {
	case data.EventRequested:
		return !s.IsRecoveryLocked && !pending && (s.LockedUntil == nil || !s.LockedUntil.After(e.MinedAt))
	case data.EventContested:
		return !s.IsRecoveryLocked && pending
	case data.EventCancelled:
		return pending
	case data.EventFinalized:
		return pending && s.ReadyAt != nil && !s.ReadyAt.After(e.MinedAt)
	case data.EventLocked:
		return !s.IsRecoveryLocked
	case data.EventUnlocked:
		return s.IsRecoveryLocked
	}
	return false
}

// Apply returns the state after e. Events must be applied in chain order.
func Apply(s data.RecoveryState, e data.RecoveryEvent) data.RecoveryState {
	switch e.Type {
	case data.EventRequested:
		s.OldIdentityKey, s.NewIdentityKey, s.DGCommit, s.ReadyAt = e.OldIdentityKey, e.NewIdentityKey, e.DGCommit, e.ReadyAt
	case data.EventContested, data.EventCancelled:
		s.OldIdentityKey, s.NewIdentityKey, s.DGCommit, s.ReadyAt = data.Numeric{}, data.Numeric{}, data.Numeric{}, nil
		s.LockedUntil = e.LockedUntil
	case data.EventFinalized:
		s.OldIdentityKey, s.NewIdentityKey, s.DGCommit, s.ReadyAt = data.Numeric{}, data.Numeric{}, data.Numeric{}, nil
	case data.EventLocked:
		s.IsRecoveryLocked = true
	case data.EventUnlocked:
		s.IsRecoveryLocked = false
	}

	s.PassportHash = e.PassportHash
	s.RecoveryEventID = e.ID
	return s
}
