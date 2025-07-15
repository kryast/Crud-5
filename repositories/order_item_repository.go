package repositories

import (
	"github.com/kryast/Crud-5.git/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(item *models.OrderItem) error
	FindAll() ([]models.OrderItem, error)
	FindByID(id uint) (models.OrderItem, error)
}

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepo{db}
}

func (r *orderItemRepo) Create(item *models.OrderItem) error {
	return r.db.Create(item).Error
}

func (r *orderItemRepo) FindAll() ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := r.db.Preload("Order").Preload("Product").Find(&items).Error
	return items, err
}

func (r *orderItemRepo) FindByID(id uint) (models.OrderItem, error) {
	var item models.OrderItem
	err := r.db.Preload("Order").Preload("Product").First(&item, id).Error
	return item, err
}
