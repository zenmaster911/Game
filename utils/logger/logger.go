package logger

import (
	"log/slog"
	"os"
)

const (
	envlocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type WriterToSlog struct {
	Logger *slog.Logger
}

func (w *WriterToSlog) Write(p []byte) (n int, err error) {
	w.Logger.Info(string(p[:len(p)-1]))
	return len(p), nil
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envlocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
