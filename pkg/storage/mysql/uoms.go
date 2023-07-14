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

func (s *Storage) CreateUOMs(list []model.UOM) ([]model.UOM, error) {
	err := s.db.Create(&list).Error
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *Storage) ReadUOMs(qp model.QueryPagination) ([]model.UOM, int64, error) {
	var list []model.UOM
	var ttl int64

	if qp.Name != "" {
		s.db.Find(&list).Count(&ttl)
		err := s.db.Where("name like ?", "%"+qp.Name+"%").Scopes(Paginate(qp)).Find(&list).Count(&ttl).Error
		if err != nil {
			return list, ttl, err
		}
	} else {
		s.db.Find(&list).Count(&ttl)
		err := s.db.Scopes(Paginate(qp)).Find(&list).Error
		if err != nil {
			return list, ttl, err
		}
	}

	return list, ttl, nil
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
