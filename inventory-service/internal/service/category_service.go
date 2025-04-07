package service

import (
	"errors"
	"inventory-service/internal/repository"
	"inventory-service/pkg/models"
)

const MaxProductsPerCategory = 100

func AddProductToCategory(categoryID int, product models.Product) error {
	category := repository.GetCategoryByID(categoryID)
	if category == nil {
		return errors.New("category not found")
	}

	products := repository.GetProductsByCategory(categoryID)
	if len(products) >= MaxProductsPerCategory {
		return errors.New("category has reached its maximum capacity")
	}

	repository.CreateProduct(product)
	return nil
}
func DeleteCategoryIfEmpty(categoryID int) error {
	products := repository.GetProductsByCategory(categoryID)
	if len(products) > 0 {
		return errors.New("category is not empty")
	}

	err := repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}
