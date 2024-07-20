package contract_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateContractUseCase(t *testing.T) {
	mockContractRepo := new(repository.MockContractRepository)
	createContract := NewCreateContractUseCase(mockContractRepo)

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    "TEST",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
	}

	mockContractRepo.On("CreateContract", mock.AnythingOfType("*entity.Contract")).Return(mockContract, nil)

	input := &CreateContractInputDTO{
		Symbol:  "TEST",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := createContract.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &CreateContractOutputDTO{
		Id:        1,
		Symbol:    "TEST",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
	}, output)

	mockContractRepo.AssertExpectations(t)
}
