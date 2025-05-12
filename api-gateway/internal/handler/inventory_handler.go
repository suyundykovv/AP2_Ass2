package handler

import (
	"api-gateway/pkg/client"
	"net/http"

	// Import the generated gRPC protobuf package

	"github.com/gin-gonic/gin"
	pb "github.com/suyundykovv/protos/gen/go/inventory/v1"
)

// OrderHandler handles HTTP requests for orders
type InventoryHandler struct {
	InventoryClient *client.InventoryClient
}

// NewOrderHandler creates a new OrderHandler
func NewInventoryHandler(InventoryClient *client.InventoryClient) *InventoryHandler {
	return &InventoryHandler{
		InventoryClient: InventoryClient,
	}
}

func (h *InventoryHandler) CreateProduct(c *gin.Context) {
	var req pb.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product, err := h.InventoryClient.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateOrder handles the request to update an order by ID
func (h *InventoryHandler) Updateproduct(c *gin.Context) {
	id := c.Param("id") // Берём ID из пути

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var req pb.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Устанавливаем ID заказа из параметра пути в запрос
	req.Id = id

	order, err := h.InventoryClient.UpdateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetOrder handles the request to get an order by ID
// GetOrder handles the request to get an order by ID
func (h *InventoryHandler) GetProduct(c *gin.Context) {
	id := c.Param("id") // <-- берём из пути

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	req := &pb.GetProductRequest{Id: id}
	order, err := h.InventoryClient.GetProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving order: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
