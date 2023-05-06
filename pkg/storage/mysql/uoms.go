package mysql

import (
	"food_delivery_api/pkg/model"
)

func (s *Storage) CreateUOM(obj model.UOM) (model.UOM, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) ReadUOMs() ([]model.UOM, error) {
	var list []model.UOM

	err := s.db.Find(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadUOM(obj model.UOM) (model.UOM, error) {
	err := s.db.First(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) UpdateUOM(obj model.UOM) (model.UOM, error) {
	err := s.db.Model(&obj).Updates(obj).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *Storage) DeleteUOM(obj model.UOM) (model.UOM, error) {
	err := s.db.Delete(&obj, obj.ID).Error
	if err != nil {
		return obj, err
	}

	return obj, nil
}
