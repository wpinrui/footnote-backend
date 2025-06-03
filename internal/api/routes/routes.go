package routes

import (
	"footnote-backend/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
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
}
