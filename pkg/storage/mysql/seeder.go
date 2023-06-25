package mysql

import (
	"errors"
	"fmt"
	"food_delivery_api/cfg"
	"food_delivery_api/pkg/model"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"time"
)

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
			Qty:         10000,
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
			Qty:         5000,
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
			Qty:         25000,
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
			Qty:         50000,
			UOMID:       int(obat.ID),
			UOM:         botol,
			Price:       33000,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created", em4)
	}

	err = s.db.First(&model.Transaction{}, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create the last 90 days transaction
		var trxs []model.Transaction
		var id uint

		// randomize user and customer
		rand.New(rand.NewSource(time.Now().UnixNano()))
		usr := []string{"Super Admin", "Ikhsan Guntara"}
		cst := []string{"Tamara Isnin", "Sekar Septi", "John Doe"}

		qp := model.QueryPagination{Page: 1, PageSize: 100}
		products, _, err := s.ReadProducts(qp)
		if err != nil {
			return err
		}

		for i := 174; i > 0; i-- {
			// max limit trx per day
			limit := rand.Intn(10) + 1

			for j := 0; j < limit; j++ {
				// create trx lines
				tl := rand.Intn(5) + 1
				var lines []model.TransactionLine
				var total float64

				date := time.Now().AddDate(0, 0, -(i))
				id += 1

				for k := 0; k < tl; k++ {
					prd := products[rand.Intn(len(products))]
					qty := rand.Intn(10) + 1
					subTotal := prd.Price * float64(qty)
					lines = append(lines, model.TransactionLine{
						Model:         model.Model{CreatedAt: date, UpdatedAt: date},
						TransactionID: id,
						ProductID:     prd.ID,
						Code:          prd.Code,
						Name:          prd.Name,
						Description:   prd.Description,
						Qty:           float64(qty),
						UOM:           prd.UOM.Name,
						Price:         prd.Price,
						SubTotal:      subTotal,
					})

					total += subTotal
				}

				// create trx header
				trxs = append(trxs, model.Transaction{
					Model:            model.Model{ID: id, CreatedAt: date, UpdatedAt: date},
					TrxID:            fmt.Sprintf("TRX-%s", strconv.FormatInt(date.UnixNano(), 10)),
					CreatedBy:        usr[rand.Intn(2)],
					Customer:         cst[rand.Intn(3)],
					Total:            total,
					TransactionLines: lines,
				})
			}
		}

		trxs, err = s.CreateTransactions(trxs)
		if err != nil {
			return err
		}
		log.Println(len(trxs), "Transactions Created")
	}

	return nil
}
