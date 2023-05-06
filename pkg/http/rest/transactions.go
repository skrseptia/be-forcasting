package rest

import (
	"errors"
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/service"
	"food_delivery_api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func addTransaction(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Transaction
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		var user model.User
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)

		id, err := util.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		user.ID = uint(id)

		acc, err := s.GetUser(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.AddTransaction(body, acc.FullName)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getTransactions(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res []model.Transaction
		var err error

		if res, err = s.GetTransactions(); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getTransaction(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res model.Transaction

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("id must be uint").Error()})
			return
		}
		res.ID = uint(id)

		res, err = s.GetTransaction(res)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func editTransaction(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Transaction
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("id must be uint").Error()})
			return
		}
		body.ID = uint(id)

		res, err := s.EditTransaction(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func removeTransaction(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Transaction

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("id must be uint").Error()})
			return
		}
		body.ID = uint(id)

		res, err := s.RemoveTransaction(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}
