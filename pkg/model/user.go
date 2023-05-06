package model

type User struct {
	Model
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
	ImageURL string `json:"image_url"`
	Phone    string `json:"phone" binding:"required" gorm:"unique"`
	Address  string `json:"address" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
