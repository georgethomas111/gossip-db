package stats

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var MetricsPort = 10000

var NodeCount = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "node_count",
	Help: "Number of nodes available as seen from this node",
})

var RowsInDatabase = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "rows_in_node",
	Help: "Number of key value pairs in the map ",
})

func New() {

	// Create non-global registry.
	registry := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		NodeCount,
		RowsInDatabase,
	)

	// Expose /metrics HTTP endpoint using the created custom registry.
	http.Handle(
		"/metrics",
		promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		),
	)

	err := http.ListenAndServe(fmt.Sprintf(":%d", MetricsPort), nil)
	if err != nil {
		fmt.Println("WARNING: Metrics could not be started ", err.Error())
	}
}
