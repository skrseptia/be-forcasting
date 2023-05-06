package model

type Transaction struct {
	Model
	Number           string            `json:"number"`
	CreatedBy        string            `json:"created_by"`
	Customer         string            `json:"customer" binding:"required"`
	Total            float64           `json:"total"`
	TransactionLines []TransactionLine `json:"transaction_lines"`
}
