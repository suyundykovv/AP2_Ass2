package repository

import (
	"fmt"
	"inventory-service/pkg/models"
)

// Хранилище продуктов
var products = []models.Product{}

// CreateProduct добавляет новый продукт в хранилище
func CreateProduct(product models.Product) {
	products = append(products, product)
}

// GetProductByID возвращает продукт по ID
func GetProductByID(id int) *models.Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

// UpdateProduct обновляет информацию о продукте
func UpdateProduct(id int, updatedProduct models.Product) error {
	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}

// DeleteProduct удаляет продукт по ID
func DeleteProduct(id int) error {
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product with ID %d not found", id)
}

// ListProducts возвращает список всех продуктов
func ListProducts() []models.Product {
	return products
}
