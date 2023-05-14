package model

type Dashboard struct {
	TotalCategory    int64 `json:"total_category"`
	TotalProduct     int64 `json:"total_product"`
	TotalTransaction int64 `json:"total_transaction"`
	TotalUOM         int64 `json:"total_uom"`
	TotalCustomer    int64 `json:"total_customer"`
}
