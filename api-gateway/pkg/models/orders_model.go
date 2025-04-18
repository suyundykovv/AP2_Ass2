package models

import (
	"time"
)

// OrderStatus represents the possible states of an order
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// OrderItem represents a single item in an order
type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int32   `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
	Name      string  `json:"name" bson:"name"`           // Product name for display
	ImageURL  string  `json:"image_url" bson:"image_url"` // Product image
}

// Order represents a customer order
type Order struct {
	ID        string       `json:"id" bson:"_id"`          // Using string ID for MongoDB compatibility
	UserID    string       `json:"user_id" bson:"user_id"` // Changed to string for flexibility
	Items     []OrderItem  `json:"items" bson:"items"`     // More structured items
	Status    OrderStatus  `json:"status" bson:"status"`   // Using typed status
	Total     float64      `json:"total" bson:"total"`
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" bson:"updated_at"`
	Shipping  ShippingInfo `json:"shipping" bson:"shipping"`
	Payment   PaymentInfo  `json:"payment" bson:"payment"`
}

// ShippingInfo contains shipping details
type ShippingInfo struct {
	Address    string `json:"address" bson:"address"`
	City       string `json:"city" bson:"city"`
	PostalCode string `json:"postal_code" bson:"postal_code"`
	Country    string `json:"country" bson:"country"`
	TrackingID string `json:"tracking_id" bson:"tracking_id"`
}

// PaymentInfo contains payment details
type PaymentInfo struct {
	Method        string    `json:"method" bson:"method"` // credit_card, paypal, etc.
	Amount        float64   `json:"amount" bson:"amount"`
	Status        string    `json:"status" bson:"status"` // paid, pending, failed
	PaidAt        time.Time `json:"paid_at" bson:"paid_at"`
	TransactionID string    `json:"transaction_id" bson:"transaction_id"`
}

// CreateOrderRequest represents the payload for creating an order
type CreateOrderRequest struct {
	Items    []OrderItem  `json:"items" binding:"required,min=1"`
	Shipping ShippingInfo `json:"shipping" binding:"required"`
	Payment  PaymentInfo  `json:"payment" binding:"required"`
}

// UpdateOrderStatusRequest represents the payload for updating order status
type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" binding:"required"`
}
