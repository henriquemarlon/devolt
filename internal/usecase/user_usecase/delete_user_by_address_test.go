package user_usecase

import (
	"testing"

	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserByAddressUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	deleteUserByAddressUseCase := NewDeleteUserByAddressUseCase(mockRepo)

	input := &DeleteUserByAddressInputDTO{
		Address: custom_type.NewAddress(common.HexToAddress("0x123")),
	}

	mockRepo.On("DeleteUserByAddress", input.Address).Return(nil)

	err := deleteUserByAddressUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
