package main

import (
	"context"
	"log"

	"{{module_name}}/server"

	"github.com/twistingmercury/utils"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	utils.ListenForInterrupt(cancel)

	if err := server.Bootstrap(ctx); err != nil {
		log.Fatal(err)
	}

	server.Start()
}
