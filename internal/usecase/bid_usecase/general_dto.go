package bid_usecase

import (
	"math/big"
)

type FindBidOutputDTO struct {
	Id        uint     `json:"id"`
	AuctionId uint     `json:"auction_id"`
	Bidder    string   `json:"bidder"`
	Credits   *big.Int `json:"credits"`
	Price     *big.Int `json:"price"`
	State     string   `json:"state"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}
