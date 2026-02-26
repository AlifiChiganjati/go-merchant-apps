package usecase

import (
	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
)

type (
	ProductUsecase interface {
		CreateProduct(userID string, payload dto.ProductRequestDto) (models.Product, error)
	}
	productUsecase struct {
		repo repository.ProductRepostory
	}
)

func NewProductUsecase(repo repository.ProductRepostory) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (uc *productUsecase) CreateProduct(userID string, payload dto.ProductRequestDto) (models.Product, error) {
	product, err := uc.repo.Create(userID, payload)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
