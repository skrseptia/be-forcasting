package rest

import (
	"errors"
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/service"
	"food_delivery_api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type JWT struct {
	Token string `json:"token"`
}

func login(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body Auth
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.GetUserByEmailPassword(model.User{Email: body.Email, Password: body.Password})
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("invalid email or password").Error()})
			return
		}

		token, err := util.GenerateJWT(res.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: JWT{Token: token}})
	}
}
