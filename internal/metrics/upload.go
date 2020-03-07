// Package metrics ...
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// UploadedOK represent the total number of success job
	UploadedOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http2-uploaderserver-upload-successfully-total",
		Help: "The total number of success upload",
	})
)

var (
	// UploadNOK represent the total number of failed job creation
	UploadNOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http2-uploaderserver-upload-failed-total",
		Help: "The total number of failed upload",
	})
)
