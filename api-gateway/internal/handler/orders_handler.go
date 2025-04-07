package handler

import (
	"api-gateway/pkg/client"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {
	Client *client.OrdersClient
}

func NewOrdersHandler(client *client.OrdersClient) *OrdersHandler {
	return &OrdersHandler{Client: client}
}

func (h *OrdersHandler) ForwardToOrders(c *gin.Context) {
	method := c.Request.Method
	path := c.Param("path")
	body, _ := ioutil.ReadAll(c.Request.Body)

	response, err := h.Client.ForwardRequest(method, path, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", response)
}
