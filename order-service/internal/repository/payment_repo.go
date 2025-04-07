package repository

import "order-service/pkg/models"

var payments = []models.Payment{}

func CreatePayment(payment models.Payment) {
	payments = append(payments, payment)
}

func GetPaymentByID(id int) *models.Payment {
	for _, payment := range payments {
		if payment.ID == id {
			return &payment
		}
	}
	return nil
}

// Добавьте другие функции для обработки платежей...
