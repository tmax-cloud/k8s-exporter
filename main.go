package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
}
