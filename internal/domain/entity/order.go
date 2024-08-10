package entity

import (
	"errors"
	"math/big"

	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidOrder  = errors.New("invalid order")
	ErrOrderNotFound = errors.New("order not found")
)

type OrderRepository interface {
	CreateOrder(order *Order) (*Order, error)
	FindOrderById(id uint) (*Order, error)
	FindOrdersByUser(buyer custom_type.Address) ([]*Order, error)
	FindAllOrders() ([]*Order, error)
	FindOrdersByTimeRange(from int64, to int64) ([]*Order, error)
	UpdateOrder(order *Order) (*Order, error)
	DeleteOrder(id uint) error
}

type Order struct {
	Id             uint                `json:"id" gorm:"primaryKey"`
	Buyer          custom_type.Address `json:"buyer,omitempty" gorm:"type:text;not null"`
	Credits        custom_type.BigInt  `json:"credits,omitempty" gorm:"type:bigint;not null"`
	StationId      uint                `json:"station_id,omitempty" gorm:"not null"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit,omitempty" gorm:"type:bigint;not null"`
	CreatedAt      int64               `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt      int64               `json:"updated_at,omitempty" gorm:"default:0"`
}

func NewOrder(buyer custom_type.Address, credits custom_type.BigInt, stationId uint, pricePerCredit *big.Int, createdAt int64) (*Order, error) {
	order := &Order{
		Buyer:          buyer,
		Credits:        credits,
		StationId:      stationId,
		PricePerCredit: custom_type.NewBigInt(pricePerCredit),
		CreatedAt:      createdAt,
	}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.Buyer.Address == (common.Address{}) || o.Credits.Cmp(big.NewInt(0)) <= 0 || o.PricePerCredit.Cmp(big.NewInt(0)) <= 0 || o.StationId == 0 {
		return ErrInvalidOrder
	}
	return nil
}
