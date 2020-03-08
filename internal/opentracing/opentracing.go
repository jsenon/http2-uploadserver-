// Package opentracing ...
package opentracing

import (
	"io"
	"net"
	"strings"

	"github.com/pkg/errors"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// ConfigureTracing configures OpenTracing client
func ConfigureTracing(jaeger string) (closer io.Closer, err error) {
	var localagenthostport string

	if strings.Contains(jaeger, "http") || strings.Contains(jaeger, ":5") ||
		strings.Contains(jaeger, ":") || strings.Contains(jaeger, "/") {
		errmsg := errors.New("wrong caracters set in jaeger url, do not set http:// or port")
		return closer, errors.Wrapf(errmsg, "Unable to launch opentracing")
	}

	if net.ParseIP(jaeger) == nil {
		localAgentHostPortIP, err2 := net.LookupHost(jaeger)
		if err2 != nil {
			errmsg := errors.New("could not resolv DNS jaeger tracer")
			return closer, errors.Wrapf(errmsg, "Unable to launch opentracing")
		}

		localagenthostport = localAgentHostPortIP[0] + ":5775"
	} else {
		localagenthostport = jaeger + ":5775"
	}

	samplingurl := "http://" + jaeger + ":5778/sampling"
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:              "const",
			Param:             0.1,
			SamplingServerURL: samplingurl,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: false,
			// LocalAgentHostPort: localagenthostport,
			LocalAgentHostPort: localagenthostport,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err = cfg.InitGlobalTracer(
		"http2-uploader",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		errmsg := errors.New("could not initialize jaeger tracer")
		return closer, errors.Wrapf(errmsg, "Unable to launch opentracing %v", localagenthostport)
	}

	return closer, nil
}
