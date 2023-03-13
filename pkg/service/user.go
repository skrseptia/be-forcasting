package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.CreateUser(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetUsers() ([]model.User, error) {
	list, err := s.rmy.ReadUsers()
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.ReadUser(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.UpdateUser(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.DeleteUser(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
