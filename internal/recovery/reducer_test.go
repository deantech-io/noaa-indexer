package recovery

import (
	"math/big"
	"testing"
	"time"

	"github.com/rarimo/noaa-indexer/internal/data"
)

func requested(key int64, readyAt uint64) data.RecoveryEvent {
	return data.RecoveryEvent{
		Type:           data.EventRequested,
		OldIdentityKey: data.NewNumeric(big.NewInt(1)),
		NewIdentityKey: data.NewNumeric(big.NewInt(key)),
		DGCommit:       data.NewNumeric(big.NewInt(5)),
		ReadyAt:        data.Time(readyAt),
	}
}

// dropping is a contest or a cancel with the lockedUntil read after it.
func dropping(t data.EventType, lockedUntil uint64) data.RecoveryEvent {
	return data.RecoveryEvent{Type: t, LockedUntil: data.Time(lockedUntil)}
}

func ev(t data.EventType) data.RecoveryEvent {
	return data.RecoveryEvent{Type: t}
}

func at(e data.RecoveryEvent, minedAt int64) data.RecoveryEvent {
	e.MinedAt = time.Unix(minedAt, 0).UTC()
	return e
}

func unix(t *time.Time) uint64 {
	if t == nil {
		return 0
	}
	return uint64(t.Unix())
}

func TestApply(t *testing.T) {
	tests := []struct {
		name        string
		events      []data.RecoveryEvent
		key         int64
		readyAt     uint64
		lockedUntil uint64
		locked      bool
	}{
		{"request", []data.RecoveryEvent{requested(7, 100)}, 7, 100, 0, false},
		{"cancel", []data.RecoveryEvent{requested(7, 100), dropping(data.EventCancelled, 500)}, 0, 0, 500, false},
		{"contest", []data.RecoveryEvent{requested(7, 100), dropping(data.EventContested, 500)}, 0, 0, 500, false},
		{"request after contest lock", []data.RecoveryEvent{requested(7, 100), dropping(data.EventContested, 500), requested(9, 600)}, 9, 600, 500, false},
		{"lock drops request", []data.RecoveryEvent{requested(7, 100), dropping(data.EventCancelled, 0), ev(data.EventLocked)}, 0, 0, 0, true},
		{"unlock keeps locked until", []data.RecoveryEvent{requested(7, 100), dropping(data.EventContested, 500), ev(data.EventLocked), ev(data.EventUnlocked)}, 0, 0, 500, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s data.RecoveryState
			for i, e := range tt.events {
				e.ID = int64(i + 1)
				s = Apply(s, e)
			}

			if !s.NewIdentityKey.Equal(data.NewNumeric(big.NewInt(tt.key))) {
				t.Errorf("new identity key = %v, want %d", s.NewIdentityKey.Int, tt.key)
			}
			if pending := tt.key != 0; pending != (s.OldIdentityKey.Int != nil && s.DGCommit.Int != nil) {
				t.Errorf("old identity key = %v, dg commit = %v, want them set only with a request", s.OldIdentityKey.Int, s.DGCommit.Int)
			}
			if unix(s.ReadyAt) != tt.readyAt {
				t.Errorf("ready at = %d, want %d", unix(s.ReadyAt), tt.readyAt)
			}
			if unix(s.LockedUntil) != tt.lockedUntil {
				t.Errorf("locked until = %d, want %d", unix(s.LockedUntil), tt.lockedUntil)
			}
			if s.IsRecoveryLocked != tt.locked {
				t.Errorf("is recovery locked = %t, want %t", s.IsRecoveryLocked, tt.locked)
			}
			if s.RecoveryEventID != int64(len(tt.events)) {
				t.Errorf("recovery event id = %d, want %d", s.RecoveryEventID, len(tt.events))
			}
		})
	}
}

func TestExpected(t *testing.T) {
	tests := []struct {
		name   string
		before []data.RecoveryEvent
		event  data.RecoveryEvent
		want   bool
	}{
		{"request while pending", []data.RecoveryEvent{requested(7, 100)}, at(requested(8, 100), 40), false},
		{"request while locked", []data.RecoveryEvent{ev(data.EventLocked)}, at(requested(7, 100), 40), false},
		{"request before locked until", []data.RecoveryEvent{dropping(data.EventCancelled, 500)}, at(requested(7, 600), 499), false},
		{"request at locked until", []data.RecoveryEvent{dropping(data.EventCancelled, 500)}, at(requested(7, 600), 500), true},
		{"finalize before ready", []data.RecoveryEvent{requested(7, 100)}, at(ev(data.EventFinalized), 99), false},
		{"cancel without request", nil, dropping(data.EventCancelled, 500), false},
		{"lock twice", []data.RecoveryEvent{ev(data.EventLocked)}, ev(data.EventLocked), false},
		{"unlock without lock", nil, ev(data.EventUnlocked), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s data.RecoveryState
			for _, e := range tt.before {
				s = Apply(s, e)
			}
			if got := Expected(s, tt.event); got != tt.want {
				t.Errorf("Expected = %t, want %t", got, tt.want)
			}
		})
	}
}
