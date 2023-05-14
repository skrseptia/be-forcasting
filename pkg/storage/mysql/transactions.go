package mysql

import (
	"fmt"
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

func (s *Storage) ReadTransactions(qp model.QueryGetTransactions) ([]model.Transaction, int64, error) {
	var list []model.Transaction
	var ttl int64
	var err error

	cust := fmt.Sprintf("%%%s%%", qp.Customer)

	err = s.db.Find(&list).Where("customer LIKE ?", cust).Count(&ttl).Error
	err = s.db.Preload("TransactionLines").
		Where("customer LIKE ?", cust).
		Scopes(Paginate(qp.QueryPagination)).
		Find(&list).Error

	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *Storage) ReadTransactionsBetweenDate(qp model.QueryGetTransactions) ([]model.Transaction, int64, error) {
	var list []model.Transaction
	var ttl int64
	var err error

	cust := fmt.Sprintf("%%%s%%", qp.Customer)

	err = s.db.Find(&list).Where("customer LIKE ? AND created_at BETWEEN ? AND ?", cust, qp.StartDate, qp.EndDate).
		Count(&ttl).Error
	err = s.db.Preload("TransactionLines").
		Where("customer LIKE ? AND created_at BETWEEN ? AND ?", cust, qp.StartDate, qp.EndDate).
		Scopes(Paginate(qp.QueryPagination)).
		Find(&list).Error

	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
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
