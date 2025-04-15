package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"order-service/pkg/models"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (int32, error)
	GetOrder(orderID int32) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(orderID int32) error
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order) (int32, error) {
	if order.UserID <= 0 {
		return 0, errors.New("invalid user ID")
	}
	if len(order.Items) == 0 {
		return 0, errors.New("order must contain items")
	}
	if order.Total <= 0 {
		return 0, errors.New("total must be positive")
	}

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return 0, err
	}

	var orderID int32
	err = r.db.QueryRow(
		`INSERT INTO orders (user_id, items, status, total) 
		 VALUES ($1, $2, $3, $4) 
		 RETURNING id`,
		order.UserID, itemsJSON, order.Status, order.Total,
	).Scan(&orderID)

	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (r *orderRepository) GetOrder(orderID int32) (*models.Order, error) {
	if orderID <= 0 {
		return nil, errors.New("invalid order ID")
	}

	var order models.Order
	var itemsJSON string

	err := r.db.QueryRow(
		`SELECT id, user_id, items, status, total 
		 FROM orders WHERE id = $1`,
		orderID,
	).Scan(&order.ID, &order.UserID, &itemsJSON, &order.Status, &order.Total)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	// Convert JSON items to slice
	if err := json.Unmarshal([]byte(itemsJSON), &order.Items); err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	if order.ID <= 0 {
		return errors.New("invalid order ID")
	}
	if len(order.Items) == 0 {
		return errors.New("order must contain items")
	}
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	result, err := r.db.Exec(
		`UPDATE orders 
		 SET user_id = $1, items = $2, status = $3, total = $4 
		 WHERE id = $5`,
		order.UserID, itemsJSON, order.Status, order.Total, order.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (r *orderRepository) DeleteOrder(orderID int32) error {
	if orderID <= 0 {
		return errors.New("invalid order ID")
	}

	result, err := r.db.Exec(
		"DELETE FROM orders WHERE id = $1",
		orderID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
