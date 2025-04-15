package handler

import (
	"api-gateway/pkg/client"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	Client *client.InventoryClient
}

func NewInventoryHandler(client *client.InventoryClient) *InventoryHandler {
	return &InventoryHandler{Client: client}
}

func (h *InventoryHandler) ForwardToInventory(c *gin.Context) {
	method := c.Request.Method
	path := strings.TrimPrefix(c.Param("path"), "/")
	body, _ := ioutil.ReadAll(c.Request.Body)

	response, err := h.Client.ForwardRequest(method, path, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", response)
}
