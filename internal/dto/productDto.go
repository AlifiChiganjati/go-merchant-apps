package dto

import "time"

type (
	ProductRequestDto struct {
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		Point       int       `json:"point"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
