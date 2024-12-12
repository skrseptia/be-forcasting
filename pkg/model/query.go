package model

type QueryPagination struct {
	Page        int    `form:"page" binding:"required"`
	PageSize    int    `form:"page_size" binding:"required"`
	Name        string `form:"name"`
	FullName    string `form:"fullname"`
	Email       string `form:"email"`
	Description string `form:"description"`
}

type QueryGetTransactions struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Customer  string `form:"customer"`
	CreatedBy string `form:"created_by"`
	QueryPagination
}

type QueryGetExponentialSmoothing struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	ProductID string `form:"product_id"`
}

type QueryGetArima struct {
	AutoRegressive   int    `form:"p" binding:"required"`
	Differencing     int    `form:"d"`
	MovingAverage    int    `form:"q" binding:"required"`
	SeasonalAR       int    `form:"P"`                     // Parameter Seasonal AR
	SeasonalDiff     int    `form:"D"`                     // Parameter Seasonal Differencing
	SeasonalMA       int    `form:"Q"`                     // Parameter Seasonal MA
	SeasonalPeriod   int    `form:"s" binding:"required"`   // Seasonal Period (s)
	PredictionLength int    `form:"pl" binding:"required"`
	ProductID        string `form:"product_id" binding:"required"`
}

type QueryGetExpo struct {
	Alpha            float64 `form:"alpha" binding:"required"`
	Beta            	float64 `form:"beta" binding:"required"`
	Gamma            float64 `form:"gamma" binding:"required"`
	SeasonLength     int `form:"seasonLength" binding:"required"`
	PredictionLength int     `form:"pl" binding:"required"`
	ProductID        string  `form:"product_id" binding:"required"`
}
