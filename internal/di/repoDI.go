package di

import "github.com/AlifiChiganjati/go-merchant-apps/internal/repository"

type (
	RepoDI interface {
		UserRepo() repository.UserRepository
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
