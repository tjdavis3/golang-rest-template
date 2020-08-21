package api

import (
	"net/http"
	"strconv"

	"../util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// InfoMetric contains system information
	InfoMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_info",
		Help: "Information about the service",
	}, []string{"service", "version"})
	opsProcessing = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "api_processing_ops_total",
		Help: "The number of events processing",
	})
	responseMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "api_responses_total",
		Help: "The number of responses by endpoint and status",
	},
		[]string{"code", "method", "endpoint"})
)

// mwMetrics is simple middleware to count ongoing requests.
func mwMetrics(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer opsProcessing.Dec()
		opsProcessing.Inc()
		lw := util.WrapWriter(w)
		next.ServeHTTP(lw, r)
		responseMetric.WithLabelValues(strconv.Itoa(lw.Status()), r.Method, r.URL.Path).Inc()
	}
	return http.HandlerFunc(f)
}
