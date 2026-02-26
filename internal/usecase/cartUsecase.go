package usecase

import (
	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
)

type (
	CartUsecase interface {
		CreateCart(payload dto.CartRequestDto) (models.Cart, error)
	}
	cartUsecase struct {
		repo repository.CartRepository
	}
)

func NewCartUsecase(repo repository.CartRepository) CartUsecase {
	return &cartUsecase{repo: repo}
}

func (uc *cartUsecase) CreateCart(payload dto.CartRequestDto) (models.Cart, error) {
	cart, err := uc.repo.Create(payload)
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}
