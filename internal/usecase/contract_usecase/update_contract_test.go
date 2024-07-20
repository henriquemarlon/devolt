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

func TestUpdateContractUseCase(t *testing.T) {
	mockContractRepo := new(repository.MockContractRepository)
	updateContract := NewUpdateContractUseCase(mockContractRepo)

	mockContract := &entity.Contract{
		Id:        1,
		Symbol:    "UPDATED",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		UpdatedAt: 1600,
	}

	mockContractRepo.On("UpdateContract", mock.AnythingOfType("*entity.Contract")).Return(mockContract, nil)

	input := &UpdateContractInputDTO{
		Id:      1,
		Symbol:  "UPDATED",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := updateContract.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &UpdateContractOutputDTO{
		Id:        1,
		Symbol:    "UPDATED",
		Address:   common.HexToAddress("0x1234567890abcdef"),
		UpdatedAt: 1600,
	}, output)

	mockContractRepo.AssertExpectations(t)
}
