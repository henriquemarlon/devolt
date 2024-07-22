package contract_usecase

import (
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindContractBySymbolUseCase(t *testing.T) {
	mockRepo := new(repository.MockContractRepository)
	findContractBySymbolUseCase := NewFindContractBySymbolUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    "VOLT",
		Address:   custom_type.NewAddress(common.HexToAddress("0x123")),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	mockRepo.On("FindContractBySymbol", "VOLT").Return(mockContract, nil)

	input := &FindContractBySymbolInputDTO{
		Symbol: "VOLT",
	}

	output, err := findContractBySymbolUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockContract.Id, output.Id)
	assert.Equal(t, mockContract.Symbol, output.Symbol)
	assert.Equal(t, mockContract.Address, output.Address)
	assert.Equal(t, mockContract.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockContract.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
