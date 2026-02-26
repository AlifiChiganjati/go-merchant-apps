package dto

type (
	MerchantRequestDto struct {
		UserID          string `json:"user_id"`
		MerchantName    string `json:"merchant_name"`
		MerchantAddress string `json:"merchant_address"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
	}
)
