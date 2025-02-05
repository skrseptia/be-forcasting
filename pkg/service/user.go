package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

func (s *service) AddUser(p model.User) (model.User, error) {
	obj, err := s.rmy.CreateUser(p)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) AddUsers(f *multipart.FileHeader) (model.Upload, error) {
	var res model.Upload
	var list []model.User
	tableName := "users"
	ttlColumn := 7

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

		list = append(list, model.User{
			FullName: row[1],
			Email:    row[2],
			Password: "password",
			ImageURL: row[3],
			Phone:    row[4],
			Address:  row[5],
			Role:     row[6],
		})
	}

	listTx, err := s.rmy.CreateUsers(list)
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

func (s *service) GetUsers(qp model.QueryPagination) ([]model.User, int64, error) {
	list, ttl, err := s.rmy.ReadUsers(qp)
	if err != nil {
		return list, ttl, err
	}

	// hide password
	for i := range list {
		list[i].Password = ""
	}

	return list, ttl, nil
}

func (s *service) GetUser(p model.User) (model.User, error) {
	obj, err := s.rmy.ReadUser(p)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) GetUserPasswordByEmail(email string) (string, error) {
    user, err := s.rmy.ReadUserByEmail(email)
    if err != nil {
        return "", err
    }
    return user.Password, nil
}

func (s *service) GetUserByEmailPassword(p model.User) (model.User, error) {
	obj, err := s.rmy.ReadUserByEmailPassword(p)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) EditUser(p model.User) (model.User, error) {
	obj, err := s.rmy.ReadUser(p)
	if err != nil {
		return obj, err
	}

	obj, err = s.rmy.UpdateUser(p)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) RemoveUser(p model.User) (model.User, error) {
	obj, err := s.rmy.ReadUser(p)
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteUser(p)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}
