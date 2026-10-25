//go:build integration

// Package integration tests the services against anvil and Postgres.
//
//	TEST_DB_URL=postgres://... go test -tags integration -p 1 ./tests/integration/
//
// Requires anvil in PATH. Every test resets the database schema.
package integration

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/deantech-io/noaa-indexer/internal/cli"
	"github.com/deantech-io/noaa-indexer/internal/config"
	"github.com/deantech-io/noaa-indexer/internal/data"
	"github.com/deantech-io/noaa-indexer/internal/data/pg"
	"github.com/deantech-io/noaa-indexer/tests/integration/mock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/kv"
)

const (
	anvilChainID = 31337
	// anvil reports finalized = latest - 64.
	finalityDepth = 64

	deployerKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // anvil #0
	relayerKey  = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" // anvil #1
	walletKey   = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a" // anvil #2
)

type env struct {
	t           *testing.T
	rpcURL      string
	ec          *ethclient.Client
	contract    *mock.MockNoAARecovery
	address     common.Address
	deployBlock uint64
	auth        *bind.TransactOpts
	dbURL       string
}

func setup(t *testing.T, recoveryDelay, contestLock uint64) *env {
	t.Helper()

	dbURL := os.Getenv("TEST_DB_URL")
	if dbURL == "" {
		t.Skip("TEST_DB_URL is not set")
	}
	if _, err := exec.LookPath("anvil"); err != nil {
		t.Skip("anvil is not in PATH")
	}

	e := &env{t: t, dbURL: dbURL}
	e.startAnvil()

	key := mustKey(t, deployerKey)
	var err error
	if e.auth, err = bind.NewKeyedTransactorWithChainID(key, big.NewInt(anvilChainID)); err != nil {
		t.Fatal(err)
	}

	address, tx, contract, err := mock.DeployMockNoAARecovery(e.auth, e.ec, recoveryDelay, contestLock)
	if err != nil {
		t.Fatal(err)
	}
	e.address, e.contract = address, contract
	e.deployBlock = e.wait(tx).BlockNumber.Uint64()

	cfg := e.config(nil, nil)
	if err = cli.MigrateDown(cfg); err != nil {
		t.Fatal(err)
	}
	if err = cli.MigrateUp(cfg); err != nil {
		t.Fatal(err)
	}
	return e
}

func (e *env) startAnvil() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		e.t.Fatal(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	_ = l.Close()

	cmd := exec.Command("anvil", "--port", strconv.Itoa(port), "--chain-id", strconv.Itoa(anvilChainID), "--silent")
	if err = cmd.Start(); err != nil {
		e.t.Fatal(err)
	}
	e.t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	})

	e.rpcURL = fmt.Sprintf("http://127.0.0.1:%d", port)
	eventually(e.t, 10*time.Second, func() bool {
		if e.ec, err = ethclient.Dial(e.rpcURL); err != nil {
			return false
		}
		_, err = e.ec.BlockNumber(context.Background())
		return err == nil
	})
}

func (e *env) config(indexer, finalizer map[string]any) config.Config {
	return e.configWith(map[string]map[string]any{"indexer": indexer, "finalizer": finalizer})
}

// configWith overrides keys of the default config sections.
func (e *env) configWith(overrides map[string]map[string]any) config.Config {
	sections := map[string]map[string]any{
		"log":           {"level": "warn", "disable_sentry": true},
		"db":            {"url": e.dbURL},
		"network":       {"chain_id": anvilChainID, "rpc": []any{e.rpcURL}, "timeout": "5s", "retries": 1, "backoff": "100ms"},
		"noaa_recovery": {"address": e.address.Hex(), "block": e.deployBlock},
		"indexer":       {"poll_period": "100ms", "chunk_size": 1000},
		"finalizer":     {"period": "200ms", "mode": config.ModeWallet, "wallet_private_key": walletKey, "wallet_allowed_chain_ids": []any{anvilChainID}},
	}
	for section, values := range overrides {
		for k, v := range values {
			sections[section][k] = v
		}
	}

	return config.New(kv.GetterFunc(func(key string) (map[string]any, error) {
		return sections[key], nil
	}))
}

func (e *env) wait(tx *types.Transaction) *types.Receipt {
	e.t.Helper()
	r, err := bind.WaitMined(context.Background(), e.ec, tx)
	if err != nil {
		e.t.Fatal(err)
	}
	if r.Status != types.ReceiptStatusSuccessful {
		e.t.Fatalf("tx %s reverted", tx.Hash().Hex())
	}
	return r
}

func (e *env) send(tx *types.Transaction, err error) {
	e.t.Helper()
	if err != nil {
		e.t.Fatal(err)
	}
	e.wait(tx)
}

