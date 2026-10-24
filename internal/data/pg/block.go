package pg

import (
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	blockTable       = "block"
	transactionTable = "transaction"
)

type BlockQ struct {
	db      *pgdb.DB
	chainID uint64
}

func NewBlockQ(db *pgdb.DB, chainID uint64) BlockQ {
	return BlockQ{db: db, chainID: chainID}
}

func (q BlockQ) Ensure(number uint64, hash data.Hash, minedAt time.Time) (int64, error) {
	var id int64
	err := q.db.Get(&id, sq.Select("block_id").From(blockTable).
		Where(sq.Eq{"chain_id": q.chainID, "block_hash": hash, "is_active": true}))
	if !errors.Is(err, sql.ErrNoRows) {
		return id, err
	}

	err = q.db.Get(&id, sq.Insert(blockTable).
		Columns("chain_id", "block_number", "block_hash", "mined_at").
		Values(q.chainID, number, hash, minedAt.UTC()).
		Suffix("RETURNING block_id"))
	return id, err
}

func (q BlockQ) EnsureTransaction(blockID int64, hash data.Hash) (int64, error) {
	var id int64
	err := q.db.Get(&id, sq.Select("transaction_id").From(transactionTable).
		Where(sq.Eq{"block_id": blockID, "tx_hash": hash, "is_active": true}))
	if !errors.Is(err, sql.ErrNoRows) {
		return id, err
	}

	err = q.db.Get(&id, sq.Insert(transactionTable).
		Columns("block_id", "tx_hash").
		Values(blockID, hash).
		Suffix("RETURNING transaction_id"))
	return id, err
}
