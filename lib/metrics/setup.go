package metrics

import (
	"fmt"
	"github.com/OAB/utils/log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var EnabledExpensive = false

// Setup starts a dedicated metrics server at the given address.
// This function enables metrics reporting separate from pprof.
func Setup(address string) *http.ServeMux {
	prometheus.DefaultRegisterer.MustRegister(defaultSet)

	prometheusMux := http.NewServeMux()
	prometheusMux.Handle("/debug/metrics/prometheus", promhttp.Handler())

	promServer := &http.Server{
		Addr:    address,
		Handler: prometheusMux,
	}

	go func() {
		if err := promServer.ListenAndServe(); err != nil {
			log.Error("Failure in running Prometheus server", "err", err)
		}
	}()

	log.Info("Enabling metrics export to prometheus", "path", fmt.Sprintf("http://%s/debug/metrics/prometheus", address))
	return prometheusMux
}