// request sends requestRecovery with dgCommit = key + 1000.
func (e *env) request(p common.Hash, key int64) {
	e.t.Helper()
	e.send(e.contract.RequestRecovery(e.auth, [32]byte{}, big.NewInt(key), big.NewInt(key+1000), p, [32]byte{}, nil))
}

// assertOnChain checks that the indexed state of p mirrors the contract storage.
func (e *env) assertOnChain(cfg config.Config, p common.Hash) {
	e.t.Helper()
	opts := &bind.CallOpts{}
	request, err := e.contract.GetRecoveryRequest(opts, p)
	if err != nil {
		e.t.Fatal(err)
	}
	lockedUntil, err := e.contract.LockedUntil(opts, p)
	if err != nil {
		e.t.Fatal(err)
	}
	locked, err := e.contract.IsRecoveryLocked(opts, p)
	if err != nil {
		e.t.Fatal(err)
	}

	s := e.state(cfg, p)
	if !s.OldIdentityKey.Equal(data.NumericFromBytes(request.OldIdentityKey)) ||
		!s.NewIdentityKey.Equal(data.NumericFromBytes(request.NewIdentityKey)) ||
		!s.DGCommit.Equal(data.NewNumeric(request.DgCommit)) ||
		!sameTime(s.ReadyAt, data.Time(request.ReadyAt)) || !sameTime(s.LockedUntil, data.Time(lockedUntil)) ||
		s.IsRecoveryLocked != locked {
		e.t.Errorf("state of %s = %+v, on chain %+v, locked until %d, locked %t", p.Hex(), s, request, lockedUntil, locked)
	}
}

func (e *env) rpc(method string, args ...any) {
	e.t.Helper()
	if err := e.ec.Client().Call(nil, method, args...); err != nil {
		e.t.Fatalf("%s: %v", method, err)
	}
}

func (e *env) mine(blocks int) {
	e.rpc("anvil_mine", fmt.Sprintf("0x%x", blocks))
}

// finalize mines past the finality depth and returns the newest finalized block.
func (e *env) finalize() uint64 {
	e.t.Helper()
	head, err := e.ec.BlockNumber(context.Background())
	if err != nil {
		e.t.Fatal(err)
	}
	e.mine(finalityDepth)
	return head
}

func (e *env) contractID(cfg config.Config) int64 {
	e.t.Helper()
	id, err := pg.NewContractQ(cfg.DB().Clone()).Ensure(anvilChainID, e.address)
	if err != nil {
		e.t.Fatal(err)
	}
	return id
}

func (e *env) checkpoint(cfg config.Config) uint64 {
	e.t.Helper()
	cp, err := pg.NewContractQ(cfg.DB()).Checkpoint(e.contractID(cfg))
	if err != nil {
		e.t.Fatal(err)
	}
	if cp == nil {
		return 0
	}
	return cp.BlockNumber
}

func (e *env) events(cfg config.Config) []data.RecoveryEvent {
	e.t.Helper()
	events, err := pg.NewEventsQ(cfg.DB(), anvilChainID, e.contractID(cfg)).Select()
	if err != nil {
		e.t.Fatal(err)
	}
	return events
}

func (e *env) state(cfg config.Config, passport common.Hash) data.RecoveryState {
	e.t.Helper()
	s, err := pg.NewStateQ(cfg.DB(), e.contractID(cfg)).Get(data.Hash(passport))
	if err != nil {
		e.t.Fatal(err)
	}
	if s == nil {
		e.t.Fatalf("no state for %s", passport.Hex())
	}
	return *s
}

func (e *env) submissions(cfg config.Config) []data.Submission {
	e.t.Helper()
	subs, err := pg.NewSubmissionsQ(cfg.DB(), e.contractID(cfg)).SelectByStatus(
		data.SubmissionSubmitting, data.SubmissionSubmitted, data.SubmissionConfirmed, data.SubmissionFailed,
		data.SubmissionRejected, data.SubmissionUnknown, data.SubmissionResolved, data.SubmissionExpired)
	if err != nil {
		e.t.Fatal(err)
	}
	return subs
}

func runInBackground(f func(ctx context.Context) error) func() error {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- f(ctx) }()
	return func() error {
		cancel()
		return <-done
	}
}

func eventually(t *testing.T, timeout time.Duration, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatal("condition not met in time")
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func mustKey(t *testing.T, hex string) *ecdsa.PrivateKey {
	t.Helper()
	key, err := crypto.HexToECDSA(hex)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func passport(n int64) common.Hash {
	return common.BigToHash(big.NewInt(n))
}
