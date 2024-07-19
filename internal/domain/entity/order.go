package entity

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type OrderRepository interface {
	CreateOrder(order *Order) (*Order, error)
	FindOrderById(id uint) (*Order, error)
	FindOrdersByUser(buyer common.Address) ([]*Order, error)
	FindAllOrders() ([]*Order, error)
	UpdateOrder(order *Order) (*Order, error)
	DeleteOrder(id uint) error
}

type Order struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	Buyer          common.Address `json:"buyer" gorm:"type:text;not null"`
	Credits        *big.Int       `json:"credits" gorm:"type:bigint;not null"`
	StationId      string         `json:"station_id" gorm:"not null"`
	PricePerCredit *big.Int       `json:"price_per_credit" gorm:"type:bigint;not null"`
	CreatedAt      int64          `json:"created_at" gorm:"not null"`
	UpdatedAt      int64          `json:"updated_at" gorm:"default:0"`
}

func NewOrder(buyer common.Address, credits *big.Int, stationId string, pricePerCredit *big.Int, createdAt int64) *Order {
	return &Order{
		Buyer:          buyer,
		Credits:        credits,
		StationId:      stationId,
		PricePerCredit: pricePerCredit,
		CreatedAt:      createdAt,
	}
}
