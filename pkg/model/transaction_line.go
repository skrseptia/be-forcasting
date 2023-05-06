package model

type TransactionLine struct {
	Model
	TransactionID uint    `json:"transaction_id" binding:"required"`
	Code          string  `json:"code" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Qty           float64 `json:"qty" binding:"required"`
	UOM           string  `json:"uom" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	SubTotal      float64 `json:"sub_total" binding:"required"`
}
