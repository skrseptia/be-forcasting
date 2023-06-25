package model

type Chart struct {
	DailyTrxAmountChart ChartData       `json:"daily_trx_amount_chart"`
	DailyTrxQtyChart    ChartData       `json:"daily_trx_qty_chart"`
	WeeklyTrxChart      WeeklyTrxChart  `json:"weekly_trx_chart"`
	MonthlyTrxChart     MonthlyTrxChart `json:"monthly_trx_chart"`
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

type WeeklyTrxChart struct {
	ChartType string    `json:"chart_type"`
	Labels    []string  `json:"labels"`
	Datasets  []Dataset `json:"datasets"`
}

type MonthlyTrxChart struct {
	ChartType string    `json:"chart_type"`
	Labels    []string  `json:"labels"`
	Datasets  []Dataset `json:"datasets"`
}

type MonthlyRow struct {
	Month  string `json:"month"`
	Year   string `json:"year"`
	Amount int64  `json:"amount"`
}

type MultiTypeRow struct {
	Month    string  `json:"month"`
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

type MultiTypeChart struct {
	Labels   []string           `json:"labels"`
	Datasets []MultiTypeDataset `json:"datasets"`
}

type MultiTypeDataset struct {
	Label string `json:"label"`
	Data  []int  `json:"data"`
}
