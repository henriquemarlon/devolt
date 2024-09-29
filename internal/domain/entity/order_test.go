package entity

import (
	"testing"
	"math/big"
	"time"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder_Success(t *testing.T) {
	buyer := custom_type.Address{Address: common.HexToAddress("0x123")}
	credits := custom_type.BigInt{Int: big.NewInt(100)}
	stationId := uint(1)
	pricePerCredit := big.NewInt(50)
	createdAt := time.Now().Unix()

	order, err := NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, buyer, order.Buyer)
	assert.Equal(t, credits, order.Credits)
	assert.Equal(t, stationId, order.StationId)
	assert.Equal(t, custom_type.NewBigInt(pricePerCredit), order.PricePerCredit)
	assert.Equal(t, createdAt, order.CreatedAt)
}

func TestNewOrder_Fail_InvalidOrder(t *testing.T) {
	buyer := custom_type.Address{Address: common.HexToAddress("0x0")}
	credits := custom_type.BigInt{Int: big.NewInt(100)}
	stationId := uint(1)
	pricePerCredit := big.NewInt(50)
	createdAt := time.Now().Unix()

	order, err := NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, ErrInvalidOrder, err)

	buyer = custom_type.Address{Address: common.HexToAddress("0x123")}
	credits = custom_type.BigInt{Int: big.NewInt(0)}

	order, err = NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, ErrInvalidOrder, err)

	credits = custom_type.BigInt{Int: big.NewInt(100)}
	stationId = uint(0)

	order, err = NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, ErrInvalidOrder, err)

	stationId = uint(1)
	pricePerCredit = big.NewInt(0)

	order, err = NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, ErrInvalidOrder, err)
}
