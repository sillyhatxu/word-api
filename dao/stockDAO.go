package dao

import (
	"bytes"
	"database/sql"
	"github.com/mitchellh/mapstructure"
	"github.com/sillyhatxu/mysql-client"
	"text/template"
)

const (
	insert_sql = `
		INSERT INTO stock_keeping_unit (id, variation_id, product_id, create_time, update_time)
		VALUES (?, ?, ?, now(), now())
	`
	delete_sql = `
		DELETE FROM stock_keeping_unit 
		WHERE 
		{{if .ProductId }}
			product_id = ? AND order_id IS NULL
		{{else}}
			variation_id = ? AND order_id IS NULL LIMIT ?
        {{end}}
	`

	consume_stock_sql = `
		UPDATE stock_keeping_unit
		SET order_id = ?
		WHERE order_id IS NULL
			and variation_id = ?
			and product_id = ?
		LIMIT ?
	`
	release_stock_by_order_id_sql = `
		UPDATE stock_keeping_unit
		SET order_id = NULL
		WHERE order_id = ?
	`
	get_all_list = `
		SELECT order_id, variation_id, product_id, count(*) as stock_amount
		FROM stock_keeping_unit
		GROUP BY order_id, variation_id, product_id
		HAVING order_id is null
	`

	get_stock_by_variation_ids_sql = `
		SELECT variation_id, product_id, count(*) as stock_amount
		FROM stock_keeping_unit
		WHERE order_id IS NULL 
		AND variation_id IN ( {{ . }} )
		GROUP BY variation_id, product_id
	`
	get_stock_by_order_id = `
		SELECT variation_id, product_id, count(*) as stock_amount
		FROM stock_keeping_unit
		WHERE order_id = ?
		GROUP BY variation_id
	`
	get_stock_by_product_ids = `
		SELECT variation_id, product_id, count(*) as stock_amount
		FROM stock_keeping_unit
		WHERE order_id IS NULL AND product_id IN ( {{ . }} )
		GROUP BY variation_id
	`
)

type Stock struct {
	Id          string `mapstructure:"id"`
	VariationId string `mapstructure:"variation_id"`
	ProductId   string `mapstructure:"product_id"`
}

func AddSKUs(stockArray []Stock) (int64, error) {
	result, err := dbclient.Client.BatchInsert(func(tx *sql.Tx) (int64, error) {
		totalCount := 0
		for _, stock := range stockArray {
			_, err := tx.Exec(insert_sql, stock.Id, stock.VariationId, stock.ProductId)
			if err != nil {
				return int64(totalCount), err
			}
			totalCount++
		}
		return int64(totalCount), nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Consume(orderId string, variationId string, productId string, quantity int) (int64, error) {
	count, err := dbclient.Client.Update(consume_stock_sql, orderId, variationId, productId, quantity)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ReleaseStockByOrderId(orderId string) (int64, error) {
	count, err := dbclient.Client.Update(release_stock_by_order_id_sql, orderId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type DeleteStock struct {
	ProductId    string `mapstructure:"product_id"`
	VariationId  string `mapstructure:"variation_id"`
	DeleteAmount int    `mapstructure:"product_id"`
}

func DeleteStockByParams(productId string, variationId string, deleteAmount int) (int64, error) {
	deleteSQL, err := getSQL(delete_sql, DeleteStock{VariationId: variationId, ProductId: productId, DeleteAmount: deleteAmount})
	if err != nil {
		return 0, err
	}
	if productId == "" {
		count, err := dbclient.Client.Delete(deleteSQL, variationId, deleteAmount)
		if err != nil {
			return 0, err
		}
		return count, nil
	}
	count, err := dbclient.Client.Delete(deleteSQL, productId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type StockAmount struct {
	OrderId     string `mapstructure:"order_id"`
	VariationId string `mapstructure:"variation_id"`
	ProductId   string `mapstructure:"product_id"`
	StockAmount int    `mapstructure:"stock_amount"`
}

func GetAllStockList() ([]StockAmount, error) {
	results, err := dbclient.Client.Find(get_all_list)
	if err != nil {
		return nil, err
	}
	var stockAmountArray []StockAmount
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &stockAmountArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(results)
	if err != nil {
		return nil, err
	}
	return stockAmountArray, nil
}

func GetStockByVariationIds(variationIds []string) ([]StockAmount, error) {
	var buffer bytes.Buffer
	for i := 0; i < len(variationIds); i++ {
		buffer.WriteString("'")
		buffer.WriteString(variationIds[i])
		buffer.WriteString("',")
	}
	var variationIdsString = buffer.String()
	if last := len(variationIdsString) - 1; last >= 0 {
		variationIdsString = variationIdsString[:last]
	}
	getStockSQL, err := getSQL(get_stock_by_variation_ids_sql, variationIdsString)
	if err != nil {
		return nil, err
	}
	results, err := dbclient.Client.Find(getStockSQL)
	if err != nil {
		return nil, err
	}
	var stockAmountArray []StockAmount
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &stockAmountArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(results)
	if err != nil {
		return nil, err
	}
	return stockAmountArray, nil
}

func GetStockByOrderId(orderId string) ([]StockAmount, error) {
	getStockSQL, err := getSQL(get_stock_by_order_id, orderId)
	if err != nil {
		return nil, err
	}
	results, err := dbclient.Client.Find(getStockSQL, orderId)
	if err != nil {
		return nil, err
	}
	var stockAmountArray []StockAmount
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &stockAmountArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(results)
	if err != nil {
		return nil, err
	}
	return stockAmountArray, nil
}

func GetStockByProductIds(productIds []string) ([]StockAmount, error) {
	var buffer bytes.Buffer
	for i := 0; i < len(productIds); i++ {
		buffer.WriteString("'")
		buffer.WriteString(productIds[i])
		buffer.WriteString("',")
	}
	var productIdsString = buffer.String()
	if last := len(productIdsString) - 1; last >= 0 {
		productIdsString = productIdsString[:last]
	}
	getStockSQL, err := getSQL(get_stock_by_product_ids, productIdsString)
	if err != nil {
		return nil, err
	}
	results, err := dbclient.Client.Find(getStockSQL)
	if err != nil {
		return nil, err
	}
	var stockAmountArray []StockAmount
	config := &mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		WeaklyTypedInput: true,
		Result:           &stockAmountArray,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(results)
	if err != nil {
		return nil, err
	}
	return stockAmountArray, nil
}

func getSQL(sql string, data interface{}) (string, error) {
	temp, err := template.New("sql").Parse(sql)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err := temp.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
