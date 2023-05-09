package service

import (
	"food_delivery_api/pkg/model"

	"github.com/gin-gonic/gin"
)

func (s *service) AddUOM(p model.UOM) (model.UOM, error) {
	obj, err := s.rmy.CreateUOM(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) GetUOMs(c *gin.Context) ([]model.UOM, int64, error) {
	list, ttl, err := s.rmy.ReadUOMs(c)
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
