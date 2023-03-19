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

	// Merchants
	CreateMerchant(model.Merchant) (model.Merchant, error)
	ReadMerchants() ([]model.Merchant, error)
	ReadMerchant(model.Merchant) (model.Merchant, error)
	UpdateMerchant(model.Merchant) (model.Merchant, error)
	DeleteMerchant(model.Merchant) (model.Merchant, error)

	// Products
	CreateProduct(model.Product) (model.Product, error)
	ReadProducts() ([]model.Product, error)
	ReadProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)

	// Categories
	CreateCategories(model.Categories) (model.Categories, error)
	ReadCategoriess() ([]model.Categories, error)
	ReadCategories(model.Categories) (model.Categories, error)
	UpdateCategories(model.Categories) (model.Categories, error)
	DeleteCategories(model.Categories) (model.Categories, error)

	// Store
	CreateStore(model.Store) (model.Store, error)
	ReadStores() ([]model.Store, error)
	ReadStore(model.Store) (model.Store, error)
	UpdateStore(model.Store) (model.Store, error)
	DeleteStore(model.Store) (model.Store, error)
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
		&model.Merchant{},
		&model.Product{},
		&model.Categories{},
		&model.Store{},
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
