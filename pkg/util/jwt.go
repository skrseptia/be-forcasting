package util

import (
	"food_delivery_api/cfg"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(email string) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		Issuer:    cfg.AppName,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(cfg.Aut.Key))
	if err != nil {
		return "", err
	}

	return ss, nil
}
