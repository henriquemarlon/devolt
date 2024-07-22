package user_usecase

import (
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByAddressUseCase(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	findUserByAddressUseCase := NewFindUserByAddressUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockUser := &entity.User{
		Id:        1,
		Role:      "admin",
		Address:   custom_type.NewAddress(common.HexToAddress("0x123")),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	input := &FindUserByAddressInputDTO{
		Address: custom_type.NewAddress(common.HexToAddress("0x123")),
	}

	mockRepo.On("FindUserByAddress", input.Address).Return(mockUser, nil)

	output, err := findUserByAddressUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockUser.Id, output.Id)
	assert.Equal(t, mockUser.Role, output.Role)
	assert.Equal(t, mockUser.Address, output.Address)
	assert.Equal(t, mockUser.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockUser.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
