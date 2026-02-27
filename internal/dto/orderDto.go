package dto

type (
	OrderItemRequest struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
	}

	OrderRequestDto struct {
		MerchantID string             `json:"merchant_id"`
		Items      []OrderItemRequest `json:"items"`
	}
)
