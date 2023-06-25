package service

import (
	"food_delivery_api/pkg/model"
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
