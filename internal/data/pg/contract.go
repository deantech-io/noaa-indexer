package pg

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/rarimo/noaa-indexer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const contractTable = "contract"

type ContractQ struct {
	db *pgdb.DB
}

func NewContractQ(db *pgdb.DB) ContractQ {
	return ContractQ{db: db}
}

// Lock serializes writers for the chain until the transaction ends.
func (q ContractQ) Lock(chainID uint64) error {
	return q.db.ExecRaw("SELECT pg_advisory_xact_lock(?)", chainID)
}

func (q ContractQ) Ensure(chainID uint64, address common.Address) (int64, error) {
	var id int64
	err := q.db.Transaction(func() error {
		if err := q.Lock(chainID); err != nil {
			return err
		}

		err := q.db.Get(&id, sq.Select("contract_id").From(contractTable).
			Where(sq.Eq{"chain_id": chainID, "address": data.Address(address), "is_active": true}))
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		return q.db.Get(&id, sq.Insert(contractTable).
			Columns("chain_id", "address").
			Values(chainID, data.Address(address)).
			Suffix("RETURNING contract_id"))
	})
	return id, err
}

// Checkpoint returns nil if nothing was indexed yet.
func (q ContractQ) Checkpoint(contractID int64) (*data.Checkpoint, error) {
	var result data.Checkpoint
	err := q.db.Get(&result, sq.Select("b.block_number", "b.block_hash").
		From(contractTable+" c").
		Join("block b ON b.block_id = c.block_id").
		Where(sq.Eq{"c.contract_id": contractID}))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, err
}

func (q ContractQ) SetCheckpoint(contractID, blockID int64) error {
	return q.db.Exec(sq.Update(contractTable).Set("block_id", blockID).Where(sq.Eq{"contract_id": contractID}))
}

func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}
