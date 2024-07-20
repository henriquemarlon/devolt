package user_usecase

import (
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserUseCase(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	createUser := NewCreateUserUseCase(mockUserRepo)

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
	}

	mockUserRepo.On("CreateUser", mock.AnythingOfType("*entity.User")).Return(mockUser, nil)

	input := &CreateUserInputDTO{
		Role:    "admin",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := createUser.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &CreateUserOutputDTO{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
	}, output)

	mockUserRepo.AssertExpectations(t)
}
