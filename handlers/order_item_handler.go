package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-5.git/models"
	"github.com/kryast/Crud-5.git/services"
)

type OrderItemHandler struct {
	service services.OrderItemService
}

func NewOrderItemHandler(s services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{s}
}

func (h *OrderItemHandler) Create(c *gin.Context) {
	var item models.OrderItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, item)
}
