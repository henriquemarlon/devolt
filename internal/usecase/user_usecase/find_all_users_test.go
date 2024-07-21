package user_usecase

import (
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindAllUsersUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	findAllUsersUseCase := NewFindAllUsersUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockUsers := []*entity.User{
		{
			Id:        1,
			Role:      "admin",
			Address:   common.HexToAddress("0x123").String(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		{
			Id:        2,
			Role:      "user",
			Address:   common.HexToAddress("0x456").String(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	mockRepo.On("FindAllUsers").Return(mockUsers, nil)

	output, err := findAllUsersUseCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockUsers), len(*output))

	for i, user := range mockUsers {
		assert.Equal(t, user.Id, (*output)[i].Id)
		assert.Equal(t, user.Role, (*output)[i].Role)
		assert.Equal(t, user.Address, (*output)[i].Address)
		assert.Equal(t, user.CreatedAt, (*output)[i].CreatedAt)
		assert.Equal(t, user.UpdatedAt, (*output)[i].UpdatedAt)
	}

	mockRepo.AssertExpectations(t)
}
