package mysql

import (
	"errors"
	"food_delivery_api/pkg/model"

	"gorm.io/gorm"
)

func (s *Storage) CreateTransactionLine(obj model.TransactionLine) (model.TransactionLine, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadTransactionLines() ([]model.TransactionLine, error) {
	var list []model.TransactionLine

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadTransactionLine(obj model.TransactionLine) (model.TransactionLine, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateTransactionLine(obj model.TransactionLine) (model.TransactionLine, error) {
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

func (s *Storage) DeleteTransactionLine(obj model.TransactionLine) (model.TransactionLine, error) {
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
