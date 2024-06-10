package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twistingmercury/heartbeat"

	"MODULE_NAME/conf"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/twistingmercury/telemetry/logging"
	"github.com/twistingmercury/telemetry/metrics"
	"github.com/twistingmercury/telemetry/tracing"
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

	logLevel, err := zerolog.ParseLevel(viper.GetString(conf.ViperLogLevelKey))
	if err != nil {
		return err
	}
	err = logging.Initialize(logLevel, os.Stdout, svcName, svcVersion, environment)
	if err != nil {
		return err
	}
	ctx = context

	logging.Info("initializing tracing")
	texporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return err
	}
	err = tracing.InitializeWithSampleRate(texporter, viper.GetFloat64(conf.ViperTraceSampleRateKey), svcName, svcVersion, namespace)
	if err != nil {
		return err
	}

	logging.Info("initializing metrics")

	err = metrics.Initialize(namespace, svcName)
	if err != nil {
		return err
	}

	// register metrics
	// reference: https://github.com/twistingmercury/telemetry/blob/main/metrics/README.md
	sampleCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: viper.GetString(conf.ViperNamespaceKey),
		Name:      "sample_counter",
		Help:      "A sample counter metric"},
		[]string{"label1", "label2", "label3"})

	metrics.RegisterMetrics(sampleCounter)
	metrics.Publish()

	return nil
}

// Start initializes the application's API service and starts the server.
func Start() {
	logging.Info("starting server")

	logging.Info("starting heartbeat")
	StartHeartbeat(ctx)

	// -->
	// do whatever is required to start the server, such as initializing the database, listening
	// to message brokers, starting HTTP or gRPC servers, etc.
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(5 * time.Second)
				sampleCounter.WithLabelValues("foo", "bar", "bas").Inc()
				logging.Info("doing some work")
			}
		}
	}()
	// <--

	logging.Info("serivce started successfully")

	<-ctx.Done()
	logging.Info("service stopped")
}

func Stop() error {
	if err := hbsvr.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}

// StartHeartbeat starts the heartbeat endpoint, and provides an endpoint for health monitoring.
// It can also be used to provide a liveness and readiness check for Kubernetes.
func StartHeartbeat(ctx context.Context) {
	hbr := gin.New()
	hbr.Use(gin.Recovery())
	hbr.GET("/heartbeat", heartbeat.Handler("BIN_NAME", CheckDeps()...))
	addr := fmt.Sprintf(":%d", conf.DefaultHeartbeatPort)

	logging.Info("starting heartbeat", logging.KeyValue{Key: "addr", Value: addr})
	hbsvr = &http.Server{Addr: addr, Handler: hbr}
	go func() {
		if err := hbsvr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logging.Fatal(err, "heartbeat endpoint failed with error")
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
