package db

import (
	"fmt"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
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

func (r *OrderRepositorySqlite) FindOrdersByTimeRange(from int64, to int64) ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.db.Where("created_at BETWEEN ? AND ?", from, to).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
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

func (r *OrderRepositorySqlite) DeleteOrder(id uint) error {
	res := r.db.Delete(&entity.Order{}, "id = ?", id)
	if res.Error != nil {
		return fmt.Errorf("failed to delete order: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.ErrOrderNotFound
	}
	return nil
}
