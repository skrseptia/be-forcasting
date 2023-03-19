package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddStore(obj model.Store) (model.Store, error) {
	obj, err := s.rmy.CreateStore(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetStores() ([]model.Store, error) {
	list, err := s.rmy.ReadStores()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetStore(obj model.Store) (model.Store, error) {
	obj, err := s.rmy.ReadStore(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditStore(obj model.Store) (model.Store, error) {
	obj, err := s.rmy.UpdateStore(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveStore(obj model.Store) (model.Store, error) {
	obj, err := s.rmy.DeleteStore(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
