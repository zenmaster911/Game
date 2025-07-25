package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zenmaster911/Game/internal/config"
)

// type Config struct {
// 	Host     string
// 	Port     string
// 	Username string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

func NewPostgresDB(cfg *config.DBConfig) (*sqlx.DB, error) {
	// sqlDriver := stdlib.GetDefaultDriver()
	// sql.Register("pgx", sqlDriver)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to conncet to DB %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}
