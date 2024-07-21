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

func TestCreateUserUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	createUserUseCase := NewCreateUserUseCase(mockRepo)

	createdAt := time.Now().Unix()

	input := &CreateUserInputDTO{
		Role:    "admin",
		Address: common.HexToAddress("0x123").String(),
	}

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   common.HexToAddress("0x123").String(),
		CreatedAt: createdAt,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*entity.User")).Return(mockUser, nil)

	output, err := createUserUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockUser.Id, output.Id)
	assert.Equal(t, mockUser.Role, output.Role)
	assert.Equal(t, mockUser.Address, output.Address)
	assert.Equal(t, mockUser.CreatedAt, output.CreatedAt)

	mockRepo.AssertExpectations(t)
}
