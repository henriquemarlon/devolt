package auction_usecase

import (
	"github.com/devolthq/devolt/pkg/custom_type"
)

type FindAuctionOutputDTO struct {
	Id         uint                       `json:"id"`
	Credits    custom_type.BigInt         `json:"credits"`
	PriceLimit custom_type.BigInt         `json:"price_limit"`
	State      string                     `json:"state"`
	Bids       []*FindAuctionOutputSubDTO `json:"bids"`
	ExpiresAt  int64                      `json:"expires_at"`
	CreatedAt  int64                      `json:"created_at"`
	UpdatedAt  int64                      `json:"updated_at"`
}

type FindAuctionOutputSubDTO struct {
	Id         uint                `json:"id"`
	AuctionId  uint                `json:"auction_id"`
	Bidder     custom_type.Address `json:"bidder"`
	Credits    custom_type.BigInt  `json:"credits"`
	Price      custom_type.BigInt  `json:"price"`
	PriceLimit custom_type.BigInt  `json:"price_limit"`
	State      string              `json:"state"`
	ExpiresAt  int64               `json:"expires_at"`
	CreatedAt  int64               `json:"created_at"`
	UpdatedAt  int64               `json:"updated_at"`
}
