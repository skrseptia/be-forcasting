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

	// Category
	AddCategory(model.Category) (model.Category, error)
	GetCategories() ([]model.Category, error)
	GetCategory(model.Category) (model.Category, error)
	EditCategory(model.Category) (model.Category, error)
	RemoveCategory(model.Category) (model.Category, error)

	// UOM
	AddUOM(model.UOM) (model.UOM, error)
	GetUOMs() ([]model.UOM, error)
	GetUOM(model.UOM) (model.UOM, error)
	EditUOM(model.UOM) (model.UOM, error)
	RemoveUOM(model.UOM) (model.UOM, error)

	// Product
	AddProduct(product model.ProductRequest) (model.Product, error)
	GetProducts() ([]model.Product, error)
	GetProduct(model.Product) (model.Product, error)
	EditProduct(model.ProductRequest) (model.Product, error)
	RemoveProduct(model.Product) (model.Product, error)

	// Transaction
	AddTransaction(model.Transaction, string) (model.Transaction, error)
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
