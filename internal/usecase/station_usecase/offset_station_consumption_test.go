package station_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOffSetStationConsumptionUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	offSetStationConsumptionUseCase := NewOffSetStationConsumptionUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	stationConsumption := big.NewInt(500)
	creditsToBeOffSet := big.NewInt(100)
	newConsumption := big.NewInt(400)

	mockStation := &entity.Station{
		Id:             "station_1",
		Consumption:    stationConsumption,
		Owner:          common.HexToAddress("0x123").String(),
		PricePerCredit: big.NewInt(10),
		State:          entity.StationStateActive,
		Latitude:       40.7128,
		Longitude:      -74.0060,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	updatedStation := *mockStation
	updatedStation.Consumption = newConsumption
	updatedStation.UpdatedAt = updatedAt

	input := &OffSetStationConsumptionInputDTO{
		Id:                "station_1",
		CreditsToBeOffSet: creditsToBeOffSet,
	}

	metadata := rollmelette.Metadata{
		MsgSender:      common.HexToAddress("0x123"),
		BlockTimestamp: updatedAt,
	}

	mockRepo.On("FindStationById", "station_1").Return(mockStation, nil)
	mockRepo.On("UpdateStation", mock.AnythingOfType("*entity.Station")).Return(&updatedStation, nil)

	output, err := offSetStationConsumptionUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, updatedStation.Id, output.Id)
	assert.Equal(t, updatedStation.Consumption, output.Consumption)
	assert.Equal(t, updatedStation.Owner, output.Owner)
	assert.Equal(t, updatedStation.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, string(updatedStation.State), output.State)
	assert.Equal(t, updatedStation.Latitude, output.Latitude)
	assert.Equal(t, updatedStation.Longitude, output.Longitude)
	assert.Equal(t, updatedStation.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestOffSetStationConsumptionUseCase_Unauthorized(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	offSetStationConsumptionUseCase := NewOffSetStationConsumptionUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	stationConsumption := big.NewInt(500)
	creditsToBeOffSet := big.NewInt(100)

	mockStation := &entity.Station{
		Id:             "station_1",
		Consumption:    stationConsumption,
		Owner:          common.HexToAddress("0x123").String(),
		PricePerCredit: big.NewInt(10),
		State:          entity.StationStateActive,
		Latitude:       40.7128,
		Longitude:      -74.0060,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	input := &OffSetStationConsumptionInputDTO{
		Id:                "station_1",
		CreditsToBeOffSet: creditsToBeOffSet,
	}

	metadata := rollmelette.Metadata{
		MsgSender:      common.HexToAddress("0x999"), // Different from station owner
		BlockTimestamp: updatedAt,
	}

	mockRepo.On("FindStationById", "station_1").Return(mockStation, nil)

	output, err := offSetStationConsumptionUseCase.Execute(input, metadata)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "can't offSet station consumption, because the station owner is not equal to the msg_sender address, expected: 0x0000000000000000000000000000000000000123, got: 0x0000000000000000000000000000000000000999", err.Error())

	mockRepo.AssertExpectations(t)
}
