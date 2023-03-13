package rest

import (
	"food_delivery_api/pkg/svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(s svc.Service) *gin.Engine {
	router := gin.Default()

	// Public API
	router.GET("/health", getHealthStatus)

	// Protected API
	v1 := router.Group("api/v1")
	{
		// Adding
		v1.POST("/users", addUser(s))

		// Editing
		v1.PUT("/users/:id", editUser(s))

		// Listing
		v1.GET("/users", getUsers(s))
		v1.GET("/users/:id", getUser(s))

		// Removing
		v1.DELETE("/users/:id", removeUser(s))
	}

	return router
}

func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
