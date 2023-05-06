package mysql

import (
	"errors"
	"food_delivery_api/pkg/model"

	"gorm.io/gorm"
)

func (s *Storage) CreateCategory(obj model.Category) (model.Category, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadCategories() ([]model.Category, error) {
	var list []model.Category

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadCategory(obj model.Category) (model.Category, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateCategory(obj model.Category) (model.Category, error) {
	err := s.db.First(&obj, obj.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return obj, errors.New("data not found")
	}

	err = s.db.Model(&obj).Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteCategory(obj model.Category) (model.Category, error) {
	err := s.db.First(&obj, obj.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return obj, errors.New("data not found")
	}

	err = s.db.Delete(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
