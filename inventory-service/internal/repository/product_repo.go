package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-service/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	GetAll(ctx context.Context) ([]*domain.Product, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}
type SQLProductRepositorry struct {
	db *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepositorry {
	return &SQLProductRepositorry{db: db}
}
func (r *SQLProductRepositorry) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *SQLProductRepositorry) Update(ctx context.Context, id string) error {
	query := `UPDATE products SET name = $1, category_id = $2, price = $3, stock = $4 WHERE id = $5`
	var product domain.Product
	result, err := r.db.ExecContext(ctx, query, product.Name, product.CategoryID, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product not founded")
	}
	return nil
}
func (r *SQLProductRepositorry) GetAll(ctx context.Context) ([]*domain.Product, error) {
	query := `SELECT id, name, category_id, price, stock FROM products`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// CreateProduct inserts a new product into the database
func (r *SQLProductRepositorry) Create(ctx context.Context, product *domain.Product) error {
	query := `INSERT INTO products (name, category_id, price, stock) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, product.Name, product.CategoryID, product.Price, product.Stock).Scan(&product.ID)

}

// GetProductByID fetches a product by ID from the database
func (r *SQLProductRepositorry) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product
	query := `SELECT id, name, category_id, price, stock FROM products WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Stock)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
