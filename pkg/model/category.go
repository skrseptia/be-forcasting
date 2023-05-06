package model

type Category struct {
	Model
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
