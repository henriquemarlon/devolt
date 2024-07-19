package auction_usecase

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type FindAuctionOutputDTO struct {
	Id         uint                       `json:"id"`
	Credits    *big.Int                   `json:"credits"`
	PriceLimit *big.Int                   `json:"price_limit"`
	State      string                     `json:"state"`
	Bids       []*FindAuctionOutputSubDTO `json:"bids"`
	ExpiresAt  int64                      `json:"expires_at"`
	CreatedAt  int64                      `json:"created_at"`
	UpdatedAt  int64                      `json:"updated_at"`
}

type FindAuctionOutputSubDTO struct {
	Id         uint           `json:"id"`
	AuctionId  uint           `json:"auction_id"`
	Bidder     common.Address `json:"bidder"`
	Credits    *big.Int       `json:"credits"`
	Price      *big.Int       `json:"price"`
	PriceLimit *big.Int       `json:"price_limit"`
	State      string         `json:"state"`
	ExpiresAt  int64          `json:"expires_at"`
	CreatedAt  int64          `json:"created_at"`
	UpdatedAt  int64          `json:"updated_at"`
}
