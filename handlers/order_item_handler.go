package handlers

import "github.com/kryast/Crud-5.git/services"

type OrderItemHandler struct {
	service services.OrderItemService
}

func NewOrderItemHandler(s services.OrderItemService) *OrderItemHandler {
	return &OrderItemHandler{s}
}
