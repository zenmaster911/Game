package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func MustLoad() *Config {

	configPath := "config/local.yaml"
	viper.SetConfigFile(configPath)
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading env file %s", err.Error())
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
	fmt.Println(cfg.DB.Password)
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	fmt.Println(cfg.DB.Password)

	return &cfg
}
