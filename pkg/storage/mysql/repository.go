package mysql

import (
	"food_delivery_api/cfg"
	"food_delivery_api/pkg/storage/mysql/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(c cfg.MySQL) (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	if err != nil {
		return s, err
	}

	// Migrate the schema
	err = s.db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		return nil, err
	}

	log.Println("MySQL connected")

	return s, nil
}

func (s *Storage) CreateUser(obj model.User) (uint, error) {
	err := s.db.Create(&obj).Error
	if err != nil {
		return 0, err
	}

	return obj.ID, nil
}

// func (s *Storage) UpdateUser(eu editing.User) (uint, error) {

// 	err := s.db.Create(&au).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	return au.ID, nil
// }

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

// func (s *Storage) DeleteUser(id int) (uint, error) {

// 	err := s.db.Ge(&au).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	return au.ID, nil
// }
