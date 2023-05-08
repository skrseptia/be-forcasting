package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *service) AddTransaction(p model.Transaction, user string) (model.Transaction, error) {
	var trx model.Transaction
	var err error

	// create transaction header
	trx.TrxID = fmt.Sprintf("TRX-%s", strconv.FormatInt(time.Now().Unix(), 10))
	trx.CreatedBy = user
	trx.Customer = p.Customer

	trx, err = s.rmy.CreateTransaction(trx)
	if err != nil {
		return trx, err
	}

	// create transaction lines
	var lines []model.TransactionLine
	var total float64
	var products []model.Product

	for _, v := range p.TransactionLines {
		product, err := s.GetProduct(model.Product{Model: model.Model{ID: v.ProductID}})
		if err != nil {
			return trx, err
		}

		// store product to the list for reducing the stock later
		products = append(products, product)

		if product.Qty == 0 {
			trx, _ = s.RemoveTransaction(trx)
			return trx, errors.New(fmt.Sprintf("%s is empty", product.Name))
		} else if v.Qty > product.Qty {
			trx, _ = s.RemoveTransaction(trx)
			return trx, errors.New(fmt.Sprintf("%s only %v left", product.Name, product.Qty))
		}

		subTotal := v.Qty * product.Price

		lines = append(lines, model.TransactionLine{
			TransactionID: trx.ID,
			ProductID:     product.ID,
			Code:          product.Code,
			Name:          product.Name,
			Description:   product.Description,
			Qty:           v.Qty,
			UOM:           product.UOM.Name,
			Price:         product.Price,
			SubTotal:      subTotal,
		})

		total += subTotal
	}

	// bulk insert transaction lines
	lines, err = s.rmy.CreateTransactionLines(lines)
	if err != nil {
		return trx, err
	}

	// update total amount in transaction header
	trx.Total = total
	trx, err = s.rmy.UpdateTransaction(trx)
	if err != nil {
		return trx, err
	}
	trx.TransactionLines = lines

	// reduce stock
	for i, v := range products {
		v.Qty -= lines[i].Qty
		_, err = s.rmy.UpdateProduct(v)
		if err != nil {
			return trx, err
		}
	}

	return trx, nil
}

func (s *service) GetTransactions(c *gin.Context) ([]model.Transaction, error) {
	list, err := s.rmy.ReadTransactions(c)
	if err != nil {
		return list, err
	}

	return list, nil
}

func (s *service) GetTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) EditTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	obj, err = s.rmy.UpdateTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func (s *service) RemoveTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	_, err = s.rmy.DeleteTransaction(p)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
