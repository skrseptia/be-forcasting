package rest

import (
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
