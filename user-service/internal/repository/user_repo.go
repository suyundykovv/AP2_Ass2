package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"user-service/internal/domain"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, order *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type SQLOrderRepository struct {
	db *sql.DB
}

func NewSQLOrderRepository(db *sql.DB) *SQLOrderRepository {
	return &SQLOrderRepository{db: db}
}

func (r *SQLOrderRepository) Create(ctx context.Context, user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (username, email,password, role, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.ExecContext(ctx, query, user.Username, user.Email, hashedPassword, user.Role, time.Unix(user.CreatedAt, 0))
	return err
}

func (r *SQLOrderRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID cannot be empty")
	}

	query := `SELECT id, username, email, role, created_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var user domain.User
	var createdAt time.Time

	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	user.CreatedAt = createdAt.Unix()
	return &user, nil
}

func (r *SQLOrderRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id, username, email, role, created_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		var createdAt time.Time
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &createdAt); err != nil {
			return nil, err
		}

		user.CreatedAt = createdAt.Unix()
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *SQLOrderRepository) Update(ctx context.Context, id string) error {

	var user domain.User
	query := `UPDATE users SET username = $1, email = $2,role = $3, created_at = $4 WHERE id = $5`
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Role, user.CreatedAt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}

func (r *SQLOrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
