package conf_test

import (
	"fmt"
	"os"
	"testing"
	conf2 "token_go_module/internal/conf"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

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
	conf2.SetBuildVersion(tVer)
	conf2.SetServiceName(tServiceName)
	conf2.SetBuildDate(tDate)
	conf2.SetLogLevel(tlogLevel)
	viper.Set(conf2.ViperOtelAddrKey, tOtelColEP)
}

func TestDefaultValues(t *testing.T) {
	defer viper.Reset()
	conf2.Initialize()

	assert.False(t, viper.GetBool(conf2.ViperShowVersionKey))
	assert.False(t, viper.GetBool(conf2.ViperShowHelpKey))
	assert.Equal(t, tlogLevel, viper.GetString(conf2.ViperLogLevelKey))
	assert.Equal(t, tVer, viper.GetString(conf2.ViperBuildVersionKey))
	assert.Equal(t, tDate, viper.GetString(conf2.ViperBuildDateKey))
	assert.Equal(t, tServiceName, viper.GetString(conf2.ViperServiceNameKey))
	assert.Equal(t, conf2.DefaultTraceSampleRate, viper.GetFloat64(conf2.ViperTraceSampleRateKey))
}

func TestEnvVars(t *testing.T) {
	defer viper.Reset()
	const ep = "http://test-host:4317"
	err := os.Setenv(conf2.OtelColletorEPEnv, ep)

	require.NoError(t, err)
	defer os.Unsetenv(conf2.OtelColletorEPEnv)

	err = os.Setenv(conf2.LogLevelEnv, "info")
	defer os.Unsetenv(conf2.LogLevelEnv)

	conf2.Initialize()

	assert.Equal(t, ep, viper.GetString(conf2.ViperOtelAddrKey))
	assert.Equal(t, "info", viper.GetString(conf2.ViperLogLevelKey))
}

func TestShowVersion(t *testing.T) {
	oldStdout := os.Stdout
	tmpStdout, _ := os.CreateTemp("", "tmpStdout")
	os.Stdout = tmpStdout

	os.Args = append(os.Args, "--version")
	conf2.Initialize()

	defer func() {
		viper.Reset()
		os.Stdout = oldStdout
		_ = tmpStdout.Close()
		os.Remove(tmpStdout.Name())
	}()

	conf2.SetExitFunc(func(code int) {})

	conf2.ShowVersion()
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
	conf2.Initialize()

	defer func() {
		viper.Reset()
		os.Stdout = oldStdout
		_ = tmpStdout.Close()
		os.Remove(tmpStdout.Name())
	}()

	conf2.SetExitFunc(func(code int) {})
	conf2.ShowHelp()

	content, err := os.ReadFile(tmpStdout.Name())
	require.NoError(t, err)

	expected := `Usage of`
	actual := string(content)
	assert.Contains(t, actual, expected)
}
