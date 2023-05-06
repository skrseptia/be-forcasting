package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddProduct(body model.ProductRequest) (model.Product, error) {
	obj := model.Product{
		Code:        body.Code,
		Name:        body.Name,
		Description: body.Description,
		ImageURL:    body.ImageURL,
		Qty:         body.Qty,
		UOMID:       body.UOMID,
		Price:       body.Price,
	}

	obj, err := s.rmy.CreateProduct(obj)
	if err != nil {
		return obj, err
	}

	uom, err := s.rmy.ReadUOM(model.UOM{Model: model.Model{ID: uint(obj.UOMID)}})
	obj.UOM = uom

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

func (s *service) EditProduct(body model.ProductRequest) (model.Product, error) {
	obj := model.Product{
		Model:       model.Model{ID: body.ID},
		Code:        body.Code,
		Name:        body.Name,
		Description: body.Description,
		ImageURL:    body.ImageURL,
		Qty:         body.Qty,
		UOMID:       body.UOMID,
		Price:       body.Price,
	}

	obj, err := s.rmy.UpdateProduct(obj)
	if err != nil {
		return obj, err
	}

	uom, err := s.rmy.ReadUOM(model.UOM{Model: model.Model{ID: uint(obj.UOMID)}})
	obj.UOM = uom

	return obj, nil
}

func (s *service) RemoveProduct(obj model.Product) (model.Product, error) {
	obj, err := s.rmy.DeleteProduct(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
