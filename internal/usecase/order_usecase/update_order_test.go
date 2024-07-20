package order_usecase

import (
	"math/big"
	"testing"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
)

func TestUpdateOrderUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	updateOrder := NewUpdateOrderUseCase(mockOrderRepo)

	mockOrder := &entity.Order{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
		UpdatedAt:      1600,
	}

	mockOrderRepo.On("UpdateOrder", mock.AnythingOfType("*entity.Order")).Return(mockOrder, nil)

	input := &UpdateOrderInputDTO{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := updateOrder.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &UpdateOrderOutputDTO{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
		UpdatedAt:      1600,
	}, output)

	mockOrderRepo.AssertExpectations(t)
}
