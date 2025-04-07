package repository

import (
	"fmt"
	"inventory-service/pkg/models"
)

var categories = []models.Category{}

// CreateCategory добавляет новую категорию
func CreateCategory(category models.Category) {
	categories = append(categories, category)
}

// GetCategoryByID возвращает категорию по ID
func GetCategoryByID(id int) *models.Category {
	for _, category := range categories {
		if category.ID == id {
			return &category
		}
	}
	return nil
}

// UpdateCategory обновляет существующую категорию по ID
func UpdateCategory(id int, updatedCategory models.Category) error {
	for i, category := range categories {
		if category.ID == id {
			categories[i] = updatedCategory
			return nil
		}
	}
	return fmt.Errorf("category with ID %d not found", id)
}

// DeleteCategory удаляет категорию по ID
func DeleteCategory(id int) error {
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("category with ID %d not found", id)
}

// ListCategories возвращает список всех категорий
func ListCategories() []models.Category {
	return categories
}
func GetProductsByCategory(categoryID int) []models.Product {
	var categoryProducts []models.Product
	for _, product := range products {
		if product.ID == categoryID { // Предполагается, что CategoryID — это int
			categoryProducts = append(categoryProducts, product)
		}
	}
	return categoryProducts
}
