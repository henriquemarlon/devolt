package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindUserByRoleUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	findUserByRoleUseCase := NewFindUserByRoleUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x123"),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	input := &FindUserByRoleInputDTO{
		Role: "admin",
	}

	mockRepo.On("FindUserByRole", input.Role).Return(mockUser, nil)

	output, err := findUserByRoleUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockUser.Id, output.Id)
	assert.Equal(t, mockUser.Role, output.Role)
	assert.Equal(t, mockUser.Address, output.Address)
	assert.Equal(t, mockUser.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockUser.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
