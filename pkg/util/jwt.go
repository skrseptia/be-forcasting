package util

import (
	"food_delivery_api/cfg"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email string) (string, error) {
	// Create the Claims
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(cfg.Aut.Key))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseJWT(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Aut.Key), nil
	})

	var userID int
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userID = claims.UserID
	} else {
		return 0, err
	}

	return userID, nil
}
