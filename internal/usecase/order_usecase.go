package usecase

import (
	"time"

	"github.com/sorroche-m/desafio-clean-architecture/internal/domain"

	"github.com/google/uuid"
)

type orderUseCase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUseCase(orderRepo domain.OrderRepository) domain.OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}

func (uc *orderUseCase) CreateOrder(customerID string, amount float64) (*domain.Order, error) {
	order := &domain.Order{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		Amount:     amount,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := uc.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (uc *orderUseCase) ListOrders() ([]*domain.Order, error) {
	return uc.orderRepo.List()
}

func (uc *orderUseCase) GetOrder(id string) (*domain.Order, error) {
	return uc.orderRepo.GetByID(id)
}
