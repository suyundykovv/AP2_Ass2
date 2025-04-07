package repository

import (
	"database/sql"
	"fmt"
	"inventory-service/pkg/models"
)

// GetCategoryByID fetches a category by ID from the database
func GetCategoryByID(db *sql.DB, id int) (*models.Category, error) {
	var category models.Category
	query := `SELECT id, name FROM categories WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&category.ID, &category.Name)
	if err == sql.ErrNoRows {
		return nil, nil // No category found
	} else if err != nil {
		return nil, err
	}
	return &category, nil
}

// UpdateCategory updates a category by ID in the database
func UpdateCategory(db *sql.DB, id int, updatedCategory models.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	result, err := db.Exec(query, updatedCategory.Name, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("category with ID %d not found", id)
	}
	return nil
}

// repository/category.go

func CreateCategory(db *sql.DB, category models.Category) error {
	query := `INSERT INTO categories (name) VALUES ($1)`
	_, err := db.Exec(query, category.Name)
	return err
}
func GetAllCategories(db *sql.DB) ([]models.Category, error) {
	query := `SELECT id, name FROM categories`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

func DeleteCategory(db *sql.DB, id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("category with ID %d not found", id)
	}
	return nil
}
