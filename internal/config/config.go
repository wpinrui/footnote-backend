package config

import (
	"flag"
	"os"
)

type Config struct {
	DSN        string
	JWT_SECRET string
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.DSN, "dsn", os.Getenv("DSN"), "Database connection string")
	flag.StringVar(&cfg.JWT_SECRET, "jwt_secret", os.Getenv("JWT_SECRET"), "JWT secret key")
	return nil
}
