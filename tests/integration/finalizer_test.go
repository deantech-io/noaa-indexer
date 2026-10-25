//go:build integration

package integration

import (
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rarimo/noaa-indexer/internal/config"
	"github.com/rarimo/noaa-indexer/internal/data"
	"github.com/rarimo/noaa-indexer/internal/data/pg"
	"github.com/rarimo/noaa-indexer/internal/service/finalizer"
	"github.com/rarimo/noaa-indexer/internal/service/indexer"
)

const recoveryDelay = 100

func (e *env) readyAndIndexed(cfg config.Config, keys map[common.Hash]int64) {
	e.t.Helper()
	for p, k := range keys {
		e.send(e.contract.RequestRecovery(e.auth, p, big.NewInt(k)))
	}
	e.rpc("evm_increaseTime", recoveryDelay+1)
	finalized := e.finalize()

	stop := runInBackground(indexer.New(cfg).Run)
	eventually(e.t, 30*time.Second, func() bool { return e.checkpoint(cfg) >= finalized })
	if err := stop(); err != nil {
		e.t.Fatal(err)
	}
}

func (e *env) identityKey(p common.Hash) int64 {
	e.t.Helper()
	k, err := e.contract.IdentityKey(&bind.CallOpts{}, p)
	if err != nil {
		e.t.Fatal(err)
	}
	return k.Int64()
}

func submissionsFor(subs []data.Submission, p common.Hash) []data.Submission {
	var res []data.Submission
	for _, s := range subs {
		if s.PassportHash == data.Hash(p) {
			res = append(res, s)
		}
	}
	return res
}

func TestFinalizerWallet(t *testing.T) {
	e := setup(t, recoveryDelay, 600)
	cfg := e.config(nil, nil)
	ready, cancelled := passport(1), passport(2)
	e.readyAndIndexed(cfg, map[common.Hash]int64{ready: 11, cancelled: 22})

	// Cancelled on chain, still pending in the indexed state.
	e.send(e.contract.CancelRecovery(e.auth, cancelled))
	if s := e.state(cfg, cancelled); s.Status != data.StatusPending {
		t.Fatalf("precondition: indexed status = %s, want pending", s.Status)
	}

	stop := runInBackground(finalizer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool {
		subs := submissionsFor(e.submissions(cfg), ready)
		return len(subs) == 1 && subs[0].Status == data.SubmissionConfirmed
	})
	// More periods: it must not submit again.
	time.Sleep(time.Second)
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	if got := e.identityKey(ready); got != 11 {
		t.Errorf("on-chain identity key = %d, want 11", got)
	}
	if subs := submissionsFor(e.submissions(cfg), ready); len(subs) != 1 {
		t.Errorf("submissions for ready passport = %d, want exactly 1", len(subs))
	}
	if subs := submissionsFor(e.submissions(cfg), cancelled); len(subs) != 0 {
		t.Errorf("submissions for cancelled passport = %d, want 0 (re-checked on chain)", len(subs))
	}
}

func TestFinalizerRelayer(t *testing.T) {
	e := setup(t, recoveryDelay, 600)
	url := e.startMockRelayer()

	cfg := e.config(nil, map[string]interface{}{"mode": config.ModeRelayer, "relayer_url": url})
	p := passport(1)
	e.readyAndIndexed(cfg, map[common.Hash]int64{p: 11})

	stop := runInBackground(finalizer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool {
		subs := e.submissions(cfg)
		return len(subs) == 1 && subs[0].Status == data.SubmissionConfirmed && subs[0].TxHash != nil
	})
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	if got := e.identityKey(p); got != 11 {
		t.Errorf("on-chain identity key = %d, want 11", got)
	}
}

