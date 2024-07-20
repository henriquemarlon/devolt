package contract_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindContractBySymbolUseCase(t *testing.T) {
	mockContractRepo := new(repository.MockContractRepository)
	findContractBySymbol := NewFindContractBySymbolUseCase(mockContractRepo)

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    "TEST",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}

	mockContractRepo.On("FindContractBySymbol", "TEST").Return(mockContract, nil)

	input := &FindContractBySymbolInputDTO{
		Symbol: "TEST",
	}

	output, err := findContractBySymbol.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &FindContractBySymbolOutputDTO{
		Id:        1,
		Symbol:    "TEST",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}, output)

	mockContractRepo.AssertExpectations(t)
}
