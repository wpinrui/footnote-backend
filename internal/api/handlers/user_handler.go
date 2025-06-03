package handlers

import (
	"encoding/json"
	"footnote-backend/internal/api/models"
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/consts"
	"footnote-backend/internal/db/repositories"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
	TokenService   *services.TokenService
}

func NewUserHandler(ur *repositories.UserRepository, tk *services.TokenService) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
		TokenService:   tk,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	id, err := uh.UserRepository.Create(&models.User{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	tokenString, err := uh.TokenService.GenerateToken(id)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     consts.AuthTokenCookieName,
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // dev
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusCreated)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := uh.UserRepository.FindByEmail(req.Email)
	if err != nil || user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(req.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := uh.TokenService.GenerateToken(user.Id)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     consts.AuthTokenCookieName,
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // dev
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}
