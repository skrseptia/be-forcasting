package model

type User struct {
	Model
	FullName   string `json:"full_name" binding:"required"`
	Email      string `json:"email" binding:"required" gorm:"unique"`
	Password   string `json:"password" binding:"required"`
	ImageURL   string `json:"image_url"`
	Phone      string `json:"phone" binding:"required" gorm:"unique"`
	Address    string `json:"address" binding:"required"`
	UserType   string `json:"user_type" binding:"required"`
	MerchantID string `json:"merchant_id" `
}
