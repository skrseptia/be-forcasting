package model

type Product struct {
	Model
	MerchantID  string  `json:"merchant_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  string  `json:"category_id" binding:"required"`
}
