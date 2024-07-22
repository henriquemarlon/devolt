package auction_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/rollmelette/rollmelette"
)

func TestFinishAuctionUseCase(t *testing.T) {
	mockAuctionRepo := new(repository.MockAuctionRepository)
	mockBidRepo := new(repository.MockBidRepository)
	finishAuctionUseCase := NewFinishAuctionUseCase(mockAuctionRepo, mockBidRepo)

	activeAuction := &entity.Auction{
		Id:         1,
		Credits:    custom_type.NewBigInt(big.NewInt(1000)),
		PriceLimit: custom_type.NewBigInt(big.NewInt(500)),
		State:      entity.AuctionOngoing,
		ExpiresAt:  time.Now().Add(-1 * time.Hour).Unix(), // expired auction
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	mockBids := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    custom_type.NewAddress(common.HexToAddress("0x1")),
			Credits:   custom_type.NewBigInt(big.NewInt(500)),
			Price:     custom_type.NewBigInt(big.NewInt(250)),
			State:     entity.BidStatePending,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		{
			Id:        2,
			AuctionId: 1,
			Bidder:    custom_type.NewAddress(common.HexToAddress("0x2")),
			Credits:   custom_type.NewBigInt(big.NewInt(700)),
			Price:     custom_type.NewBigInt(big.NewInt(350)),
			State:     entity.BidStatePending,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}

	mockAuctionRepo.On("FindActiveAuction").Return(activeAuction, nil)
	mockBidRepo.On("FindBidsByAuctionId", uint(1)).Return(mockBids, nil)
	mockBidRepo.On("CreateBid", mock.Anything).Return(&entity.Bid{}, nil)
	mockBidRepo.On("DeleteBid", mock.AnythingOfType("uint")).Return(nil)
	mockBidRepo.On("UpdateBid", mock.Anything).Return(nil, nil)
	mockAuctionRepo.On("UpdateAuction", mock.Anything).Return(nil, nil)

	metadata := rollmelette.Metadata{
		BlockTimestamp: time.Now().Unix(),
	}

	output, err := finishAuctionUseCase.Execute(metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, activeAuction.Id, output.Id)

	mockAuctionRepo.AssertExpectations(t)
	mockBidRepo.AssertExpectations(t)
}
