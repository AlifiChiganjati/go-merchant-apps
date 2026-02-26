package di

import "github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"

type (
	UsecaseDI interface {
		AuthUsecase() usecase.AuthUsecase
		MerchantUsecase() usecase.MerchantUsecase
	}

	usecaseDI struct {
		repo RepoDI
	}
)

func NewUseCaseDI(repo RepoDI) UsecaseDI {
	return &usecaseDI{repo: repo}
}

func (uc *usecaseDI) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(uc.repo.UserRepo())
}

func (uc *usecaseDI) MerchantUsecase() usecase.MerchantUsecase {
	return usecase.NewMerchantUsecase(uc.repo.MerchantRepo())
}
