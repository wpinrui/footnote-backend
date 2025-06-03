package middleware

import (
	"footnote-backend/internal/api/services"
)

type Middleware struct {
	AuthMiddleware *AuthMiddleware
}

func NewMiddleware(ts *services.TokenService) *Middleware {
	return &Middleware{
		AuthMiddleware: NewAuthMiddleware(ts),
	}
}
