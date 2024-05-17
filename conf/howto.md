# How to use and update the conf package.

The `conf` package is responsible for handling the configuration of the Go application. It uses the `viper` package for configuration management and the `pflag` package for command-line flag parsing.

## Configuration Keys

The following configuration keys are defined in the `conf` package:

- `ViperOtelAddrKey`: OpenTelemetry collector endpoint
- `ViperLogLevelKey`: Log level (debug, info, warn, error, fatal, panic)
- `ViperEnviormentKey`: Environment (dev, test, stage, prod)
- `ViperNamespaceKey`: Namespace
- `ViperTraceSampleRateKey`: Trace sample rate
- `ViperBuildVersionKey`: Build version
- `ViperBuildDateKey`: Build date
- `ViperCommitHashKey`: Commit hash
- `ViperServiceNameKey`: Service name
- `ViperShowHelpKey`: Show help flag
- `ViperShowVersionKey`: Show version flag

## Command-Line Flags

The following command-line flags are defined in the `conf` package:

- `--help`: Print the help information
- `--version`: Print the version information
- `--otel-endpoint`: OpenTelemetry collector endpoint
- `--log-level`: Log level (debug, info, warn, error, fatal, panic)
- `--environment`: Environment (dev, test, stage, prod)

## Environment Variables

The following environment variables are used by the `conf` package:

- `OTEL_EXPORTER_OTLP_ENDPOINT`: OpenTelemetry collector endpoint
- `LOG_LEVEL`: Log level (debug, info, warn, error, fatal, panic)
- `ENVIRONMENT`: Environment (dev, test, stage, prod)
- `OTEL_TRACE_SAMPLE_RATE`: Trace sample rate
- `HEARTBEAT_PORT`: Heartbeat port

## Default Values

The `conf` package defines default values for the following configuration keys:

- `DefaultHelp`: false
- `DefaultVersion`: false
- `DefaultLogLevel`: zerolog.DebugLevel
- `DefaultEnv`: "dev"
- `DefaultTraceSampleRate`: 0.25
- `DefaultHeartbeatPort`: 8081

## Utilizing Configuration Values

To access and utilize the configuration values in your Go application, follow these steps:

1. Import the `conf` package in the file where you want to use the configuration values:

```go
import "path/to/conf"
```

2. Call the `conf.Initialize()` function to initialize the configuration. This should be done in the `main` function or during the application's startup:

```go
func main() {
    conf.Initialize()
    // ...
}
```

3. Access the configuration values using the `viper.Get*` functions, based on the data type of the value you want to retrieve. Use the constants defined in the `conf` package as keys:

```go
logLevel := viper.GetString(conf.ViperLogLevelKey)
traceSampleRate := viper.GetFloat64(conf.ViperTraceSampleRateKey)
environment := viper.GetString(conf.ViperEnviormentKey)
```

4. Use the retrieved configuration values as needed in your application logic.

For example, you can use the `logLevel` value to set the log level of your logger, or the `traceSampleRate` value to configure the sampling rate for distributed tracing.

Remember that the configuration values can be set through command-line flags, environment variables, or configuration files (if supported by your application). The `viper` package automatically handles the precedence of these sources.

## Extending the Configuration

To add new configuration keys, command-line flags, or environment variables, follow these steps:

1. Open the `conf.go` file in a text editor.

2. To add a new configuration key, define a new constant in the `// viper keys` section:

```go
const (
    // ...
    ViperNewKeyKey = "new_key"
)
```

3. To add a new command-line flag, define a new constant in the `// flags` section and add a new `pflag.String` or `pflag.Bool` call in the `init` function:

```go
const (
    // ...
    NewFlag = "new-flag"
)

func init() {
    // ...
    _ = pflag.String(NewFlag, "", "Description of the new flag")
}
```

4. To add a new environment variable, define a new constant in the `// env vars` section:

```go
const (
    // ...
    NewEnv = "NEW_ENV"
)
```

5. In the `Initialize` function, bind the new configuration key to the corresponding command-line flag and environment variable using `viper.BindPFlag` and `viper.BindEnv`:

```go
func Initialize() {
    // ...
    _ = viper.BindPFlag(ViperNewKeyKey, pflag.Lookup(NewFlag))
    _ = viper.BindEnv(ViperNewKeyKey, NewEnv)
}
```

6. If needed, set a default value for the new configuration key using `viper.SetDefault`:

```go
func Initialize() {
    // ...
    viper.SetDefault(ViperNewKeyKey, "default_value")
}
```

7. Save the changes to the `conf.go` file.

Now, the new configuration key, command-line flag, and environment variable will be available for use in your Go application. You can access the value of the new configuration key using `viper.GetString(conf.ViperNewKeyKey)` or the appropriate `viper.Get*` function based on the data type.

Remember to update the README and other relevant documentation to include information about the new configuration options you've added.