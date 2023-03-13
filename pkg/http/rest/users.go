package rest

import (
	"errors"
	"food_delivery_api/pkg/http/res"
	"food_delivery_api/pkg/storage/mysql/model"
	"food_delivery_api/pkg/svc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addUser(s svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.User
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{Error: err.Error()})
			return
		}

		obj, err := s.AddUser(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, res.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			Success: true,
			Data:    obj,
		})

	}
}

func editUser(s svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func getUsers(s svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var list []model.User
		var err error

		if list, err = s.GetUsers(); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			Success: true,
			Data:    list,
		})

	}
}

func getUser(s svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.User

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, res.Response{Error: errors.New("id must be uint").Error()})
			return
		}
		body.ID = uint(id)

		obj, err := s.GetUser(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, res.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			Success: true,
			Data:    obj,
		})

	}
}

func removeUser(s svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}
