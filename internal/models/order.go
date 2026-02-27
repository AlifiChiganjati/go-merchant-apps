package models

import "time"

type (
	Order struct {
		ID            string      `json:"id"`
		UserID        string      `json:"user_id"`
		MerchantID    string      `json:"merchant_id"`
		TotalPoin     int         `json:"point_total"`
		SubTotal      float64     `json:"subtotal"`
		TransactionNo string      `json:"transaction_no"`
		CreatedAt     time.Time   `json:"created_at"`
		Item          []OrderItem `json:"items"`
	}

	OrderItem struct {
		ID            string    `json:"id"`
		TransactionID string    `json:"transaction_id"`
		ProductID     string    `json:"product_id"`
		Qty           int       `json:"qty"`
		Point         int       `json:"point"`
		Total         float64   `json:"total"`
		CreatedAt     time.Time `json:"created_at"`
	}
)
