package repository

import (
	"database/sql"
	"order-service/pkg/models"
)

// CreateOrder inserts a new order into the database
func CreateOrder(db *sql.DB, order *models.Order) error {
	query := `INSERT INTO orders (user_id, items, status, total) 
			  VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, order.UserID, order.Items, order.Status, order.Total).Scan(&order.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderByID fetches an order by ID from the database
func GetOrderByID(db *sql.DB, id int) (*models.Order, error) {
	var order models.Order
	query := `SELECT id, user_id, items, status, total FROM orders WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.Items, &order.Status, &order.Total)
	if err == sql.ErrNoRows {
		return nil, nil // No order found
	} else if err != nil {
		return nil, err
	}
	return &order, nil
}
