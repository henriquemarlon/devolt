package entity

// import (
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/stretchr/testify/assert"
// 	"math/big"
// 	"testing"
// 	"time"
// )

// func TestNewStation(t *testing.T) {
// 	id := "station1"
// 	owner := common.HexToAddress("0x123")
// 	pricePerCredit := big.NewInt(100)
// 	latitude := 40.7128
// 	longitude := -74.0060
// 	createdAt := time.Now().Unix()

// 	station := NewStation(id, owner, pricePerCredit, latitude, longitude, createdAt)

// 	assert.NotNil(t, station)
// 	assert.Equal(t, id, station.Id)
// 	assert.Equal(t, owner, station.Owner)
// 	assert.Equal(t, pricePerCredit, station.PricePerCredit)
// 	assert.Equal(t, latitude, station.Latitude)
// 	assert.Equal(t, longitude, station.Longitude)
// 	assert.Equal(t, createdAt, station.CreatedAt)
// 	assert.Equal(t, "pending", station.State)
// 	assert.Nil(t, station.Consumption)
// 	assert.Empty(t, station.Orders)
// 	assert.Equal(t, int64(0), station.UpdatedAt)
// }

// func TestStation_Validate(t *testing.T) {
// 	station := &Station{
// 		Id:             "station1",
// 		Owner:          common.HexToAddress("0x123"),
// 		PricePerCredit: big.NewInt(100),
// 		Latitude:       40.7128,
// 		Longitude:      -74.0060,
// 		CreatedAt:      time.Now().Unix(),
// 	}

// 	assert.True(t, station.Latitude >= -90 && station.Latitude <= 90, "latitude must be between -90 and 90")
// 	assert.True(t, station.Longitude >= -180 && station.Longitude <= 180, "longitude must be between -180 and 180")

// 	station.Latitude = -91
// 	assert.False(t, station.Latitude >= -90 && station.Latitude <= 90, "latitude must be between -90 and 90")

// 	station.Longitude = -181
// 	assert.False(t, station.Longitude >= -180 && station.Longitude <= 180, "longitude must be between -180 and 180")
// }

// func TestStation_UpdateState(t *testing.T) {
// 	station := &Station{
// 		Id:             "station1",
// 		Owner:          common.HexToAddress("0x123"),
// 		PricePerCredit: big.NewInt(100),
// 		Latitude:       40.7128,
// 		Longitude:      -74.0060,
// 		CreatedAt:      time.Now().Unix(),
// 		State:          "pending",
// 	}

// 	station.State = "active"
// 	assert.Equal(t, "active", station.State)

// 	station.State = "inactive"
// 	assert.Equal(t, "inactive", station.State)
// }

// func TestStation_AddOrder(t *testing.T) {
// 	station := &Station{
// 		Id:             "station1",
// 		Owner:          common.HexToAddress("0x123"),
// 		PricePerCredit: big.NewInt(100),
// 		Latitude:       40.7128,
// 		Longitude:      -74.0060,
// 		CreatedAt:      time.Now().Unix(),
// 		Orders:         []*Order{},
// 	}

// 	order := &Order{
// 		Id:        1,
// 		StationId: "station1",
// 	}

// 	station.Orders = append(station.Orders, order)
// 	assert.Equal(t, 1, len(station.Orders))
// 	assert.Equal(t, order, station.Orders[0])
// 	assert.Equal(t, "station1", station.Orders[0].StationId)
// }
