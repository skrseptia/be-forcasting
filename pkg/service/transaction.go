package service

import (
	"errors"
	"fmt"
	"food_delivery_api/pkg/model"
	"strconv"
	"time"
)

func (s *service) AddTransaction(p model.Transaction, user string) (model.Transaction, error) {
	var trx model.Transaction
	var err error

	rollback := false

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

	for _, v := range p.TransactionLines {
		product, err := s.GetProduct(model.Product{Model: model.Model{ID: v.ProductID}})
		if err != nil {
			rollback = true
			return trx, err
		}

		if v.Qty > product.Qty {
			rollback = true
			return trx, errors.New(fmt.Sprintf("%s only %v left", product.Name, product.Qty))
		} else if product.Qty == 0 {
			rollback = true
			return trx, errors.New(fmt.Sprintf("%s is empty", product.Name))
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

		// reduce stock
		product.Qty -= v.Qty
		product, err = s.rmy.UpdateProduct(product)
		if err != nil {
			return trx, err
		}
	}

	// bulk insert transaction lines
	lines, err = s.rmy.CreateTransactionLines(lines)
	if err != nil {
		rollback = true
		return trx, err
	}

	if rollback {
		trx, err = s.RemoveTransaction(trx)
		if err != nil {
			return trx, err
		}
	}

	// update total amount in transaction header
	trx.Total = total
	trx, err = s.rmy.UpdateTransaction(trx)
	if err != nil {
		return trx, err
	}
	trx.TransactionLines = lines

	return trx, nil
}

func (s *service) GetTransactions() ([]model.Transaction, error) {
	list, err := s.rmy.ReadTransactions()
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
