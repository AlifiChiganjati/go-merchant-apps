package dto

import "time"

type (
	CartRequestDto struct {
		UserID    string    `json:"user_id"`
		ProductID string    `json:"product_id"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
