package usecase

import (
	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
)

type (
	MerchantUsecase interface {
		RegisterMerchant(userID string, payload dto.MerchantRequestDto) (models.Merchant, error)
	}

	merchantUsecase struct {
		repo repository.MerchantRepository
	}
)

func NewMerchantUsecase(repo repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{repo: repo}
}

func (u *merchantUsecase) RegisterMerchant(
	userID string,
	payload dto.MerchantRequestDto,
) (models.Merchant, error) {
	merchant, err := u.repo.Create(userID, payload)
	if err != nil {
		return models.Merchant{}, err
	}

	return merchant, nil
}
