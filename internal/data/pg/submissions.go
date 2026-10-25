package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const submissionTable = "finalizer_submission"

var submissionColumns = []string{"s.finalizer_submission_id", "s.recovery_event_id", "e.passport_hash",
	"s.status", "s.tx_hash", "s.error", "s.updated_at"}

type SubmissionsQ struct {
	db         *pgdb.DB
	contractID int64
}

func NewSubmissionsQ(db *pgdb.DB, contractID int64) SubmissionsQ {
	return SubmissionsQ{db: db, contractID: contractID}
}

// Claim returns 0 if the request has a blocking attempt or the passport has an open one.
func (q SubmissionsQ) Claim(requestID int64, passportHash data.Hash) (int64, error) {
	var id int64
	err := q.db.Transaction(func() error {
		// Serializes claims for the passport.
		var stateID int64
		err := q.db.Get(&stateID, sq.Select("recovery_state_id").From(stateTable).
			Where(sq.Eq{"contract_id": q.contractID, "passport_hash": passportHash, "is_active": true}).
			Suffix("FOR UPDATE"))
		if err != nil {
			return err
		}

		var blocking int
		err = q.db.Get(&blocking, q.selector("COUNT(*)").Where(sq.Or{
			sq.And{
				sq.Eq{"s.recovery_event_id": requestID},
				sq.NotEq{"s.status": []data.SubmissionStatus{data.SubmissionFailed, data.SubmissionRejected, data.SubmissionExpired}},
			},
			sq.Eq{
				"e.passport_hash": passportHash,
				"s.status":        []data.SubmissionStatus{data.SubmissionSubmitting, data.SubmissionSubmitted, data.SubmissionUnknown},
			},
		}))
		if err != nil || blocking > 0 {
			return err
		}

		return q.db.Get(&id, sq.Insert(submissionTable).
			Columns("recovery_event_id", "status").
			Values(requestID, data.SubmissionSubmitting).
			Suffix("RETURNING finalizer_submission_id"))
	})
	return id, err
}

func (q SubmissionsQ) SelectByStatus(statuses ...data.SubmissionStatus) ([]data.Submission, error) {
	var result []data.Submission
	err := q.db.Select(&result, q.selector(submissionColumns...).Where(sq.Eq{"s.status": statuses}).
		OrderBy("s.finalizer_submission_id"))
	return result, err
}

// SelectByRequest returns the attempts for the request, oldest first.
func (q SubmissionsQ) SelectByRequest(requestID int64) ([]data.Submission, error) {
	var result []data.Submission
	err := q.db.Select(&result, q.selector(submissionColumns...).Where(sq.Eq{"s.recovery_event_id": requestID}).
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

func (q SubmissionsQ) selector(columns ...string) sq.SelectBuilder {
	return sq.Select(columns...).
		From(submissionTable + " s").
		Join(eventTable + " e ON e.recovery_event_id = s.recovery_event_id").
		Where(sq.Eq{"e.contract_id": q.contractID, "s.is_active": true})
}
