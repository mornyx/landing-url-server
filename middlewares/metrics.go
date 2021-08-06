// Package middlewares provides all gin middleware.
package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mornyx/landing-url-server/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

// MetricsMiddleware is responsible for reporting prometheus metrics.
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.HTTPRequestCount.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Inc()
		begin := time.Now()
		c.Next()
		metrics.HTTPRequestLatencyHistogramMilliseconds.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Observe(float64(time.Now().Sub(begin).Milliseconds()))
		metrics.HTTPRequestLatencySummaryMilliseconds.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Observe(float64(time.Now().Sub(begin).Milliseconds()))
	}
}
