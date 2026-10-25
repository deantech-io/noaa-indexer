package pg

import (
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const stateTable = "recovery_state"

var stateColumns = []string{"recovery_state_id", "recovery_event_id", "passport_hash", "old_identity_key",
	"new_identity_key", "dg_commit", "ready_at", "locked_until", "is_recovery_locked"}

type StateQ struct {
	db         *pgdb.DB
	contractID int64
}

func NewStateQ(db *pgdb.DB, contractID int64) StateQ {
	return StateQ{db: db, contractID: contractID}
}

// Get returns nil if the passport has no events yet.
func (q StateQ) Get(passportHash data.Hash) (*data.RecoveryState, error) {
	var result data.RecoveryState
	err := q.db.Get(&result, q.selector().Where(sq.Eq{"passport_hash": passportHash}))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, err
}

// SelectReady returns pending requests with ready_at at or before now.
func (q StateQ) SelectReady(now time.Time) ([]data.RecoveryState, error) {
	var result []data.RecoveryState
	err := q.db.Select(&result, q.selector().
		Where(sq.LtOrEq{"ready_at": now.UTC()}).
		OrderBy("ready_at"))
	return result, err
}

// Save inserts the state if ID is 0, otherwise updates it.
func (q StateQ) Save(s data.RecoveryState) error {
	if s.ID == 0 {
		return q.db.Exec(sq.Insert(stateTable).
			Columns("contract_id", "recovery_event_id", "passport_hash", "old_identity_key", "new_identity_key",
				"dg_commit", "ready_at", "locked_until", "is_recovery_locked").
			Values(q.contractID, s.RecoveryEventID, s.PassportHash, s.OldIdentityKey, s.NewIdentityKey,
				s.DGCommit, s.ReadyAt, s.LockedUntil, s.IsRecoveryLocked))
	}

	return q.db.Exec(sq.Update(stateTable).
		Set("recovery_event_id", s.RecoveryEventID).
		Set("old_identity_key", s.OldIdentityKey).
		Set("new_identity_key", s.NewIdentityKey).
		Set("dg_commit", s.DGCommit).
		Set("ready_at", s.ReadyAt).
		Set("locked_until", s.LockedUntil).
		Set("is_recovery_locked", s.IsRecoveryLocked).
		Where(sq.Eq{"recovery_state_id": s.ID}))
}

func (q StateQ) selector() sq.SelectBuilder {
	return sq.Select(stateColumns...).From(stateTable).
		Where(sq.Eq{"contract_id": q.contractID, "is_active": true})
}
