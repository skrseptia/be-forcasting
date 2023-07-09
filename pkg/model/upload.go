package model

type Upload struct {
	File         string `json:"file"`
	Table        string `json:"table"`
	RowsAffected int64  `json:"rows_affected"`
}
