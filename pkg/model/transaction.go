package model

type Transaction struct {
	Model
	TrxDate          string            `json:"trx_date"`
	TrxID            string            `json:"trx_id"`
	CreatedBy        string            `json:"created_by"`
	Customer         string            `json:"customer" binding:"required"`
	Total            float64           `json:"total"`
	TransactionLines []TransactionLine `json:"transaction_lines"`
}
