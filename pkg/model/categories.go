package model

type Categories struct {
	Model
	Name string `json:"name" binding:"required" gorm:"unique"`
}
