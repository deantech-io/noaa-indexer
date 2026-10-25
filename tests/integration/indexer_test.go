//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/deantech-io/noaa-indexer/internal/data"
	"github.com/deantech-io/noaa-indexer/internal/data/pg"
	"github.com/deantech-io/noaa-indexer/internal/service/indexer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	p1, p2, p3, p4, p5 := passport(1), passport(2), passport(3), passport(4), passport(5)

	e.request(p1, 11)
	e.mine(5)
	e.request(p2, 22)
	e.send(e.contract.CancelRecovery(e.auth, p2, nil))
	e.mine(5)
	e.request(p3, 33)
	e.request(p3, 34) // contests
	e.request(p4, 44)
	e.send(e.contract.LockRecovery(e.auth, p4, nil))
	e.request(p5, 55)
	e.rpc("evm_increaseTime", 61)
	e.send(e.contract.FinalizeRecovery(e.auth, p5))
	// Lock and unlock while the cancel's time lock is active.
	e.send(e.contract.LockRecovery(e.auth, p2, nil))
	e.send(e.contract.UnlockRecovery(e.auth, p2, nil))
	finalized := e.finalize()

	// Small chunks force many eth_getLogs ranges.
	cfg := e.config(map[string]interface{}{"chunk_size": 3}, nil)
	stop := runInBackground(indexer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	// The lock also emits RecoveryCancelled for the dropped request.
	if n := len(e.events(cfg)); n != 12 {
		t.Fatalf("events = %d, want 12", n)
	}
	for _, p := range []common.Hash{p1, p2, p3, p4, p5} {
		e.assertOnChain(cfg, p)
	}
	assertState(t, e.state(cfg, p1), 11, false, false)
	assertState(t, e.state(cfg, p2), 0, true, false)
	assertState(t, e.state(cfg, p3), 0, true, false)
	assertState(t, e.state(cfg, p4), 0, false, true)
	assertState(t, e.state(cfg, p5), 0, false, false)
	if s := e.state(cfg, p1); s.DGCommit.Int64() != 1011 || s.OldIdentityKey.Int64() != 1 {
		t.Errorf("state of p1 = %+v, want dg commit 1011, old identity key 1", s)
	}

	t.Run("restart from checkpoint", func(t *testing.T) {
		checkpoint := e.checkpoint(cfg)
		e.send(e.contract.CancelRecovery(e.auth, p1, nil))
		finalized := e.finalize()

		stop := runInBackground(indexer.New(cfg).Run)
		eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
		if err := stop(); err != nil {
			t.Fatal(err)
		}

		if e.checkpoint(cfg) <= checkpoint {
			t.Fatal("checkpoint did not advance")
		}
		if n := len(e.events(cfg)); n != 13 {
			t.Fatalf("events = %d, want 13", n)
		}
		e.assertOnChain(cfg, p1)
		assertState(t, e.state(cfg, p1), 0, true, false)
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

		if n := len(e.events(cfg)); n != 13 {
			t.Fatalf("events = %d, want 13", n)
		}
		// Stored events must not be applied again.
		for _, p := range []common.Hash{p1, p2, p3, p4, p5} {
			e.assertOnChain(cfg, p)
		}
	})

	t.Run("rebuild from events", func(t *testing.T) {
		before := map[common.Hash]data.RecoveryState{}
		for _, p := range []common.Hash{p1, p2, p3, p4, p5} {
			before[p] = e.state(cfg, p)
		}
		if err := cfg.DB().ExecRaw("UPDATE recovery_state SET old_identity_key = NULL, new_identity_key = NULL, dg_commit = NULL, " +
			"ready_at = NULL, locked_until = NULL, is_recovery_locked = FALSE"); err != nil {
			t.Fatal(err)
		}

		if err := indexer.Rebuild(cfg); err != nil {
			t.Fatal(err)
		}

		for p, want := range before {
			got := e.state(cfg, p)
			want.ID = got.ID
			if !got.OldIdentityKey.Equal(want.OldIdentityKey) || !got.NewIdentityKey.Equal(want.NewIdentityKey) ||
				!got.DGCommit.Equal(want.DGCommit) || !sameTime(got.ReadyAt, want.ReadyAt) ||
				!sameTime(got.LockedUntil, want.LockedUntil) || got.IsRecoveryLocked != want.IsRecoveryLocked ||
				got.RecoveryEventID != want.RecoveryEventID {
				t.Errorf("state of %s = %+v, want %+v", p.Hex(), got, want)
			}
		}
	})
}

