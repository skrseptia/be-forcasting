package service

import (
	"errors"
	"fmt"
	"food_delivery_api/cfg"
	"food_delivery_api/pkg/model"
	"strconv"
	"time"
)

func (s *service) AddTransaction(p model.Transaction, user string) (model.Transaction, error) {
	var trx model.Transaction
	var err error

	// create transaction header
	trx.TrxID = fmt.Sprintf("TRX-%s", strconv.FormatInt(time.Now().Unix(), 10))
	trx.CreatedBy = user
	trx.Customer = p.Customer

	if p.TrxDate != "" { // Jika TrxDate dikirim dari frontend
    trx.CreatedAt, err = time.Parse("2006-01-02", p.TrxDate)
	trx.UpdatedAt, err = time.Parse("2006-01-02", p.TrxDate)

    if err != nil {
        return trx, errors.New("invalid date format, use YYYY-MM-DD")
    }
		parsedDate, err := time.Parse("2006-01-02", p.TrxDate)
    if err != nil {
        return trx, errors.New("invalid date format, use YYYY-MM-DD")
    }
    trx.TrxDate = parsedDate.Format("2006-01-02") // Konversi kembali ke string

	} 
	trx, err = s.rmy.CreateTransaction(trx)
	if err != nil {
		return trx, err
	}

	// create transaction lines
	var lines []model.TransactionLine
	var total float64
	var products []model.Product
	parsedDate, err := time.Parse("2006-01-02", p.TrxDate)

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
			CreatedAt:     parsedDate,
			UpdatedAt:     parsedDate,	
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
	for _, v := range products {
		v.Qty -= 0
		// v.Qty -= lines[i].Qty
		_, err = s.rmy.UpdateProduct(v)
		if err != nil {
			return trx, err
		}
	}

	return trx, nil
}

func (s *service) GetTransactions(qp model.QueryGetTransactions) ([]model.Transaction, int64, error) {
	var list []model.Transaction
	var ttl int64
	var err error

	if qp.StartDate != "" && qp.EndDate != "" {
		list, ttl, err = s.rmy.ReadTransactionsBetweenDate(qp)
	} else {
		list, ttl, err = s.rmy.ReadTransactions(qp)
	}

	if err != nil {
		return list, ttl, err
	}

	for i, v := range list {
		list[i].TrxDate = v.CreatedAt.Format(cfg.AppTLayout)
	}

	return list, ttl, nil
}

func (s *service) GetTransaction(p model.Transaction) (model.Transaction, error) {
	obj, err := s.rmy.ReadTransaction(p)
	if err != nil {
		return obj, err
	}

	obj.TrxDate = obj.CreatedAt.Format(cfg.AppTLayout)

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
