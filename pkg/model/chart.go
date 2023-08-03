package model

type Chart struct {
	DailyTrxAmountChart   ChartData `json:"daily_trx_amount_chart"`
	DailyTrxQtyChart      ChartData `json:"daily_trx_qty_chart"`
	MonthlyTrxAmountChart ChartData `json:"monthly_trx_amount_chart"`
}

type Dataset struct {
	Label string      `json:"label"`
	UOM   string      `json:"uom"`
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
	UOM       string `json:"uom"`
	Qty       int    `json:"qty"`
}

type ExponentialSmoothingData struct {
	ProductID       int                           `json:"product_id"`
	Name            string                        `json:"name"`
	UOM             string                        `json:"uom"`
	SmoothingFactor float64                       `json:"smoothing_factor"`
	Dataset         []ExponentialSmoothingDataset `json:"dataset"`
}

type MonthlyExponentialSmoothingDataset struct {
	Name        string `json:"name"`
	UOM         string `json:"uom"`
	Period      string `json:"period"`
	Actual      int    `json:"actual"`
	Forecast    int    `json:"forecast"`
	Formulation string `json:"formulation"`
}

type ExponentialSmoothingDataset struct {
	Period      string `json:"period"`
	Actual      int    `json:"actual"`
	Forecast    int    `json:"forecast"`
	Formulation string `json:"formulation"`
}

type ArimaChart struct {
	ChartType         string    `json:"chart_type"`
	Labels            []string  `json:"labels"`
	Datasets          []Dataset `json:"datasets"`
	Actual            []float64 `json:"actual"`
	Predicted         []float64 `json:"predicted"`
	MeanAbsoluteError float64   `json:"mean_absolute_error"`
	MSE               float64   `json:"mse"`
	MAPE              float64   `json:"mape"`
}

type ArimaRow struct {
	StartOfWeek string `json:"start_of_week"`
	EndOfWeek   string `json:"end_of_week"`
	TotalQty    int    `json:"total_qty"`
}

type ExpoChart struct {
	ChartType         string    `json:"chart_type"`
	Labels            []string  `json:"labels"`
	Datasets          []Dataset `json:"datasets"`
	Actual            []float64 `json:"actual"`
	Smoothed          []float64 `json:"smoothed"`
	Prediction        []float64 `json:"prediction"`
	MeanAbsoluteError float64   `json:"mean_absolute_error"`
	MSE               float64   `json:"mse"`
	MAPE              float64   `json:"mape"`
}

type ExpoRow struct {
	StartOfWeek string `json:"start_of_week"`
	EndOfWeek   string `json:"end_of_week"`
	TotalQty    int    `json:"total_qty"`
}
