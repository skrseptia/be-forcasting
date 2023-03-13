package model

type Product struct {
	Model
	MerchantID  uint    `json:"merchant_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required"`
}
