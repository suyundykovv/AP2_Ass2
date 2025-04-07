package repository

import (
	"database/sql"
	"inventory-service/pkg/models"
)

// CreateProduct inserts a new product into the database
func CreateProduct(db *sql.DB, product *models.Product) error {
	query := `INSERT INTO products (name, category_id, price, stock) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.QueryRow(query, product.Name, product.CategoryID, product.Price, product.Stock).Scan(&product.ID)

}

// GetProductByID fetches a product by ID from the database
func GetProductByID(db *sql.DB, id int) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, name, category_id, price, stock FROM products WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
