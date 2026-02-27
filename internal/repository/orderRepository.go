package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
)

type (
	OrderRepository interface {
		Create(order models.Order) (models.Order, error)
	}
	orderRepository struct {
		db *sql.DB
	}
)

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(order models.Order) (models.Order, error) {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Order{}, err
	}

	defer tx.Rollback()

	queryOrder := `INSERT INTO orders (id, user_id, merchant_id, point_total, subtotal, transaction_no, created_at) 
                   VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.ExecContext(ctx, queryOrder,
		order.ID, order.UserID, order.MerchantID, order.TotalPoin, order.SubTotal, order.TransactionNo, order.CreatedAt)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to insert order: %v", err)
	}

	queryItem := `INSERT INTO order_items (id, transaction_id, product_id, qty, point, total, created_at) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for _, item := range order.Item {
		_, err = tx.ExecContext(ctx, queryItem,
			item.ID, order.ID, item.ProductID, item.Qty, item.Point, item.Total, item.CreatedAt)
		if err != nil {
			return models.Order{}, fmt.Errorf("failed to insert item %s: %v", item.ProductID, err)
		}

		queryUpdateStock := `UPDATE users SET point = point + $1 WHERE id =$2`
		res, err := tx.ExecContext(ctx, queryUpdateStock, item.Point, order.UserID)
		if err != nil {
			return models.Order{}, err
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			return models.Order{}, fmt.Errorf("stok produk %s tidak mencukupi atau tidak ditemukan", item.ProductID)
		}
	}

	if err := tx.Commit(); err != nil {
		return models.Order{}, err
	}

	return order, nil
}
