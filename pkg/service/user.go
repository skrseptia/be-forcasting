package service

import (
	"food_delivery_api/pkg/model"
)

func (s *service) AddUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.CreateUser(obj)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) GetUsers() ([]model.User, error) {
	list, err := s.rmy.ReadUsers()
	if err != nil {
		return list, err
	}

	// hide password
	for i := range list {
		list[i].Password = ""
	}

	return list, nil
}

func (s *service) GetUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.ReadUser(obj)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) GetUserByEmailPassword(obj model.User) (model.User, error) {
	obj, err := s.rmy.ReadUserByEmailPassword(obj)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) EditUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.UpdateUser(obj)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}

func (s *service) RemoveUser(obj model.User) (model.User, error) {
	obj, err := s.rmy.DeleteUser(obj)
	if err != nil {
		return obj, err
	}

	// hide password
	obj.Password = ""

	return obj, nil
}
