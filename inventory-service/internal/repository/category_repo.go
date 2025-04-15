package repository

import (
	"database/sql"
	"fmt"
	"inventory-service/pkg/models"
)

func GetCategoryByID(db *sql.DB, id int) (*models.Category, error) {
	var category models.Category
	query := `SELECT id, name FROM categories WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&category.ID, &category.Name)
	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}
	return &category, nil
}

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

// func CreateDiscount(db *sql.DB, discount *models.Discount) error {
// 	products, _ := json.Marshal(discount.ApplicableProducts)
// 	query := `INSERT INTO discounts (id, name, description, discount_percentage, applicable_products, start_date, end_date, is_active)
// 			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
// 	_, err := db.Exec(query, discount.ID, discount.Name, discount.Description, discount.DiscountPercentage, products, discount.StartDate, discount.EndDate, discount.IsActive)
// 	return err
// }

// func DeleteDiscount(db *sql.DB, id string) error {
// 	query := `DELETE FROM discounts WHERE id = $1`
// 	_, err := db.Exec(query, id)
// 	return err
// }

// func GetAllProductsWithPromotion(db *sql.DB) ([]models.Product, error) {
// 	query := `
// 	SELECT p.id, p.name, p.category_id, p.price, p.stock
// 	FROM products p
// 	JOIN discounts d ON d.is_active = true AND d.start_date <= NOW() AND d.end_date >= NOW()
// 	WHERE p.id = ANY (SELECT jsonb_array_elements_text(d.applicable_products)::int)
// 	`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var products []models.Product
// 	for rows.Next() {
// 		var p models.Product
// 		err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock)
// 		if err != nil {
// 			return nil, err
// 		}
// 		products = append(products, p)
// 	}
// 	return products, nil
// }
