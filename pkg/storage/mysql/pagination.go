package mysql

import (
	"food_delivery_api/pkg/model"

	"gorm.io/gorm"
)

func Paginate(qp model.QueryPagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := qp.Page
		pageSize := qp.PageSize

		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
