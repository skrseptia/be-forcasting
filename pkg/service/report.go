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
