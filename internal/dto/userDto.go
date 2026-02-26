package dto

import "time"

type (
	UserResponseDto struct {
		Fullname    string    `json:"fullname"`
		Role        string    `json:"role"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	UserRequestDto struct {
		Fullname    string    `json:"fullname"`
		Password    string    `json:"password"`
		Role        string    `json:"role"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
