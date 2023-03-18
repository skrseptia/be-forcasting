package model

type Merchant struct {
	Model
	Name      string  `json:"name" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	ImageURL  string  `json:"image_url"`
	Phone     string  `json:"phone" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
