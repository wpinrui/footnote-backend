package config

import (
	"fmt"
	"os"
)

type Config struct {
	Dsn       string
	JwtSecret string
	Env       string
}

func NewConfig() *Config {
	env := os.Getenv("ENV") // e.g. "local" or "cloud"

	var dsn string
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	return &Config{
		Dsn:       dsn,
		JwtSecret: os.Getenv("JWT_SECRET"),
		Env:       env,
	}
}
