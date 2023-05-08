package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success  bool        `json:"success"`
	Error    string      `json:"error,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
}

func successResponse(c *gin.Context, data interface{}) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	c.JSON(http.StatusOK, Response{Success: true, Data: data, Page: page, PageSize: pageSize})
}
