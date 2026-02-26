package di

import "github.com/AlifiChiganjati/go-merchant-apps/internal/usecase"

type (
	UsecaseDI interface {
		AuthUsecase() usecase.AuthUsecase
	}

	usecaseDI struct {
		repo RepoDI
	}
)

func NewUseCaseDI(repo RepoDI) UsecaseDI {
	return &usecaseDI{repo: repo}
}

func (u *usecaseDI) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(u.repo.UserRepo())
}
