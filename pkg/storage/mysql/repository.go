package mysql

import (
	"food_delivery_api/config"
	"food_delivery_api/pkg/adding"
	"food_delivery_api/pkg/listing"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(cfg config.MySQL) (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return s, err
	}

	// Migrate the schema
	s.db.AutoMigrate(
		&adding.User{},
	)

	log.Println("MySQL connected")

	return s, nil
}

func (s *Storage) CreateUser(au adding.User) (uint, error) {

	err := s.db.Create(&au).Error
	if err != nil {
		return 0, err
	}

	return au.ID, nil
}

// func (s *Storage) UpdateUser(eu editing.User) (uint, error) {

// 	err := s.db.Create(&au).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	return au.ID, nil
// }

func (s *Storage) ReadUser(lu listing.User) (listing.User, error) {

	err := s.db.First(&lu, lu.ID).Error
	if err != nil {
		return lu, err
	}

	return lu, nil
}

// func (s *Storage) DeleteUser(id int) (uint, error) {

// 	err := s.db.Ge(&au).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	return au.ID, nil
// }

