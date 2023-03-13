package model

type Merchant struct {
	Model
	UserID   uint   `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	ImageURL string `json:"image_url"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}
