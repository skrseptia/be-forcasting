package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddUom(obj model.Uom) (model.Uom, error) {
	obj, err := s.rmy.CreateUom(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetUoms() ([]model.Uom, error) {
	list, err := s.rmy.ReadUoms()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetUom(obj model.Uom) (model.Uom, error) {
	obj, err := s.rmy.ReadUom(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditUom(obj model.Uom) (model.Uom, error) {
	obj, err := s.rmy.UpdateUom(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveUom(obj model.Uom) (model.Uom, error) {
	obj, err := s.rmy.DeleteUom(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
