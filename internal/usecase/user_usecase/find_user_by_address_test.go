package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUserByAddressUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	findUserByAddress := NewFindUserByAddressUseCase(mockUserRepo)

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}

	mockUserRepo.On("FindUserByAddress", common.HexToAddress("0x1234567890abcdef")).Return(mockUser, nil)

	input := &FindUserByAddressInputDTO{
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	output, err := findUserByAddress.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &FindUserOutputDTO{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}, output)

	mockUserRepo.AssertExpectations(t)
}
