package station_usecase

import (
	"math/big"
)

type FindStationOutputDTO struct {
	Id             string                     `json:"id"`
	Consumption    *big.Int                   `json:"consumption"`
	Owner          string                     `json:"owner"`
	PricePerCredit *big.Int                   `json:"price_per_credit"`
	State          string                     `json:"state"`
	Orders         []*FindStationOutputSubDTO `json:"orders"`
	Latitude       float64                    `json:"latitude"`
	Longitude      float64                    `json:"longitude"`
	CreatedAt      int64                      `json:"created_at"`
	UpdatedAt      int64                      `json:"updated_at"`
}

type FindStationOutputSubDTO struct {
	Id             uint     `json:"id"`
	Buyer          string   `json:"buyer"`
	Credits        *big.Int `json:"credits"`
	StationId      string   `json:"station_id"`
	PricePerCredit *big.Int `json:"price_per_credit"`
	CreatedAt      int64    `json:"created_at"`
	UpdatedAt      int64    `json:"updated_at"`
}
