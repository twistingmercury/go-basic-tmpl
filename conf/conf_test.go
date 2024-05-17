package conf_test

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twistingmercury/tmpl/conf"
)

const (
	tCommit      = "1a2b3c4d"
	tDate        = "2021-01-01T00:00:00Z"
	tVer         = "1.0.0"
	tlogLevel    = "debug"
	tServiceName = "tunnelvision"
	tOtelColEP   = "http://localhost:4317"
)

func TestDefaultValues(t *testing.T) {
	defer viper.Reset()
	conf.SetBuildCommit(tCommit)
	conf.SetBuildDate(tDate)
	conf.SetBuildVersion(tVer)
	conf.SetServiceName(tServiceName)
	conf.SetLogLevel(tlogLevel)
	viper.Set(conf.ViperOtelAddrKey, tOtelColEP)
	conf.Initialize()

	assert.False(t, viper.GetBool(conf.ViperShowVersionKey))
	assert.False(t, viper.GetBool(conf.ViperShowHelpKey))
	assert.Equal(t, tlogLevel, viper.GetString(conf.ViperLogLevelKey))
	assert.Equal(t, tVer, viper.GetString(conf.ViperBuildVersionKey))
	assert.Equal(t, tDate, viper.GetString(conf.ViperBuildDateKey))
	assert.Equal(t, tCommit, viper.GetString(conf.ViperCommitHashKey))
	assert.Equal(t, tServiceName, viper.GetString(conf.ViperServiceNameKey))
}

func TestEnvVars(t *testing.T) {
	defer viper.Reset()
	const ep = "http://test-host:4317"
	err := os.Setenv(conf.OtelColletorEPEnv, ep)

	require.NoError(t, err)
	defer os.Unsetenv(conf.OtelColletorEPEnv)

	err = os.Setenv(conf.LogLevelEnv, "info")
	defer os.Unsetenv(conf.LogLevelEnv)

	conf.Initialize()

	assert.Equal(t, ep, viper.GetString(conf.ViperOtelAddrKey))
	assert.Equal(t, "info", viper.GetString(conf.ViperLogLevelKey))
}

