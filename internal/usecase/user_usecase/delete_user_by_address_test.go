package user_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUserByAddressUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	deleteUserByAddressUseCase := NewDeleteUserByAddressUseCase(mockRepo)

	input := &DeleteUserByAddressInputDTO{
		Address: common.HexToAddress("0x123").String(),
	}

	mockRepo.On("DeleteUserByAddress", input.Address).Return(nil)

	err := deleteUserByAddressUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
