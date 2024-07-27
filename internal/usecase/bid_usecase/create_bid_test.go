package bid_usecase

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

func TestCreateBidUseCase(t *testing.T) {
	mockBidRepo := new(repository.MockBidRepository)
	mockContractRepo := new(repository.MockContractRepository)
	mockAuctionRepo := new(repository.MockAuctionRepository)
	createBidUseCase := NewCreateBidUseCase(mockBidRepo, mockContractRepo, mockAuctionRepo)

	credits := custom_type.NewBigInt(big.NewInt(1000))
	priceLimit := custom_type.NewBigInt(big.NewInt(500))
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      entity.AuctionOngoing,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	mockContract := &entity.Contract{
		Id:      1,
		Symbol:  "VOLT",
		Address: custom_type.NewAddress(common.HexToAddress("0x123")),
	}

	mockBid := &entity.Bid{
		Id:        1,
		AuctionId: mockAuction.Id,
		Bidder:    custom_type.NewAddress(common.HexToAddress("0x1")),
		Credits:   credits,
		Price:     custom_type.NewBigInt(big.NewInt(500)),
		State:     entity.BidStatePending,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	mockAuctionRepo.On("FindActiveAuction").Return(mockAuction, nil)
	mockContractRepo.On("FindContractBySymbol", "VOLT").Return(mockContract, nil)
	mockBidRepo.On("CreateBid", mock.AnythingOfType("*entity.Bid")).Return(mockBid, nil)

	input := &CreateBidInputDTO{
		Price: custom_type.NewBigInt(big.NewInt(500)),
	}

	deposit := &rollmelette.ERC20Deposit{
		Sender: common.HexToAddress("0x125"),
		Token:  common.HexToAddress("0x123"),
		Amount: credits.Int,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	output, err := createBidUseCase.Execute(input, deposit, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockBid.Id, output.Id)
	assert.Equal(t, mockBid.AuctionId, output.AuctionId)
	assert.Equal(t, mockBid.Bidder, output.Bidder)
	assert.Equal(t, mockBid.Credits, output.Credits)
	assert.Equal(t, mockBid.Price, output.Price)
	assert.Equal(t, string(mockBid.State), output.State)
	assert.Equal(t, mockBid.CreatedAt, output.CreatedAt)

	mockAuctionRepo.AssertExpectations(t)
	mockContractRepo.AssertExpectations(t)
	mockBidRepo.AssertExpectations(t)
}
