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

func (s *Storage) ReadProducts() ([]model.Product, error) {
	var list []model.Product

	err := s.db.Preload("UOM").Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
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
