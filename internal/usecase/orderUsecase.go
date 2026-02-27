package usecase

import (
	"fmt"
	"time"

	"github.com/AlifiChiganjati/go-merchant-apps/internal/dto"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/models"
	"github.com/AlifiChiganjati/go-merchant-apps/internal/repository"
)

type (
	OrderUsecase interface {
		CreateOrder(userID string, payload dto.OrderRequestDto) (models.Order, error)
	}
	orderUsecase struct {
		repo        repository.OrderRepository
		repoProduct repository.ProductRepostory
	}
)

func NewOrderUsecase(repo repository.OrderRepository, repoProduct repository.ProductRepostory) OrderUsecase {
	return &orderUsecase{repo: repo, repoProduct: repoProduct}
}

func (uc *orderUsecase) CreateOrder(userID string, payload dto.OrderRequestDto) (models.Order, error) {
	if payload.MerchantID == "" {
		return models.Order{}, fmt.Errorf("merchant id tidak boleh kosong")
	}
	if len(payload.Items) == 0 {
		return models.Order{}, fmt.Errorf("product tidak boleh kosong")
	}

	var totalPoin int
	var subTotal float64
	var orderItems []models.OrderItem

	for _, item := range payload.Items {
		if item.Qty <= 0 {
			return models.Order{}, fmt.Errorf("qty product %s tidak boleh minus %d ", item.ProductID, item.Qty)
		}

		product, err := uc.repoProduct.GetByID(item.ProductID)
		if err != nil {
			return models.Order{}, fmt.Errorf("produk %s gak ketemu", item.ProductID)
		}

		itemTotal := float64(item.Qty) * product.Price
		itemPoin := item.Qty * product.Point

		subTotal += itemTotal
		totalPoin += itemPoin

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Qty:       item.Qty,
			Point:     itemPoin,
			Total:     itemTotal,
			CreatedAt: time.Now(),
		})
	}

	newOrder := models.Order{
		UserID:        userID,
		MerchantID:    payload.MerchantID,
		TotalPoin:     totalPoin,
		SubTotal:      subTotal,
		TransactionNo: fmt.Sprintf("TRX-%d", time.Now().UnixNano()), // Simple Trx No
		CreatedAt:     time.Now(),
		Item:          orderItems,
	}

	order, err := uc.repo.Create(newOrder)
	if err != nil {
		return models.Order{}, fmt.Errorf("gagal bikin order: %w", err)
	}

	return order, nil
}
