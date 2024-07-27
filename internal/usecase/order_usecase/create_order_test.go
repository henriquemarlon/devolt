package order_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	mockStationRepo := new(repository.MockStationRepository)
	mockContractRepo := new(repository.MockContractRepository)
	createOrderUseCase := NewCreateOrderUseCase(mockOrderRepo, mockStationRepo, mockContractRepo)

	input := &CreateOrderInputDTO{
		StationId: "station_1",
	}

	createdAt := time.Now().Unix()

	mockStation := &entity.Station{
		Id:             "station_1",
		Owner:          custom_type.NewAddress(common.HexToAddress("0x123")),
		PricePerCredit: custom_type.NewBigInt(big.NewInt(10)),
	}

	mockContract := &entity.Contract{
		Id:      1,
		Symbol:  "STABLECOIN",
		Address: custom_type.NewAddress(common.HexToAddress("0x789")),
	}

	mockOrder := &entity.Order{
		Id:             1,
		Buyer:          custom_type.NewAddress(common.HexToAddress("0x123")),
		Credits:        custom_type.NewBigInt(big.NewInt(100)),
		StationId:      "station_1",
		PricePerCredit: custom_type.NewBigInt(big.NewInt(10)),
		CreatedAt:      createdAt,
	}

	deposit := &rollmelette.ERC20Deposit{
		Sender: common.HexToAddress("0x123"),
		Token:  common.HexToAddress("0x789"),
		Amount: big.NewInt(1000),
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	mockStationRepo.On("FindStationById", "station_1").Return(mockStation, nil)
	mockContractRepo.On("FindContractBySymbol", "STABLECOIN").Return(mockContract, nil)
	mockOrderRepo.On("CreateOrder", mock.AnythingOfType("*entity.Order")).Return(mockOrder, nil)

	output, err := createOrderUseCase.Execute(input, deposit, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockOrder.Id, output.Id)
	assert.Equal(t, mockOrder.Buyer, output.Buyer)
	assert.Equal(t, mockOrder.Credits, output.Credits)
	assert.Equal(t, mockOrder.StationId, output.StationId)
	assert.Equal(t, mockStation.Owner, output.StationOwner)
	assert.Equal(t, mockOrder.PricePerCredit, output.PricePerCredit)
	assert.Equal(t, mockOrder.CreatedAt, output.CreatedAt)

	mockStationRepo.AssertExpectations(t)
	mockContractRepo.AssertExpectations(t)
	mockOrderRepo.AssertExpectations(t)
}
