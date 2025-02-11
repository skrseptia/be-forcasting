package service

import (
	"fmt"
	"food_delivery_api/pkg/model"
	"math"
	"math/rand"
	"strconv"
	"time"
)
func generateDateLabels(startDate time.Time, weeks int) []string {
	var labels []string
	for i := 0; i < weeks; i++ {
			// Hitung tanggal awal dan akhir minggu
			weekStart := startDate.AddDate(0, 0, 7*i)
			weekEnd := weekStart.AddDate(0, 0, 6)

			// Format tanggal menjadi "1 Jan-7 Jan"
			label := fmt.Sprintf("%d %s %d - %d %s %d",
				weekStart.Day(), weekStart.Month().String()[:3], weekStart.Year(),
				weekEnd.Day(), weekEnd.Month().String()[:3], weekEnd.Year())

			labels = append(labels, label)
	}
	return labels
}


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
	beta := qp.Beta
	gamma := qp.Gamma
	pl := qp.PredictionLength
	seasonLength := qp.SeasonLength

	// calculate smoothed data
	// smoothedData := expoSmoothing(actual, alpha)
		// calculate smoothed data using triple exponential smoothing
	level, trend, seasonal, smoothedData := tripleExponentialSmoothing(actual, alpha, beta, gamma, seasonLength)

	// #OLD# //
	// // calculate prediction based on prediction length
	// predictionData := make([]float64, pl)
	// predictionData[0] = alpha*actual[len(actual)-1] + (1-alpha)*smoothedData[len(smoothedData)-1]
	// for i := 1; i < pl; i++ {
	// 	predictionData[i] = alpha*actual[len(actual)-1] + (1-alpha)*predictionData[i-1]
	// }

	// // round smoothedData and predictionData to remove decimals
	// for i := 0; i < len(smoothedData); i++ {
	// 	smoothedData[i] = math.Round(smoothedData[i])
	// }
	// for i := 0; i < len(predictionData); i++ {
	// 	predictionData[i] = math.Round(predictionData[i])
	// }

	// // Menghitung MAE, MSE, dan MAPE untuk hasil peramalan
	// mae, mse := countMAEAndMSE(actual, smoothedData)
	// mape := countMAPE(actual, smoothedData)

	// fmt.Println("\nMean Absolute Error (MAE) untuk hasil peramalan:", mae)
	// fmt.Println("Mean Squared Error (MSE) untuk hasil peramalan:", mse)
	// fmt.Println("Mean Absolute Percentage Error (MAPE) untuk hasil peramalan:", mape)

	// // combine smoothed and prediction
	// var combined []int

	// // add actual and convert to int
	// for _, v := range smoothedData {
	// 	combined = append(combined, int(v))
	// }

	// // add predictions and convert to int
	// for _, v := range predictionData {
	// 	combined = append(combined, int(v))
	// }

	// var labels []string
	// for i := range combined {
	// 	labels = append(labels, fmt.Sprintf("Week-%d", i+1))
	// }

	// productID, _ := strconv.Atoi(qp.ProductID)
	// product, _ := s.rmy.ReadProduct(model.Product{Model: model.Model{ID: uint(productID)}})

	// obj.ChartType = "Multi Type Chart"
	// obj.Labels = labels
	// obj.Datasets = []model.Dataset{
	// 	{Label: product.Name, UOM: product.UOM.Name, Data: actual},
	// 	{Label: fmt.Sprintf("Forecast - %s", product.Name), UOM: product.UOM.Name, Data: combined},
	// }
	// obj.Actual = actual
	// obj.Smoothed = smoothedData
	// obj.Prediction = predictionData
	// obj.MeanAbsoluteError = mae
	// obj.MSE = mse
	// obj.MAPE = mape

	// return obj, nil
	// #OLD# //
	rand.Seed(time.Now().UnixNano())
	factor := 0.4 + rand.Float64()*(0.9-0.4)

	fmt.Println("\nfactor", factor)

	historicalAverage := 0.0
	for _, value := range actual {
		historicalAverage += value
	}
	historicalAverage /= float64(len(actual))
	
	// calculate predictions based on prediction length
	predictionData := make([]float64, pl)
	for i := 0; i < pl; i++ {
		idx := len(actual) + i
		basePrediction := level[len(level)-1] + float64(i+1)*trend[len(trend)-1] + seasonal[idx%seasonLength]+ safeLast(actual, i)

		// Cegah prediksi negatif
		if basePrediction < 0 {
			basePrediction = 0
		}
	
		// Cegah seasonal negatif
		if seasonal[idx%seasonLength] < 0 {
			seasonal[idx%seasonLength] = 0
		}
	
		// Cegah prediksi menurun terlalu tajam
		if i > 0 && basePrediction < predictionData[i-1]* 0.7 { // Tidak boleh turun lebih dari 10% dibandingkan sebelumnya
			basePrediction = predictionData[i-1] * 0.88
		}
	
		// Jaga prediksi tidak kurang dari rata-rata historis
		if basePrediction < historicalAverage* 0.5 { // Minimum 70% dari rata-rata historis
			basePrediction = historicalAverage * 0.77
		}
	
		predictionData[i] = math.Round(basePrediction)

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

	startDate := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC) 
	labels := generateDateLabels(startDate, len(combined))

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

// func expoSmoothing(data []float64, alpha float64) []float64 {
func tripleExponentialSmoothing(data []float64, alpha, beta, gamma float64, seasonLength int) ([]float64, []float64, []float64, []float64) {
	//  // #old#  //
	// // Inisialisasi
	// smoothedData := make([]float64, len(data))
	// smoothedData[0] = data[0]

	// // Peramalan dengan metode Exponential Smoothing
	// for i := 1; i < len(data); i++ {
	// 	smoothedData[i] = alpha*data[i] + (1-alpha)*smoothedData[i-1]
	// }

	// return smoothedData
	// // #old#  //

	if len(data) < seasonLength {
		panic("Data length must be at least as long as the season length")
	}

	level := make([]float64, len(data))
	trend := make([]float64, len(data))
	seasonal := make([]float64, len(data)+seasonLength)
	smoothedData := make([]float64, len(data))

	// Initialize level, trend, and seasonal components
	level[0] = data[0]
	trend[0] = data[1] - data[0]
	for i := 0; i < seasonLength; i++ {
		seasonal[i] = data[i] - level[0]
	}

	// Apply triple exponential smoothing
	for i := 1; i < len(data); i++ {
		level[i] = alpha*(data[i]-seasonal[i%seasonLength]) + (1-alpha)*(level[i-1]+trend[i-1])
		trend[i] = beta*(level[i]-level[i-1]) + (1-beta)*trend[i-1]
		seasonal[i%seasonLength] = gamma*(data[i]-level[i]) + (1-gamma)*seasonal[i%seasonLength]
		smoothedData[i] = level[i] + trend[i] + seasonal[i%seasonLength]
	}

	return level, trend, seasonal[:seasonLength], smoothedData
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

func safeLast(data []float64, index int) float64 {
	if index == 0 {
		return 0
	}
	if len(data) > 0 {
		return data[len(data)-1] / 5
	}
	return 0
}

