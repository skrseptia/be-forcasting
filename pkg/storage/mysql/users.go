package mysql

import (
	"food_delivery_api/pkg/model"
)

func (s *Storage) CreateUser(obj model.User) (model.User, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadUsers(qp model.QueryPagination) ([]model.User, int64, error) {
	var list []model.User
	var ttl int64

	s.db.Find(&list).Count(&ttl)
	err := s.db.Scopes(Paginate(qp)).Find(&list).Error
	if err != nil {
		return list, ttl, err
	}

	return list, ttl, nil
}

func (s *Storage) ReadUser(obj model.User) (model.User, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadUserByEmailPassword(obj model.User) (model.User, error) {
	err := s.db.Where(&obj).First(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateUser(obj model.User) (model.User, error) {
	err := s.db.Model(&obj).Select("*").Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteUser(obj model.User) (model.User, error) {
	err := s.db.Delete(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
