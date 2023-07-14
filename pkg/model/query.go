package model

type QueryPagination struct {
	Page        int    `form:"page" binding:"required"`
	PageSize    int    `form:"page_size" binding:"required"`
	Name        string `form:"name"`
	FullName    string `form:"fullname"`
	Email       string `form:"email"`
	Description string `form:"description"`
}

type QueryGetTransactions struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Customer  string `form:"customer"`
	CreatedBy  string `form:"created_by"`
	QueryPagination
}

type QueryGetExponentialSmoothing struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	ProductID string `form:"product_id"`
}
