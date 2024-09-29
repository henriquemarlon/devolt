package entity

import (
	"testing"
	"math/big"
	"time"

	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewStation_Success(t *testing.T) {
	owner := custom_type.Address{Address: common.HexToAddress("0x123")}
	consumption := custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	latitude := 37.7749
	longitude := -122.4194
	createdAt := time.Now().Unix()

	station, err := NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.NoError(t, err)
	assert.NotNil(t, station)
	assert.Equal(t, owner, station.Owner)
	assert.Equal(t, consumption, station.Consumption)
	assert.Equal(t, pricePerCredit, station.PricePerCredit)
	assert.Equal(t, latitude, station.Latitude)
	assert.Equal(t, longitude, station.Longitude)
	assert.Equal(t, createdAt, station.CreatedAt)
	assert.Equal(t, StationStateActive, station.State)
}

func TestNewStation_Fail_InvalidStation(t *testing.T) {
	owner := custom_type.Address{Address: common.HexToAddress("0x0")}
	consumption := custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit := custom_type.BigInt{Int: big.NewInt(50)}
	latitude := 37.7749
	longitude := -122.4194
	createdAt := time.Now().Unix()

	station, err := NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)

	owner = custom_type.Address{Address: common.HexToAddress("0x123")}
	consumption = custom_type.BigInt{Int: big.NewInt(0)}

	station, err = NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)

	consumption = custom_type.BigInt{Int: big.NewInt(100)}
	pricePerCredit = custom_type.BigInt{Int: big.NewInt(0)}

	station, err = NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)

	pricePerCredit = custom_type.BigInt{Int: big.NewInt(50)}
	latitude = 0

	station, err = NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)

	latitude = 37.7749
	longitude = 0

	station, err = NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)

	longitude = -122.4194
	createdAt = 0

	station, err = NewStation(owner, consumption, pricePerCredit, latitude, longitude, createdAt)
	assert.Error(t, err)
	assert.Nil(t, station)
	assert.Equal(t, ErrInvalidStation, err)
}
