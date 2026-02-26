package di

import (
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"
	"github.com/AlifiChiganjati/go-merchant-apps/pkg/jwttoken"
)

type (
	UsecaseDI interface {
		AuthUsecase() usecase.AuthUsecase
		MerchantUsecase() usecase.MerchantUsecase
		ProductUsecase() usecase.ProductUsecase
		CartUsecase() usecase.CartUsecase
	}

	usecaseDI struct {
		repo       RepoDI
		jwtService *jwttoken.JWTService
		tokenTTL   time.Duration
	}
)

func NewUseCaseDI(
	repo RepoDI,
	jwtService *jwttoken.JWTService,
	tokenTTL time.Duration,
) UsecaseDI {
	return &usecaseDI{
		repo:       repo,
		jwtService: jwtService,
		tokenTTL:   tokenTTL,
	}
}

func (uc *usecaseDI) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(
		uc.repo.UserRepo(),
		uc.jwtService,
		uc.tokenTTL,
	)
}

func (uc *usecaseDI) MerchantUsecase() usecase.MerchantUsecase {
	return usecase.NewMerchantUsecase(uc.repo.MerchantRepo())
}

func (uc *usecaseDI) ProductUsecase() usecase.ProductUsecase {
	return usecase.NewProductUsecase(uc.repo.ProductRepo())
}

func (uc *usecaseDI) CartUsecase() usecase.CartUsecase {
	return usecase.NewCartUsecase(uc.repo.CartRepo())
}
