package usecase

import (
	"errors"
	"fmt"
	"regexp"
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
		repo       repository.UserRepository
		jwtService *jwttoken.JWTService
		tokenTTL   time.Duration
	}
)

func NewAuthUsecase(
	repo repository.UserRepository,
	jwtService *jwttoken.JWTService,
	tokenTTL time.Duration,
) AuthUsecase {
	return &authUsecase{
		repo:       repo,
		jwtService: jwtService,
		tokenTTL:   tokenTTL,
	}
}

func (uc *authUsecase) CreateUser(payload dto.UserRequestDto) (models.User, error) {
	if len(payload.Fullname) < 3 {
		return models.User{}, fmt.Errorf("fullname minimal 3 karakter")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(payload.Email) {
		return models.User{}, fmt.Errorf("format email salah %s ?", payload.Email)
	}

	if len(payload.Password) < 6 {
		return models.User{}, fmt.Errorf("password terlalu lemah, minimal 6 karakter!")
	}

	phoneRegex := regexp.MustCompile(`^[0-9]+$`)
	if !phoneRegex.MatchString(payload.PhoneNumber) {
		return models.User{}, fmt.Errorf("nomor telepon harus number!")
	}
	existingUser, _ := uc.repo.GetByEmail(payload.Email)
	if existingUser.ID != "" {
		return models.User{}, errors.New("email already registered")
	}

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

	expiredAt := time.Now().Add(uc.tokenTTL).Unix()

	accessToken, err := uc.jwtService.GenerateToken(
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
