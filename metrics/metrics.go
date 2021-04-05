package metrics

import (
	"net/http"
	"strconv"

	"github.com/ibrokethecloud/kube-bench-metrics/wrapper"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	cisScore = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cis_results",
			Help: "Details of CIS benchmarks for cluster",
		},
		[]string{"test_number", "type", "status", "scored", "hostname"},
	)
)

// GenerateMetrics will populate Prometheus metrics by parsing the
// metrics available from the ParseResults method.
// Predefined numeric values will be assigned based on the status of the check
// PASS: 1
// NOTPASSED: 0
// Please refer to value in label to identify the actual result
// of the test
func GenerateMetrics(w wrapper.Wrapper) {
	for _, control := range w.Results.Controls {
		for _, group := range control.Groups {
			for _, check := range group.Checks {
				status := 1.0
				if check.State != "PASS" {
					status = 0.0
				}
				if check.Type == "" {
					check.Type = "automated"
				}
				logrus.Debug(check.ID, " ", check.Type, " ", string(check.State))
				cisScore.With(prometheus.Labels{"test_number": check.ID,
					"type":     check.Type,
					"status":   string(check.State),
					"scored":   strconv.FormatBool(check.Scored),
					"hostname": w.NodeName}).Set(status)
			}
		}
	}
}

// ServeMetrics will simply start up the prom http endpoint to
// serve the metrics
func ServeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	logrus.Fatal(http.ListenAndServe(":8000", nil))
}
