package main

import (
	"log/slog"

	"github.com/zenmaster911/Game/internal/config"
	"github.com/zenmaster911/Game/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("Initializing server", slog.String("address,", cfg.Address))
	log.Debug("logger debug mode enabled")

}
