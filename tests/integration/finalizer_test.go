//go:build integration

package integration

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/deantech-io/noaa-indexer/internal/config"
	"github.com/deantech-io/noaa-indexer/internal/data"
	"github.com/deantech-io/noaa-indexer/internal/data/pg"
	"github.com/deantech-io/noaa-indexer/internal/metrics"
	"github.com/deantech-io/noaa-indexer/internal/service/finalizer"
	"github.com/deantech-io/noaa-indexer/internal/service/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

const recoveryDelay = 100

func (e *env) readyAndIndexed(cfg config.Config, keys map[common.Hash]int64) {
	e.t.Helper()
	for p, k := range keys {
		e.request(p, k)
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
	k, err := e.contract.ActiveIdentity(&bind.CallOpts{}, p)
	if err != nil {
		e.t.Fatal(err)
	}
	return data.NumericFromBytes(k).Int64()
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
	e.send(e.contract.CancelRecovery(e.auth, cancelled, nil))
	if s := e.state(cfg, cancelled); s.NewIdentityKey.IsZero() {
		t.Fatal("precondition: indexed request is not pending")
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

	if subs := e.submissions(cfg); len(subs) != 1 || subs[0].Status != data.SubmissionUnknown || subs[0].Error != "relayer request timed out" {
		t.Errorf("submissions = %+v, want a single unknown one without the relayer URL", subs)
	}

	t.Run("expires after unknown_retry_after", func(t *testing.T) {
		cfg := e.config(nil, map[string]interface{}{
			"mode":                config.ModeRelayer,
			"relayer_url":         "http://" + l.Addr().String() + "/register",
			"relayer_timeout":     "300ms",
			"unknown_retry_after": "10s",
			"retry_base":          "1s",
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
	id, err := subs.Claim(e.events(cfg)[0].ID, data.Hash(passport(1)))
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

func TestFinalizerRetries(t *testing.T) {
	e := setup(t, recoveryDelay, 600)

	const reason = "execution reverted: not ready"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, `{"errors":[{"title":%q}]}`, reason)
	}))
	t.Cleanup(srv.Close)

	cfg := e.config(nil, map[string]interface{}{
		"mode":               config.ModeRelayer,
		"relayer_url":        srv.URL,
		"wallet_private_key": nil,
		"retry_base":         "1s",
		"retry_max":          "2s",
		"max_attempts":       3,
	})
	e.readyAndIndexed(cfg, map[common.Hash]int64{passport(1): 11})
	exhausted := counterValue(t, metrics.ExhaustedRequests)

	stop := runInBackground(finalizer.New(cfg).Run)
	defer func() { _ = stop() }()

	eventually(t, 10*time.Second, func() bool { return len(e.submissions(cfg)) == 1 })
	time.Sleep(500 * time.Millisecond)
	subs := e.submissions(cfg)
	if len(subs) != 1 || subs[0].Status != data.SubmissionRejected || !strings.Contains(subs[0].Error, reason) {
		t.Fatalf("submissions = %+v, want one rejected with the revert reason, no retry before retry_base", subs)
	}

	eventually(t, 15*time.Second, func() bool { return len(e.submissions(cfg)) == 3 })
	// The backoff alone would allow a fourth attempt after 2s.
	time.Sleep(3 * time.Second)
	subs = e.submissions(cfg)
	if len(subs) != 3 {
		t.Fatalf("submissions = %d, want 3 (max_attempts)", len(subs))
	}
	for i, want := range []time.Duration{time.Second, 2 * time.Second} {
		if gap := subs[i+1].UpdatedAt.Sub(subs[i].UpdatedAt); gap < want {
			t.Errorf("gap before attempt %d = %s, want at least %s", i+2, gap, want)
		}
	}
	if got := counterValue(t, metrics.ExhaustedRequests) - exhausted; got != 1 {
		t.Errorf("exhausted requests = %v, want 1", got)
	}
}

func TestOneOpenSubmissionPerPassport(t *testing.T) {
	e := setup(t, recoveryDelay, 600)
	cfg := e.config(nil, nil)
	p := passport(1)

	e.request(p, 11)
	e.send(e.contract.CancelRecovery(e.auth, p, nil))
	e.rpc("evm_increaseTime", 601) // past the lock set by the cancel
	e.request(p, 12)
	e.readyAndIndexed(cfg, nil)

	var requests []int64
	for _, ev := range e.events(cfg) {
		if ev.Type == data.EventRequested {
			requests = append(requests, ev.ID)
		}
	}
	if len(requests) != 2 {
		t.Fatalf("requests = %d, want 2", len(requests))
	}

	subs := pg.NewSubmissionsQ(cfg.DB(), e.contractID(cfg))
	claim := func(request int64) int64 {
		id, err := subs.Claim(request, data.Hash(p))
		if err != nil {
			t.Fatal(err)
		}
		return id
	}
	update := func(id int64, to, from data.SubmissionStatus) {
		if ok, err := subs.Update(id, to, nil, "", from); err != nil || !ok {
			t.Fatalf("Update(%d, %s) = %t, %v", id, to, ok, err)
		}
	}

	first := claim(requests[0])
	if first == 0 {
		t.Fatal("first claim refused")
	}
	if claim(requests[1]) != 0 {
		t.Error("claimed while another request of the passport is submitting")
	}
	update(first, data.SubmissionUnknown, data.SubmissionSubmitting)
	if claim(requests[1]) != 0 {
		t.Error("claimed while another request of the passport is unknown")
	}
	update(first, data.SubmissionExpired, data.SubmissionUnknown)
	if claim(requests[1]) == 0 {
		t.Error("claim refused after the other attempt expired")
	}
}

func TestConfigChecks(t *testing.T) {
	e := setup(t, recoveryDelay, 600)

	t.Run("chain ID mismatch", func(t *testing.T) {
		cfg := e.configWith(map[string]map[string]interface{}{"network": {"chain_id": 1}})
		mustPanic(t, "rpc chain ID mismatch", func() { cfg.Network() })
	})
	t.Run("wallet mode on an unlisted chain", func(t *testing.T) {
		cfg := e.config(nil, map[string]interface{}{"wallet_allowed_chain_ids": []interface{}{1}})
		mustPanic(t, "wallet mode is not allowed on this chain", func() { cfg.Finalizer() })
	})
}

func mustPanic(t *testing.T, want string, f func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil || !strings.Contains(fmt.Sprint(r), want) {
			t.Errorf("panic = %v, want %q", r, want)
		}
	}()
	f()
}

func counterValue(t *testing.T, c prometheus.Counter) float64 {
	t.Helper()
	var m dto.Metric
	if err := c.Write(&m); err != nil {
		t.Fatal(err)
	}
	return m.GetCounter().GetValue()
}
