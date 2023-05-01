package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddTransaction(obj model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.CreateTransaction(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetTransactions() ([]model.Transaction, error) {
	list, err := s.rmy.ReadTransactions()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetTransaction(obj model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditTransaction(obj model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.UpdateTransaction(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveTransaction(obj model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.DeleteTransaction(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
