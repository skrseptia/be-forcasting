package mysql

import (
	"errors"
	"food_delivery_api/cfg"
	"food_delivery_api/pkg/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryMySQL interface {
	// Users
	CreateUser(model.User) (model.User, error)
	ReadUsers() ([]model.User, error)
	ReadUser(model.User) (model.User, error)
	ReadUserByEmailPassword(model.User) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)

	// Products
	CreateProduct(model.Product) (model.Product, error)
	ReadProducts() ([]model.Product, error)
	ReadProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)

	// Products
	CreateUom(model.Uom) (model.Uom, error)
	ReadUoms() ([]model.Uom, error)
	ReadUom(model.Uom) (model.Uom, error)
	UpdateUom(model.Uom) (model.Uom, error)
	DeleteUom(model.Uom) (model.Uom, error)

	// Categories
	CreateCategories(model.Categories) (model.Categories, error)
	ReadCategoriess() ([]model.Categories, error)
	ReadCategories(model.Categories) (model.Categories, error)
	UpdateCategories(model.Categories) (model.Categories, error)
	DeleteCategories(model.Categories) (model.Categories, error)

	// Products
	CreateTransaction(model.Transaction) (model.Transaction, error)
	ReadTransactions() ([]model.Transaction, error)
	ReadTransaction(model.Transaction) (model.Transaction, error)
	UpdateTransaction(model.Transaction) (model.Transaction, error)
	DeleteTransaction(model.Transaction) (model.Transaction, error)
}

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

	if err = autoMigrateDB(s); err != nil {
		return nil, err
	}

	if err = seedDB(s); err != nil {
		return nil, err
	}

	log.Println("MySQL connected")

	return s, nil
}

func autoMigrateDB(s *Storage) error {
	// Migrate the schema
	err := s.db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Uom{},
		&model.Categories{},
		&model.Transaction{},
	)

	return err
}

func seedDB(s *Storage) error {
	var user model.User
	err := s.db.First(&user, 1).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		user, err = s.CreateUser(model.User{
			Model:    model.Model{},
			FullName: "Super Admin",
			Email:    "admin@mail.com",
			Password: "password",
			ImageURL: "www.image.com",
			Phone:    "+6281234567890",
			Address:  "Jakarta",
			UserType: "Admin",
		})

		if err != nil {
			return err
		}

		log.Println("Super Admin Created", user)
	}

	return err
}