func TestFinalizerUnknownOutcomeIsNotRetried(t *testing.T) {
	e := setup(t, recoveryDelay, 600)

	// Never answers, so the outcome is unknown.
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	srv := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() })}
	go func() { _ = srv.Serve(l) }()
	t.Cleanup(func() { _ = srv.Close() })

	cfg := e.config(nil, map[string]interface{}{
		"mode":            config.ModeRelayer,
		"relayer_url":     "http://" + l.Addr().String() + "/register",
		"relayer_timeout": "300ms",
	})
	p := passport(1)
	e.readyAndIndexed(cfg, map[common.Hash]int64{p: 11})

	stop := runInBackground(finalizer.New(cfg).Run)
	eventually(t, 30*time.Second, func() bool {
		subs := e.submissions(cfg)
		return len(subs) == 1 && subs[0].Status == data.SubmissionUnknown
	})
	time.Sleep(2 * time.Second)
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	if subs := e.submissions(cfg); len(subs) != 1 || subs[0].Status != data.SubmissionUnknown {
		t.Errorf("submissions = %+v, want a single unknown one", subs)
	}

	t.Run("expires after unknown_retry_after", func(t *testing.T) {
		cfg := e.config(nil, map[string]interface{}{
			"mode":                config.ModeRelayer,
			"relayer_url":         "http://" + l.Addr().String() + "/register",
			"relayer_timeout":     "300ms",
			"unknown_retry_after": "10s",
		})

		stop := runInBackground(finalizer.New(cfg).Run)
		defer func() { _ = stop() }()

		// Not before the retry window: a time zone mix-up would expire it at once.
		time.Sleep(3 * time.Second)
		if subs := e.submissions(cfg); len(subs) != 1 {
			t.Fatalf("submissions = %d before the retry window, want 1", len(subs))
		}

		eventually(t, 30*time.Second, func() bool {
			subs := e.submissions(cfg)
			return len(subs) == 2 && subs[0].Status == data.SubmissionExpired
		})
	})
}

func (e *env) startMockRelayer() string {
	e.t.Helper()
	bin := filepath.Join(e.t.TempDir(), "mockrelayer")
	build := exec.Command("go", "build", "-o", bin, "../../devtools/mockrelayer")
	build.Env = append(os.Environ(), "GOTOOLCHAIN=local")
	if out, err := build.CombinedOutput(); err != nil {
		e.t.Fatalf("build mockrelayer: %v\n%s", err, out)
	}

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		e.t.Fatal(err)
	}
	addr := l.Addr().String()
	_ = l.Close()

	cmd := exec.Command(bin)
	cmd.Env = append(os.Environ(),
		"MOCK_RELAYER_ADDR="+addr,
		"MOCK_RELAYER_RPC="+e.rpcURL,
		"MOCK_RELAYER_PRIVATE_KEY="+relayerKey,
		"MOCK_RELAYER_WHITELIST="+e.address.Hex(),
	)
	if err = cmd.Start(); err != nil {
		e.t.Fatal(err)
	}
	e.t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	})

	eventually(e.t, 10*time.Second, func() bool {
		c, err := net.Dial("tcp", addr)
		if err == nil {
			_ = c.Close()
		}
		return err == nil
	})
	return fmt.Sprintf("http://%s/integrations/registration-relayer/v1/register", addr)
}

func TestSubmissionConditionalUpdate(t *testing.T) {
	e := setup(t, recoveryDelay, 600)
	cfg := e.config(nil, nil)
	e.readyAndIndexed(cfg, map[common.Hash]int64{passport(1): 11})

	subs := pg.NewSubmissionsQ(cfg.DB(), e.contractID(cfg))
	id, err := subs.Claim(e.events(cfg)[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	txHash := data.Hash(common.Hash{1})
	for _, c := range []struct {
		to   data.SubmissionStatus
		from []data.SubmissionStatus
		want bool
	}{
		{data.SubmissionConfirmed, []data.SubmissionStatus{data.SubmissionSubmitted}, false},
		{data.SubmissionUnknown, []data.SubmissionStatus{data.SubmissionSubmitting}, true},
		{data.SubmissionUnknown, []data.SubmissionStatus{data.SubmissionSubmitting}, false},
		{data.SubmissionSubmitted, []data.SubmissionStatus{data.SubmissionSubmitting, data.SubmissionUnknown}, true},
	} {
		ok, err := subs.Update(id, c.to, &txHash, "", c.from...)
		if err != nil {
			t.Fatal(err)
		}
		if ok != c.want {
			t.Errorf("Update(%s from %v) = %t, want %t", c.to, c.from, ok, c.want)
		}
	}
	if got := e.submissions(cfg); len(got) != 1 || got[0].Status != data.SubmissionSubmitted {
		t.Errorf("submissions = %+v, want a single submitted one", got)
	}
}
