package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/deantech-io/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const eventTable = "recovery_event"

type EventsQ struct {
	db         *pgdb.DB
	blocks     BlockQ
	contractID int64
}

func NewEventsQ(db *pgdb.DB, chainID uint64, contractID int64) EventsQ {
	return EventsQ{db: db, blocks: NewBlockQ(db, chainID), contractID: contractID}
}

// Insert stores the event with its block and transaction and reports whether it was new.
func (q EventsQ) Insert(e *data.RecoveryEvent) (bool, error) {
	blockID, err := q.blocks.Ensure(e.BlockNumber, e.BlockHash, e.MinedAt)
	if err != nil {
		return false, err
	}
	txID, err := q.blocks.EnsureTransaction(blockID, e.TxHash)
	if err != nil {
		return false, err
	}

	err = q.db.Get(&e.ID, sq.Select("recovery_event_id").From(eventTable).
		Where(sq.Eq{"transaction_id": txID, "log_index": e.LogIndex, "is_active": true}))
	if err == nil {
		return false, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	err = q.db.Get(&e.ID, sq.Insert(eventTable).
		Columns("contract_id", "transaction_id", "log_index", "event_type", "passport_hash",
			"old_identity_key", "new_identity_key", "dg_commit", "ready_at", "locked_until").
		Values(q.contractID, txID, e.LogIndex, e.Type, e.PassportHash,
			e.OldIdentityKey, e.NewIdentityKey, e.DGCommit, e.ReadyAt, e.LockedUntil).
		Suffix("RETURNING recovery_event_id"))
	return err == nil, err
}

// Select returns all events in chain order.
func (q EventsQ) Select() ([]data.RecoveryEvent, error) {
	var result []data.RecoveryEvent
	err := q.db.Select(&result, q.selector().OrderBy("b.block_number", "e.log_index"))
	return result, err
}

// LatestRequest returns the last RecoveryRequested for the passport, or nil.
func (q EventsQ) LatestRequest(passportHash data.Hash) (*data.RecoveryEvent, error) {
	var result data.RecoveryEvent
	err := q.db.Get(&result, q.selector().
		Where(sq.Eq{"e.passport_hash": passportHash, "e.event_type": data.EventRequested}).
		OrderBy("b.block_number DESC", "e.log_index DESC").
		Limit(1))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, err
}

func (q EventsQ) selector() sq.SelectBuilder {
	return sq.Select("e.recovery_event_id", "t.tx_hash", "e.log_index", "b.block_number", "b.block_hash",
		"b.mined_at", "e.event_type", "e.passport_hash", "e.old_identity_key", "e.new_identity_key", "e.dg_commit",
		"e.ready_at", "e.locked_until").
		From(eventTable + " e").
		Join(transactionTable + " t ON t.transaction_id = e.transaction_id").
		Join(blockTable + " b ON b.block_id = t.block_id").
		Where(sq.Eq{"e.contract_id": q.contractID, "e.is_active": true})
}
