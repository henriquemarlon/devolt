package db

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"gorm.io/gorm"
)

type OrderRepositorySqlite struct {
	db *gorm.DB
}

func NewOrderRepositorySqlite(db *gorm.DB) *OrderRepositorySqlite {
	return &OrderRepositorySqlite{
		db: db,
	}
}

func (r *OrderRepositorySqlite) CreateOrder(input *entity.Order) (*entity.Order, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *OrderRepositorySqlite) FindAllOrders() ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepositorySqlite) FindOrderById(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepositorySqlite) FindOrdersByUser(buyer custom_type.Address) ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Where("buyer = ?", buyer).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepositorySqlite) UpdateOrder(order *entity.Order) (*entity.Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepositorySqlite) DeleteOrder(id uint) error {
	err := r.db.Delete(&entity.Order{}, "order_id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
