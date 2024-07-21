package order_usecase

import (
	"math/big"
)

type FindOrderOutputDTO struct {
	Id             uint     `json:"id"`
	Buyer          string   `json:"buyer"`
	Credits        *big.Int `json:"credits"`
	StationId      string   `json:"station_id"`
	PricePerCredit *big.Int `json:"price_per_credit"`
	CreatedAt      int64    `json:"created_at"`
	UpdatedAt      int64    `json:"updated_at"`
}
