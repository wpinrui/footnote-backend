package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenHandler struct {
	Secret string
}

func NewTokenHandler(secret string) *TokenHandler {
	return &TokenHandler{Secret: secret}
}

func (tm *TokenHandler) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.Secret))
}

func (tm *TokenHandler) ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(tm.Secret), nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(float64); ok {
			return int(userID), nil
		}
	}
	return 0, jwt.ErrInvalidKeyType
}

func (tm *TokenHandler) RefreshToken(tokenString string) (string, error) {
	userID, err := tm.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return tm.GenerateToken(userID)
}
