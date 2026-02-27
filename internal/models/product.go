package models

import "time"

type (
	Product struct {
		ID          string    `json:"id"`
		MerchantID  string    `json:"merchat_id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		Point       int       `json:"point"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
