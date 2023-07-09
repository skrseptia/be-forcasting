package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

func (s *service) AddCategory(p model.Category) (model.Category, error) {
	obj, err := s.rmy.CreateCategory(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) AddCategories(f *multipart.FileHeader) (model.Upload, error) {
	var res model.Upload
	var list []model.Category
	tableName := "categories"
	ttlColumn := 3

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

		list = append(list, model.Category{
			Code: row[1],
			Name: row[2],
		})
	}

	listTx, err := s.rmy.CreateCategories(list)
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

func (s *service) GetCategories(qp model.QueryPagination) ([]model.Category, int64, error) {
	list, ttl, err := s.rmy.ReadCategories(qp)
	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *service) GetCategory(p model.Category) (model.Category, error) {
	obj, err := s.rmy.ReadCategory(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditCategory(p model.Category) (model.Category, error) {
	obj, err := s.rmy.ReadCategory(p)
	if err != nil {
		return obj, err
	}

	obj, err = s.rmy.UpdateCategory(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveCategory(p model.Category) (model.Category, error) {
	obj, err := s.rmy.ReadCategory(p)
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteCategory(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
