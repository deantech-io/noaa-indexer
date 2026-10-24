//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/noaa-indexer/internal/data"
	"github.com/rarimo/noaa-indexer/internal/service/indexer"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Prague headers include requestsHash, which older go-ethereum versions hash differently.
func TestHeaderHashMatchesNode(t *testing.T) {
	e := setup(t, 60, 600)
	e.mine(3)

	var raw struct {
		Hash         common.Hash  `json:"hash"`
		RequestsHash *common.Hash `json:"requestsHash"`
	}
	var header types.Header
	for _, out := range []interface{}{&raw, &header} {
		var msg json.RawMessage
		if err := e.ec.Client().Call(&msg, "eth_getBlockByNumber", "latest", false); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(msg, out); err != nil {
			t.Fatal(err)
		}
	}

	if raw.RequestsHash == nil {
		t.Fatal("node did not return a Prague header (no requestsHash)")
	}
	if header.Hash() != raw.Hash {
		t.Fatalf("computed header hash %s, node returned %s", header.Hash().Hex(), raw.Hash.Hex())
	}
}

func TestIndexerBackfillRestartIdempotency(t *testing.T) {
	e := setup(t, 60, 600)
	p1, p2, p3, p4 := passport(1), passport(2), passport(3), passport(4)

	e.send(e.contract.RequestRecovery(e.auth, p1, big.NewInt(11)))
	e.mine(5)
	e.send(e.contract.RequestRecovery(e.auth, p2, big.NewInt(22)))
	e.send(e.contract.CancelRecovery(e.auth, p2))
	e.mine(5)
	e.send(e.contract.RequestRecovery(e.auth, p3, big.NewInt(33)))
	e.send(e.contract.RequestRecovery(e.auth, p3, big.NewInt(34))) // contests
	e.send(e.contract.RequestRecovery(e.auth, p4, big.NewInt(44)))
	e.send(e.contract.LockRecovery(e.auth, p4))
	finalized := e.finalize()

	// Small chunks force many eth_getLogs ranges.
	cfg := e.config(map[string]interface{}{"chunk_size": 3}, nil)
	stop := runInBackground(indexer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	if n := len(e.events(cfg)); n != 7 {
		t.Fatalf("events = %d, want 7", n)
	}
	assertState(t, e.state(cfg, p1), data.StatusPending, 11)
	assertState(t, e.state(cfg, p2), data.StatusActive, 0)
	assertState(t, e.state(cfg, p3), data.StatusContested, 0)
	assertState(t, e.state(cfg, p4), data.StatusLocked, 44)
	if s := e.state(cfg, p3); s.LockedUntil == nil {
		t.Error("contested state has no locked_until")
	}

	t.Run("restart from checkpoint", func(t *testing.T) {
		checkpoint := e.checkpoint(cfg)
		e.send(e.contract.CancelRecovery(e.auth, p1))
		finalized := e.finalize()

		stop := runInBackground(indexer.New(cfg).Run)
		eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
		if err := stop(); err != nil {
			t.Fatal(err)
		}

		if e.checkpoint(cfg) <= checkpoint {
			t.Fatal("checkpoint did not advance")
		}
		if n := len(e.events(cfg)); n != 8 {
			t.Fatalf("events = %d, want 8", n)
		}
		assertState(t, e.state(cfg, p1), data.StatusActive, 0)
	})

	t.Run("re-indexing is idempotent", func(t *testing.T) {
		checkpoint := e.checkpoint(cfg)
		if err := cfg.DB().ExecRaw("UPDATE contract SET block_id = NULL"); err != nil {
			t.Fatal(err)
		}

		stop := runInBackground(indexer.New(cfg).Run)
		eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= checkpoint })
		if err := stop(); err != nil {
			t.Fatal(err)
		}

		if n := len(e.events(cfg)); n != 8 {
			t.Fatalf("events = %d, want 8", n)
		}
		// Stored events must not be applied again.
		assertState(t, e.state(cfg, p1), data.StatusActive, 0)
		assertState(t, e.state(cfg, p4), data.StatusLocked, 44)
	})

	t.Run("rebuild from events", func(t *testing.T) {
		before := map[common.Hash]data.RecoveryState{}
		for _, p := range []common.Hash{p1, p2, p3, p4} {
			before[p] = e.state(cfg, p)
		}
		if err := cfg.DB().ExecRaw("UPDATE recovery_state SET status = 'active', new_identity_key = NULL, ready_at = NULL, locked_until = NULL"); err != nil {
			t.Fatal(err)
		}

		if err := indexer.Rebuild(cfg); err != nil {
			t.Fatal(err)
		}

		for p, want := range before {
			got := e.state(cfg, p)
			if got.Status != want.Status || !got.NewIdentityKey.Equal(want.NewIdentityKey) ||
				!sameTime(got.ReadyAt, want.ReadyAt) || !sameTime(got.LockedUntil, want.LockedUntil) ||
				got.RecoveryEventID != want.RecoveryEventID {
				t.Errorf("state of %s = %+v, want %+v", p.Hex(), got, want)
			}
		}
	})
}

func TestIndexerHaltsOnParentHashMismatch(t *testing.T) {
	e := setup(t, 60, 600)

	var snapshot string
	if err := e.ec.Client().Call(&snapshot, "evm_snapshot"); err != nil {
		t.Fatal(err)
	}

	e.send(e.contract.RequestRecovery(e.auth, passport(1), big.NewInt(11)))
	finalized := e.finalize()

	cfg := e.config(nil, nil)
	stop := runInBackground(indexer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	// Replace history below the checkpoint.
	e.rpc("evm_revert", snapshot)
	e.send(e.contract.RequestRecovery(e.auth, passport(2), big.NewInt(22)))
	e.mine(2 * finalityDepth)

	ix := indexer.New(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := ix.Run(ctx)
	if errors.Cause(err) != indexer.ErrParentHashMismatch {
		t.Fatalf("Run() error = %v, want parent hash mismatch", err)
	}
	if ix.Healthy() {
		t.Error("indexer reports healthy after halting")
	}
	if n := len(e.events(cfg)); n != 1 {
		t.Errorf("events = %d, want 1 (nothing indexed after the mismatch)", n)
	}
}

func sameTime(a, b *time.Time) bool {
	return (a == nil && b == nil) || (a != nil && b != nil && a.Equal(*b))
}

func assertState(t *testing.T, s data.RecoveryState, status data.Status, key int64) {
	t.Helper()
	if s.Status != status || !s.NewIdentityKey.Equal(data.NewNumeric(big.NewInt(key))) {
		t.Errorf("state of %s = %s/%v, want %s/%d", s.PassportHash.Hex(), s.Status, s.NewIdentityKey.Int, status, key)
	}
}
