package station_usecase

import (
	"testing"

	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteStationUseCase(t *testing.T) {
	mockRepo := new(repository.MockStationRepository)
	deleteStationUseCase := NewDeleteStationUseCase(mockRepo)

	input := &DeleteStationInputDTO{
		Id: "station_1",
	}

	mockRepo.On("DeleteStation", input.Id).Return(nil)

	err := deleteStationUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
