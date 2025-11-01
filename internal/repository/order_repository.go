package repository

import (
	"database/sql"

	"github.com/sorroche-m/desafio-clean-architecture/internal/domain"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	query := `
        INSERT INTO orders (id, customer_id, amount, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	_, err := r.db.Exec(query, order.ID, order.CustomerID, order.Amount,
		order.Status, order.CreatedAt, order.UpdatedAt)

	return err
}

func (r *orderRepository) List() ([]*domain.Order, error) {
	query := `
        SELECT id, customer_id, amount, status, created_at, updated_at
        FROM orders
        ORDER BY created_at DESC
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Amount,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) GetByID(id string) (*domain.Order, error) {
	query := `
        SELECT id, customer_id, amount, status, created_at, updated_at
        FROM orders
        WHERE id = $1
    `

	order := &domain.Order{}
	err := r.db.QueryRow(query, id).Scan(
		&order.ID,
		&order.CustomerID,
		&order.Amount,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}
