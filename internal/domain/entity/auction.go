package entity

import (
	"errors"
	"github.com/devolthq/devolt/pkg/custom_type"
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
	Id         uint               `json:"id" gorm:"primaryKey"`
	Credits    custom_type.BigInt `json:"credits" gorm:"type:bigint;not null"`
	PriceLimit custom_type.BigInt `json:"price_limit" gorm:"type:bigint;not null"`
	State      AuctionState       `json:"state" gorm:"type:text;not null"`
	Bids       []*Bid             `json:"bids" gorm:"foreignKey:AuctionId;constraint:OnDelete:CASCADE"`
	ExpiresAt  int64              `json:"expires_at" gorm:"not null"`
	CreatedAt  int64              `json:"created_at" gorm:"not null"`
	UpdatedAt  int64              `json:"updated_at" gorm:"default:0"`
}

func NewAuction(credits custom_type.BigInt, priceLimit custom_type.BigInt, expiresAt int64, createdAt int64) (*Auction, error) {
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
	if a.Credits.Int == nil || a.PriceLimit.Int == nil || a.ExpiresAt == 0 || a.CreatedAt == 0 || a.CreatedAt >= a.ExpiresAt {
		return ErrInvalidAuction
	}
	return nil
}