// The request is dropped within its block, so its old identity key and dg commit cannot be read.
func TestIndexerRequestContestedInSameBlock(t *testing.T) {
	e := setup(t, 60, 600)
	p := passport(1)

	e.rpc("evm_setAutomine", false)
	auth := *e.auth
	auth.GasLimit = 500_000
	txs := make([]*types.Transaction, 0, 2)
	for _, key := range []int64{11, 12} {
		tx, err := e.contract.RequestRecovery(&auth, [32]byte{}, big.NewInt(key), big.NewInt(key+1000), p, [32]byte{}, nil)
		if err != nil {
			t.Fatal(err)
		}
		txs = append(txs, tx)
	}
	e.rpc("evm_mine")
	e.rpc("evm_setAutomine", true)
	if e.wait(txs[0]).BlockNumber.Cmp(e.wait(txs[1]).BlockNumber) != 0 {
		t.Fatal("transactions were mined in different blocks")
	}
	finalized := e.finalize()

	cfg := e.config(nil, nil)
	stop := runInBackground(indexer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	events := e.events(cfg)
	if len(events) != 2 || events[0].Type != data.EventRequested || events[1].Type != data.EventContested {
		t.Fatalf("events = %+v, want requested and contested", events)
	}
	if events[0].OldIdentityKey.Int != nil || events[0].DGCommit.Int != nil {
		t.Errorf("requested event = %+v, want no old identity key and dg commit", events[0])
	}
	e.assertOnChain(cfg, p)
}

func TestIndexerHaltsOnParentHashMismatch(t *testing.T) {
	e := setup(t, 60, 600)

	var snapshot string
	if err := e.ec.Client().Call(&snapshot, "evm_snapshot"); err != nil {
		t.Fatal(err)
	}

	e.request(passport(1), 11)
	finalized := e.finalize()

	cfg := e.config(nil, nil)
	stop := runInBackground(indexer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	// Replace history below the checkpoint.
	e.rpc("evm_revert", snapshot)
	e.request(passport(2), 22)
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

func assertState(t *testing.T, s data.RecoveryState, key int64, lockedUntil, locked bool) {
	t.Helper()
	if !s.NewIdentityKey.Equal(data.NewNumeric(big.NewInt(key))) || (s.LockedUntil != nil) != lockedUntil || s.IsRecoveryLocked != locked {
		t.Errorf("state of %s = %+v, want key %d, locked until set %t, locked %t", s.PassportHash.Hex(), s, key, lockedUntil, locked)
	}
}

func TestCheckpointCompareAndSet(t *testing.T) {
	e := setup(t, 60, 600)
	cfg := e.config(nil, nil)
	contractID := e.contractID(cfg)
	contracts, blocks := pg.NewContractQ(cfg.DB()), pg.NewBlockQ(cfg.DB(), anvilChainID)

	b1, err := blocks.Ensure(1, data.Hash(common.Hash{1}), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	b2, err := blocks.Ensure(2, data.Hash(common.Hash{2}), time.Now())
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range []struct {
		prev  *int64
		block int64
		want  bool
	}{
		{nil, b1, true},
		{nil, b2, false},
		{&b2, b2, false},
		{&b1, b2, true},
	} {
		moved, err := contracts.SetCheckpoint(contractID, c.prev, c.block)
		if err != nil {
			t.Fatal(err)
		}
		if moved != c.want {
			t.Errorf("SetCheckpoint(prev %v, %d) = %t, want %t", c.prev, c.block, moved, c.want)
		}
	}
	if got := e.checkpoint(cfg); got != 2 {
		t.Errorf("checkpoint = %d, want 2", got)
	}
}
