package entity

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewBid(t *testing.T) {
	auctionId := uint(1)
	bidder := custom_type.NewAddress(common.HexToAddress("0x123"))
	credits := custom_type.NewBigInt(big.NewInt(1000))
	price := custom_type.NewBigInt(big.NewInt(500))
	createdAt := time.Now().Unix()

	bid, err := NewBid(auctionId, bidder, credits, price, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, bid)
	assert.Equal(t, auctionId, bid.AuctionId)
	assert.Equal(t, bidder, bid.Bidder)
	assert.Equal(t, credits, bid.Credits)
	assert.Equal(t, price, bid.PricePerCredit)
	assert.Equal(t, BidStatePending, bid.State)
	assert.NotZero(t, bid.CreatedAt)
}

func TestBid_Validate(t *testing.T) {
	auctionId := uint(1)
	bidder := common.HexToAddress("0x123")
	createdAt := time.Now().Unix()

	// Invalid credits
	bid := &Bid{
		AuctionId:      auctionId,
		Bidder:         custom_type.NewAddress(bidder),
		Credits:        custom_type.NewBigInt(big.NewInt(-1)),
		PricePerCredit: custom_type.NewBigInt(big.NewInt(500)),
		State:          BidStatePending,
		CreatedAt:      createdAt,
	}
	err := bid.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidBid, err)

	// Invalid price
	bid.Credits = custom_type.NewBigInt(big.NewInt(1000))
	bid.PricePerCredit = custom_type.NewBigInt(big.NewInt(-1))
	err = bid.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidBid, err)

	// Invalid bidder
	bid.PricePerCredit = custom_type.NewBigInt(big.NewInt(500))
	bid.Bidder = custom_type.NewAddress(common.Address{})
	err = bid.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidBid, err)

	// Valid bid
	bid.Bidder = custom_type.NewAddress(common.HexToAddress("0x123"))
	err = bid.Validate()
	assert.Nil(t, err)
}
