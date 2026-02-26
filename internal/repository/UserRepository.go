package repository

import (
	"database/sql"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
)

type (
	UserRepository interface {
		Create(payload dto.UserRequestDto) (models.User, error)
	}
	userRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(payload dto.UserRequestDto) (models.User, error) {
	var user models.User

	err := u.db.QueryRow(`
     INSERT INTO users 
			(fullname,password,email,phone_number,role,created_at,updated_at) 
		VALUES
		($1,$2,$3,$4,$5,$6,$7)
		RETURNING  fullname, password, email, phone_number, role, created_at, updated_at
		`,
		payload.Fullname,
		payload.Password,
		payload.Email,
		payload.PhoneNumber,
		payload.Role,
		time.Now(),
		time.Now()).Scan(
		&user.Fullname,
		&user.Password,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
