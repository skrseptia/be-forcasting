package model

type QueryPagination struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

type QueryGetTransactions struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Customer  string `form:"customer"`
	QueryPagination
}

type QueryGetExponentialSmoothing struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	ProductID string `form:"product_id"`
}
