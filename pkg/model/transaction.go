package model

type Transaction struct {
	Model
	NoTransaction string  `json:"no_transaction" `
	Username      string  `json:"username" binding:"required"`
	Total         string  `json:"Total"`
	Items         []Items `json:"items"`
}

type Items struct {
	Model
	TransactionId int     `json:"TransactionId" binding:"required"`
	Code          string  `json:"code" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Qty           float64 `json:"qty" binding:"qty" binding:"required"`
	Price         float64 `json:"Price" binding:"Price" binding:"required"`
	Total         float64 `json:"total" binding:"Price" binding:"required"`
	UOM           string  `json:"uom" binding:"required"`
}
