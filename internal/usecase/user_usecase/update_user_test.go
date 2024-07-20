package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUpdateUserUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	updateUser := NewUpdateUserUseCase(mockUserRepo)

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		UpdatedAt: 1600,
	}

	mockUserRepo.On("UpdateUser", mock.AnythingOfType("*entity.User")).Return(mockUser, nil)

	input := &UpdateUserInputDTO{
		Id:      1,
		Role:    "admin",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := updateUser.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &UpdateUserOutputDTO{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		UpdatedAt: 1600,
	}, output)

	mockUserRepo.AssertExpectations(t)
}
