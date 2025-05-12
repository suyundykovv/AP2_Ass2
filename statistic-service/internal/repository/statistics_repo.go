package repository

import (
	"context"
	"database/sql"
	"time"
)

type StatisticRepository interface {
	SaveOrderEvent(ctx context.Context, event *OrderEvent) error
	GetUserOrderStatistics(ctx context.Context, userID string, timePeriod string) (*UserOrderStatistics, error)
	GetGeneralUserStatistics(ctx context.Context, timePeriod string) (*GeneralUserStatistics, error)
	SaveInventoryEvent(ctx context.Context, event *InventoryEvent) error
}

type SQLStatisticRepository struct {
	db *sql.DB
}

func NewSQLStatisticRepository(db *sql.DB) *SQLStatisticRepository {
	return &SQLStatisticRepository{db: db}
}

type OrderEvent struct {
	EventType  string
	OrderID    string
	UserID     string
	Timestamp  time.Time
	ItemsCount int
	Total      float64
}

type InventoryEvent struct {
	EventType   string
	ProductID   string
	CategoryID  string
	Timestamp   time.Time
	StockChange int
}

type UserOrderStatistics struct {
	TotalOrders       int
	AverageOrderValue float64
	OrdersByHour      map[int]int
	OrdersByDay       map[int]int
	FavoriteCategory  string
}

type GeneralUserStatistics struct {
	TotalUsers           int
	ActiveUsers          int
	AverageOrdersPerUser float64
	MostActiveTime       string
}

func (r *SQLStatisticRepository) SaveOrderEvent(ctx context.Context, event *OrderEvent) error {
	query := `INSERT INTO order_statistics 
		(user_id, order_id, event_type, items_count, total_amount, event_time, hour_of_day, day_of_week) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		event.UserID,
		event.OrderID,
		event.EventType,
		event.ItemsCount,
		event.Total,
		event.Timestamp,
		event.Timestamp.Hour(),
		int(event.Timestamp.Weekday()),
	)
	return err
}

func (r *SQLStatisticRepository) GetUserOrderStatistics(ctx context.Context, userID string, timePeriod string) (*UserOrderStatistics, error) {
	// Implementation to query user-specific statistics
	// This would involve complex SQL queries to calculate the required metrics
	// ...
	return &UserOrderStatistics{}, nil
}

func (r *SQLStatisticRepository) GetGeneralUserStatistics(ctx context.Context, timePeriod string) (*GeneralUserStatistics, error) {
	// Implementation to query general user statistics
	// ...
	return &GeneralUserStatistics{}, nil
}

func (r *SQLStatisticRepository) SaveInventoryEvent(ctx context.Context, event *InventoryEvent) error {
	query := `INSERT INTO inventory_statistics 
		(product_id, category_id, event_type, stock_changes, last_updated) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (product_id) 
		DO UPDATE SET 
			stock_changes = inventory_statistics.stock_changes + EXCLUDED.stock_changes,
			last_updated = EXCLUDED.last_updated`

	_, err := r.db.ExecContext(ctx, query,
		event.ProductID,
		event.CategoryID,
		event.EventType,
		event.StockChange,
		event.Timestamp,
	)
	return err
}
