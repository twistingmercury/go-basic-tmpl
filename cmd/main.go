package main

import (
	"context"
	"github.com/spf13/viper"
	"github.com/twistingmercury/telemetry/v2/logging"
	"github.com/twistingmercury/utils"
	"token_go_module/internal/conf"
	"token_go_module/internal/server"
)

func main() {
	conf.Initialize()

	ctx, cancel := context.WithCancel(context.Background())
	utils.ListenForInterrupt(cancel)

	if err := server.Bootstrap(ctx,
		viper.GetString(conf.ViperServiceNameKey),
		viper.GetString(conf.ViperBuildVersionKey),
		viper.GetString(conf.ViperNamespaceKey),
		viper.GetString(conf.ViperEnviormentKey),
	); err != nil {
		logging.Fatal(context.Background(), err, "failed to bootstrap the server")
	}

	server.Start()
}
