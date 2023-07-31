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

	// count MAE
	mae := countMAE(actual, smoothedData)

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

func countMAE(actualData, predictedData []float64) float64 {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	totalAbsoluteError := 0.0
	n := float64(len(actualData))

	for i := 0; i < len(actualData); i++ {
		totalAbsoluteError += math.Abs(actualData[i] - predictedData[i])
	}

	mae := totalAbsoluteError / n
	return mae
}
