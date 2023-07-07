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
		var productID int
		var name, month, formulation string
		var prevQty, prevForecast float64
		var dataset []model.ExponentialSmoothingDataset

		for _, row := range rows {
			productID = row.ProductID
			name = row.Name
			month = row.Month
			prevQty = float64(row.Qty)

			m := convertMonth(row.Month)
			if !contains(labels, m) {
				labels = append(labels, m)
			}

			// set prediction equal with actual on the first data
			if prevForecast == 0 {
				prevForecast = prevQty
				formulation = fmt.Sprintf("formulation: %v = %v", prevQty, prevQty)
			} else {
				prevForecast, formulation = exponentialSmoothing(prevQty, prevForecast, smoothingFactor)
			}

			dataset = append(dataset, model.ExponentialSmoothingDataset{
				Period:      row.Month,
				Actual:      prevQty,
				Forecast:    prevForecast,
				Formulation: formulation,
			})
		}

		// add prediction into dataset
		date, err := time.Parse("2006-01", month)
		if err != nil {
			return obj, err
		}
		oneMonthLater := date.AddDate(0, 1, 0)
		month = oneMonthLater.Format("2006-01")

		prevForecast, formulation = exponentialSmoothing(prevQty, prevForecast, smoothingFactor)
		dataset = append(dataset, model.ExponentialSmoothingDataset{
			Period:      month,
			Actual:      0,
			Forecast:    prevForecast,
			Formulation: formulation,
		})

		m := convertMonth(month)
		if !contains(labels, m) {
			labels = append(labels, m)
		}

		esds = append(esds, model.ExponentialSmoothingData{
			ProductID:       productID,
			Name:            name,
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
			actuals = append(actuals, data.Actual)
			forecast = append(forecast, data.Forecast)
			formulations = append(formulations, data.Formulation)
		}

		mescd = append(mescd, model.Dataset{
			Label: v.Name,
			Data:  actuals,
		})

		mescd = append(mescd, model.Dataset{
			Label: fmt.Sprintf("Forecast - %s", v.Name),
			Data:  forecast,
		})

		mescd = append(mescd, model.Dataset{
			Label: fmt.Sprintf("Formulation - %s", v.Name),
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

func exponentialSmoothing(prevQty float64, prevForecast float64, smoothingFactor float64) (float64, string) {
	forecast := prevForecast + smoothingFactor*(prevQty-prevForecast)

	// format two decimal places
	formatted := math.Round(forecast*100) / 100
	formulation := fmt.Sprintf("formulation: %v + %v * (%v - %v) = %v", prevForecast, smoothingFactor, prevQty, prevForecast, formatted)

	return formatted, formulation
}
