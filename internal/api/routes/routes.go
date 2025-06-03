package routes

import (
	"footnote-backend/internal/api/handlers"
	"footnote-backend/internal/api/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers, m *middleware.Middleware) {
	// r.Get("/health", h.HealthCheck)
	// r.Route("/users", func(r chi.Router) {
	// 	r.Get("/", h.GetUsers)
	// 	r.Post("/", h.CreateUser)
	// 	r.Get("/{id}", h.GetUserByID)
	// 	r.Put("/{id}", h.UpdateUser)
	// 	r.Delete("/{id}", h.DeleteUser)
	// })
	// r.Route("/notes", func(r chi.Router) {
	// 	r.Get("/", h.GetNotes)
	// 	r.Post("/", h.CreateNote)
	// 	r.Get("/{id}", h.GetNoteByID)
	// 	r.Put("/{id}", h.UpdateNote)
	// 	r.Delete("/{id}", h.DeleteNote)
	// })
	r.Route("/users", func(r chi.Router) {
		r.Post("/create", h.UserHandler.CreateUser)
		r.Post("/login", h.UserHandler.Login)
	})
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
