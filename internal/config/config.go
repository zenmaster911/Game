package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string `mapstructure:"env"`
	StoragePath string `mapstructure:"storage_path"`
	HTTPServer  `mapstructure:"http_server"`
}

type HTTPServer struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

func MustLoad() *Config {
	viper.SetDefault("env", "development")
	viper.SetDefault("http_server.address", "0.0.0.0:8080")
	viper.SetDefault("http_server.timeout", "5s")
	viper.SetDefault("http_server.idle_timeout", "60s")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}

	if cfg.StoragePath == "" {
		log.Fatal("storage path required")
	}

	return &cfg
}
