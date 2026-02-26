package usecase

import (
	"fmt"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/encryption"
)

type (
	AuthUsecase interface {
		CreateUser(payload dto.UserRequestDto) (models.User, error)
	}
	authUsecase struct {
		repo repository.UserRepository
	}
)

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (au *authUsecase) CreateUser(payload dto.UserRequestDto) (models.User, error) {
	hashPassword, err := encryption.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, err
	}
	newUser := dto.UserRequestDto{
		Fullname:    payload.Fullname,
		Email:       payload.Email,
		Password:    hashPassword,
		Role:        "Customer",
		PhoneNumber: payload.PhoneNumber,
		CreatedAt:   payload.CreatedAt,
		UpdatedAt:   payload.UpdatedAt,
	}
	user, err := au.repo.Create(newUser)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %v", err.Error())
	}
	return user, nil
}
