package entity

import (
	"math/big"
)

type AuctionRepository interface {
	CreateAuction(auction *Auction) (*Auction, error)
	FindActiveAuction() (*Auction, error)
	FindAuctionById(id uint) (*Auction, error)
	FindAllAuctions() ([]*Auction, error)
	UpdateAuction(auction *Auction) (*Auction, error)
	DeleteAuction(id uint) error
}

type Auction struct {
	Id         uint     `json:"id" gorm:"primaryKey"`
	Credits    *big.Int `json:"credits" gorm:"type:bigint;not null"`
	PriceLimit *big.Int `json:"price_limit" gorm:"type:bigint;not null"`
	State      string   `json:"state" gorm:"type:text;default:'ongoing'"`
	Bids       []*Bid   `json:"bids" gorm:"foreignKey:AuctionId;constraint:OnDelete:CASCADE"`
	ExpiresAt  int64    `json:"expires_at" gorm:"not null"`
	CreatedAt  int64    `json:"created_at" gorm:"not null"`
	UpdatedAt  int64    `json:"updated_at" gorm:"default:0"`
}

func NewAuction(credits *big.Int, priceLimit *big.Int, state string, expires_at int64, createdAt int64) *Auction {
	return &Auction{
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      state,
		ExpiresAt:  expires_at,
		CreatedAt:  createdAt,
	}
}
