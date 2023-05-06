package middleware

import (
	"errors"
	"food_delivery_api/cfg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `json:"authorization"`
}

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := Header{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
		}

		bearer := strings.Replace(header.Authorization, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(bearer, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Aut.Key), nil
		}, jwt.WithLeeway(5*time.Second))

		if token == nil {
			c.JSON(http.StatusUnauthorized, Response{Error: errors.New("invalid token").Error()})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(*jwt.RegisteredClaims); ok && !token.Valid {
			c.JSON(http.StatusUnauthorized, Response{Error: err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
