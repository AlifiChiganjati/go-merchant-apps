package di

import "github.com/AlifiChiganjati/go-merchant-apps/internal/repository"

type (
	RepoDI interface {
		UserRepo() repository.UserRepository
		MerchantRepo() repository.MerchantRepository
		ProductRepo() repository.ProductRepostory
	}
	repoDI struct {
		infra InfraDI
	}
)

func NewRepoDI(infra InfraDI) RepoDI {
	return &repoDI{infra: infra}
}

func (r *repoDI) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoDI) MerchantRepo() repository.MerchantRepository {
	return repository.NewMerchantRepository(r.infra.Conn())
}

func (r *repoDI) ProductRepo() repository.ProductRepostory {
	return repository.NewProductRepository(r.infra.Conn())
}
