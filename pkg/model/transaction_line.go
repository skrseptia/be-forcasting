package model

type TransactionLine struct {
	Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id" binding:"required"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Qty           float64 `json:"qty" binding:"required"`
	UOM           string  `json:"uom"`
	Price         float64 `json:"price"`
	SubTotal      float64 `json:"sub_total"`
}
