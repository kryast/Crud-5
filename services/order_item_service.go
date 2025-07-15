package services

import (
	"github.com/kryast/Crud-5.git/models"
	"github.com/kryast/Crud-5.git/repositories"
)

type OrderItemService interface {
	Create(item *models.OrderItem) error
}

type orderItemService struct {
	repo repositories.OrderItemRepository
}

func NewOrderItemService(r repositories.OrderItemRepository) OrderItemService {
	return &orderItemService{r}
}

func (s *orderItemService) Create(item *models.OrderItem) error {
	item.Subtotal = float64(item.Quantity) * item.Product.Price
	return s.repo.Create(item)
}
