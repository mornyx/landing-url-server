// Package metrics provides all prometheus metrics used within the project.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_count",
	}, []string{"method", "path"})

	HTTPRequestLatencyHistogramMilliseconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_latency_histogram_milliseconds",
		Buckets: []float64{
			1, 2, 3, 4, 5, 7, 10, 12, 15, 17, 20, 25, 30, 50, 70, 100, 200, 500, 1000, 2000, 5000, 10000,
		},
	}, []string{"method", "path"})

	HTTPRequestLatencySummaryMilliseconds = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_latency_summary_milliseconds",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	}, []string{"method", "path"})
)
