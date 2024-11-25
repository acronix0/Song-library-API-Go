package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	AppEnv              string             `env:"APP_ENV" envDefault:"local"`
	DatabaseConndection DatabaseConnection `envPrefix:"DB_"`
	HTTPConfig          HTTPServer         `envPrefix:"HTTP_"`
	MigrationsPath      string             `env:"MIGRATIONS_PATH" envDefault:"../../migrations"`
	ExternalAPIBaseURL  string             `env:"EXTERNAL_API_BASE_URL"`
}

type DatabaseConnection struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"secret"`
	Name     string `env:"NAME" envDefault:"songs"`
}

type HTTPServer struct {
	Host               string        `env:"HOST" envDefault:"0.0.0.0"`
	Port               string        `env:"PORT" envDefault:"8080"`
	ReadTimeout        time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout       time.Duration `env:"WRITE_TIMEOUT" envDefault:"5s"`
	MaxHeaderMegabytes int           `env:"MAX_HEADER_MEGABYTES" envDefault:"1"`
}

func Load() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: No .env file found, using environment variables")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
