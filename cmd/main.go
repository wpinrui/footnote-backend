// @title Footnote API
// @version 1.0
// @description API for managing footnotes
// @BasePath /
// @securityDefinitions.apikey cookieAuth
// @in cookie
// @name auth_token

package main

import (
	"footnote-backend/internal/api"
	"footnote-backend/internal/api/handlers"
	"footnote-backend/internal/api/middleware"
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/config"
	"footnote-backend/internal/db"
	"footnote-backend/internal/db/repositories"

	_ "footnote-backend/docs"
)

func main() {
	cfg := config.NewConfig()
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
