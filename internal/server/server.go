package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"token_go_module/internal/conf"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/twistingmercury/heartbeat"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/twistingmercury/telemetry/v2/logging"
	"github.com/twistingmercury/telemetry/v2/metrics"
	"github.com/twistingmercury/telemetry/v2/tracing"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

var (
	ctx           context.Context
	hbsvr         *http.Server
	sampleCounter *prometheus.CounterVec
)

// Bootstrap initializes the application's telemetry components, logging, tracing, and metrics.
func Bootstrap(context context.Context, svcName, svcVersion, namespace, environment string) error {
	conf.ShowVersion()
	conf.ShowHelp()

	// init logging
	logLevel, err := zerolog.ParseLevel(viper.GetString(conf.ViperLogLevelKey))
	if err != nil {
		return err
	}
	err = logging.Initialize(logLevel, os.Stdout, svcName, svcVersion, environment)
	if err != nil {
		return err
	}
	ctx = context

	// init tracing
	logging.Info(ctx, "initializing tracing")
	texporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return err
	}
	err = tracing.InitializeWithSampleRate(texporter, viper.GetFloat64(conf.ViperTraceSampleRateKey), svcName, svcVersion, namespace)
	if err != nil {
		return err
	}

	// init metrics
	logging.Info(ctx, "initializing metrics")
	err = metrics.Initialize(context, namespace, svcName)
	if err != nil {
		return err
	}

	// ---> register metrics
	// reference: https://github.com/twistingmercury/telemetry/blob/main/metrics/README.md
	sampleCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: viper.GetString(conf.ViperNamespaceKey),
		Name:      "sample_counter",
		Help:      "A sample counter metric"},
		[]string{"label1", "label2", "label3"})

	metrics.RegisterMetrics(sampleCounter)
	metrics.Publish()
	//<---

	logging.Info(ctx, "starting heartbeat")
	StartHeartbeat(ctx)

	return nil
}

// Start initializes the application's API service and starts the cmd.
func Start() {
	logging.Info(ctx, "starting cmd")

	// -->
	// do whatever is required to start the cmd, such as initializing the database, listening
	// to message brokers, starting HTTP or gRPC servers, etc.
	// <--

	for {
		select {
		case <-ctx.Done():
			logging.Info(ctx, "stopping cmd")
			_ = stop()
			return
		default:
			time.Sleep(1 * time.Second)
			sampleCounter.WithLabelValues("foo", "bar", "bas").Inc()
			logging.Info(ctx, "doing some work")
		}
	}
}

func stop() error {
	if hbsvr == nil {
		return nil
	}

	if err := hbsvr.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// StartHeartbeat starts the heartbeat endpoint, and provides an endpoint for health monitoring.
// It can also be used to provide a liveness and readiness check for Kubernetes.
func StartHeartbeat(ctx context.Context) {
	hbr := gin.New()
	hbr.Use(gin.Recovery())
	hbr.GET("/heartbeat", heartbeat.Handler("BIN_NAME", CheckDeps()...))
	addr := fmt.Sprintf(":%d", conf.DefaultHeartbeatPort)

	logging.Info(ctx, "starting heartbeat", logging.KeyValue{Key: "addr", Value: addr})
	hbsvr = &http.Server{Addr: addr, Handler: hbr}
	go func() {
		if err := hbsvr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logging.Fatal(ctx, err, "heartbeat endpoint failed with error")
		}
	}()
}

// CheckDeps provides a list of dependencies to be checked by the heartbeat endpoint.
// reference: https://github.com/twistingmercury/heartbeat/blob/main/readme.md
func CheckDeps() []heartbeat.DependencyDescriptor {
	deps := []heartbeat.DependencyDescriptor{
		{
			Name: "desc 01",
			Type: "http/rest", // or whatever makes sense for your service
			HandlerFunc: func() heartbeat.StatusResult {
				hsr := heartbeat.StatusResult{Status: heartbeat.StatusOK, Message: "ok"}
				// for a REST apo, you'd create a func that checks if the REST api is reachable,
				// perhaps invoking its health endpoint (if it has one).
				return hsr
			},
		},
		{
			Name: "desc 02",
			Type: "database",
			HandlerFunc: func() heartbeat.StatusResult {
				hsr := heartbeat.StatusResult{Status: heartbeat.StatusOK, Message: "ok"}
				// for a database, you'd create a func that checks if the database is up and running
				return hsr
			},
		},
	}

	return deps
}
