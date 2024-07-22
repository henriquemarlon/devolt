package station_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateStationUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	updateStationUseCase := NewUpdateStationUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockStation := &entity.Station{
		Id:             "station_1",
		Consumption:    custom_type.NewBigInt(big.NewInt(500)),
		Owner:          custom_type.NewAddress(common.HexToAddress("0x123")),
		PricePerCredit: custom_type.NewBigInt(big.NewInt(10)),
		State:          entity.StationStateActive,
		Latitude:       40.7128,
		Longitude:      -74.0060,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	input := &UpdateStationInputDTO{
		Id:             mockStation.Id,
		Consumption:    mockStation.Consumption.Int,
		Owner:          mockStation.Owner.Address,
		PricePerCredit: mockStation.PricePerCredit.Int,
		State:          string(mockStation.State),
		Latitude:       mockStation.Latitude,
		Longitude:      mockStation.Longitude,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: updatedAt,
	}

	mockRepo.On("UpdateStation", mock.AnythingOfType("*entity.Station")).Return(mockStation, nil)

	output, err := updateStationUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockStation.Id, output.Id)
	assert.Equal(t, mockStation.Consumption, output.Consumption)
	assert.Equal(t, mockStation.Owner, output.Owner)
	assert.Equal(t, mockStation.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, string(mockStation.State), output.State)
	assert.Equal(t, mockStation.Latitude, output.Latitude)
	assert.Equal(t, mockStation.Longitude, output.Longitude)
	assert.Equal(t, mockStation.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
