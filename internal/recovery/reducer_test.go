package recovery

import (
	"math/big"
	"testing"
	"time"

	"github.com/rarimo/noaa-indexer/internal/data"
)

func requested(key int64, readyAt uint64) data.RecoveryEvent {
	return data.RecoveryEvent{Type: data.EventRequested, NewIdentityKey: data.NewNumeric(big.NewInt(key)), ReadyAt: data.Time(readyAt)}
}

func contested(lockedUntil uint64) data.RecoveryEvent {
	return data.RecoveryEvent{Type: data.EventContested, LockedUntil: data.Time(lockedUntil)}
}

func ev(t data.EventType) data.RecoveryEvent {
	return data.RecoveryEvent{Type: t}
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
		status      data.Status
		key         int64
		readyAt     uint64
		lockedUntil uint64
	}{
		{"request", []data.RecoveryEvent{requested(7, 100)}, data.StatusPending, 7, 100, 0},
		{"cancel", []data.RecoveryEvent{requested(7, 100), ev(data.EventCancelled)}, data.StatusActive, 0, 0, 0},
		{"finalize", []data.RecoveryEvent{requested(7, 100), ev(data.EventFinalized)}, data.StatusActive, 0, 0, 0},
		{"contest", []data.RecoveryEvent{requested(7, 100), contested(500)}, data.StatusContested, 0, 0, 500},
		{"request after contest lock", []data.RecoveryEvent{requested(7, 100), contested(500), requested(9, 600)}, data.StatusPending, 9, 600, 0},
		{"lock active", []data.RecoveryEvent{ev(data.EventLocked)}, data.StatusLocked, 0, 0, 0},
		{"unlock active", []data.RecoveryEvent{ev(data.EventLocked), ev(data.EventUnlocked)}, data.StatusActive, 0, 0, 0},
		{"lock keeps request", []data.RecoveryEvent{requested(7, 100), ev(data.EventLocked)}, data.StatusLocked, 7, 100, 0},
		{"unlock restores request", []data.RecoveryEvent{requested(7, 100), ev(data.EventLocked), ev(data.EventUnlocked)}, data.StatusPending, 7, 100, 0},
		{"unlock restores contest", []data.RecoveryEvent{requested(7, 100), contested(500), ev(data.EventLocked), ev(data.EventUnlocked)}, data.StatusContested, 0, 0, 500},
		{"cancel while locked", []data.RecoveryEvent{requested(7, 100), ev(data.EventLocked), ev(data.EventCancelled)}, data.StatusLocked, 0, 0, 0},
		{"unexpected finalize on empty state", []data.RecoveryEvent{ev(data.EventFinalized)}, data.StatusActive, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s data.RecoveryState
			for i, e := range tt.events {
				e.ID = int64(i + 1)
				s = Apply(s, e)
			}

			if s.Status != tt.status {
				t.Errorf("status = %s, want %s", s.Status, tt.status)
			}
			if !s.NewIdentityKey.Equal(data.NewNumeric(big.NewInt(tt.key))) {
				t.Errorf("new identity key = %v, want %d", s.NewIdentityKey.Int, tt.key)
			}
			if unix(s.ReadyAt) != tt.readyAt {
				t.Errorf("ready at = %d, want %d", unix(s.ReadyAt), tt.readyAt)
			}
			if unix(s.LockedUntil) != tt.lockedUntil {
				t.Errorf("locked until = %d, want %d", unix(s.LockedUntil), tt.lockedUntil)
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
		{"request", nil, requested(7, 100), true},
		{"request while pending", []data.RecoveryEvent{requested(7, 100)}, requested(8, 100), false},
		{"request while locked", []data.RecoveryEvent{ev(data.EventLocked)}, requested(7, 100), false},
		{"request after contest", []data.RecoveryEvent{requested(7, 100), contested(500)}, requested(9, 600), true},
		{"contest", []data.RecoveryEvent{requested(7, 100)}, contested(500), true},
		{"contest without request", nil, contested(500), false},
		{"finalize", []data.RecoveryEvent{requested(7, 100)}, ev(data.EventFinalized), true},
		{"finalize without request", nil, ev(data.EventFinalized), false},
		{"finalize while locked", []data.RecoveryEvent{requested(7, 100), ev(data.EventLocked)}, ev(data.EventFinalized), false},
		{"cancel while locked", []data.RecoveryEvent{requested(7, 100), ev(data.EventLocked)}, ev(data.EventCancelled), true},
		{"cancel without request", nil, ev(data.EventCancelled), false},
		{"lock twice", []data.RecoveryEvent{ev(data.EventLocked)}, ev(data.EventLocked), false},
		{"unlock", []data.RecoveryEvent{ev(data.EventLocked)}, ev(data.EventUnlocked), true},
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
