package user_usecase

import (
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByRoleUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	findUserByRole := NewFindUserByRoleUseCase(mockUserRepo)

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}

	mockUserRepo.On("FindUserByRole", "admin").Return(mockUser, nil)

	input := &FindUserByRoleInputDTO{
		Role: "admin",
	}

	output, err := findUserByRole.Execute(input)
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
