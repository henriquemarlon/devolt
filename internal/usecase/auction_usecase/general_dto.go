package auction_usecase

import (
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
)

type FindAuctionOutputDTO struct {
	Id                  uint                       `json:"id"`
	RequiredCredits             custom_type.BigInt         `json:"required_credits"`
	PriceLimitPerCredit custom_type.BigInt         `json:"price_limit_per_credit"`
	State               string                     `json:"state"`
	Bids                []*FindAuctionOutputSubDTO `json:"bids"`
	ExpiresAt           int64                      `json:"expires_at"`
	CreatedAt           int64                      `json:"created_at"`
	UpdatedAt           int64                      `json:"updated_at"`
}

type FindAuctionOutputSubDTO struct {
	Id             uint                `json:"id"`
	AuctionId      uint                `json:"auction_id"`
	Bidder         custom_type.Address `json:"bidder"`
	Credits        custom_type.BigInt  `json:"credits"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	State          string              `json:"state"`
	ExpiresAt      int64               `json:"expires_at"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      int64               `json:"updated_at"`
}
