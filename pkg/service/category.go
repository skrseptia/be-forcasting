package service

import (
	"food_delivery_api/pkg/model"

	"github.com/gin-gonic/gin"
)

func (s *service) AddCategory(p model.Category) (model.Category, error) {
	obj, err := s.rmy.CreateCategory(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetCategories(c *gin.Context) ([]model.Category, int64, error) {
	list, ttl, err := s.rmy.ReadCategories(c)
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
