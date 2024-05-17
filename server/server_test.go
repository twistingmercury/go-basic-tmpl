package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"{{module_name}}/server"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twistingmercury/heartbeat"
)

func TestBootstrap(t *testing.T) {
	// Create a new context
	ctx := context.Background()

	// Call the Bootstrap function
	err := server.Bootstrap(ctx)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestStart(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Call the Bootstrap function
	err := server.Bootstrap(ctx)
	require.NoError(t, err)

	// Start the server in a goroutine
	go server.Start()

	// Perform any necessary assertions or checks
	// ...

	// Cancel the context to stop the server
	cancel()
}

func TestStartHeartbeat(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the heartbeat endpoint in a goroutine
	go server.StartHeartbeat(ctx)

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
	expectedBody := `{"status":"OK","dependencies":[{"name":"desc 01","type":"http/rest","status":"NOT_SET","message":"unknown","timestamp":""},{"name":"desc 02","type":"database","status":"NOT_SET","message":"unknown","timestamp":""}],"timestamp":""}`
	assert.JSONEq(t, expectedBody, resp.Body.String())
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
