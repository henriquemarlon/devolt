package station_usecase

import "github.com/devolthq/devolt/pkg/custom_type"

type FindStationOutputDTO struct {
	Id             uint                       `json:"id"`
	Consumption    custom_type.BigInt         `json:"consumption"`
	Owner          custom_type.Address        `json:"owner"`
	PricePerCredit custom_type.BigInt         `json:"price_per_credit"`
	State          string                     `json:"state"`
	Orders         []*FindStationOutputSubDTO `json:"orders"`
	Latitude       float64                    `json:"latitude"`
	Longitude      float64                    `json:"longitude"`
	CreatedAt      int64                      `json:"created_at"`
	UpdatedAt      int64                      `json:"updated_at"`
}

type FindStationOutputSubDTO struct {
	Id             uint                `json:"id"`
	Buyer          custom_type.Address `json:"buyer"`
	Credits        custom_type.BigInt  `json:"credits"`
	StationId      uint                `json:"station_id"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      int64               `json:"updated_at"`
}
