package entity

import (
	"testing"
	"math/big"
	"time"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewBid_Success(t *testing.T) {
	auctionId := uint(1)
	bidder := custom_type.Address{Address: common.HexToAddress("0x123")}
	credits := custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	createdAt := time.Now().Unix()

	bid, err := NewBid(auctionId, bidder, credits, pricePerCredit, createdAt)
	assert.NoError(t, err)
	assert.NotNil(t, bid)
	assert.Equal(t, auctionId, bid.AuctionId)
	assert.Equal(t, bidder, bid.Bidder)
	assert.Equal(t, credits, bid.Credits)
	assert.Equal(t, pricePerCredit, bid.PricePerCredit)
	assert.Equal(t, BidStatePending, bid.State)
	assert.Equal(t, createdAt, bid.CreatedAt)
}

func TestNewBid_Fail_InvalidBid(t *testing.T) {
	auctionId := uint(0)
	bidder := custom_type.Address{Address: common.HexToAddress("0x123")}
	credits := custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	createdAt := time.Now().Unix()

	bid, err := NewBid(auctionId, bidder, credits, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, bid)
	assert.Equal(t, ErrInvalidBid, err)

	auctionId = uint(1)
	bidder = custom_type.Address{Address: common.Address{}}

	bid, err = NewBid(auctionId, bidder, credits, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, bid)
	assert.Equal(t, ErrInvalidBid, err)

	bidder = custom_type.Address{Address: common.HexToAddress("0x123")}
	credits = custom_type.BigInt{Int: big.NewInt(0)}

	bid, err = NewBid(auctionId, bidder, credits, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, bid)
	assert.Equal(t, ErrInvalidBid, err)

	credits = custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit = custom_type.BigInt{Int: big.NewInt(0)}

	bid, err = NewBid(auctionId, bidder, credits, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, bid)
	assert.Equal(t, ErrInvalidBid, err)
}
