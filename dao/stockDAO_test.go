package dao

import (
	"github.com/sillyhatxu/mysql-client"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const (
	dataSourceName = `sillyhat:sillyhat@tcp(127.0.0.1:3308)/sillyhat`
	maxIdleConns   = 5
	maxOpenConns   = 10
)

func TestAddSKUs(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	stockArray := make([]Stock, 1000)
	for i := range stockArray {
		stockArray[i].Id = "ID_" + strconv.Itoa(i)
		stockArray[i].VariationId = "VARIATION_ID_" + strconv.Itoa(i)
		stockArray[i].ProductId = "PRODUCT_ID_" + strconv.Itoa(i)
	}
	total, err := AddSKUs(stockArray)
	assert.Nil(t, err)
	assert.EqualValues(t, total, 1000)
}

func TestConsume(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count, err := Consume("order_111", "VARIATION_ID_111", "PRODUCT_ID_111", 1)
	assert.Nil(t, err)
	assert.EqualValues(t, count, 1)
}

func TestReleaseStockByOrderId(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count, err := ReleaseStockByOrderId("order_111")
	assert.Nil(t, err)
	assert.EqualValues(t, count, 1)
}

func TestDeleteStockByParams(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	productCount, err := DeleteStockByParams("PRODUCT_ID_0", "", 0)
	assert.Nil(t, err)
	assert.EqualValues(t, productCount, 1)

	variationCount, err := DeleteStockByParams("", "VARIATION_ID_1", 2)
	assert.Nil(t, err)
	assert.EqualValues(t, variationCount, 2)
}

func TestGetAllStockList(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	stockAmountArray, err := GetAllStockList()
	assert.Nil(t, err)
	assert.EqualValues(t, len(stockAmountArray), 989)
	for _, stockAmount := range stockAmountArray {
		assert.NotNil(t, stockAmount.OrderId)
		assert.NotNil(t, stockAmount.VariationId)
		assert.NotNil(t, stockAmount.VariationId)
		assert.NotEqual(t, stockAmount.StockAmount, 0)
	}
}

func TestGetStockByVariationIds(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	variationIds := []string{"VARIATION_ID_101", "VARIATION_ID_102", "VARIATION_ID_103"}
	stockAmountArray, err := GetStockByVariationIds(variationIds)
	assert.Nil(t, err)
	assert.EqualValues(t, len(stockAmountArray), 3)
	assert.EqualValues(t, stockAmountArray[0].VariationId, "VARIATION_ID_101")
	assert.EqualValues(t, stockAmountArray[0].ProductId, "PRODUCT_ID_101")
	assert.EqualValues(t, stockAmountArray[0].StockAmount, 5)
	assert.EqualValues(t, stockAmountArray[1].VariationId, "VARIATION_ID_102")
	assert.EqualValues(t, stockAmountArray[1].ProductId, "PRODUCT_ID_102")
	assert.EqualValues(t, stockAmountArray[1].StockAmount, 2)
	assert.EqualValues(t, stockAmountArray[2].VariationId, "VARIATION_ID_103")
	assert.EqualValues(t, stockAmountArray[2].ProductId, "PRODUCT_ID_103")
	assert.EqualValues(t, stockAmountArray[2].StockAmount, 1)
}

func TestGetStockByOrderId(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	stockAmountArray, err := GetStockByOrderId("11111")
	assert.Nil(t, err)
	assert.EqualValues(t, len(stockAmountArray), 2)
	assert.EqualValues(t, stockAmountArray[0].VariationId, "VARIATION_ID_183")
	assert.EqualValues(t, stockAmountArray[0].ProductId, "PRODUCT_ID_183")
	assert.EqualValues(t, stockAmountArray[0].StockAmount, 1)
	assert.EqualValues(t, stockAmountArray[1].VariationId, "VARIATION_ID_184")
	assert.EqualValues(t, stockAmountArray[1].ProductId, "PRODUCT_ID_184")
	assert.EqualValues(t, stockAmountArray[1].StockAmount, 1)
}

func TestGetStockByProductIds(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	productIds := []string{"PRODUCT_ID_102", "PRODUCT_ID_103"}
	stockAmountArray, err := GetStockByProductIds(productIds)
	assert.Nil(t, err)
	assert.EqualValues(t, len(stockAmountArray), 2)
	assert.EqualValues(t, stockAmountArray[0].VariationId, "VARIATION_ID_102")
	assert.EqualValues(t, stockAmountArray[0].ProductId, "PRODUCT_ID_102")
	assert.EqualValues(t, stockAmountArray[0].StockAmount, 2)
	assert.EqualValues(t, stockAmountArray[1].VariationId, "VARIATION_ID_103")
	assert.EqualValues(t, stockAmountArray[1].ProductId, "PRODUCT_ID_103")
	assert.EqualValues(t, stockAmountArray[1].StockAmount, 1)
}
