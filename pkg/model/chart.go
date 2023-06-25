package model

type Chart struct {
	DailyTrxAmountChart   ChartData `json:"daily_trx_amount_chart"`
	DailyTrxQtyChart      ChartData `json:"daily_trx_qty_chart"`
	MonthlyTrxAmountChart ChartData `json:"monthly_trx_amount_chart"`
}

type Dataset struct {
	Label string      `json:"label"`
	Data  interface{} `json:"data"`
}

type ChartData struct {
	ChartType string    `json:"chart_type"`
	Labels    []string  `json:"labels"`
	Datasets  []Dataset `json:"datasets"`
}

type DailyRow struct {
	Product string `json:"product"`
	Qty     int    `json:"qty"`
	Amount  int64  `json:"amount"`
}

type MonthlyRow struct {
	Month    string `json:"month"`
	Category string `json:"category"`
	Product  string `json:"product"`
	Qty      int    `json:"qty"`
	Amount   int64  `json:"amount"`
}

type ExponentialSmoothingChart struct {
	ChartType       string    `json:"chart_type"`
	Labels          []string  `json:"labels"`
	Datasets        []Dataset `json:"datasets"`
	SmoothingFactor float64   `json:"smoothing_factor"`
}

type ExponentialSmoothingRow struct {
	Month     string `json:"month"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Qty       int    `json:"qty"`
}

type ExponentialSmoothingData struct {
	ProductID       int                           `json:"product_id"`
	Name            string                        `json:"name"`
	SmoothingFactor float64                       `json:"smoothing_factor"`
	Dataset         []ExponentialSmoothingDataset `json:"dataset"`
}

type ExponentialSmoothingDataset struct {
	Period   string  `json:"period"`
	Actual   float64 `json:"actual"`
	Forecast float64 `json:"forecast"`
}
