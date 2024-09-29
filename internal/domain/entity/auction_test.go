package entity

import (
	"testing"
	"time"
	"math/big"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/stretchr/testify/assert"
)

func TestNewAuction(t *testing.T) {
	requiredCredits := custom_type.BigInt{Int: big.NewInt(100)}
	priceLimitPerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	createdAt := time.Now().Unix()
	expiresAt := createdAt + 3600

	auction, err := NewAuction(requiredCredits, priceLimitPerCredit, expiresAt, createdAt)
	assert.NoError(t, err)
	assert.NotNil(t, auction)
	assert.Equal(t, requiredCredits, auction.RequiredCredits)
	assert.Equal(t, priceLimitPerCredit, auction.PriceLimitPerCredit)
	assert.Equal(t, AuctionOngoing, auction.State)
	assert.Equal(t, expiresAt, auction.ExpiresAt)
	assert.Equal(t, createdAt, auction.CreatedAt)
}

func TestNewAuction_Fail_InvalidAuction(t *testing.T) {
	requiredCredits := custom_type.BigInt{Int: big.NewInt(0)}
	priceLimitPerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	createdAt := time.Now().Unix()
	expiresAt := createdAt + 3600

	auction, err := NewAuction(requiredCredits, priceLimitPerCredit, expiresAt, createdAt)
	assert.Error(t, err)
	assert.Nil(t, auction)
	assert.Equal(t, ErrInvalidAuction, err)

	requiredCredits = custom_type.BigInt{Int: big.NewInt(100)}
	priceLimitPerCredit = custom_type.BigInt{Int: big.NewInt(0)}

	auction, err = NewAuction(requiredCredits, priceLimitPerCredit, expiresAt, createdAt)
	assert.Error(t, err)
	assert.Nil(t, auction)
	assert.Equal(t, ErrInvalidAuction, err)

	requiredCredits = custom_type.BigInt{Int: big.NewInt(100)}
	priceLimitPerCredit = custom_type.BigInt{Int: big.NewInt(50)}
	expiresAt = createdAt

	auction, err = NewAuction(requiredCredits, priceLimitPerCredit, expiresAt, createdAt)
	assert.Error(t, err)
	assert.Nil(t, auction)
	assert.Equal(t, ErrInvalidAuction, err)
}
