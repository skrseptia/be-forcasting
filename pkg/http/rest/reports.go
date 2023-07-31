package rest

import (
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getReportDashboard(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := s.GetReportDashboard()
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getReportChart(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := s.GetReportChart()
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getReportExponentialSmoothing(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		qp := model.QueryGetExponentialSmoothing{}
		err := c.ShouldBindQuery(&qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.GetReportExponentialSmoothing(qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getReportMonthlyExponentialSmoothing(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		qp := model.QueryGetExponentialSmoothing{}
		err := c.ShouldBindQuery(&qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.GetReportMonthlyExponentialSmoothing(qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}

func getReportArima(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		qp := model.QueryGetArima{}
		err := c.ShouldBindQuery(&qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		res, err := s.GetReportArima(qp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, Response{Success: true, Data: res})
	}
}
