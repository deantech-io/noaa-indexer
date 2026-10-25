package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const submissionTable = "finalizer_submission"

type SubmissionsQ struct {
	db         *pgdb.DB
	contractID int64
}

func NewSubmissionsQ(db *pgdb.DB, contractID int64) SubmissionsQ {
	return SubmissionsQ{db: db, contractID: contractID}
}

func (q SubmissionsQ) Blocked(requestID int64) (bool, error) {
	var n int
	err := q.db.Get(&n, sq.Select("COUNT(*)").From(submissionTable).
		Where(sq.Eq{"recovery_event_id": requestID, "is_active": true}).
		Where(sq.NotEq{"status": []data.SubmissionStatus{data.SubmissionFailed, data.SubmissionRejected, data.SubmissionExpired}}))
	return n > 0, err
}

// Claim returns 0 if another attempt for the request is already claimed.
func (q SubmissionsQ) Claim(requestID int64) (int64, error) {
	var id int64
	err := q.db.Get(&id, sq.Insert(submissionTable).
		Columns("recovery_event_id", "status").
		Values(requestID, data.SubmissionSubmitting).
		Suffix("RETURNING finalizer_submission_id"))
	if isUniqueViolation(err) {
		return 0, nil
	}
	return id, err
}

func (q SubmissionsQ) SelectByStatus(statuses ...data.SubmissionStatus) ([]data.Submission, error) {
	var result []data.Submission
	err := q.db.Select(&result, sq.Select("s.finalizer_submission_id", "s.recovery_event_id", "e.passport_hash",
		"s.status", "s.tx_hash", "s.error", "s.updated_at").
		From(submissionTable+" s").
		Join(eventTable+" e ON e.recovery_event_id = s.recovery_event_id").
		Where(sq.Eq{"e.contract_id": q.contractID, "s.status": statuses, "s.is_active": true}).
		OrderBy("s.finalizer_submission_id"))
	return result, err
}

// Update reports false if the submission is no longer in one of the from statuses.
func (q SubmissionsQ) Update(id int64, to data.SubmissionStatus, txHash *data.Hash, errMsg string, from ...data.SubmissionStatus) (bool, error) {
	upd := sq.Update(submissionTable).
		Set("status", to).
		Set("error", errMsg).
		Where(sq.Eq{"finalizer_submission_id": id, "status": from}).
		Suffix("RETURNING finalizer_submission_id")
	if txHash != nil {
		upd = upd.Set("tx_hash", *txHash)
	}
	err := q.db.Get(&id, upd)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil, err
}
