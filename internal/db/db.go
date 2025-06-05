package db

import (
	"database/sql"
	"footnote-backend/internal/cloudsql"
	"footnote-backend/internal/config"
	"os"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	if os.Getenv("ENV") == "local" {
		db, err := sql.Open("postgres", cfg.Dsn)
		if err != nil {
			return nil, err
		}
		if err := db.Ping(); err != nil {
			return nil, err
		}
		return db, nil
	}
	return cloudsql.ConnectWithConnector()
}
