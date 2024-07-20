package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestFindOrdersByUserUseCase(t *testing.T) {
	mockRepo := new(repository.MockOrderRepository)
	findOrdersByUserUseCase := NewFindOrdersByUserUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockOrders := []*entity.Order{
		{
			Id:             1,
			Buyer:          common.HexToAddress("0x123"),
			Credits:        big.NewInt(100),
			StationId:      "station_1",
			PricePerCredit: big.NewInt(10),
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		},
		{
			Id:             2,
			Buyer:          common.HexToAddress("0x123"),
			Credits:        big.NewInt(200),
			StationId:      "station_2",
			PricePerCredit: big.NewInt(20),
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		},
	}

	mockRepo.On("FindOrdersByUser", common.HexToAddress("0x123")).Return(mockOrders, nil)

	input := &FindOrderByUserInputDTO{
		User: common.HexToAddress("0x123"),
	}

	output, err := findOrdersByUserUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockOrders), len(output))

	for i, order := range mockOrders {
		assert.Equal(t, order.Id, output[i].Id)
		assert.Equal(t, order.Buyer, output[i].Buyer)
		assert.Equal(t, order.Credits, output[i].Credits)
		assert.Equal(t, order.StationId, output[i].StationId)
		assert.Equal(t, order.PricePerCredit, output[i].PricePerCredit)
		assert.Equal(t, order.CreatedAt, output[i].CreatedAt)
		assert.Equal(t, order.UpdatedAt, output[i].UpdatedAt)
	}

	mockRepo.AssertExpectations(t)
}
