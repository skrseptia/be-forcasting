package model

type Store struct {
	Model
	Name       string  `json:"name" binding:"required"`
	Phone      string  `json:"phone" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	Latitude   float64 `json:"latitude" binding:"required"`
	Longitude  float64 `json:"longitude" binding:"required"`
	MerchantID int     `json:"merchant_id"`
}
