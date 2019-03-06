package dto

type Products struct {
	ProductArray []Product `json:"products" binding:"required"`
}

type Product struct {
	ProductId string `json:"productId" binding:"required"`

	VariationId string `json:"variationId" binding:"required"`
}
