package entity

import (
	"errors"
	"math/big"
)

var (
	ErrInvalidOrder  = errors.New("invalid order")
	ErrOrderNotFound = errors.New("order not found")
)

type OrderRepository interface {
	CreateOrder(order *Order) (*Order, error)
	FindOrderById(id uint) (*Order, error)
	FindOrdersByUser(buyer string) ([]*Order, error)
	FindAllOrders() ([]*Order, error)
	UpdateOrder(order *Order) (*Order, error)
	DeleteOrder(id uint) error
}

type Order struct {
	Id             uint     `json:"id" gorm:"primaryKey"`
	Buyer          string   `json:"buyer" gorm:"type:text;not null"`
	Credits        *big.Int `json:"credits" gorm:"type:bigint;not null"`
	StationId      string   `json:"station_id" gorm:"not null"`
	PricePerCredit *big.Int `json:"price_per_credit" gorm:"type:bigint;not null"`
	CreatedAt      int64    `json:"created_at" gorm:"not null"`
	UpdatedAt      int64    `json:"updated_at" gorm:"default:0"`
}

func NewOrder(buyer string, credits *big.Int, stationId string, pricePerCredit *big.Int, createdAt int64) (*Order, error) {
	order := &Order{
		Buyer:          buyer,
		Credits:        credits,
		StationId:      stationId,
		PricePerCredit: pricePerCredit,
		CreatedAt:      createdAt,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.Buyer == "" || o.Credits.Cmp(big.NewInt(0)) <= 0 || o.PricePerCredit.Cmp(big.NewInt(0)) <= 0 || o.StationId == "" {
		return ErrInvalidOrder
	}
	return nil
}
