package model

type TransactionDetail struct {
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
