package data

import (
	"time"
)

type EventType string

const (
	EventRequested EventType = "requested"
	EventCancelled EventType = "cancelled"
	EventContested EventType = "contested"
	EventFinalized EventType = "finalized"
	EventLocked    EventType = "locked"
	EventUnlocked  EventType = "unlocked"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusPending   Status = "pending"
	StatusContested Status = "contested"
	StatusLocked    Status = "locked"
)

type SubmissionStatus string

const (
	SubmissionSubmitting SubmissionStatus = "submitting"
	SubmissionSubmitted  SubmissionStatus = "submitted"
	SubmissionConfirmed  SubmissionStatus = "confirmed"
	SubmissionFailed     SubmissionStatus = "failed"
	SubmissionRejected   SubmissionStatus = "rejected"
	SubmissionUnknown    SubmissionStatus = "unknown"
	SubmissionResolved   SubmissionStatus = "resolved"
	SubmissionExpired    SubmissionStatus = "expired"
)

type Checkpoint struct {
	BlockNumber uint64 `db:"block_number"`
	BlockHash   Hash   `db:"block_hash"`
}

// RecoveryEvent is a recovery_event row joined with its block.
type RecoveryEvent struct {
	ID             int64      `db:"recovery_event_id"`
	TxHash         Hash       `db:"tx_hash"`
	LogIndex       uint64     `db:"log_index"`
	BlockNumber    uint64     `db:"block_number"`
	BlockHash      Hash       `db:"block_hash"`
	MinedAt        time.Time  `db:"mined_at"`
	Type           EventType  `db:"event_type"`
	PassportHash   Hash       `db:"passport_hash"`
	NewIdentityKey Numeric    `db:"new_identity_key"`
	ReadyAt        *time.Time `db:"ready_at"`
	LockedUntil    *time.Time `db:"locked_until"`
}

type RecoveryState struct {
	ID              int64      `db:"recovery_state_id"`
	RecoveryEventID int64      `db:"recovery_event_id"`
	PassportHash    Hash       `db:"passport_hash"`
	Status          Status     `db:"status"`
	NewIdentityKey  Numeric    `db:"new_identity_key"`
	ReadyAt         *time.Time `db:"ready_at"`
	LockedUntil     *time.Time `db:"locked_until"`
}

// Submission is a finalizer_submission row with the passport of its request event.
type Submission struct {
	ID              int64            `db:"finalizer_submission_id"`
	RecoveryEventID int64            `db:"recovery_event_id"`
	PassportHash    Hash             `db:"passport_hash"`
	Status          SubmissionStatus `db:"status"`
	TxHash          *Hash            `db:"tx_hash"`
	Error           string           `db:"error"`
	UpdatedAt       time.Time        `db:"updated_at"`
}
