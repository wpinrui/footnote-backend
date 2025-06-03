package handlers

import (
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/db/repositories"
)

type Handlers struct {
	UserHandler *UserHandler
}

func NewHandlers(ur *repositories.UserRepository, services *services.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(ur, services.TokenService),
	}
}
