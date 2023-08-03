package service

import (
	"fmt"
	"food_delivery_api/pkg/model"
	"math"
	"strconv"
)

func (s *service) GetReportExpo(qp model.QueryGetExpo) (model.ExpoChart, error) {
	var obj model.ExpoChart

	// get total qty from database for all products
	list, err := s.rmy.ReadReportExpo(qp)
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

	alpha := qp.Alpha
	pl := qp.PredictionLength

	// calculate smoothed data
	smoothedData := expoSmoothing(actual, alpha)

	// calculate prediction based on prediction length
	predictionData := make([]float64, pl)
	predictionData[0] = alpha*actual[len(actual)-1] + (1-alpha)*smoothedData[len(smoothedData)-1]
	for i := 1; i < pl; i++ {
		predictionData[i] = alpha*actual[len(actual)-1] + (1-alpha)*predictionData[i-1]
	}

	// round smoothedData and predictionData to remove decimals
	for i := 0; i < len(smoothedData); i++ {
		smoothedData[i] = math.Round(smoothedData[i])
	}
	for i := 0; i < len(predictionData); i++ {
		predictionData[i] = math.Round(predictionData[i])
	}

	// Menghitung MAE, MSE, dan MAPE untuk hasil peramalan
	mae, mse := countMAEAndMSE(actual, smoothedData)
	mape := countMAPE(actual, smoothedData)

	fmt.Println("\nMean Absolute Error (MAE) untuk hasil peramalan:", mae)
	fmt.Println("Mean Squared Error (MSE) untuk hasil peramalan:", mse)
	fmt.Println("Mean Absolute Percentage Error (MAPE) untuk hasil peramalan:", mape)

	// combine smoothed and prediction
	var combined []int

	// add actual and convert to int
	for _, v := range smoothedData {
		combined = append(combined, int(v))
	}

	// add predictions and convert to int
	for _, v := range predictionData {
		combined = append(combined, int(v))
	}

	var labels []string
	for i := range combined {
		labels = append(labels, fmt.Sprintf("Week-%d", i+1))
	}

	productID, _ := strconv.Atoi(qp.ProductID)
	product, _ := s.rmy.ReadProduct(model.Product{Model: model.Model{ID: uint(productID)}})

	obj.ChartType = "Multi Type Chart"
	obj.Labels = labels
	obj.Datasets = []model.Dataset{
		{Label: product.Name, UOM: product.UOM.Name, Data: actual},
		{Label: fmt.Sprintf("Forecast - %s", product.Name), UOM: product.UOM.Name, Data: combined},
	}
	obj.Actual = actual
	obj.Smoothed = smoothedData
	obj.Prediction = predictionData
	obj.MeanAbsoluteError = mae
	obj.MSE = mse
	obj.MAPE = mape

	return obj, nil
}

func expoSmoothing(data []float64, alpha float64) []float64 {
	// Inisialisasi
	smoothedData := make([]float64, len(data))
	smoothedData[0] = data[0]

	// Peramalan dengan metode Exponential Smoothing
	for i := 1; i < len(data); i++ {
		smoothedData[i] = alpha*data[i] + (1-alpha)*smoothedData[i-1]
	}

	return smoothedData
}

func countMAEAndMSE(actualData, predictedData []float64) (float64, float64) {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	n := float64(len(actualData))
	totalAbsoluteError := 0.0
	totalSquaredError := 0.0

	for i := 0; i < len(actualData); i++ {
		absoluteError := math.Abs(actualData[i] - predictedData[i])
		totalAbsoluteError += absoluteError
		squaredError := math.Pow(absoluteError, 2)
		totalSquaredError += squaredError
	}

	mae := totalAbsoluteError / n
	mse := totalSquaredError / n
	return mae, mse
}

func countMAPE(actualData, predictedData []float64) float64 {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	n := float64(len(actualData))
	totalAbsolutePercentageError := 0.0

	for i := 0; i < len(actualData); i++ {
		percentageError := math.Abs((actualData[i] - predictedData[i]) / actualData[i])
		totalAbsolutePercentageError += percentageError
	}

	mape := (totalAbsolutePercentageError / n) * 100.0
	return mape
}
