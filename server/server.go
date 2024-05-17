package server

import (
	"context"
	"fmt"
	"os"

	"{{module_name}}/conf"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/twistingmercury/heartbeat"
	"github.com/twistingmercury/telemetry/attributes"
	"github.com/twistingmercury/telemetry/logging"
	"github.com/twistingmercury/telemetry/metrics"
	"github.com/twistingmercury/telemetry/tracing"
	"github.com/twistingmercury/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

var (
	ctx     context.Context
	attribs attributes.Attributes
)

// Bootstrap initializes the application's telemetry components, logging, tracing, and metrics.
func Bootstrap(context context.Context) error {
	conf.Initialize()
	conf.ShowVersion()
	conf.ShowHelp()

	ctx = context
	attribs = attributes.New(
		viper.GetString(conf.ViperNamespaceKey),
		viper.GetString(conf.ViperServiceNameKey),
		viper.GetString(conf.ViperBuildVersionKey),
		viper.GetString(conf.ViperEnviormentKey),
		attribute.String("commit", viper.GetString(conf.ViperCommitHashKey)),
	)

	logging.Info("initializing logging")
	logLevel, err := zerolog.ParseLevel(viper.GetString(conf.ViperLogLevelKey))
	if err != nil {
		return err
	}
	err = logging.Initialize(logLevel, attribs, os.Stdout)
	if err != nil {
		return err
	}

	logging.Info("initializing tracing")
	texporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return err
	}
	err = tracing.Initialize(texporter, viper.GetFloat64(conf.ViperTraceSampleRateKey), attribs)
	if err != nil {
		return err
	}

	logging.Info("initializing metrics")
	mexexporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return err
	}
	return metrics.Initialize(mexexporter, attribs)
}

// Start initializes the application's API service and starts the server.
func Start() {
	logging.Info("starting server")

	// ->
	// do whatever is required to start the server, such as initializing the database, listening
	// to message brokers, starting HTTP or gRPC servers, etc.
	// <-

	logging.Info("starting heartbeat")
	go startHeartbeat(ctx)

	logging.Info("serivce started successfully")

	<-ctx.Done()
	logging.Info("service stopped")
}

// startHeartbeat starts the heartbeat endpoint, and provides an endpoint for health monitoring.
// It can also be used to provide a liveness and readiness check for Kubernetes.
func startHeartbeat(ctx context.Context) {
	hbr := gin.New()
	hbr.Use(gin.Recovery())
	hbr.GET("/heartbeat", heartbeat.Handler("tunnelvisioin", checkDeps()...))
	addr := fmt.Sprintf(":%d", conf.DefaultHeartbeatPort)

	logging.Info("starting heartbeat", logging.KeyValue{Key: "addr", Value: addr})

	go func() {
		if err := hbr.Run(addr); err != nil {
			utils.FailFast(err, "failed to initialize heartbeat")
		}
	}()

	<-ctx.Done()
}

// checkDeps provides a list of dependencies to be checked by the heartbeat endpoint.
func checkDeps() []heartbeat.DependencyDescriptor {
	deps := []heartbeat.DependencyDescriptor{
		{
			Name: "desc 01",
			Type: "http/rest", // or whatever makes sense for your service
			HandlerFunc: func() heartbeat.StatusResult {
				hsr := heartbeat.StatusResult{Status: heartbeat.StatusNotSet, Message: "unknown"}
				// for a REST apo, you'd create a func that checks if the REST api is reachable,
				// perhaps invoking its health endpoint (if it has one).
				return hsr
			},
		},
		{
			Name: "desc 02",
			Type: "database",
			HandlerFunc: func() heartbeat.StatusResult {
				hsr := heartbeat.StatusResult{Status: heartbeat.StatusNotSet, Message: "unknown"}
				// for a database, you'd create a func that checks if the database is up and running
				return hsr
			},
		},
	}

	return deps
}
