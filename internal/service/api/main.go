package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rarimo/noaa-indexer/internal/config"
)

func Run(ctx context.Context, cfg config.Config, healthy func() bool) {
	mux := http.NewServeMux()
	mux.Handle("GET /metrics", promhttp.Handler())
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		if !healthy() {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})

	srv := &http.Server{Handler: mux}
	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(context.Background())
	}()

	log := cfg.Log().WithField("service", "api")
	log.Info("api started")
	if err := srv.Serve(cfg.Listener()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithError(err).Error("api stopped")
	}
}
