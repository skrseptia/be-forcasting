package rest

import (
	"food_delivery_api/pkg/adding"
	"food_delivery_api/pkg/editing"
	"food_delivery_api/pkg/listing"
	"food_delivery_api/pkg/removing"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(a adding.Service, e editing.Service, l listing.Service, r removing.Service) *gin.Engine {
	router := gin.Default()

	// Public API
	router.GET("/health", getHealthStatus)

	// Protected API
	v1 := router.Group("api/v1")
	{
		// Adding
		v1.POST("/users", addUser(a))

		// Editing
		v1.PUT("/users", editUser(e))

		// Listing
		v1.GET("/users/:id", getUser(l))

		// Removing
		v1.DELETE("/users", removeUser(r))
	}

	return router
}

func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
