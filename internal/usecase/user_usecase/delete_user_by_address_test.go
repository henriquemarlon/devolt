package user_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUserByAddressUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	deleteUser := NewDeleteUserByAddressUseCase(mockUserRepo)

	mockUserRepo.On("DeleteUserByAddress", common.HexToAddress("0x1234567890abcdef")).Return(nil)

	input := &DeleteUserByAddressInputDTO{
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	err := deleteUser.Execute(input)
	assert.Nil(t, err)

	mockUserRepo.AssertExpectations(t)
}
