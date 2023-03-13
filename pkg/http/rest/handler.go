package rest

import (
	"food_delivery_api/pkg/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(s svc.Service) *gin.Engine {
	r := gin.Default()

	// Public API
	r.GET("/health", getHealthStatus)

	// Protected API
	v1 := r.Group("/api/v1")
	{
		// Users
		v1.POST("/users", addUser(s))
		v1.GET("/users", getUsers(s))
		v1.GET("/users/:id", getUser(s))
		v1.PUT("/users/:id", editUser(s))
		v1.DELETE("/users/:id", removeUser(s))

		// Merchants
		v1.POST("/merchants", addMerchant(s))
		v1.GET("/merchants", getMerchants(s))
		v1.GET("/merchants/:id", getMerchant(s))
		v1.PUT("/merchants/:id", editMerchant(s))
		v1.DELETE("/merchants/:id", removeMerchant(s))
	}

	return r
}

func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
