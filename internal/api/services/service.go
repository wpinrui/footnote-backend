package services

import "footnote-backend/internal/config"

type Services struct {
	TokenService *TokenService
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		TokenService: NewTokenService(cfg.JwtSecret),
	}
}
