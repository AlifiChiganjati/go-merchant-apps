package repository

import (
	"database/sql"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
)

type (
	CartRepository interface {
		Create(payload dto.CartRequestDto) (models.Cart, error)
	}
	cartRepository struct {
		db *sql.DB
	}
)

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepository{db: db}
}

func (repo *cartRepository) Create(payload dto.CartRequestDto) (models.Cart, error) {
	var cart models.Cart
	err := repo.db.QueryRow(`
		INSERT INTO cart (user_id, product_id, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id, product_id) 
DO UPDATE SET 
    quantity = cart.quantity + EXCLUDED.quantity,
    updated_at = EXCLUDED.updated_at
RETURNING user_id, product_id, quantity, created_at, updated_at`,
		payload.UserID,
		payload.ProductID,
		payload.Quantity,
		time.Now(),
		time.Now(),
	).Scan(
		&cart.UserID,
		&cart.ProductID,
		&cart.Quantity,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}
