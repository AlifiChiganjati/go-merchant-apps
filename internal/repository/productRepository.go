package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
)

type (
	ProductRepostory interface {
		Create(userID string, payload dto.ProductRequestDto) (models.Product, error)
		GetByName(name string, limit, offset int) ([]models.Product, error)
		CountByName(name string) (int, error)
		GetByID(id string) (models.Product, error)
	}
	productRepository struct {
		db *sql.DB
	}
)

func NewProductRepository(db *sql.DB) ProductRepostory {
	return &productRepository{db: db}
}

func (repo *productRepository) Create(userID string, payload dto.ProductRequestDto) (models.Product, error) {
	var merchantID string

	err := repo.db.
		QueryRow(`SELECT id FROM merchants WHERE user_id = $1`, userID).
		Scan(&merchantID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, fmt.Errorf("merchant not found for this user")
		}
		return models.Product{}, err
	}

	var product models.Product

	err = repo.db.QueryRow(`
		INSERT INTO products 
			(merchant_id, name, price, description, point, created_at, updated_at)
		VALUES 
			($1,$2,$3,$4,$5,$6,$7)
		RETURNING id, merchant_id, name, price, description, point, created_at, updated_at
	`,
		merchantID,
		payload.Name,
		payload.Price,
		payload.Description,
		payload.Point,
		time.Now(),
		time.Now(),
	).Scan(
		&product.ID,
		&product.MerchantID,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Point,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (repo *productRepository) GetByName(name string, limit, offset int) ([]models.Product, error) {
	rows, err := repo.db.Query(`
		SELECT 
			id, merchant_id, name, price, description, point, created_at, updated_at
		FROM products
		WHERE name ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`,
		"%"+name+"%",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.MerchantID,
			&p.Name,
			&p.Price,
			&p.Description,
			&p.Point,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *productRepository) CountByName(name string) (int, error) {
	var total int

	err := repo.db.QueryRow(`
		SELECT COUNT(*)
		FROM products
		WHERE name ILIKE $1
	`,
		"%"+name+"%",
	).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (db *productRepository) GetByID(id string) (models.Product, error) {
	var product models.Product
	err := db.db.QueryRow(`
SELECT price,point FROM products WHERE id=$1
		`, id).Scan(&product.Price, &product.Point)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
