package service

import (
	"errors"
	"word-api/dao"
	"word-api/dto"
	"word-api/utils"
)

func AddSKUs(products []dto.Product) error {
	stockArray := make([]dao.Stock, len(products))
	for i, product := range products {
		stockArray[i] = dao.Stock{Id: utils.UUID(), VariationId: product.VariationId, ProductId: product.ProductId}
	}
	count, err := dao.AddSKUs(stockArray)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Add stock error. No one save success.")
	}
	return nil
}
