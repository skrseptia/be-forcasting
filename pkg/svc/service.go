package svc

import "food_delivery_api/pkg/storage/mysql/model"

type RepositoryMySQL interface {
	CreateUser(model.User) (uint, error)
	ReadUsers() ([]model.User, error)
	ReadUser(model.User) (model.User, error)
}

type Service interface {
	AddUser(model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUser(model.User) (model.User, error)
}

type service struct {
	rmy RepositoryMySQL
}

func NewService(rmy RepositoryMySQL) Service {
	return &service{rmy}
}

func (s *service) AddUser(obj model.User) (model.User, error) {
	var err error
	obj.ID, err = s.rmy.CreateUser(obj)
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
	var err error
	obj, err = s.rmy.ReadUser(obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
