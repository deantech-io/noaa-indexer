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
	UnexpectedTransitions = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "noaa_indexer_unexpected_transitions_total",
		Help: "Events the contract should not have emitted in the indexed state, by type.",
	}, []string{"type"})
	Halted = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_indexer_halted",
		Help: "1 if the indexer stopped on a parent hash mismatch.",
	})

	Submissions = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "noaa_finalizer_submissions_total",
		Help: "Finalizer submission outcomes.",
	}, []string{"outcome"})
	ExhaustedRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_finalizer_exhausted_requests",
		Help: "Pending requests that reached max_attempts and are no longer retried.",
	})
	UnknownSubmissions = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "noaa_finalizer_unknown_submissions",
		Help: "Submissions whose outcome is unknown; each blocks further attempts for its passport.",
	})
)

func RegisterIndexer() {
	prometheus.MustRegister(IndexedBlock, HeadBlock, LagBlocks, Events, UnexpectedTransitions, Halted)
}

func RegisterFinalizer() {
	prometheus.MustRegister(Submissions, ExhaustedRequests, UnknownSubmissions)
}
