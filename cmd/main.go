// @title Footnote API
// @version 1.0
// @description API for managing footnotes
// @BasePath /

package main

import (
	"flag"
	"footnote-backend/internal/api"
	"footnote-backend/internal/api/handlers"
	"footnote-backend/internal/api/middleware"
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/config"
	"footnote-backend/internal/db"
	"footnote-backend/internal/db/repositories"
	"path/filepath"

	_ "footnote-backend/docs"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		panic("Error loading .env file")
	}
	cfg := config.NewConfig()
	err = cfg.ParseFlags()
	if err != nil {
		panic("Error parsing flags: " + err.Error())
	}
	flag.Parse()

	db, err := db.Connect(cfg)
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}
	defer db.Close()

	repos := repositories.NewRepositories(db)
	services := services.NewServices(cfg)
	handlers := handlers.NewHandlers(repos.UserRepository, repos.FootnoteRepository, services)
	middleware := middleware.NewMiddleware(services.TokenService)
	srv := api.NewAPI(cfg, handlers, middleware)
	err = srv.Run()
	if err != nil {
		panic("Error running server: " + err.Error())
	}
}
