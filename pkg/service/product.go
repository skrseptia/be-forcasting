package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddProduct(obj model.Product) (model.Product, error) {
	obj, err := s.rmy.CreateProduct(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetProducts() ([]model.Product, error) {
	list, err := s.rmy.ReadProducts()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetProduct(obj model.Product) (model.Product, error) {
	obj, err := s.rmy.ReadProduct(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditProduct(obj model.Product) (model.Product, error) {
	obj, err := s.rmy.UpdateProduct(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveProduct(obj model.Product) (model.Product, error) {
	obj, err := s.rmy.DeleteProduct(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
