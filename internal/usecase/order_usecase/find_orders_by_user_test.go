package order_usecase

import (
	"math/big"
	"testing"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
)

func TestFindOrdersByUserUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	findOrdersByUser := NewFindOrdersByUserUseCase(mockOrderRepo)

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
			Buyer:          common.HexToAddress("0xabcdef"),
			Credits:        big.NewInt(200),
			StationId:      "station2",
			PricePerCredit: big.NewInt(20),
			CreatedAt:      1700,
			UpdatedAt:      1700,
		},
	}

	mockOrderRepo.On("FindOrdersByUser", common.HexToAddress("0xabcdef")).Return(mockOrders, nil)

	input := &FindOrderByUserInputDTO{
		User: common.HexToAddress("0xabcdef"),
	}

	output, err := findOrdersByUser.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output, 2)

	expectedOutput := FindOrderByUserOutputDTO{
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
			Buyer:          common.HexToAddress("0xabcdef"),
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
