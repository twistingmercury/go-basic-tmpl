# Package: server

The `server` package is responsible for initializing and starting the application server. It sets up the necessary components, such as logging, tracing, and metrics, and starts the server along with a heartbeat endpoint for health monitoring.

## Functions

### Bootstrap

The `Bootstrap` function initializes the application's telemetry components, including logging, tracing, and metrics. It performs the following steps:

1. Initializes the configuration using `conf.Initialize()`.
2. Shows the version information if the `ViperShowVersionKey` flag is set.
3. Shows the help information if the `ViperShowHelpKey` flag is set.
4. Initializes the `zerolog` logger with the specified log level and attaches it to `os.Stdout`.
5. Initializes the OpenTelemetry tracer with the specified trace sampler and exporter.
6. Initializes the OpenTelemetry meter with the specified exporter.

### Start

The `Start` function initializes the application's API service and starts the server. It performs the following steps:

1. Initializes the server components and starts listening for incoming requests.
2. Starts the heartbeat endpoint for health monitoring by calling the `startHeartbeat` function.
3. Waits for the context to be done (e.g., when the server is terminated).

### startHeartbeat

The `startHeartbeat` function starts the heartbeat endpoint using the `gin` web framework. It exposes an endpoint at `/heartbeat` that returns the health status of the application and its dependencies.

### checkDeps

The `checkDeps` function returns a list of `heartbeat.DependencyDescriptor` that describes the dependencies of the application. Each descriptor includes the name, type, and a handler function to check the health of the dependency.

## Extending the Server

To extend the functionality of the `server` package, you can modify the `Bootstrap` and `Start` functions to include additional initialization steps or start additional services.

### Extending the Bootstrap Function

1. Open the `server.go` file in a text editor.

2. Locate the `Bootstrap` function.

3. Add any additional initialization steps after the existing ones. For example, you can initialize a database connection, set up message queue consumers, or initialize other third-party libraries.

```go
func Bootstrap(ctx context.Context) error {
    // ...

    // Initialize database connection
    db, err := initializeDatabase()
    if err != nil {
        return err
    }

    // Initialize message queue consumer
    mqConsumer, err := initializeMessageQueueConsumer()
    if err != nil {
        return err
    }

    // ...
}
```

4. Make sure to handle any errors that may occur during the initialization process and return them from the `Bootstrap` function.

### Extending the Start Function

1. Open the `server.go` file in a text editor.

2. Locate the `Start` function.

3. Add any additional server initialization steps or start additional services within the function. For example, you can start gRPC servers, background workers, or periodic tasks.

```go
func Start() {
    // ...

    // Start gRPC server
    go startGRPCServer()

    // Start background worker
    go startBackgroundWorker()

    // Start periodic task
    go startPeriodicTask()

    // ...
}
```

4. Make sure to handle any errors that may occur during the startup process and log them appropriately.

### Adding Dependencies

To add new dependencies to the health check endpoint:

1. Open the `server.go` file in a text editor.

2. Locate the `checkDeps` function.

3. Add a new `heartbeat.DependencyDescriptor` to the `deps` slice for each new dependency you want to monitor.

```go
func checkDeps() []heartbeat.DependencyDescriptor {
    deps := []heartbeat.DependencyDescriptor{
        // ...

        {
            Name: "New Dependency",
            Type: "database",
            HandlerFunc: func() heartbeat.StatusResult {
                // Check the health of the new dependency
                // ...
            },
        },
    }

    return deps
}
```

4. Implement the `HandlerFunc` for each new dependency descriptor to check the health of the respective dependency and return the appropriate `heartbeat.StatusResult`.
   
5. In all cases, update the unit tests in [server_test.go](./server_test.go). The tests have been stubbed out, but you will need to finish them.

Remember to also update the README and other relevant documentation to reflect any changes or extensions made to the `server` package.