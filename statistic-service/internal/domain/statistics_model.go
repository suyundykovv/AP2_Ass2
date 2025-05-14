package domain

import "time"

type Order struct {
	ID        string
	UserID    string
	Items     []string // просто список ID товаров
	Total     float64
	Status    string
	CreatedAt int64
}

const (
	OrderStatusPending   = "PENDING"
	OrderStatusCanceled  = "CANCELED"
	OrderStatusCompleted = "COMPLETED"
)

type OrderEvent struct {
	EventType  string
	OrderID    string
	UserID     string
	Timestamp  time.Time
	ItemsCount int
	Total      float64
}
