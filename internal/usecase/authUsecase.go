package usecase

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/encryption"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
)

type (
	AuthUsecase interface {
		CreateUser(payload dto.UserRequestDto) (models.User, error)
		LoginUser(payload dto.LoginRequestDto) (dto.LoginResponseDto, error)
	}
	authUsecase struct {
		repo repository.UserRepository
	}
)

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (uc *authUsecase) CreateUser(payload dto.UserRequestDto) (models.User, error) {
	hashPassword, err := encryption.HashPassword(payload.Password)
	if err != nil {
		return models.User{}, err
	}
	newUser := dto.UserRequestDto{
		Fullname:    payload.Fullname,
		Email:       payload.Email,
		Password:    hashPassword,
		Role:        "customer",
		PhoneNumber: payload.PhoneNumber,
		CreatedAt:   payload.CreatedAt,
		UpdatedAt:   payload.UpdatedAt,
	}
	user, err := uc.repo.Create(newUser)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to create user: %v", err.Error())
	}
	return user, nil
}

func (uc *authUsecase) LoginUser(payload dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	userData, err := uc.repo.GetByEmail(payload.Email)
	if err != nil {
		return dto.LoginResponseDto{}, errors.New("invalid email or password")
	}

	isValid := encryption.CheckPasswordHash(payload.Password, userData.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("invalid email or password")
	}

	loginExpDuration, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	expiredAt := time.Now().
		Add(time.Duration(loginExpDuration) * time.Minute).
		Unix()

	accessToken, err := jwttoken.GenerateTokenJwt(
		userData.ID,
		userData.Fullname,
		userData.Role,
		expiredAt,
	)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	return dto.LoginResponseDto{
		AccessToken: accessToken,
		UserID:      userData.ID,
	}, nil
}
