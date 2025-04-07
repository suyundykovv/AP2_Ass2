package service

import (
	"database/sql"
	"errors"
	"inventory-service/internal/repository"
	"inventory-service/pkg/models"
)

const MaxProductsPerCategory = 100

func AddProductToCategory(db *sql.DB, categoryID int, product models.Product) error {
	category, err := repository.GetCategoryByID(db, categoryID)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("category not found")
	}

	products, err := repository.GetProductsByCategory(db, categoryID)
	if err != nil {
		return err
	}

	if len(products) >= MaxProductsPerCategory {
		return errors.New("category has reached its maximum capacity")
	}

	err = repository.CreateProduct(db, product)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategoryIfEmpty(db *sql.DB, categoryID int) error {
	products, err := repository.GetProductsByCategory(db, categoryID)
	if err != nil {
		return err
	}

	if len(products) > 0 {
		return errors.New("category is not empty")
	}

	err = repository.DeleteCategory(db, categoryID)
	if err != nil {
		return err
	}
	return nil
}
