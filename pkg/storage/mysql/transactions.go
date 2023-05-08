package mysql

import (
	"food_delivery_api/pkg/model"
)

func (s *Storage) CreateTransaction(obj model.Transaction) (model.Transaction, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) CreateTransactions(list []model.Transaction) ([]model.Transaction, error) {
	err := s.db.Create(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadTransactions() ([]model.Transaction, error) {
	var list []model.Transaction

	err := s.db.Preload("TransactionLines").Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadTransaction(obj model.Transaction) (model.Transaction, error) {
	err := s.db.Preload("TransactionLines").First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateTransaction(obj model.Transaction) (model.Transaction, error) {
	err := s.db.Model(&obj).Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteTransaction(obj model.Transaction) (model.Transaction, error) {
	err := s.db.Delete(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
