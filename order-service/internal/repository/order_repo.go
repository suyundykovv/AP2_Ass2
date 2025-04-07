package repository

import "order-service/pkg/models"

var orders = []models.Order{}

func CreateOrder(order models.Order) {
	orders = append(orders, order)
}

func GetOrderByID(id int) *models.Order {
	for _, order := range orders {
		if order.ID == id {
			return &order
		}
	}
	return nil
}

// Добавьте другие CRUD функции: UpdateOrder, DeleteOrder, ListOrders...
