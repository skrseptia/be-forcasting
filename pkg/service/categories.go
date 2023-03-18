package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddCategories(obj model.Categories) (model.Categories, error) {
	obj, err := s.rmy.CreateCategories(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetCategoriess() ([]model.Categories, error) {
	list, err := s.rmy.ReadCategoriess()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetCategories(obj model.Categories) (model.Categories, error) {
	obj, err := s.rmy.ReadCategories(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditCategories(obj model.Categories) (model.Categories, error) {
	obj, err := s.rmy.UpdateCategories(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveCategories(obj model.Categories) (model.Categories, error) {
	obj, err := s.rmy.DeleteCategories(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
