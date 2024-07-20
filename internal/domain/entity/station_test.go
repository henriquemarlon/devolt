package entity

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewStation(t *testing.T) {
	id := "station-1"
	owner := common.HexToAddress("0x123")
	pricePerCredit := big.NewInt(50)
	latitude := 40.7128
	longitude := -74.0060
	createdAt := time.Now().Unix()

	station, err := NewStation(id, owner, pricePerCredit, latitude, longitude, createdAt)
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
	owner := common.HexToAddress("0x123")
	pricePerCredit := big.NewInt(50)
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
	station.Id = "station-1"
	station.Owner = common.Address{}
	err = station.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidStation, err)

	// Invalid price per credit
	station.Owner = owner
	station.PricePerCredit = nil
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
