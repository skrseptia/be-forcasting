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

func (s *Storage) ReadUsers() ([]model.User, error) {
	var list []model.User

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
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
	err := s.db.Model(&obj).Updates(obj).Error
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
