package handler

import (
	"api-gateway/pkg/client"
	"net/http"

	// Import the generated gRPC protobuf package

	"github.com/gin-gonic/gin"
	pb "github.com/suyundykovv/protos/gen/go/order/v1"
)

// OrderHandler handles HTTP requests for orders
type OrderHandler struct {
	orderClient *client.OrderClient
}

// NewOrderHandler creates a new OrderHandler
func NewOrderHandler(orderClient *client.OrderClient) *OrderHandler {
	return &OrderHandler{
		orderClient: orderClient,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req pb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order, err := h.orderClient.CreateOrder(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder handles the request to update an order by ID
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id") // Берём ID из пути

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID is required"})
		return
	}

	var req pb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Устанавливаем ID заказа из параметра пути в запрос
	req.Id = id

	order, err := h.orderClient.UpdateOrderStatus(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrder handles the request to get an order by ID
// GetOrder handles the request to get an order by ID
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id") // <-- берём из пути

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order ID is required"})
		return
	}

	req := &pb.GetOrderRequest{Id: id}
	order, err := h.orderClient.GetOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
