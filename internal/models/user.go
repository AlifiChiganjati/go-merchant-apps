package models

import "time"

type (
	User struct {
		ID          string    `json:"id"`
		Fullname    string    `json:"fullname"`
		Password    string    `json:"Password"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		Role        string    `json:"role"`
		Point       int       `json:"point"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
