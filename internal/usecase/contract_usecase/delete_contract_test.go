package contract_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteContractUseCase(t *testing.T) {
	mockRepo := new(repository.MockContractRepository)
	deleteContractUseCase := NewDeleteContractUseCase(mockRepo)

	input := &DeleteContractInputDTO{
		Symbol: "VOLT",
	}

	mockRepo.On("DeleteContract", input.Symbol).Return(nil)

	err := deleteContractUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
