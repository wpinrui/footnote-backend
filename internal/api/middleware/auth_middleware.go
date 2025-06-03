package middleware

import (
	"context"
	"errors"
	"footnote-backend/internal/api/services"
	"footnote-backend/internal/consts"
	"net/http"
)

type AuthMiddleware struct {
	TokenService *services.TokenService
}

func NewAuthMiddleware(ts *services.TokenService) *AuthMiddleware {
	return &AuthMiddleware{
		TokenService: ts,
	}
}

func (am *AuthMiddleware) AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(consts.AuthTokenCookieName)
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		uid, err := am.TokenService.ValidateToken(cookie.Value)
		// store user ID in context
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, consts.UidContextKey, uid)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func UserIdFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(consts.UidContextKey).(int)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}
	return userID, nil
}
