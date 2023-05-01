package model

type Categories struct {
	Model
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}
