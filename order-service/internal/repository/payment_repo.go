package repository

import (
	"database/sql"
	"order-service/pkg/models"
)

// CreatePayment inserts a new payment into the database
func CreatePayment(db *sql.DB, payment *models.Payment) error {
	query := `INSERT INTO payments (order_id, amount, method, status) 
			  VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, payment.OrderID, payment.Amount, payment.Method, payment.Status).Scan(&payment.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetPaymentByID fetches a payment by ID from the database
func GetPaymentByID(db *sql.DB, id int) (*models.Payment, error) {
	var payment models.Payment
	query := `SELECT id, order_id, amount, method, status FROM payments WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Method, &payment.Status)
	if err == sql.ErrNoRows {
		return nil, nil // No payment found
	} else if err != nil {
		return nil, err
	}
	return &payment, nil
}
