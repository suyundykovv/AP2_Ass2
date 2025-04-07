package repository

import "inventory-service/pkg/models"

var products = []models.Product{}

func CreateProduct(product models.Product) {
	products = append(products, product)
}

func GetProductByID(id int) *models.Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

// Добавьте другие функции: UpdateProduct, DeleteProduct, ListProducts...
