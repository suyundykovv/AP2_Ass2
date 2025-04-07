package repository

import "inventory-service/pkg/models"

var categories = []models.Category{}

func CreateCategory(category models.Category) {
	categories = append(categories, category)
}

func GetCategoryByID(id int) *models.Category {
	for _, category := range categories {
		if category.ID == id {
			return &category
		}
	}
	return nil
}

// Добавьте другие функции: UpdateCategory, DeleteCategory, ListCategories...
