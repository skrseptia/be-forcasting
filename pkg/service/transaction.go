package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.CreateTransaction(p)
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

func (s *service) GetTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	obj, err = s.rmy.UpdateTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
