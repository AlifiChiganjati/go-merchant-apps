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
		GetByEmail(email string) (models.User, error)
	}
	userRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(payload dto.UserRequestDto) (models.User, error) {
	var user models.User

	err := repo.db.QueryRow(`
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

func (repo *userRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.db.QueryRow(` 
SELECT 
		id, fullname, email, password, role 
FROM 
		users 
WHERE 
		email=$1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Fullname,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
