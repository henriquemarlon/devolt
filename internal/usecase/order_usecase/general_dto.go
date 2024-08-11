package order_usecase

import "github.com/Mugen-Builders/devolt/pkg/custom_type"

type FindOrderOutputDTO struct {
	Id             uint                `json:"id"`
	Buyer          custom_type.Address `json:"buyer"`
	Credits        custom_type.BigInt  `json:"credits"`
	StationId      uint                `json:"station_id"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      int64               `json:"updated_at"`
}
