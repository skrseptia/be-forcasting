package mysql

import (
	"food_delivery_api/pkg/model"
)

func (s *Storage) CreateProduct(obj model.Product) (model.Product, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) CreateProducts(list []model.Product) ([]model.Product, error) {
	err := s.db.Create(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadProducts(qp model.QueryPagination) ([]model.Product, int64, error) {
	var list []model.Product
	var ttl int64

	s.db.Find(&list).Count(&ttl)
	err := s.db.Preload("UOM").Scopes(Paginate(qp)).Find(&list).Error
	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *Storage) ReadProduct(obj model.Product) (model.Product, error) {
	err := s.db.Preload("UOM").First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateProduct(obj model.Product) (model.Product, error) {
	err := s.db.Model(&obj).Select("*").Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteProduct(obj model.Product) (model.Product, error) {
	err := s.db.Delete(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
