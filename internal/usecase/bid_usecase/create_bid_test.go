package bid_usecase

import (
	"math/big"
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBidUseCase(t *testing.T) {
	mockBidRepo := new(repository.MockBidRepository)
	mockContractRepo := new(repository.MockContractRepository)
	mockAuctionRepo := new(repository.MockAuctionRepository)
	createBid := NewCreateBidUseCase(mockBidRepo, mockContractRepo, mockAuctionRepo)

	voltContract := &entity.Contract{
		Symbol:  "VOLT",
		Address: common.HexToAddress("0x1234567890abcdef"),
	}

	mockContractRepo.On("FindContractBySymbol", "VOLT").Return(voltContract, nil)

	activeAuction := &entity.Auction{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(2000),
		State:      "active",
		ExpiresAt:  1000,
		CreatedAt:  900,
		UpdatedAt:  900,
	}

	mockAuctionRepo.On("FindActiveAuction").Return(activeAuction, nil)

	mockBid := &entity.Bid{
		Id:        1,
		AuctionId: activeAuction.Id,
		Bidder:    common.HexToAddress("0x1234567890abcdef"),
		Credits:   big.NewInt(500),
		Price:     big.NewInt(1000),
		State:     "pending",
		CreatedAt: 1100,
	}

	mockBidRepo.On("CreateBid", mock.AnythingOfType("*entity.Bid")).Return(mockBid, nil)

	input := &CreateBidInputDTO{
		Bidder: common.HexToAddress("0x1234567890abcdef"),
		Price:  big.NewInt(1000),
	}

	deposit := &rollmelette.ERC20Deposit{
		Token:  voltContract.Address,
		Amount: big.NewInt(500),
	}

	metadata := rollmelette.Metadata{BlockTimestamp: 1100}

	output, err := createBid.Execute(input, deposit, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, uint(1), output.Id)
	assert.Equal(t, uint(1), output.AuctionId)
	assert.Equal(t, common.HexToAddress("0x1234567890abcdef"), output.Bidder)
	assert.Equal(t, big.NewInt(500), output.Credits)
	assert.Equal(t, big.NewInt(1000), output.Price)
	assert.Equal(t, "pending", output.State)
	assert.Equal(t, int64(1100), output.CreatedAt)

	mockBidRepo.AssertExpectations(t)
	mockContractRepo.AssertExpectations(t)
	mockAuctionRepo.AssertExpectations(t)
}
