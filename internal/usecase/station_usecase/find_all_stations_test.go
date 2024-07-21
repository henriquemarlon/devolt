package station_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindAllStationsUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	findAllStationsUseCase := NewFindAllStationsUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockStations := []*entity.Station{
		{
			Id:             "station_1",
			Consumption:    big.NewInt(500),
			Owner:          common.HexToAddress("0x123").String(),
			PricePerCredit: big.NewInt(10),
			State:          entity.StationStateActive,
			Latitude:       40.7128,
			Longitude:      -74.0060,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Orders: []*entity.Order{
				{
					Id:             1,
					Buyer:          common.HexToAddress("0x456").String(),
					StationId:      "station_1",
					Credits:        big.NewInt(100),
					PricePerCredit: big.NewInt(10),
					CreatedAt:      createdAt,
					UpdatedAt:      updatedAt,
				},
			},
		},
		{
			Id:             "station_2",
			Consumption:    big.NewInt(300),
			Owner:          common.HexToAddress("0x789").String(),
			PricePerCredit: big.NewInt(20),
			State:          entity.StationStateInactive,
			Latitude:       34.0522,
			Longitude:      -118.2437,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Orders:         []*entity.Order{},
		},
	}

	mockRepo.On("FindAllStations").Return(mockStations, nil)

	output, err := findAllStationsUseCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockStations), len(output))

	for i, station := range mockStations {
		assert.Equal(t, station.Id, output[i].Id)
		assert.Equal(t, station.Consumption, output[i].Consumption)
		assert.Equal(t, station.Owner, output[i].Owner)
		assert.Equal(t, station.PricePerCredit, output[i].PricePerCredit)
		assert.Equal(t, string(station.State), output[i].State)
		assert.Equal(t, station.Latitude, output[i].Latitude)
		assert.Equal(t, station.Longitude, output[i].Longitude)
		assert.Equal(t, station.CreatedAt, output[i].CreatedAt)
		assert.Equal(t, station.UpdatedAt, output[i].UpdatedAt)

		for j, order := range station.Orders {
			assert.Equal(t, order.Id, output[i].Orders[j].Id)
			assert.Equal(t, order.Buyer, output[i].Orders[j].Buyer)
			assert.Equal(t, order.StationId, output[i].Orders[j].StationId)
			assert.Equal(t, order.Credits, output[i].Orders[j].Credits)
			assert.Equal(t, order.PricePerCredit, output[i].Orders[j].PricePerCredit)
			assert.Equal(t, order.CreatedAt, output[i].Orders[j].CreatedAt)
			assert.Equal(t, order.UpdatedAt, output[i].Orders[j].UpdatedAt)
		}
	}

	mockRepo.AssertExpectations(t)
}
