package entity

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	buyer := common.HexToAddress("0x123").String()
	credits := big.NewInt(1000)
	stationId := "station-1"
	pricePerCredit := big.NewInt(50)
	createdAt := time.Now().Unix()

	order, err := NewOrder(buyer, credits, stationId, pricePerCredit, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, buyer, order.Buyer)
	assert.Equal(t, credits, order.Credits)
	assert.Equal(t, stationId, order.StationId)
	assert.Equal(t, pricePerCredit, order.PricePerCredit)
	assert.NotZero(t, order.CreatedAt)
}

func TestOrder_Validate(t *testing.T) {
	buyer := common.HexToAddress("0x123").String()
	createdAt := time.Now().Unix()

	// Invalid credits
	order := &Order{
		Buyer:          buyer,
		Credits:        big.NewInt(-1),
		StationId:      "station-2",
		PricePerCredit: big.NewInt(50),
		CreatedAt:      createdAt,
	}
	err := order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid price per credit
	order.Credits = big.NewInt(1000)
	order.PricePerCredit = big.NewInt(-1)
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid buyer
	order.PricePerCredit = big.NewInt(50)
	order.Buyer = ""
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Invalid station id
	order.Buyer = common.HexToAddress("0x123").String()
	order.StationId = ""
	err = order.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidOrder, err)

	// Valid order
	order.StationId = "station-3"
	err = order.Validate()
	assert.Nil(t, err)
}
