package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddCategory(obj model.Category) (model.Category, error) {
	obj, err := s.rmy.CreateCategory(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetCategories() ([]model.Category, error) {
	list, err := s.rmy.ReadCategories()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetCategory(obj model.Category) (model.Category, error) {
	obj, err := s.rmy.ReadCategory(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditCategory(obj model.Category) (model.Category, error) {
	obj, err := s.rmy.UpdateCategory(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveCategory(obj model.Category) (model.Category, error) {
	obj, err := s.rmy.DeleteCategory(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
