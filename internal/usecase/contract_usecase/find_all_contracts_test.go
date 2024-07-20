package contract_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllContractsUseCase(t *testing.T) {
	mockContractRepo := new(repository.MockContractRepository)
	findAllContracts := NewFindAllContractsUseCase(mockContractRepo)

	mockContracts := []*entity.Contract{
		{
			Id:        1,
			Symbol:    "TEST1",
			Address:   common.HexToAddress("0x1234567890abcdef"),
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			Symbol:    "TEST2",
			Address:   common.HexToAddress("0xabcdef1234567890"),
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	mockContractRepo.On("FindAllContracts").Return(mockContracts, nil)

	output, err := findAllContracts.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output, 2)

	expectedOutput := FindAllContractsOutputDTO{
		{
			Id:        1,
			Symbol:    "TEST1",
			Address:   common.HexToAddress("0x1234567890abcdef"),
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			Symbol:    "TEST2",
			Address:   common.HexToAddress("0xabcdef1234567890"),
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	assert.Equal(t, expectedOutput, output)

	mockContractRepo.AssertExpectations(t)
}
