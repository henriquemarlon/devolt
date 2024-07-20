package contract_usecase

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

func TestCreateContractUseCase(t *testing.T) {
	mockRepo := new(repository.MockContractRepository)
	createContractUseCase := NewCreateContractUseCase(mockRepo)

	input := &CreateContractInputDTO{
		Symbol:  "VOLT",
		Address: common.HexToAddress("0x123"),
	}

	createdAt := time.Now().Unix()

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    input.Symbol,
		Address:   input.Address,
		CreatedAt: createdAt,
	}

	mockRepo.On("CreateContract", mock.AnythingOfType("*entity.Contract")).Return(mockContract, nil)

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	output, err := createContractUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockContract.Id, output.Id)
	assert.Equal(t, mockContract.Symbol, output.Symbol)
	assert.Equal(t, mockContract.Address, output.Address)
	assert.Equal(t, mockContract.CreatedAt, output.CreatedAt)

	mockRepo.AssertExpectations(t)
}
