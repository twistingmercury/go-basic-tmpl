package conf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"MODULE_NAME/conf"

	"github.com/stretchr/testify/assert"
)

const (
	tDate        = "2021-01-01T00:00:00Z"
	tVer         = "1.0.0"
	tlogLevel    = "debug"
	tServiceName = "tunnelvision"
	tOtelColEP   = "http://localhost:4317"
)

func init() {
	conf.SetBuildVersion(tVer)
	conf.SetServiceName(tServiceName)
	conf.SetBuildDate(tDate)
	conf.SetLogLevel(tlogLevel)
	viper.Set(conf.ViperOtelAddrKey, tOtelColEP)
}

func TestDefaultValues(t *testing.T) {
	defer viper.Reset()
	conf.Initialize()

	assert.False(t, viper.GetBool(conf.ViperShowVersionKey))
	assert.False(t, viper.GetBool(conf.ViperShowHelpKey))
	assert.Equal(t, tlogLevel, viper.GetString(conf.ViperLogLevelKey))
	assert.Equal(t, tVer, viper.GetString(conf.ViperBuildVersionKey))
	assert.Equal(t, tDate, viper.GetString(conf.ViperBuildDateKey))
	assert.Equal(t, tServiceName, viper.GetString(conf.ViperServiceNameKey))
	assert.Equal(t, conf.DefaultTraceSampleRate, viper.GetFloat64(conf.ViperTraceSampleRateKey))
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

func TestShowVersion(t *testing.T) {
	oldStdout := os.Stdout
	tmpStdout, _ := os.CreateTemp("", "tmpStdout")
	os.Stdout = tmpStdout

	os.Args = append(os.Args, "--version")
	conf.Initialize()

	defer func() {
		viper.Reset()
		os.Stdout = oldStdout
		_ = tmpStdout.Close()
		os.Remove(tmpStdout.Name())
	}()

	conf.SetExitFunc(func(code int) {})

	conf.ShowVersion()
	content, err := os.ReadFile(tmpStdout.Name())
	require.NoError(t, err)
	//expected := `version: 0.0.0; build date: 2021-01-01T00:00:00Z; commit: fake`
	expected := fmt.Sprintf("version: %s; build date: %s", tVer, tDate)
	actual := string(content)
	assert.Contains(t, actual, expected)
}

func TestShowHelp(t *testing.T) {
	oldStdout := os.Stdout
	tmpStdout, _ := os.CreateTemp("", "tmpStdout")
	os.Stdout = tmpStdout

	os.Args = append(os.Args, "--help")
	conf.Initialize()

	defer func() {
		viper.Reset()
		os.Stdout = oldStdout
		_ = tmpStdout.Close()
		os.Remove(tmpStdout.Name())
	}()

	conf.SetExitFunc(func(code int) {})
	conf.ShowHelp()

	content, err := os.ReadFile(tmpStdout.Name())
	require.NoError(t, err)

	expected := `Usage of`
	actual := string(content)
	assert.Contains(t, actual, expected)
}
