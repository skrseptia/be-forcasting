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

	// Product
	AddProduct(model.Product) (model.Product, error)
	GetProducts() ([]model.Product, error)
	GetProduct(model.Product) (model.Product, error)
	EditProduct(model.Product) (model.Product, error)
	RemoveProduct(model.Product) (model.Product, error)

	// Uom
	AddUom(model.Uom) (model.Uom, error)
	GetUoms() ([]model.Uom, error)
	GetUom(model.Uom) (model.Uom, error)
	EditUom(model.Uom) (model.Uom, error)
	RemoveUom(model.Uom) (model.Uom, error)

	// Categories
	AddCategories(model.Categories) (model.Categories, error)
	GetCategoriess() ([]model.Categories, error)
	GetCategories(model.Categories) (model.Categories, error)
	EditCategories(model.Categories) (model.Categories, error)
	RemoveCategories(model.Categories) (model.Categories, error)

	// Transaction
	AddTransaction(model.Transaction) (model.Transaction, error)
	GetTransactions() ([]model.Transaction, error)
	GetTransaction(model.Transaction) (model.Transaction, error)
	EditTransaction(model.Transaction) (model.Transaction, error)
	RemoveTransaction(model.Transaction) (model.Transaction, error)
}

type service struct {
	rmy mysql.RepositoryMySQL
}

func NewService(rmy mysql.RepositoryMySQL) Service {
	return &service{rmy}
}
