package handlers

import (
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/db/repositories"
)

type Handlers struct {
	UserHandler     *UserHandler
	FootnoteHandler *FootnoteHandler
}

func NewHandlers(ur *repositories.UserRepository, fr *repositories.FootnoteRepository, services *services.Services) *Handlers {
	return &Handlers{
		UserHandler:     NewUserHandler(ur, services.TokenService),
		FootnoteHandler: NewFootnoteHandler(fr),
	}
}
