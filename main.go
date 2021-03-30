package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	k8s := newK8sCollector()
	prometheus.MustRegister(k8s)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
}
