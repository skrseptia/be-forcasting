package service

import (
	"fmt"
	"food_delivery_api/pkg/model"
	"math"
	"strconv"
)

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
