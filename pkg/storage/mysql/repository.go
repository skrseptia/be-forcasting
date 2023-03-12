package mysql

import (
	"food_delivery_api/config"
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

	log.Println("MySQL connected")

	return s, nil
}
