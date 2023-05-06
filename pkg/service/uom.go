package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddUOM(obj model.UOM) (model.UOM, error) {
	obj, err := s.rmy.CreateUOM(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetUOMs() ([]model.UOM, error) {
	list, err := s.rmy.ReadUOMs()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetUOM(obj model.UOM) (model.UOM, error) {
	obj, err := s.rmy.ReadUOM(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditUOM(obj model.UOM) (model.UOM, error) {
	obj, err := s.rmy.UpdateUOM(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveUOM(obj model.UOM) (model.UOM, error) {
	obj, err := s.rmy.DeleteUOM(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
