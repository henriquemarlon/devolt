package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllUsersUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	findAllUsers := NewFindAllUsersUseCase(mockUserRepo)

	mockUsers := []*entity.User{
		{
			Id:        1,
			Role:      "admin",
			Address:   common.HexToAddress("0x1234567890abcdef"),
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			Role:      "user",
			Address:   common.HexToAddress("0xabcdef1234567890"),
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	mockUserRepo.On("FindAllUsers").Return(mockUsers, nil)

	output, err := findAllUsers.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, *output, 2)

	expectedOutput := FindAllUsersOutputDTO{
		{
			Id:        1,
			Role:      "admin",
			Address:   common.HexToAddress("0x1234567890abcdef"),
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			Role:      "user",
			Address:   common.HexToAddress("0xabcdef1234567890"),
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	assert.Equal(t, &expectedOutput, output)

	mockUserRepo.AssertExpectations(t)
}
