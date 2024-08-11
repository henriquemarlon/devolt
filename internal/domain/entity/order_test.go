package entity

import (
	"math/big"
	"testing"
	"time"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	buyer := common.HexToAddress("0x123")
	credits := big.NewInt(1000)
	stationId := uint(1)
	pricePerCredit := big.NewInt(50)
	createdAt := time.Now().Unix()

	order, err := NewOrder(custom_type.NewAddress(buyer), custom_type.NewBigInt(credits), stationId, pricePerCredit, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, buyer, order.Buyer.Address)
	assert.Equal(t, credits, order.Credits.Int)
	assert.Equal(t, stationId, order.StationId)
	assert.Equal(t, pricePerCredit, order.PricePerCredit.Int)
	assert.NotZero(t, order.CreatedAt)
}

func TestOrder_Validate(t *testing.T) {
	buyer := common.HexToAddress("0x123")
	createdAt := time.Now().Unix()

	// Invalid credits
	order := &Order{
		Buyer:          custom_type.NewAddress(buyer),
		Credits:        custom_type.NewBigInt(big.NewInt(-1)),
		StationId:      2,
		PricePerCredit: custom_type.NewBigInt(big.NewInt(50)),
		CreatedAt:      createdAt,
	}
	err := order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid price per credit
	order.Credits = custom_type.NewBigInt(big.NewInt(1000))
	order.PricePerCredit = custom_type.NewBigInt(big.NewInt(-1))
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid buyer
	order.PricePerCredit = custom_type.NewBigInt(big.NewInt(50))
	order.Buyer = custom_type.NewAddress(common.Address{})
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid station id
	order.Buyer = custom_type.NewAddress(common.HexToAddress("0x123"))
	order.StationId = 0
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Valid order
	order.StationId = 1
	err = order.Validate()
	assert.Nil(t, err)
}
