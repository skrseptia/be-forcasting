package model

type Uom struct {
	Model
	Name string `json:"name" binding:"required"`
}
