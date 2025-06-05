package api

import (
	"footnote-backend/internal/api/handlers"
	"footnote-backend/internal/api/middleware"
	"footnote-backend/internal/api/routes"
	"footnote-backend/internal/config"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
)

type API struct {
	Router *chi.Mux
	Config *config.Config
	Wg     *sync.WaitGroup
}

func NewAPI(cfg *config.Config, handlers *handlers.Handlers, middleware *middleware.Middleware) *API {
	router := chi.NewRouter()
	// router.Use(middleware.something) // Add any middleware you need here
	routes.SetupRoutes(router, handlers, middleware)
	return &API{
		Router: router,
		Config: cfg,
	}
}

func (api *API) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: api.Router,
	}
	return srv.ListenAndServe()

}
