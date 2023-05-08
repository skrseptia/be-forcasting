package model

type Product struct {
	Model
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Qty         float64 `json:"qty" binding:"required"`
	UOMID       int     `json:"-"`
	UOM         UOM     `json:"uom"`
	Price       float64 `json:"price" binding:"required"`
}

type ProductRequest struct {
	Model
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	ImageURL    string  `json:"image_url"`
	Qty         float64 `json:"qty"  binding:"required"`
	UOMID       int     `json:"uom_id"`
	Price       float64 `json:"price" binding:"required"`
}
