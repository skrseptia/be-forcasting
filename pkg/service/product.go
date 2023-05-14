package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddProduct(p model.ProductRequest) (model.Product, error) {
	pr := model.Product{
		Code:        p.Code,
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		Qty:         p.Qty,
		UOMID:       p.UOMID,
		Price:       p.Price,
	}

	obj, err := s.rmy.CreateProduct(pr)
	if err != nil {
		return obj, err
	}

	uom, err := s.rmy.ReadUOM(model.UOM{Model: model.Model{ID: uint(p.UOMID)}})
	obj.UOM = uom

	return obj, nil
}

func (s *service) GetProducts(qp model.QueryPagination) ([]model.Product, int64, error) {
	list, ttl, err := s.rmy.ReadProducts(qp)
	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *service) GetProduct(p model.Product) (model.Product, error) {
	obj, err := s.rmy.ReadProduct(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditProduct(p model.ProductRequest) (model.Product, error) {
	obj, err := s.rmy.ReadProduct(model.Product{Model: model.Model{ID: p.ID}})
	if err != nil {
		return obj, err
	}

	pr := model.Product{
		Model:       model.Model{ID: p.ID},
		Code:        p.Code,
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		Qty:         p.Qty,
		UOMID:       p.UOMID,
		Price:       p.Price,
	}

	obj, err = s.rmy.UpdateProduct(pr)
	if err != nil {
		return obj, err
	}

	uom, err := s.rmy.ReadUOM(model.UOM{Model: model.Model{ID: uint(obj.UOMID)}})
	obj.UOM = uom

	return obj, nil
}

func (s *service) RemoveProduct(p model.Product) (model.Product, error) {
	obj, err := s.rmy.ReadProduct(model.Product{Model: model.Model{ID: p.ID}})
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteProduct(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
