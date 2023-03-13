package rest

import (
	"errors"
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addMerchant(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Merchant
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.AddMerchant(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getMerchants(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res []model.Merchant
		var err error

		if res, err = s.GetMerchants(); err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getMerchant(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var res model.Merchant

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("id must be uint").Error()})
			return
		}
		res.ID = uint(id)

		res, err = s.GetMerchant(res)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func editMerchant(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Merchant
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

		res, err := s.EditMerchant(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func removeMerchant(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Merchant

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: errors.New("id must be uint").Error()})
			return
		}
		body.ID = uint(id)

		res, err := s.RemoveMerchant(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}
