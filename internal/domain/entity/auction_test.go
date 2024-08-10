package entity

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/stretchr/testify/assert"
)

func TestNewAuction(t *testing.T) {
	credits := custom_type.NewBigInt(big.NewInt(1000))
	priceLimitPerCredit := custom_type.NewBigInt(big.NewInt(500))
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	createdAt := time.Now().Unix()

	auction, err := NewAuction(credits, priceLimitPerCredit, expiresAt, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, auction)
	assert.Equal(t, credits, auction.Credits)
	assert.Equal(t, priceLimitPerCredit, auction.PriceLimitPerCredit)
	assert.Equal(t, AuctionOngoing, auction.State)
	assert.NotZero(t, auction.ExpiresAt)
	assert.NotZero(t, auction.CreatedAt)
}

func TestAuction_Validate(t *testing.T) {
	createdAt := time.Now().Unix()
	expiresAt := time.Now().Add(-24 * time.Hour).Unix() // Past time
	auction := &Auction{
		Credits:             custom_type.NewBigInt(big.NewInt(1000)),
		PriceLimitPerCredit: custom_type.NewBigInt(big.NewInt(0)), // Invalid price limit
		ExpiresAt:           expiresAt,
		CreatedAt:           createdAt,
	}

	err := auction.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidAuction, err)

	// Correct the validation errors
	auction.PriceLimitPerCredit = custom_type.NewBigInt(big.NewInt(500))
	auction.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	err = auction.Validate()
	assert.Nil(t, err)
}

func TestAuctionStateTransition(t *testing.T) {
	createdAt := time.Now().Unix()
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	auction, _ := NewAuction(custom_type.NewBigInt(big.NewInt(1000)), custom_type.NewBigInt(big.NewInt(500)), expiresAt, createdAt)
	assert.Equal(t, AuctionOngoing, auction.State)

	// Transition to finished
	auction.State = AuctionFinished
	assert.Equal(t, AuctionFinished, auction.State)

	// Transition to cancelled
	auction.State = AuctionCancelled
	assert.Equal(t, AuctionCancelled, auction.State)
}

func TestAuctionExpiration(t *testing.T) {
	createdAt := time.Now().Unix()
	expiresAt := createdAt + 3600 // 1 hour later

	auction := &Auction{
		Id:                  1,
		Credits:             custom_type.NewBigInt(big.NewInt(1000)),
		PriceLimitPerCredit: custom_type.NewBigInt(big.NewInt(500)),
		State:               AuctionOngoing,
		ExpiresAt:           expiresAt,
		CreatedAt:           createdAt,
		UpdatedAt:           createdAt,
	}

	err := auction.Validate()
	assert.Nil(t, err)

	auction.ExpiresAt = createdAt - 3600 // 1 hour before creation time
	err = auction.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidAuction, err)
}
