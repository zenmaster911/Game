package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/zenmaster911/Game/internal/config"
	"github.com/zenmaster911/Game/internal/db"
	"github.com/zenmaster911/Game/pkg/handler"
	"github.com/zenmaster911/Game/pkg/repository"
	"github.com/zenmaster911/Game/pkg/service"
)

func main() {
	cfg := config.MustLoad()

	// logg := logger.SetupLogger(cfg.Env)
	// logg = logg.With(slog.String("env", cfg.Env))
	// slog.SetDefault(logg)

	// log.SetFlags(0)
	// log.SetOutput(&logger.WriterToSlog{
	// 	Logger: slog.Default().With(
	// 		slog.String("source", "chi-middleware"),
	// 	),
	// })
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	hand := slog.NewJSONHandler(os.Stdout, opts) // можно использовать TextHandler
	logger := slog.New(hand)
	slog.SetDefault(logger)

	slog.Info("Starting server", "component", "main")

	dbConn, err := db.NewPostgresDB(cfg.DB)
	if err != nil {
		slog.Error("faile to connect to database", "error", err)
		log.Fatal(err)
	}
	defer dbConn.Close()

	userRepo := repository.NewUserRepository(dbConn)
	UserService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(UserService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/users/signup", userHandler.SignUp)

	//logger.Info("Initializing server", slog.String("address,", cfg.Address))
	//logg.Debug("logger debug mode enabled")

	http.ListenAndServe(":8080", router)
}
