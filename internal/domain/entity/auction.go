package entity

import (
	"errors"
	"math/big"
)

var (
	ErrExpired         = errors.New("auction expired")
	ErrAuctionNotFound = errors.New("auction not found")
	ErrInvalidAuction  = errors.New("invalid price")
)

type AuctionRepository interface {
	DeleteAuction(id uint) error
	FindActiveAuction() (*Auction, error)
	FindAllAuctions() ([]*Auction, error)
	FindAuctionById(id uint) (*Auction, error)
	CreateAuction(auction *Auction) (*Auction, error)
	UpdateAuction(auction *Auction) (*Auction, error)
}

type AuctionState string

const (
	AuctionOngoing   AuctionState = "ongoing"
	AuctionFinished  AuctionState = "finished"
	AuctionCancelled AuctionState = "cancelled"
)

type Auction struct {
	Id         uint         `json:"id" gorm:"primaryKey"`
	Credits    *big.Int     `json:"credits" gorm:"type:bigint;not null"`
	PriceLimit *big.Int     `json:"price_limit" gorm:"type:bigint;not null"`
	State      AuctionState `json:"state" gorm:"type:text;not null"`
	Bids       []*Bid       `json:"bids" gorm:"foreignKey:AuctionId;constraint:OnDelete:CASCADE"`
	ExpiresAt  int64        `json:"expires_at" gorm:"not null"`
	CreatedAt  int64        `json:"created_at" gorm:"not null"`
	UpdatedAt  int64        `json:"updated_at" gorm:"default:0"`
}

func NewAuction(credits *big.Int, priceLimit *big.Int, expiresAt int64, createdAt int64) (*Auction, error) {
	auction := &Auction{
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      AuctionOngoing,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
	}
	if err := auction.Validate(); err != nil {
		return nil, err
	}
	return auction, nil
}

func (a *Auction) Validate() error {
	if a.Credits == nil || a.PriceLimit == nil {
		return ErrInvalidAuction
	}
	if a.ExpiresAt <= a.CreatedAt {
		return ErrInvalidAuction
	}
	return nil
}
