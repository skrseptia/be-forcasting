package rest

import (
	"errors"
	"food_delivery_api/pkg/adding"
	"food_delivery_api/pkg/editing"
	"food_delivery_api/pkg/listing"
	"food_delivery_api/pkg/removing"
	"food_delivery_api/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addUser(a adding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var au adding.User

		err := c.ShouldBindJSON(&au)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Error: err.Error()})
			return
		}

		au, err = a.AddUser(au)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, util.Response{
			Success: true,
			Data:    au,
		})

	}
}

func editUser(e editing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})

	}
}

func getUser(l listing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lu listing.User

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Error: errors.New("id must be uint").Error()})
			return
		}
		lu.ID = uint(id)

		lu, err = l.GetUser(lu)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.Response{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, util.Response{
			Success: true,
			Data:    lu,
		})

	}
}

func removeUser(r removing.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})

	}
}
