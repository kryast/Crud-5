package services

import "github.com/kryast/Crud-5.git/repositories"

type OrderItemService interface {
}

type orderItemService struct {
	repo repositories.OrderItemRepository
}

func NewOrderItemService(r repositories.OrderItemRepository) OrderItemService {
	return &orderItemService{r}
}
