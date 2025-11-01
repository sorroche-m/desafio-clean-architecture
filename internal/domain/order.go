package domain

import "time"

type Order struct {
    ID          string    `json:"id" db:"id"`
    CustomerID  string    `json:"customer_id" db:"customer_id"`
    Amount      float64   `json:"amount" db:"amount"`
    Status      string    `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type OrderRepository interface {
    Create(order *Order) error
    List() ([]*Order, error)
    GetByID(id string) (*Order, error)
}

type OrderUseCase interface {
    CreateOrder(customerID string, amount float64) (*Order, error)
    ListOrders() ([]*Order, error)
    GetOrder(id string) (*Order, error)
}