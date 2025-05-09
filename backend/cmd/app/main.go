package main

import (
	"github.com/mikemonzo/goshare/internal/shared/logger"
	"github.com/mikemonzo/goshare/pkg/config"
)

func main() {
	logger.InitLogger()
	// Load .env
	cfg := config.LoadConfig()

	logger.Log.Infof("Starting server on port %s...", cfg.AppPort)

	// TODO: Initialize routes, servicies, DB, middlewares, etc.

}
