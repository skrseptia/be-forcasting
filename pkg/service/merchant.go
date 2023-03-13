package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddMerchant(obj model.Merchant) (model.Merchant, error) {
	obj, err := s.rmy.CreateMerchant(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetMerchants() ([]model.Merchant, error) {
	list, err := s.rmy.ReadMerchants()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetMerchant(obj model.Merchant) (model.Merchant, error) {
	obj, err := s.rmy.ReadMerchant(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditMerchant(obj model.Merchant) (model.Merchant, error) {
	obj, err := s.rmy.UpdateMerchant(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveMerchant(obj model.Merchant) (model.Merchant, error) {
	obj, err := s.rmy.DeleteMerchant(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
