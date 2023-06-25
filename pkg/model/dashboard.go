package model

import "time"

type Dashboard struct {
	Summary     Summary       `json:"summary"`
	CustomerTrx []CustomerTrx `json:"customer_trx"`
	StockAlert  []StockAlert  `json:"stock_alert"`
	Top10Trx    []Top10Trx    `json:"top_10_trx"`
	Top5Product []Top5Product `json:"top_5_product"`
}

type Summary struct {
	TotalCategory    int64 `json:"total_category"`
	TotalProduct     int64 `json:"total_product"`
	TotalTransaction int64 `json:"total_transaction"`
	TotalUOM         int64 `json:"total_uom"`
	TotalCustomer    int64 `json:"total_customer"`
}

type CustomerTrx struct {
	Customer   string  `json:"customer"`
	TotalTrx   int     `json:"total_trx"`
	AmountTrx  float64 `json:"amount_trx"`
	AverageTrx float64 `json:"average_trx"`
}

type StockAlert struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Qty  int    `json:"qty"`
	UOM  string `json:"uom"`
}

type Top10Trx struct {
	TrxID     string    `json:"trx_id"`
	CreatedAt time.Time `json:"-"`
	TrxDate   string    `json:"trx_date"`
	CreatedBy string    `json:"created_by"`
	Customer  string    `json:"customer"`
	Total     float64   `json:"total"`
}

type Top5Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	TotalQty      int     `json:"total_qty"`
	AverageAmount float64 `json:"average_amount"`
	TotalAmount   float64 `json:"total_amount"`
}
