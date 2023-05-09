package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Success    bool        `json:"success"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"page_total"`
	TotalRows  int64       `json:"total_rows"`
}

func paginate(c *gin.Context, data interface{}, total int64) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	if page <= 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	c.JSON(http.StatusOK, Pagination{
		Success:    true,
		Data:       data,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(total/int64(pageSize)) + 1,
		TotalRows:  total,
	})
}
