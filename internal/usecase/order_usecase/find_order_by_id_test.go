package order_usecase

import (
	"math/big"
	"testing"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
)

func TestFindOrderByIdUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	findOrderById := NewFindOrderByIdUseCase(mockOrderRepo)

	mockOrder := &entity.Order{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
		CreatedAt:      1600,
		UpdatedAt:      1600,
	}

	mockOrderRepo.On("FindOrderById", uint(1)).Return(mockOrder, nil)

	input := &FindOrderByIdInputDTO{
		Id: 1,
	}

	output, err := findOrderById.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &FindOrderOutputDTO{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
		CreatedAt:      1600,
		UpdatedAt:      1600,
	}, output)

	mockOrderRepo.AssertExpectations(t)
}
