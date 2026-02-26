package models

import "time"

type (
	Merchant struct {
		ID              string    `json:"id"`
		UserID          string    `json:"user_id"`
		MerchantName    string    `json:"merchant_name"`
		MerchantAddress string    `json:"merchant_address"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}
)
