package entity

import (
	"errors"
	"math/big"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
)

var (
	ErrExpired         = errors.New("auction expired")
	ErrAuctionNotFound = errors.New("auction not found")
	ErrInvalidAuction  = errors.New("invalid auction")
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
	AuctionOngoing          AuctionState = "ongoing"
	AuctionFinished         AuctionState = "finished"
	AuctionCancelled        AuctionState = "cancelled"
)

type Auction struct {
	Id                  uint               `json:"id" gorm:"primaryKey"`
	RequiredCredits     custom_type.BigInt `json:"required_credits,omitempty" gorm:"type:bigint;not null"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit,omitempty" gorm:"type:bigint;not null"`
	State               AuctionState       `json:"state,omitempty" gorm:"type:text;not null"`
	Bids                []*Bid             `json:"bids,omitempty" gorm:"foreignKey:AuctionId;constraint:OnDelete:CASCADE"`
	ExpiresAt           int64              `json:"expires_at,omitempty" gorm:"not null"`
	CreatedAt           int64              `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt           int64              `json:"updated_at,omitempty" gorm:"default:0"`
}

func NewAuction(requiredCredits custom_type.BigInt, priceLimitPerCredit custom_type.BigInt, expiresAt int64, createdAt int64) (*Auction, error) {
	auction := &Auction{
		RequiredCredits:     requiredCredits,
		PriceLimitPerCredit: priceLimitPerCredit,
		State:               AuctionOngoing,
		ExpiresAt:           expiresAt,
		CreatedAt:           createdAt,
	}
	if err := auction.Validate(); err != nil {
		return nil, err
	}
	return auction, nil
}

func (a *Auction) Validate() error {
	if a.RequiredCredits.Cmp(big.NewInt(0)) <= 0 || a.PriceLimitPerCredit.Cmp(big.NewInt(0)) <= 0 || a.ExpiresAt == 0 || a.CreatedAt == 0 || a.CreatedAt >= a.ExpiresAt {
		return ErrInvalidAuction
	}
	return nil
}
