package entity

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidBid  = errors.New("invalid bid")
	ErrBidNotFound = errors.New("bid not found")
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

type BidState string

const (
	BidStatePending  BidState = "pending"
	BidStateAccepted BidState = "accepted"
	BidStateRejected BidState = "rejected"
	BidStateExpired  BidState = "expired"
)

type Bid struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	AuctionId uint           `json:"auction_id" gorm:"not null;index"`
	Bidder    common.Address `json:"bidder" gorm:"not null"`
	Credits   *big.Int       `json:"credits" gorm:"type:bigint;not null"`
	Price     *big.Int       `json:"price" gorm:"type:bigint;not null"`
	State     BidState       `json:"state" gorm:"not null"`
	CreatedAt int64          `json:"created_at" gorm:"not null"`
	UpdatedAt int64          `json:"updated_at" gorm:"default:0"`
}

func NewBid(auctionId uint, bidder common.Address, credits *big.Int, price *big.Int, createdAt int64) (*Bid, error) {
	bid := &Bid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Credits:   credits,
		Price:     price,
		State:     BidStatePending,
		CreatedAt: createdAt,
	}
	if err := bid.Validate(); err != nil {
		return nil, err
	}
	return bid, nil
}

func (b *Bid) Validate() error {
	if b.AuctionId == 0 || b.Bidder == (common.Address{}) || b.Credits.Cmp(big.NewInt(0)) <= 0 || b.Price.Cmp(big.NewInt(0)) <= 0 {
		return ErrInvalidBid
	}
	return nil
}
