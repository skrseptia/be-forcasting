package model

type Upload struct {
	File         string `json:"file"`
	Table        string `json:"table"`
	RowsAffected int    `json:"rows_affected"`
}
