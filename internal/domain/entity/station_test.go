package entity

import (
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestNewStation(t *testing.T) {
	id := "station-1"
	owner := custom_type.NewAddress(common.HexToAddress("0x123"))
	consumption := custom_type.NewBigInt(big.NewInt(1000))
	pricePerCredit := custom_type.NewBigInt(big.NewInt(50))
	latitude := 40.7128
	longitude := -74.0060
	createdAt := time.Now().Unix()

	station, err := NewStation(id, owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, station)
	assert.Equal(t, id, station.Id)
	assert.Equal(t, owner, station.Owner)
	assert.Equal(t, pricePerCredit, station.PricePerCredit)
	assert.Equal(t, latitude, station.Latitude)
	assert.Equal(t, longitude, station.Longitude)
	assert.NotZero(t, station.CreatedAt)
}

func TestStation_Validate(t *testing.T) {
	owner := custom_type.NewAddress(common.HexToAddress("0x123"))
	pricePerCredit := custom_type.NewBigInt(big.NewInt(50))
	createdAt := time.Now().Unix()

	// Invalid ID
	station := &Station{
		Id:             "",
		Owner:          owner,
		PricePerCredit: pricePerCredit,
		Latitude:       40.7128,
		Longitude:      -74.0060,
		CreatedAt:      createdAt,
	}
	err := station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid owner
	station.Id = "station-3"
	station.Owner = custom_type.NewAddress(common.Address{})
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid price per credit
	station.Owner = owner
	station.PricePerCredit.Int = nil
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid latitude
	station.PricePerCredit = pricePerCredit
	station.Latitude = 0
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid longitude
	station.Latitude = 40.7128
	station.Longitude = 0
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid createdAt
	station.Longitude = -74.0060
	station.CreatedAt = 0
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Valid station
	station.CreatedAt = createdAt
	err = station.Validate()
	assert.Nil(t, err)
}
