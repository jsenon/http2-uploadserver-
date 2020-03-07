// Package metrics ...
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// UploadOK represent the total number of success job
	UploadOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http2_uploaderserver_upload_success_total",
		Help: "The total number of success upload",
	})
)

var (
	// UploadNOK represent the total number of failed job creation
	UploadNOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http2_uploaderserver_upload_failed_total",
		Help: "The total number of failed upload",
	})
)
