package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

func (s *service) AddUOM(p model.UOM) (model.UOM, error) {
	obj, err := s.rmy.CreateUOM(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) AddUOMs(f *multipart.FileHeader) (model.Upload, error) {
	var res model.Upload
	var list []model.UOM
	tableName := "uoms"
	ttlColumn := 2

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

		list = append(list, model.UOM{
			Name: row[1],
		})
	}

	listTx, err := s.rmy.CreateUOMs(list)
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

func (s *service) GetUOMs(qp model.QueryPagination) ([]model.UOM, int64, error) {
	list, ttl, err := s.rmy.ReadUOMs(qp)
	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *service) GetUOM(p model.UOM) (model.UOM, error) {
	obj, err := s.rmy.ReadUOM(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditUOM(p model.UOM) (model.UOM, error) {
	obj, err := s.rmy.ReadUOM(p)
	if err != nil {
		return obj, err
	}

	obj, err = s.rmy.UpdateUOM(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveUOM(p model.UOM) (model.UOM, error) {
	obj, err := s.rmy.ReadUOM(p)
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteUOM(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
