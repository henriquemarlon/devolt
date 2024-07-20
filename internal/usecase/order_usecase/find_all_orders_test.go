package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestFindAllOrderUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	findAllOrders := NewFindAllOrderUseCase(mockOrderRepo)

	mockOrders := []*entity.Order{
		{
			Id:             1,
			Buyer:          common.HexToAddress("0xabcdef"),
			Credits:        big.NewInt(100),
			StationId:      "station1",
			PricePerCredit: big.NewInt(10),
			CreatedAt:      1600,
			UpdatedAt:      1600,
		},
		{
			Id:             2,
			Buyer:          common.HexToAddress("0x123456"),
			Credits:        big.NewInt(200),
			StationId:      "station2",
			PricePerCredit: big.NewInt(20),
			CreatedAt:      1700,
			UpdatedAt:      1700,
		},
	}

	mockOrderRepo.On("FindAllOrders").Return(mockOrders, nil)

	output, err := findAllOrders.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output, 2)

	expectedOutput := FindAllOrdersOutputDTO{
		{
			Id:             1,
			Buyer:          common.HexToAddress("0xabcdef"),
			Credits:        big.NewInt(100),
			StationId:      "station1",
			PricePerCredit: big.NewInt(10),
			CreatedAt:      1600,
			UpdatedAt:      1600,
		},
		{
			Id:             2,
			Buyer:          common.HexToAddress("0x123456"),
			Credits:        big.NewInt(200),
			StationId:      "station2",
			PricePerCredit: big.NewInt(20),
			CreatedAt:      1700,
			UpdatedAt:      1700,
		},
	}

	assert.Equal(t, expectedOutput, output)

	mockOrderRepo.AssertExpectations(t)
}
