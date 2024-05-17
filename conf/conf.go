package conf

import (
	"fmt"
	"os"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ( // viper keys
	ViperOtelAddrKey        = "otel_collector_endpoint"
	ViperLogLevelKey        = "log_level"
	ViperEnviormentKey      = "environment"
	ViperNamespaceKey       = "namespace"
	ViperTraceSampleRateKey = "trace_sample_rate"
	ViperBuildVersionKey    = "build_version"
	ViperBuildDateKey       = "build_date"
	ViperCommitHashKey      = "commit_hash"
	ViperServiceNameKey     = "service_name"
	ViperShowHelpKey        = "show_help"
	ViperShowVersionKey     = "show_version"
)

const ( // flags
	HelpFlg            = "help"
	VersionFlg         = "version"
	OtelColletorEPFlag = "otel-endpoint"
	LogLevelFlag       = "log-level"
	EnvFlag            = "environment"
)

const ( // env vars
	OtelColletorEPEnv  = "OTEL_EXPORTER_OTLP_ENDPOINT"
	LogLevelEnv        = "LOG_LEVEL"
	EnvEnv             = "ENVIRONMENT"
	TraceSampleRateEvn = "OTEL_TRACE_SAMPLE_RATE"
	HeartbeatPortEnv   = "HEARTBEAT_PORT"
)

const ( // Default values
	DefaultHelp                          = false
	DefaultVersion                       = false
	DefaultLogLevel        zerolog.Level = zerolog.DebugLevel
	DefaultEnv                           = "dev"
	DefaultTraceSampleRate               = 0.25
	DefaultHeartbeatPort                 = 8081
)

var exitFunc = os.Exit

var (
	// build information: this is set at compile time using LDFlags
	buildVer    = "n/a"
	buildDate   = "n/a"
	buildCommit = "n/a"
	serviceName = "{{service_name}}"
)

var (
	_ = pflag.Bool(VersionFlg, DefaultVersion, "Print the version information")
	_ = pflag.Bool(HelpFlg, DefaultHelp, "Print the help information")
	_ = pflag.String(OtelColletorEPFlag, "", "OpenTelemetry collector endpoint")
	_ = pflag.String(LogLevelFlag, DefaultLogLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	_ = pflag.String(EnvFlag, DefaultEnv, "Environment (dev, test, stage, prod)")
)

// Initialize initializes the configuration
func Initialize() {
	pflag.Parse()
	_ = viper.BindPFlag(ViperShowVersionKey, pflag.Lookup(VersionFlg))
	_ = viper.BindPFlag(ViperShowHelpKey, pflag.Lookup(HelpFlg))
	_ = viper.BindPFlag(ViperOtelAddrKey, pflag.Lookup(OtelColletorEPFlag))
	_ = viper.BindPFlag(ViperLogLevelKey, pflag.Lookup(LogLevelFlag))
	_ = viper.BindPFlag(ViperEnviormentKey, pflag.Lookup(EnvFlag))

	_ = viper.BindEnv(ViperOtelAddrKey, OtelColletorEPEnv)
	_ = viper.BindEnv(ViperLogLevelKey, LogLevelEnv)
	_ = viper.BindEnv(ViperEnviormentKey, EnvEnv)
	_ = viper.BindEnv(ViperTraceSampleRateKey, TraceSampleRateEvn)

	viper.Set(ViperBuildVersionKey, buildVer)
	viper.Set(ViperBuildDateKey, buildDate)
	viper.Set(ViperCommitHashKey, buildCommit)
	viper.Set(ViperServiceNameKey, serviceName)

	viper.SetDefault(ViperLogLevelKey, DefaultLogLevel.String())
	viper.SetDefault(ViperTraceSampleRateKey, DefaultTraceSampleRate)
}


// ShowVersion prints the version information and exits the program.
func ShowVersion() {
	if !viper.GetBool(ViperShowVersionKey) {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s; commit: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
		viper.GetString(ViperCommitHashKey),
	)
	exitFunc(0)
}

// ShowHelp prints the help information and exits the program.
func ShowHelp() {
	if !viper.GetBool(ViperShowHelpKey) {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s; commit: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
		viper.GetString(ViperCommitHashKey),
	)
	fmt.Printf("Usage of %s:\n", viper.GetString(ViperServiceNameKey))
	pflag.PrintDefaults()
	println()
	exitFunc(0)
}