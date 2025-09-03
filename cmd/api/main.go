package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/kanevidzro/gokernel/api"
	"github.com/kanevidzro/gokernel/pkg/config"
)

func main() {
	cfg := config.Load()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to init zap logger: %v", err)
	}
	defer logger.Sync()

	server, err := api.NewServer(logger, cfg)
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

	if err := server.Run(); err != nil {
		logger.Fatal("server exited with error", zap.Error(err))
	}
}
