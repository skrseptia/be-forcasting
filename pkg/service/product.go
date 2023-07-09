package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"strconv"
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

func (s *service) AddProducts(f *multipart.FileHeader) (model.Upload, error) {
	var res model.Upload
	var list []model.Product
	tableName := "products"
	ttlColumn := 8

	src, err := f.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	xlsx, err := excelize.OpenReader(src)
	if err != nil {
		return res, err
	}

	rows, err := xlsx.GetRows(xlsx.GetSheetName(0))
	if err != nil {
		return res, err
	}

	for i, row := range rows {
		// skip header row
		if i == 0 {
			continue
		}

		if len(row) < ttlColumn {
			return res, errors.New(fmt.Sprintf("total column in %s should be: %d", tableName, ttlColumn))
		}

		qty, _ := strconv.ParseFloat(row[5], 64)
		uomID, _ := strconv.Atoi(row[6])
		price, _ := strconv.ParseFloat(row[7], 64)

		list = append(list, model.Product{
			Code:        row[1],
			Name:        row[2],
			Description: row[3],
			ImageURL:    row[4],
			Qty:         qty,
			UOMID:       uomID,
			Price:       price,
		})
	}

	listTx, err := s.rmy.CreateProducts(list)
	if err != nil {
		return res, err
	}

	res = model.Upload{
		File:         f.Filename,
		Table:        tableName,
		RowsAffected: len(listTx),
	}

	return res, nil
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
