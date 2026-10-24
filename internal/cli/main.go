package cli

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/alecthomas/kingpin"
	"github.com/rarimo/noaa-indexer/internal/config"
	"github.com/rarimo/noaa-indexer/internal/metrics"
	"github.com/rarimo/noaa-indexer/internal/service/api"
	"github.com/rarimo/noaa-indexer/internal/service/finalizer"
	"github.com/rarimo/noaa-indexer/internal/service/indexer"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	defer func() {
		if rvr := recover(); rvr != nil {
			logan.New().WithRecover(rvr).Error("app panicked")
		}
	}()

	var (
		cfg            = config.New(kv.MustFromEnv())
		log            = cfg.Log()
		app            = kingpin.New("noaa-indexer", "")
		runCmd         = app.Command("run", "run command")
		indexerCmd     = runCmd.Command("indexer", "run NoAARecovery event indexer")
		finalizerCmd   = runCmd.Command("finalizer", "run recovery finalizer")
		migrateCmd     = app.Command("migrate", "migrate command")
		migrateUpCmd   = migrateCmd.Command("up", "migrate db up")
		migrateDownCmd = migrateCmd.Command("down", "migrate db down")
		rebuildCmd     = app.Command("rebuild", "rebuild recovery_state from recovery_event")
	)

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	run := func(f func()) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}

	var runErr error
	switch cmd {
	case indexerCmd.FullCommand():
		metrics.RegisterIndexer()
		ix := indexer.New(cfg)
		run(func() { api.Run(ctx, cfg, ix.Healthy) })
		// The API keeps serving after a halt so health and metrics show it.
		run(func() { _ = ix.Run(ctx) })
	case finalizerCmd.FullCommand():
		metrics.RegisterFinalizer()
		f := finalizer.New(cfg)
		run(func() { api.Run(ctx, cfg, func() bool { return true }) })
		run(func() {
			if runErr = f.Run(ctx); runErr != nil {
				log.WithError(runErr).Error("finalizer stopped")
				stop()
			}
		})
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	case rebuildCmd.FullCommand():
		err = indexer.Rebuild(cfg)
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}

	wg.Wait()
	return runErr == nil
}
