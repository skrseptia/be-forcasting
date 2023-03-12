package rest

import (
	"food_delivery_api/pkg/listing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(l listing.Service) *gin.Engine {
	r := gin.Default()

	r.GET("/health", getHealthStatus)

	return r
}

func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
