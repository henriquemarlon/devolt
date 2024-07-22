package order_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindOrderByIdUseCase(t *testing.T) {
	mockRepo := new(repository.MockOrderRepository)
	findOrderByIdUseCase := NewFindOrderByIdUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockOrder := &entity.Order{
		Id:             1,
		Buyer:          custom_type.NewAddress(common.HexToAddress("0x123")),
		Credits:        custom_type.NewBigInt(big.NewInt(100)),
		StationId:      "station_1",
		PricePerCredit: custom_type.NewBigInt(big.NewInt(10)),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	mockRepo.On("FindOrderById", uint(1)).Return(mockOrder, nil)

	input := &FindOrderByIdInputDTO{
		Id: 1,
	}

	output, err := findOrderByIdUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockOrder.Id, output.Id)
	assert.Equal(t, mockOrder.Buyer, output.Buyer)
	assert.Equal(t, mockOrder.Credits, output.Credits)
	assert.Equal(t, mockOrder.StationId, output.StationId)
	assert.Equal(t, mockOrder.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, mockOrder.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockOrder.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
