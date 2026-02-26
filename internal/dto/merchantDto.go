package dto

import "time"

type (
	MerchantRequestDto struct {
		UserID          string    `json:"user_id"`
		MerchantName    string    `json:"merchant_name"`
		MerchantAddress string    `json:"merchant_address"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	}
	MerchantResponseDto struct {
		MerchantName    string `json:"merchant_name"`
		MerchantAddress string `json:"merchant_address"`
	}
)
