package entity

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BidRepository interface {
	CreateBid(bid *Bid) (*Bid, error)
	FindBidsByState(auctionId uint, state string) ([]*Bid, error)
	FindBidById(id uint) (*Bid, error)
	FindBidsByAuctionId(id uint) ([]*Bid, error)
	FindAllBids() ([]*Bid, error)
	UpdateBid(bid *Bid) (*Bid, error)
	DeleteBid(id uint) error
}

type Bid struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	AuctionId uint           `json:"auction_id" gorm:"not null;index"`
	Bidder    common.Address `json:"bidder" gorm:"not null"`
	Credits   *big.Int       `json:"credits" gorm:"type:bigint;not null"`
	Price     *big.Int       `json:"price" gorm:"type:bigint;not null"`
	State     string         `json:"state" gorm:"not null"`
	CreatedAt int64          `json:"created_at" gorm:"not null"`
	UpdatedAt int64          `json:"updated_at" gorm:"default:0"`
}

func NewBid(auctionId uint, bidder common.Address, credits *big.Int, price *big.Int, state string, createdAt int64) *Bid {
	return &Bid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Credits:   credits,
		Price:     price,
		State:     state,
		CreatedAt: createdAt,
	}
}
