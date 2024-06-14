package server_test

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/twistingmercury/telemetry/metrics"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"token_go_module/internal/conf"
	"token_go_module/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/twistingmercury/heartbeat"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	svcName    = "test"
	svcVersion = "0.0.0"
	nspace     = "unit"
	environ    = "unittesting"
)

func init() {
	viper.Set(conf.ViperTraceSampleRateKey, conf.DefaultTraceSampleRate)
}
func TestBootstrap(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Call the Bootstrap function
	err := server.Bootstrap(ctx, svcName, svcVersion, nspace, environ)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestStart(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	server.SetContext(ctx)
	defer cancel()

	err := metrics.Initialize(ctx, nspace, svcName)
	require.NoError(t, err)

	utCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "unit",
		Name:      "test_counter",
		Help:      "unit test counter"},
		[]string{"label1", "label2", "label3"})

	server.SetCounter(utCounter)
	metrics.RegisterMetrics(utCounter)

	server.Start()
}

func TestStartHeartbeat(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(2 * time.Second)
	}()

	// Start the heartbeat endpoint in a goroutine
	server.StartHeartbeat(ctx)

	// Create a new HTTP request for the heartbeat endpoint
	req, err := http.NewRequest("GET", "/heartbeat", nil)
	require.NoError(t, err)

	// Create a new Gin router
	router := gin.New()
	router.GET("/heartbeat", heartbeat.Handler("test", server.CheckDeps()...))

	// Perform the request and record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Assert the response body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	result := string(body)
	expectedKeys := []string{"status", "request_duration_ms", "resource", "utc_DateTime", "dependencies"}

	for _, key := range expectedKeys {
		assert.Contains(t, result, key)
	}
}

func TestCheckDeps(t *testing.T) {
	// Call the CheckDeps function
	deps := server.CheckDeps()

	// Assert the number of dependencies returned
	assert.Len(t, deps, 2)

	// Assert the properties of each dependency
	assert.Equal(t, "desc 01", deps[0].Name)
	assert.Equal(t, "http/rest", deps[0].Type)
	assert.NotNil(t, deps[0].HandlerFunc)

	assert.Equal(t, "desc 02", deps[1].Name)
	assert.Equal(t, "database", deps[1].Type)
	assert.NotNil(t, deps[1].HandlerFunc)
}
