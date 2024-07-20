package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
)

func TestCreateOrderUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	mockStationRepo := new(repository.MockStationRepository)
	mockContractRepo := new(repository.MockContractRepository)
	createOrder := NewCreateOrderUseCase(mockOrderRepo, mockStationRepo, mockContractRepo)

	mockStation := &entity.Station{
		Id:             "station1",
		Owner:          common.HexToAddress("0xabcdef"),
		PricePerCredit: big.NewInt(10),
	}

	mockContract := &entity.Contract{
		Symbol:  "USDC",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	mockOrder := &entity.Order{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		PricePerCredit: big.NewInt(10),
		CreatedAt:      1600,
	}

	mockStationRepo.On("FindStationById", "station1").Return(mockStation, nil)
	mockContractRepo.On("FindContractBySymbol", "USDC").Return(mockContract, nil)
	mockOrderRepo.On("CreateOrder", mock.AnythingOfType("*entity.Order")).Return(mockOrder, nil)

	input := &CreateOrderInputDTO{
		Buyer:     common.HexToAddress("0xabcdef"),
		Credits:   big.NewInt(100),
		StationId: "station1",
	}

	deposit := &rollmelette.ERC20Deposit{
		Token:  mockContract.Address,
		Amount: big.NewInt(1000),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1600}

	output, err := createOrder.Execute(input, deposit, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &CreateOrderOutputDTO{
		Id:             1,
		Buyer:          common.HexToAddress("0xabcdef"),
		Credits:        big.NewInt(100),
		StationId:      "station1",
		StationOwner:   common.HexToAddress("0xabcdef"),
		PricePerCredit: big.NewInt(10),
		CreatedAt:      1600,
	}, output)

	mockOrderRepo.AssertExpectations(t)
	mockStationRepo.AssertExpectations(t)
	mockContractRepo.AssertExpectations(t)
}
