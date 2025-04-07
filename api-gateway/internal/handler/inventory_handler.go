package handler

import (
	"api-gateway/pkg/client"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	Client *client.InventoryClient
}

func NewInventoryHandler(client *client.InventoryClient) *InventoryHandler {
	return &InventoryHandler{Client: client}
}

func (h *InventoryHandler) ForwardToInventory(c *gin.Context) {
	// Extract request details
	method := c.Request.Method
	path := c.Param("path") // Assuming path params for endpoint like `/inventory/:path`
	body, _ := ioutil.ReadAll(c.Request.Body)

	// Forward request to inventory-service
	response, err := h.Client.ForwardRequest(method, path, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond back to client
	c.Data(http.StatusOK, "application/json", response)
}
