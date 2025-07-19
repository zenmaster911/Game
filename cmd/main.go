package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zenmaster911/Game/internal/config"
	"github.com/zenmaster911/Game/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	logg := logger.SetupLogger(cfg.Env)
	logg = logg.With(slog.String("env", cfg.Env))
	slog.SetDefault(logg)

	log.SetFlags(0)
	log.SetOutput(&logger.WriterToSlog{
		Logger: slog.Default().With(
			slog.String("source", "chi-middleware"),
		),
	})

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handling request", "method", r.Method, "url", r.URL.Path)
	})

	logg.Info("Initializing server", slog.String("address,", cfg.Address))
	logg.Debug("logger debug mode enabled")

	http.ListenAndServe(":8080", router)
}
