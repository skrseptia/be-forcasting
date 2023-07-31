package service

import (
	"fmt"
	"food_delivery_api/pkg/model"
	"math"
	"strconv"
	"strings"
	"time"
)

func (s *service) GetReportDashboard() (model.Dashboard, error) {
	obj, err := s.rmy.ReadReportDashboard()
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetReportChart() (model.Chart, error) {
	obj, err := s.rmy.ReadReportChart()
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetReportExponentialSmoothing(qp model.QueryGetExponentialSmoothing) (model.ExponentialSmoothingChart, error) {
	var obj model.ExponentialSmoothingChart

	// get total qty from database for all products
	list, err := s.rmy.ReadReportExponentialSmoothing(qp)
	if err != nil {
		return obj, err
	}

	// filter data by query params input
	ids := convertQueryParamID(qp.ProductID)

	// forming data exponential smoothing
	var esds []model.ExponentialSmoothingData
	var labels []string
	var smoothingFactor float64

	for _, id := range ids {
		// forming selected data by product id
		var rows []model.ExponentialSmoothingRow
		for _, v := range list {
			if v.ProductID == id {
				rows = append(rows, v)
			}
		}

		// count alpha
		smoothingFactor = 2 / (float64(len(rows)) + 1)

		// forming exponential smoothing data
		var productID, fCount int
		var name, month, uom, formulation string
		var prevQty, prevForecast float64
		var dataset []model.ExponentialSmoothingDataset

		// handle if no transaction found
		if len(rows) == 0 {
			continue
		}

		for _, row := range rows {
			productID = row.ProductID
			name = row.Name
			month = row.Month
			uom = row.UOM
			prevQty = float64(row.Qty)

			m := convertMonth(row.Month)
			if !contains(labels, m) {
				labels = append(labels, m)
			}

			// set prediction equal with actual on the first data
			if prevForecast == 0 {
				prevForecast = prevQty
				fCount += 1
				formulation = fmt.Sprintf("F%d: %v = %v", fCount, prevQty, prevQty)
			}

			dataset = append(dataset, model.ExponentialSmoothingDataset{
				Period:      row.Month,
				Actual:      int(math.Round(prevQty)),
				Forecast:    int(math.Round(prevForecast)),
				Formulation: formulation,
			})

			fCount += 1
			prevForecast, formulation = exponentialSmoothing(fCount, uom, prevQty, prevForecast, smoothingFactor)
		}

		// add prediction into dataset
		date, err := time.Parse("2006-01", month)
		if err != nil {
			return obj, err
		}
		oneMonthLater := date.AddDate(0, 1, 0)
		month = oneMonthLater.Format("2006-01")

		dataset = append(dataset, model.ExponentialSmoothingDataset{
			Period:      month,
			Actual:      0,
			Forecast:    int(math.Round(prevForecast)),
			Formulation: formulation,
		})

		prevForecast, formulation = exponentialSmoothing(fCount, uom, prevQty, prevForecast, smoothingFactor)

		m := convertMonth(month)
		if !contains(labels, m) {
			labels = append(labels, m)
		}

		esds = append(esds, model.ExponentialSmoothingData{
			ProductID:       productID,
			Name:            name,
			UOM:             uom,
			SmoothingFactor: smoothingFactor,
			Dataset:         dataset,
		})
	}

	var mescd []model.Dataset
	for _, v := range esds {
		var actuals []float64
		var forecast []float64
		var formulations []string
		for _, data := range v.Dataset {
			actuals = append(actuals, float64(data.Actual))
			forecast = append(forecast, float64(data.Forecast))
			formulations = append(formulations, data.Formulation)
		}

		mescd = append(mescd, model.Dataset{
			Label: v.Name,
			UOM:   v.UOM,
			Data:  actuals,
		})

		mescd = append(mescd, model.Dataset{
			Label: fmt.Sprintf("Forecast - %s", v.Name),
			UOM:   v.UOM,
			Data:  forecast,
		})

		mescd = append(mescd, model.Dataset{
			Label: fmt.Sprintf("Formulation - %s", v.Name),
			UOM:   v.UOM,
			Data:  formulations,
		})
	}

	obj = model.ExponentialSmoothingChart{
		ChartType:       "Multi Type Chart",
		Labels:          labels,
		Datasets:        mescd,
		SmoothingFactor: smoothingFactor,
	}

	return obj, nil
}

func (s *service) GetReportMonthlyExponentialSmoothing(qp model.QueryGetExponentialSmoothing) ([]model.MonthlyExponentialSmoothingDataset, error) {
	var obj []model.MonthlyExponentialSmoothingDataset

	// get total qty from database for all products
	list, err := s.rmy.ReadReportExponentialSmoothing(qp)
	if err != nil {
		return obj, err
	}

	// filter data by query params input
	ids := convertQueryParamID(qp.ProductID)

	// forming data exponential smoothing
	var smoothingFactor float64
	var name string

	for _, id := range ids {
		// forming selected data by product id
		var rows []model.ExponentialSmoothingRow
		for _, v := range list {
			if v.ProductID == id {
				rows = append(rows, v)
			}
		}

		// count alpha
		smoothingFactor = 2 / (float64(len(rows)) + 1)

		// forming exponential smoothing data
		var fCount int
		var month, uom, formulation string
		var prevQty, prevForecast float64

		// handle if no transaction found
		if len(rows) == 0 {
			continue
		}

		for _, row := range rows {
			name = row.Name
			uom = row.UOM
			month = row.Month
			prevQty = float64(row.Qty)

			m := convertMonth(row.Month)

			// set prediction equal with actual on the first data
			if prevForecast == 0 {
				prevForecast = prevQty
				fCount += 1
				formulation = fmt.Sprintf("F%d: %v = %v", fCount, prevQty, prevQty)
			}

			obj = append(obj, model.MonthlyExponentialSmoothingDataset{
				Name:        name,
				UOM:         uom,
				Period:      m,
				Actual:      int(math.Round(prevQty)),
				Forecast:    int(math.Round(prevForecast)),
				Formulation: formulation,
			})

			fCount += 1
			prevForecast, formulation = exponentialSmoothing(fCount, uom, prevQty, prevForecast, smoothingFactor)
		}

		// add prediction into dataset
		date, err := time.Parse("2006-01", month)
		if err != nil {
			return obj, err
		}
		oneMonthLater := date.AddDate(0, 1, 0)
		month = oneMonthLater.Format("2006-01")
		m := convertMonth(month)

		obj = append(obj, model.MonthlyExponentialSmoothingDataset{
			Name:        name,
			UOM:         uom,
			Period:      m,
			Actual:      0,
			Forecast:    int(math.Round(prevForecast)),
			Formulation: formulation,
		})

		prevForecast, formulation = exponentialSmoothing(fCount, uom, prevQty, prevForecast, smoothingFactor)
	}

	return obj, nil
}

func convertQueryParamID(s string) []int {
	values := strings.Split(s, ",")
	list := make([]int, len(values))

	for i, value := range values {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			// Handle error, such as skipping the invalid value
			continue
		}
		list[i] = intValue
	}

	return list
}

func convertMonth(dateString string) string {
	date, _ := time.Parse("2006-01", dateString)
	monthName := date.Month().String()

	return monthName
}

func contains(list []string, value string) bool {
	for _, str := range list {
		if str == value {
			return true
		}
	}
	return false
}

func exponentialSmoothing(fCount int, uom string, prevQty float64, prevForecast float64, smoothingFactor float64) (float64, string) {
	forecast := prevForecast + smoothingFactor*(prevQty-prevForecast)
	sf := fmt.Sprintf("%.2f", smoothingFactor)

	// format two decimal places
	formatted := math.Round(forecast*100) / 100
	formulation := fmt.Sprintf("F%d: %d + %s * (%d - %d) = %d %s", fCount, int(math.Round(prevForecast)), sf, int(math.Round(prevQty)), int(math.Round(prevForecast)), int(math.Round(formatted)), uom)

	return formatted, formulation
}

func (s *service) GetReportArima(qp model.QueryGetArima) (model.ArimaChart, error) {
	var obj model.ArimaChart

	// get total qty from database for all products
	list, err := s.rmy.ReadReportArima(qp)
	if err != nil {
		return obj, err
	}

	// handle if return empty data from database
	if list == nil {
		return obj, nil
	}

	var actual []float64
	for _, v := range list {
		actual = append(actual, float64(v.TotalQty))
	}

	p := qp.AutoRegressive
	d := qp.Differencing
	q := qp.MovingAverage
	pl := qp.PredictionLength

	// create prediction
	predictions := predictARIMA(actual, p, d, q, pl)

	// calculate MAE
	mae := calculateMAE(actual[len(actual)-pl:], predictions)

	// combine actual and prediction
	var combined []int

	// add actual and convert to int
	for _, v := range actual {
		combined = append(combined, int(v))
	}

	// add predictions and convert to int
	for _, v := range predictions {
		combined = append(combined, int(v))
	}

	var labels []string
	for i := range combined {
		labels = append(labels, fmt.Sprintf("Week-%d", i+1))
	}

	productID, _ := strconv.Atoi(qp.ProductID)
	product, _ := s.rmy.ReadProduct(model.Product{Model: model.Model{ID: uint(productID)}})

	obj.ChartType = "Line Chart"
	obj.Labels = labels
	obj.Datasets = []model.Dataset{{Label: product.Name, UOM: product.UOM.Name, Data: combined}}
	obj.Actual = actual
	obj.Predicted = predictions
	obj.MeanAbsoluteError = mae

	return obj, nil
}

// Fungsi untuk melakukan prediksi menggunakan model ARIMA
func predictARIMA(data []float64, p, d, q, predictionLength int) []float64 {
	// Lakukan differencing jika d > 0
	for i := 0; i < d; i++ {
		diffData := make([]float64, len(data)-1)
		for j := 1; j < len(data); j++ {
			diffData[j-1] = data[j] - data[j-1]
		}
		data = diffData
	}

	arCoefficients := calculateARCoefficients(data, p)

	// Lakukan residual calculation
	residual := calculateResidual(data, arCoefficients)

	// Lakukan moving average jika q > 0
	if q > 0 {
		maCoefficients := calculateMACoefficients(residual, q)

		for i := 0; i < predictionLength; i++ {
			prediction := float64(0)
			for j := 1; j <= q; j++ {
				if len(residual) >= j {
					prediction += maCoefficients[j-1] * residual[len(residual)-j]
				}
			}
			residual = append(residual, prediction)
		}
	}

	// Lakukan inverse differencing jika d > 0
	if d > 0 {
		for i := 0; i < d; i++ {
			inverseDiffData := make([]float64, len(residual))
			inverseDiffData[0] = data[len(data)-1]
			for j := 1; j < len(residual); j++ {
				inverseDiffData[j] = inverseDiffData[j-1] + residual[j-1]
			}
			residual = inverseDiffData
		}
	}

	// Lakukan inverse AR jika p > 0
	if p > 0 {
		residual = inverseARIMA(residual, p, d, arCoefficients)
	}

	return residual[len(residual)-predictionLength:]
}

// Fungsi untuk menghitung koefisien autoregressive (AR)
func calculateARCoefficients(data []float64, p int) []float64 {
	coefficients := make([]float64, p)

	for i := 1; i <= p; i++ {
		coefficients[i-1] = autocorrelation(data, i)
	}

	return coefficients
}

// Fungsi untuk menghitung autokorelasi
func autocorrelation(data []float64, lag int) float64 {
	n := len(data)
	meanData := mean(data)
	varianceData := variance(data)

	var numerator, denominator float64
	for i := 0; i < n-lag; i++ {
		numerator += (data[i] - meanData) * (data[i+lag] - meanData)
	}
	denominator = float64(n-lag) * varianceData

	return numerator / denominator
}

// Fungsi untuk menghitung rata-rata dari sebuah slice float64
func mean(data []float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

// Fungsi untuk menghitung varians dari sebuah slice float64
func variance(data []float64) float64 {
	meanValue := mean(data)
	sumSquaredDiff := 0.0
	for _, val := range data {
		diff := val - meanValue
		sumSquaredDiff += diff * diff
	}
	return sumSquaredDiff / float64(len(data)-1)
}

// Fungsi untuk menghitung residual dari model ARIMA
func calculateResidual(data []float64, arCoefficients []float64) []float64 {
	n := len(data)
	p := len(arCoefficients)
	residual := make([]float64, n-p)

	copy(residual, data[:p])

	for i := p; i < n; i++ {
		prediction := float64(0)
		for j := 0; j < p; j++ {
			prediction += arCoefficients[j] * data[i-j-1]
		}
		residual = append(residual, data[i]-prediction)
	}

	return residual
}

// Fungsi untuk menghitung moving average (MA)
func calculateMACoefficients(residual []float64, q int) []float64 {
	coefficients := make([]float64, q)

	for i := 1; i <= q; i++ {
		coefficients[i-1] = autocorrelation(residual, i)
	}

	return coefficients
}

// Fungsi untuk melakukan inverse ARIMA
func inverseARIMA(residual []float64, p, d int, arCoefficients []float64) []float64 {
	n := len(residual)
	data := make([]float64, n)

	copy(data, residual)

	// Lakukan inverse differencing jika d > 0
	if d > 0 {
		for i := n - 1; i >= d; i-- {
			data[i] = data[i] + data[i-d]
		}
	}

	// Lakukan inverse AR jika p > 0
	if p > 0 {
		for i := n - 1; i >= p; i-- {
			prediction := float64(0)
			for j := 0; j < p; j++ {
				prediction += arCoefficients[j] * data[i-j-1]
			}
			data[i-p] = data[i-p] + prediction
		}
	}

	return data
}

func calculateMAE(actual, predicted []float64) float64 {
	if len(actual) != len(predicted) {
		panic("Length of actual and predicted slices should be the same.")
	}

	sum := 0.0
	for i := 0; i < len(actual); i++ {
		sum += math.Abs(actual[i] - predicted[i])
	}

	return sum / float64(len(actual))
}
