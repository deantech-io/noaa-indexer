package pg

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

// LockChain serializes writers for the chain until the transaction ends.
func LockChain(db *pgdb.DB, chainID uint64) error {
	return db.ExecRaw("SELECT pg_advisory_xact_lock(?)", chainID)
}
