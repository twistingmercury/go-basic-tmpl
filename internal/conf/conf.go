package conf

import (
	"fmt"
	"os"

	"github.com/twistingmercury/telemetry/logging"

	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ( // viper keys
	ViperOtelAddrKey        = "otel_collector_endpoint"
	ViperLogLevelKey        = "telemetry.log_level"
	ViperEnviormentKey      = "telemetry.environment"
	ViperTraceSampleRateKey = "telemetry.trace_sample_rate"
	ViperPromPortKey        = "telemetry.prom_port"
	ViperHeartbeatPortKey   = "telemetry.heartbeat_port"
	ViperBuildVersionKey    = "build_version"
	ViperBuildDateKey       = "build_date"
	ViperServiceNameKey     = "service_name"
	ViperShowHelpKey        = "show_help"
	ViperShowVersionKey     = "show_version"
	ViperNamespaceKey       = "namespace"
)

const ( // flags
	HelpFlg            = "help"
	VersionFlg         = "version"
	OtelColletorEPFlag = "otel-endpoint"
	LogLevelFlag       = "log-level"
	EnvFlag            = "environment"
	NamespaceFlag      = "namespace"
	HeartbeatPortFlag  = "heartbeat-port"
	PromPortFlag       = "prom-port"
	TOMLFileFlag       = "toml-file"
)

const ( // env vars
	OtelColletorEPEnv  = "OTEL_EXPORTER_OTLP_ENDPOINT"
	LogLevelEnv        = "LOG_LEVEL"
	EnvEnv             = "ENVIRONMENT"
	TraceSampleRateEvn = "OTEL_TRACE_SAMPLE_RATE"
	NamespaceEnv       = "NAMESPACE"
	HeartbeatPortEnv   = "HEARTBEAT_PORT"
	PromPortEnv        = "PROM_PORT"
)

const ( // Default values
	DefaultHelp                          = false
	DefaultVersion                       = false
	DefaultLogLevel        zerolog.Level = zerolog.DebugLevel
	DefaultEnv                           = "dev"
	DefaultTraceSampleRate float64       = 0.25
	DefaultHeartbeatPort                 = 8181
	DefaultPromPort                      = 9090
	DefaultTOMLfile                      = ""
	DefaultNamespace                     = ""
)

var exitFunc = os.Exit

var (
	buildVer    = "0.0.1"
	buildDate   = "n/a"
	serviceName = "token_go_bin"
)

var (
	// These are only used at startup, within this package.
	versionFlagValue  = pflag.Bool(VersionFlg, DefaultVersion, "Print the version information")
	helpFlagValue     = pflag.Bool(HelpFlg, DefaultHelp, "Print the help information")
	tomlPathFlagValue = pflag.String(TOMLFileFlag, DefaultTOMLfile, "Uses the provided filepath for configuration")

	// Viper will be used to access these by using [viper]
	_ = pflag.String(OtelColletorEPFlag, "", "OpenTelemetry collector endpoint")
	_ = pflag.String(LogLevelFlag, DefaultLogLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
	_ = pflag.String(EnvFlag, DefaultEnv, "Environment (dev, test, stage, prod)")
	_ = pflag.Int(PromPortFlag, DefaultPromPort, "Prometheus port")
	_ = pflag.Int(HeartbeatPortFlag, DefaultHeartbeatPort, "Heartbeat port")
	_ = pflag.String(NamespaceFlag, DefaultNamespace, "Prometheus namespace")
)

// Initialize initializes the configuration
func Initialize() {
	pflag.Parse()
	cfile := *tomlPathFlagValue

	if len(cfile) > 0 {
		if _, err := os.Stat(cfile); err == nil {
			viper.SetConfigFile(cfile)
			viper.SetConfigType("toml")
			if err = viper.ReadInConfig(); err != nil {
				logging.Fatal(err, "failed to read configuration file")
			}
		}
	}
	viper.AutomaticEnv()
	bind()

}

func bind() {
	_ = viper.BindPFlag(ViperOtelAddrKey, pflag.Lookup(OtelColletorEPFlag))
	_ = viper.BindPFlag(ViperLogLevelKey, pflag.Lookup(LogLevelFlag))
	_ = viper.BindPFlag(ViperEnviormentKey, pflag.Lookup(EnvFlag))
	_ = viper.BindPFlag(ViperNamespaceKey, pflag.Lookup(NamespaceFlag))
	_ = viper.BindPFlag(ViperHeartbeatPortKey, pflag.Lookup(HeartbeatPortFlag))
	_ = viper.BindPFlag(ViperPromPortKey, pflag.Lookup(PromPortFlag))

	_ = viper.BindEnv(ViperOtelAddrKey, OtelColletorEPEnv)
	_ = viper.BindEnv(ViperLogLevelKey, LogLevelEnv)
	_ = viper.BindEnv(ViperEnviormentKey, EnvEnv)
	_ = viper.BindEnv(ViperTraceSampleRateKey, TraceSampleRateEvn)
	_ = viper.BindEnv(ViperNamespaceKey, NamespaceEnv)
	_ = viper.BindEnv(ViperHeartbeatPortKey, HeartbeatPortEnv)
	_ = viper.BindEnv(ViperPromPortKey, PromPortEnv)

	viper.Set(ViperBuildVersionKey, buildVer)
	viper.Set(ViperBuildDateKey, buildDate)
	viper.Set(ViperServiceNameKey, serviceName)

	viper.SetDefault(ViperLogLevelKey, DefaultLogLevel.String())
	viper.SetDefault(ViperTraceSampleRateKey, DefaultTraceSampleRate)
	viper.SetDefault(ViperNamespaceKey, DefaultNamespace)
}

// ShowVersion prints the version information and exits the program.
func ShowVersion() {
	if !*versionFlagValue {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
	)
	exitFunc(0)
}

// ShowHelp prints the help information and exits the program.
func ShowHelp() {
	if !*helpFlagValue {
		return
	}

	fmt.Printf("%s\nversion: %s; build date: %s\n",
		viper.GetString(ViperServiceNameKey),
		viper.GetString(ViperBuildVersionKey),
		viper.GetString(ViperBuildDateKey),
	)
	fmt.Printf("Usage of %s:\n", viper.GetString(ViperServiceNameKey))
	pflag.PrintDefaults()
	println()
	exitFunc(0)
}
