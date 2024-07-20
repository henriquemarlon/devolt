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

func TestUpdateContractUseCase(t *testing.T) {
	mockRepo := new(repository.MockContractRepository)
	updateContractUseCase := NewUpdateContractUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    "VOLT",
		Address:   common.HexToAddress("0x123"),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	input := &UpdateContractInputDTO{
		Id:      mockContract.Id,
		Address: common.HexToAddress("0x123"),
		Symbol:  "VOLT",
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: updatedAt,
	}

	mockRepo.On("UpdateContract", mock.AnythingOfType("*entity.Contract")).Return(mockContract, nil)

	output, err := updateContractUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockContract.Id, output.Id)
	assert.Equal(t, mockContract.Symbol, output.Symbol)
	assert.Equal(t, mockContract.Address, output.Address)
	assert.Equal(t, mockContract.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
