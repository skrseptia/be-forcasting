package mysql

import (
	"errors"
	"food_delivery_api/cfg"
	"food_delivery_api/pkg/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepositoryMySQL interface {
	// Users
	CreateUser(model.User) (model.User, error)
	ReadUsers() ([]model.User, error)
	ReadUser(model.User) (model.User, error)
	ReadUserByEmailPassword(model.User) (model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)

	// Categories
	CreateCategory(model.Category) (model.Category, error)
	ReadCategories() ([]model.Category, error)
	ReadCategory(model.Category) (model.Category, error)
	UpdateCategory(model.Category) (model.Category, error)
	DeleteCategory(model.Category) (model.Category, error)

	// UOMs
	CreateUOM(model.UOM) (model.UOM, error)
	ReadUOMs() ([]model.UOM, error)
	ReadUOM(model.UOM) (model.UOM, error)
	UpdateUOM(model.UOM) (model.UOM, error)
	DeleteUOM(model.UOM) (model.UOM, error)

	// Products
	CreateProduct(model.Product) (model.Product, error)
	ReadProducts() ([]model.Product, error)
	ReadProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)

	// Transactions
	CreateTransaction(model.Transaction) (model.Transaction, error)
	ReadTransactions() ([]model.Transaction, error)
	ReadTransaction(model.Transaction) (model.Transaction, error)
	UpdateTransaction(model.Transaction) (model.Transaction, error)
	DeleteTransaction(model.Transaction) (model.Transaction, error)

	// Transaction Lines
	CreateTransactionLine(model.TransactionLine) (model.TransactionLine, error)
	ReadTransactionLines() ([]model.TransactionLine, error)
	ReadTransactionLine(model.TransactionLine) (model.TransactionLine, error)
	UpdateTransactionLine(model.TransactionLine) (model.TransactionLine, error)
	DeleteTransactionLine(model.TransactionLine) (model.TransactionLine, error)
}

type Storage struct {
	db *gorm.DB
}

func NewStorage(c cfg.MySQL, goEnv string) (*Storage, error) {
	var err error

	s := new(Storage)

	var ll logger.LogLevel
	if goEnv == "local" {
		ll = logger.Info
	} else {
		ll = logger.Silent
	}

	s.db, err = gorm.Open(mysql.Open(c.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(ll),
	})
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
		&model.Category{},
		&model.UOM{},
		&model.Product{},
		&model.Transaction{},
		&model.TransactionLine{},
	)

	return err
}

func seedDB(s *Storage) error {
	err := s.db.First(&model.User{}, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create super admin
		admin, err := s.CreateUser(model.User{
			Model:    model.Model{},
			FullName: "Super Admin",
			Email:    "admin@mail.com",
			Password: "password",
			ImageURL: "www.image.com",
			Phone:    "+6281234567890",
			Address:  "Karawang",
			Role:     cfg.RoleAdministrator,
		})
		if err != nil {
			return err
		}
		log.Println("Super Admin Created", admin)

		// create user
		user, err := s.CreateUser(model.User{
			FullName: "Ikhsan Guntara",
			Email:    "ikhsanguntara22@gmail.com",
			Password: "password",
			ImageURL: "www.image.com",
			Phone:    "+6285927405167",
			Address:  "Klari",
			Role:     cfg.RoleUser,
		})
		if err != nil {
			return err
		}
		log.Println("User Created", user)

		// create category
		pupuk, err := s.CreateCategory(model.Category{
			Code: "PPK",
			Name: "Pupuk",
		})
		if err != nil {
			return err
		}
		log.Println("Category Created", pupuk)

		obat, err := s.CreateCategory(model.Category{
			Code: "OBT",
			Name: "Obat - Obatan",
		})
		if err != nil {
			return err
		}
		log.Println("Category Created", obat)

		// create uom
		karung, err := s.CreateUOM(model.UOM{
			Name: "Karung",
		})
		if err != nil {
			return err
		}
		log.Println("UOM Created", karung)

		botol, err := s.CreateUOM(model.UOM{
			Name: "Botol",
		})
		if err != nil {
			return err
		}
		log.Println("UOM Created", botol)

		// create product
		kompos, err := s.CreateProduct(model.Product{
			Code:        pupuk.Code,
			Name:        "Pupuk Kompos",
			Description: "Pupuk Kompos 1 Karung",
			ImageURL:    "https://images.tokopedia.net/img/cache/500-square/product-1/2019/12/24/2626509/2626509_6d2cc34f-e163-4d77-a494-3dd669483f99_720_720.jpg",
			Qty:         100,
			UOMID:       int(pupuk.ID),
			UOM:         karung,
			Price:       20000,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created", kompos)

		kandang, err := s.CreateProduct(model.Product{
			Code:        pupuk.Code,
			Name:        "Pupuk Kandang",
			Description: "Pupuk Kandang 1 Karung",
			ImageURL:    "https://sikumis.com/media/frontend/products/pupuk(1)1.jpg",
			Qty:         50,
			UOMID:       int(pupuk.ID),
			UOM:         karung,
			Price:       25000,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created", kandang)

		pestina, err := s.CreateProduct(model.Product{
			Code:        obat.Code,
			Name:        "Pestina MSG 3",
			Description: "Pestisida Nabati 1 Botol",
			ImageURL:    "https://s2.bukalapak.com/img/79689491992/large/data.jpeg",
			Qty:         250,
			UOMID:       int(obat.ID),
			UOM:         botol,
			Price:       49000,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created", pestina)

		em4, err := s.CreateProduct(model.Product{
			Code:        obat.Code,
			Name:        "EM4 Pertanian",
			Description: "EM4 Pertanian 1 Botol",
			ImageURL:    "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2021/6/10/c2b626fc-1e59-499b-80bc-7a10d9b55b29.jpg.webp?ect=4g",
			Qty:         500,
			UOMID:       int(obat.ID),
			UOM:         botol,
			Price:       33000,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created", em4)
	}

	return nil
}
