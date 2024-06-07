package conf

import "github.com/spf13/viper"

func SetBuildVersion(v string) {
	buildVer = v
}

func SetBuildDate(d string) {
	buildDate = d
}

func SetServiceName(s string) {
	serviceName = s
}

func SetLogLevel(l string) {
	viper.Set(ViperLogLevelKey, l)
}

func SetExitFunc(f func(int)) {
	exitFunc = f
}
