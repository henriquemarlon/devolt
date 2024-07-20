package user_usecase

import (
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUserUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	updateUserUseCase := NewUpdateUserUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockUser := &entity.User{
		Id:        1,
		Role:      "user",
		Address:   common.HexToAddress("0x123"),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	input := &UpdateUserInputDTO{
		Id:      mockUser.Id,
		Role:    "admin",
		Address: mockUser.Address,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: updatedAt,
	}

	updatedUser := &entity.User{
		Id:        mockUser.Id,
		Role:      input.Role,
		Address:   input.Address,
		CreatedAt: mockUser.CreatedAt,
		UpdatedAt: metadata.BlockTimestamp,
	}

	mockRepo.On("UpdateUser", mock.AnythingOfType("*entity.User")).Return(updatedUser, nil)

	output, err := updateUserUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, updatedUser.Id, output.Id)
	assert.Equal(t, updatedUser.Role, output.Role)
	assert.Equal(t, updatedUser.Address, output.Address)
	assert.Equal(t, updatedUser.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
