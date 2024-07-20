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

func TestCreateStationUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	createStationUseCase := NewCreateStationUseCase(mockRepo)

	createdAt := time.Now().Unix()

	input := &CreateStationInputDTO{
		Id:             "station_1",
		Owner:          common.HexToAddress("0x123"),
		PricePerCredit: big.NewInt(100),
		Latitude:       40.7128,
		Longitude:      -74.0060,
	}

	mockStation := &entity.Station{
		Id:             "station_1",
		Owner:          input.Owner,
		PricePerCredit: input.PricePerCredit,
		State:          entity.StationStateActive,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		CreatedAt:      createdAt,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	mockRepo.On("CreateStation", mock.AnythingOfType("*entity.Station")).Return(mockStation, nil)

	output, err := createStationUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockStation.Id, output.Id)
	assert.Equal(t, mockStation.Owner, output.Owner)
	assert.Equal(t, mockStation.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, string(mockStation.State), output.State)
	assert.Equal(t, mockStation.Latitude, output.Latitude)
	assert.Equal(t, mockStation.Longitude, output.Longitude)
	assert.Equal(t, mockStation.CreatedAt, output.CreatedAt)

	mockRepo.AssertExpectations(t)
}
