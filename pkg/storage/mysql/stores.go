package mysql

import (
	"errors"
	"food_delivery_api/pkg/model"
	"gorm.io/gorm"
)

func (s *Storage) CreateStore(obj model.Store) (model.Store, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadStores() ([]model.Store, error) {
	var list []model.Store

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadStore(obj model.Store) (model.Store, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateStore(obj model.Store) (model.Store, error) {
	err := s.db.Model(&obj).Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteStore(obj model.Store) (model.Store, error) {
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
