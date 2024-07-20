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

func TestFindStationByIdUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	findStationByIdUseCase := NewFindStationByIdUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockStation := &entity.Station{
		Id:             "station_1",
		Consumption:    big.NewInt(500),
		Owner:          common.HexToAddress("0x123"),
		PricePerCredit: big.NewInt(10),
		State:          entity.StationStateActive,
		Latitude:       40.7128,
		Longitude:      -74.0060,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Orders: []*entity.Order{
			{
				Id:             1,
				Buyer:          common.HexToAddress("0x456"),
				StationId:      "station_1",
				Credits:        big.NewInt(100),
				PricePerCredit: big.NewInt(10),
				CreatedAt:      createdAt,
				UpdatedAt:      updatedAt,
			},
		},
	}

	mockRepo.On("FindStationById", "station_1").Return(mockStation, nil)

	input := &FindStationByIdInputDTO{
		Id: "station_1",
	}

	output, err := findStationByIdUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockStation.Id, output.Id)
	assert.Equal(t, mockStation.Consumption, output.Consumption)
	assert.Equal(t, mockStation.Owner, output.Owner)
	assert.Equal(t, mockStation.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, string(mockStation.State), output.State)
	assert.Equal(t, mockStation.Latitude, output.Latitude)
	assert.Equal(t, mockStation.Longitude, output.Longitude)
	assert.Equal(t, mockStation.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockStation.UpdatedAt, output.UpdatedAt)

	for i, order := range mockStation.Orders {
		assert.Equal(t, order.Id, output.Orders[i].Id)
		assert.Equal(t, order.Buyer, output.Orders[i].Buyer)
		assert.Equal(t, order.StationId, output.Orders[i].StationId)
		assert.Equal(t, order.Credits, output.Orders[i].Credits)
		assert.Equal(t, order.PricePerCredit, output.Orders[i].PricePerCredit)
		assert.Equal(t, order.CreatedAt, output.Orders[i].CreatedAt)
		assert.Equal(t, order.UpdatedAt, output.Orders[i].UpdatedAt)
	}

	mockRepo.AssertExpectations(t)
}
