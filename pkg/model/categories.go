package model

type Categories struct {
	Model
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required" gorm:"unique"`
}
