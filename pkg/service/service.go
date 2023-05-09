package service

import (
	"food_delivery_api/pkg/model"
	"food_delivery_api/pkg/storage/mysql"

	"github.com/gin-gonic/gin"
)

type Service interface {
	// User
	AddUser(model.User) (model.User, error)
	GetUsers(*gin.Context) ([]model.User, int64, error)
	GetUser(model.User) (model.User, error)
	GetUserByEmailPassword(model.User) (model.User, error)
	EditUser(model.User) (model.User, error)
	RemoveUser(model.User) (model.User, error)

	// Category
	AddCategory(model.Category) (model.Category, error)
	GetCategories(*gin.Context) ([]model.Category, int64, error)
	GetCategory(model.Category) (model.Category, error)
	EditCategory(model.Category) (model.Category, error)
	RemoveCategory(model.Category) (model.Category, error)

	// UOM
	AddUOM(model.UOM) (model.UOM, error)
	GetUOMs(*gin.Context) ([]model.UOM, int64, error)
	GetUOM(model.UOM) (model.UOM, error)
	EditUOM(model.UOM) (model.UOM, error)
	RemoveUOM(model.UOM) (model.UOM, error)

	// Product
	AddProduct(model.ProductRequest) (model.Product, error)
	GetProducts(*gin.Context) ([]model.Product, int64, error)
	GetProduct(model.Product) (model.Product, error)
	EditProduct(model.ProductRequest) (model.Product, error)
	RemoveProduct(model.Product) (model.Product, error)

	// Transaction
	AddTransaction(model.Transaction, string) (model.Transaction, error)
	GetTransactions(*gin.Context) ([]model.Transaction, int64, error)
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
