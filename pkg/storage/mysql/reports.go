package mysql

import (
	"food_delivery_api/pkg/model"
)

func (s *Storage) ReadReportDashboard() (model.Dashboard, error) {
	var obj model.Dashboard
	var err error
	var ttlCtg, ttlPrd, ttlTrx, ttlUom, ttlCst int64

	err = s.db.Model(&model.Category{}).Count(&ttlCtg).Error
	err = s.db.Model(&model.Product{}).Count(&ttlPrd).Error
	err = s.db.Model(&model.Transaction{}).Count(&ttlTrx).Error
	err = s.db.Model(&model.UOM{}).Count(&ttlUom).Error
	err = s.db.Model(&model.Transaction{}).Distinct("customer").Count(&ttlCst).Error

	if err != nil {
		return obj, err
	}

	obj = model.Dashboard{
		TotalCategory:    ttlCtg,
		TotalProduct:     ttlPrd,
		TotalTransaction: ttlTrx,
		TotalUOM:         ttlUom,
		TotalCustomer:    ttlCst,
	}

	return obj, nil
}
