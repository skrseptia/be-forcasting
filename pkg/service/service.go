package service

import (
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/storage/mysql"
)

type Service interface {
	// User
	AddUser(model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUser(model.User) (model.User, error)
	GetUserByEmailPassword(model.User) (model.User, error)
	EditUser(model.User) (model.User, error)
	RemoveUser(model.User) (model.User, error)

	// Merchant
	AddMerchant(model.Merchant) (model.Merchant, error)
	GetMerchants() ([]model.Merchant, error)
	GetMerchant(model.Merchant) (model.Merchant, error)
	EditMerchant(model.Merchant) (model.Merchant, error)
	RemoveMerchant(model.Merchant) (model.Merchant, error)

	// Product
	AddProduct(model.Product) (model.Product, error)
	GetProducts() ([]model.Product, error)
	GetProduct(model.Product) (model.Product, error)
	EditProduct(model.Product) (model.Product, error)
	RemoveProduct(model.Product) (model.Product, error)

	// Categories
	AddCategories(model.Categories) (model.Categories, error)
	GetCategoriess() ([]model.Categories, error)
	GetCategories(model.Categories) (model.Categories, error)
	EditCategories(model.Categories) (model.Categories, error)
	RemoveCategories(model.Categories) (model.Categories, error)
}

type service struct {
	rmy mysql.RepositoryMySQL
}

func NewService(rmy mysql.RepositoryMySQL) Service {
	return &service{rmy}
}
