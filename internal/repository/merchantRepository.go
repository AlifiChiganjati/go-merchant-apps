package repository

import (
	"database/sql"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
)

type (
	MerchantRepository interface {
		Create(userID string, payload dto.MerchantRequestDto) (models.Merchant, error)
	}

	merchantRepository struct {
		db *sql.DB
	}
)

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (repo *merchantRepository) Create(userID string, payload dto.MerchantRequestDto) (models.Merchant, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return models.Merchant{}, err
	}
	defer tx.Rollback()

	// update role
	_, err = tx.Exec(
		`UPDATE users SET role=$1 WHERE id=$2`,
		"merchant",
		userID,
	)
	if err != nil {
		return models.Merchant{}, err
	}

	// insert merchant
	_, err = tx.Exec(`
		INSERT INTO merchants (user_id, merchant_name, merchant_address)
		VALUES ($1,$2,$3)
	`,
		userID,
		payload.MerchantName,
		payload.MerchantAddress,
	)
	if err != nil {
		return models.Merchant{}, err
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		return models.Merchant{}, err
	}

	return models.Merchant{
		UserID:          userID,
		MerchantName:    payload.MerchantName,
		MerchantAddress: payload.MerchantAddress,
	}, nil
}
