package repository

import (
	"context"
	"database/sql"
	"order-service/internal/domain"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // Импорт драйвера для PostgreSQL (вы можете изменить на другой драйвер)
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	GetByID(ctx context.Context, id string) (*domain.Order, error)
	GetByUser(ctx context.Context, userID string) ([]*domain.Order, error)
	GetAll(ctx context.Context) ([]*domain.Order, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	Delete(ctx context.Context, id string) error
}

type SQLOrderRepository struct {
	db *sql.DB
}

func NewSQLOrderRepository(db *sql.DB) *SQLOrderRepository {
	return &SQLOrderRepository{db: db}
}

func (r *SQLOrderRepository) Create(ctx context.Context, order *domain.Order) error {
	query := `INSERT INTO orders (user_id, items, total, status, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, order.UserID, pq.Array(order.Items), order.Total, order.Status, time.Unix(order.CreatedAt, 0))
	return err
}

func (r *SQLOrderRepository) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	query := `SELECT id, user_id, items, total, status, created_at FROM orders WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var order domain.Order
	var items []string
	var createdAt time.Time
	if err := row.Scan(&order.ID, &order.UserID, pq.Array(&items), &order.Total, &order.Status, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	order.Items = items
	order.CreatedAt = createdAt.Unix()
	return &order, nil
}

func (r *SQLOrderRepository) GetByUser(ctx context.Context, userID string) ([]*domain.Order, error) {
	query := `SELECT id, user_id, items, total, status, created_at FROM orders WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var order domain.Order
		var items []string
		var createdAt time.Time
		if err := rows.Scan(&order.ID, &order.UserID, pq.Array(&items), &order.Total, &order.Status, &createdAt); err != nil {
			return nil, err
		}

		order.Items = items
		order.CreatedAt = createdAt.Unix()
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *SQLOrderRepository) GetAll(ctx context.Context) ([]*domain.Order, error) {
	query := `SELECT id, user_id, items, total, status, created_at FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var order domain.Order
		var items []string
		var createdAt time.Time
		if err := rows.Scan(&order.ID, &order.UserID, pq.Array(&items), &order.Total, &order.Status, &createdAt); err != nil {
			return nil, err
		}

		order.Items = items
		order.CreatedAt = createdAt.Unix()
		orders = append(orders, &order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *SQLOrderRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *SQLOrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
