package contract_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteContractUseCase(t *testing.T) {
	mockContractRepo := new(repository.MockContractRepository)
	deleteContract := NewDeleteContractUseCase(mockContractRepo)

	mockContractRepo.On("DeleteContract", "TEST").Return(nil)

	input := &DeleteContractInputDTO{
		Symbol: "TEST",
	}

	err := deleteContract.Execute(input)
	assert.Nil(t, err)

	mockContractRepo.AssertExpectations(t)
}
