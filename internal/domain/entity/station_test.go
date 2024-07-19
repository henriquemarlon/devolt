package entity

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
	"time"
)

// StationRepositoryMock is a mock implementation of the StationRepository interface
type StationRepositoryMock struct {
	mock.Mock
}

func (m *StationRepositoryMock) CreateStation(station *Station) (*Station, error) {
	args := m.Called(station)
	return args.Get(0).(*Station), args.Error(1)
}

func (m *StationRepositoryMock) FindStationById(id string) (*Station, error) {
	args := m.Called(id)
	return args.Get(0).(*Station), args.Error(1)
}

func (m *StationRepositoryMock) FindAllStations() ([]*Station, error) {
	args := m.Called()
	return args.Get(0).([]*Station), args.Error(1)
}

func (m *StationRepositoryMock) UpdateStation(station *Station) (*Station, error) {
	args := m.Called(station)
	return args.Get(0).(*Station), args.Error(1)
}

func (m *StationRepositoryMock) DeleteStation(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewStation(t *testing.T) {
	id := "station1"
	owner := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	pricePerCredit := big.NewInt(100)
	latitude := 37.7749
	longitude := -122.4194
	state := "pending"
	createdAt := time.Now().Unix()

	station := NewStation(id, owner, pricePerCredit, latitude, longitude, state, createdAt)

	assert.NotNil(t, station)
	assert.Equal(t, id, station.Id)
	assert.Equal(t, owner, station.Owner)
	assert.Equal(t, state, station.State)
	assert.Equal(t, pricePerCredit, station.PricePerCredit)
	assert.Equal(t, latitude, station.Latitude)
	assert.Equal(t, longitude, station.Longitude)
	assert.Equal(t, createdAt, station.CreatedAt)
	assert.Empty(t, station.Orders)
}

func TestStationRepository(t *testing.T) {
	mockRepo := new(StationRepositoryMock)

	t.Run("CreateStation", func(t *testing.T) {
		station := &Station{
			Id:             "station1",
			Consumption:    big.NewInt(500),
			Owner:          common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			State:          "pending",
			PricePerCredit: big.NewInt(100),
			Latitude:       37.7749,
			Longitude:      -122.4194,
			CreatedAt:      time.Now().Unix(),
		}

		mockRepo.On("CreateStation", station).Return(station, nil)

		createdStation, err := mockRepo.CreateStation(station)

		assert.Nil(t, err)
		assert.Equal(t, station, createdStation)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindStationById", func(t *testing.T) {
		station := &Station{
			Id:             "station1",
			Consumption:    big.NewInt(500),
			Owner:          common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			State:          "pending",
			PricePerCredit: big.NewInt(100),
			Latitude:       37.7749,
			Longitude:      -122.4194,
			CreatedAt:      time.Now().Unix(),
		}

		mockRepo.On("FindStationById", "station1").Return(station, nil)

		foundStation, err := mockRepo.FindStationById("station1")

		assert.Nil(t, err)
		assert.Equal(t, station, foundStation)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindAllStations", func(t *testing.T) {
		station1 := &Station{
			Id:             "station1",
			Consumption:    big.NewInt(500),
			Owner:          common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			State:          "pending",
			PricePerCredit: big.NewInt(100),
			Latitude:       37.7749,
			Longitude:      -122.4194,
			CreatedAt:      time.Now().Unix(),
		}

		station2 := &Station{
			Id:             "station2",
			Consumption:    big.NewInt(700),
			Owner:          common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			State:          "active",
			PricePerCredit: big.NewInt(150),
			Latitude:       40.7128,
			Longitude:      -74.0060,
			CreatedAt:      time.Now().Unix(),
		}

		stations := []*Station{station1, station2}

		mockRepo.On("FindAllStations").Return(stations, nil)

		foundStations, err := mockRepo.FindAllStations()

		assert.Nil(t, err)
		assert.Equal(t, stations, foundStations)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateStation", func(t *testing.T) {
		station := &Station{
			Id:             "station1",
			Consumption:    big.NewInt(500),
			Owner:          common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			State:          "pending",
			PricePerCredit: big.NewInt(100),
			Latitude:       37.7749,
			Longitude:      -122.4194,
			CreatedAt:      time.Now().Unix(),
		}

		updatedStation := &Station{
			Id:             "station1",
			Consumption:    big.NewInt(800),
			Owner:          station.Owner,
			State:          "active",
			PricePerCredit: big.NewInt(150),
			Latitude:       station.Latitude,
			Longitude:      station.Longitude,
			CreatedAt:      station.CreatedAt,
			UpdatedAt:      time.Now().Unix(),
		}

		mockRepo.On("UpdateStation", updatedStation).Return(updatedStation, nil)

		resultStation, err := mockRepo.UpdateStation(updatedStation)

		assert.Nil(t, err)
		assert.Equal(t, updatedStation, resultStation)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteStation", func(t *testing.T) {
		mockRepo.On("DeleteStation", "station1").Return(nil)

		err := mockRepo.DeleteStation("station1")

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})
}
