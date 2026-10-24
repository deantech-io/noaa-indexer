// Package metrics defines Prometheus metrics. Each subcommand registers only its own.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	IndexedBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_indexer_indexed_block",
		Help: "Last indexed (finalized) block.",
	})
	HeadBlock = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_indexer_head_block",
		Help: "Latest block reported by the RPC.",
	})
	LagBlocks = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_indexer_lag_blocks",
		Help: "Head block minus indexed block.",
	})
	Events = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "noaa_indexer_events_total",
		Help: "Indexed events by type.",
	}, []string{"type"})
	Halted = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_indexer_halted",
		Help: "1 if the indexer stopped on a parent hash mismatch.",
	})

	Submissions = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "noaa_finalizer_submissions_total",
		Help: "Finalizer submission outcomes.",
	}, []string{"outcome"})
	UnknownSubmissions = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_finalizer_unknown_submissions",
		Help: "Submissions whose outcome is unknown; they block further attempts for their passport.",
	})
)

func RegisterIndexer() {
	prometheus.MustRegister(IndexedBlock, HeadBlock, LagBlocks, Events, Halted)
}

func RegisterFinalizer() {
	prometheus.MustRegister(Submissions, UnknownSubmissions)
}
