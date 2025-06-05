package routes

import (
	"footnote-backend/internal/api/handlers"
	"footnote-backend/internal/api/middleware"

	_ "footnote-backend/docs" // Swagger docs

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers, m *middleware.Middleware) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/create", h.UserHandler.CreateUser)
		r.Post("/login", h.UserHandler.Login)
	})

	// Footnote routes
	r.Route("/footnote", func(r chi.Router) {
		r.Use(m.AuthMiddleware.AuthenticateToken)
		r.Get("/", h.FootnoteHandler.GetFootnotes)
		r.Post("/", h.FootnoteHandler.CreateFootnote)
		r.Get("/{id}", h.FootnoteHandler.GetFootnoteByID)
		r.Put("/{id}", h.FootnoteHandler.UpdateFootnote)
		r.Delete("/{id}", h.FootnoteHandler.DeleteFootnote)
		r.Get("/search", h.FootnoteHandler.SearchFootnotes)
	})
}
