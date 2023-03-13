package mysql

import (
	"errors"
	"food_delivery_api/pkg/model"
	"gorm.io/gorm"
)

func (s *Storage) CreateMerchant(obj model.Merchant) (model.Merchant, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadMerchants() ([]model.Merchant, error) {
	var list []model.Merchant

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadMerchant(obj model.Merchant) (model.Merchant, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateMerchant(obj model.Merchant) (model.Merchant, error) {
	err := s.db.Model(&obj).Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteMerchant(obj model.Merchant) (model.Merchant, error) {
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
