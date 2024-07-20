package auction_usecase

import (
	"math/big"
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindActiveAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	findActiveAuction := NewFindActiveAuctionUseCase(mockRepo)

	mockBids := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			Credits:   big.NewInt(500),
			Price:     big.NewInt(500),
			State:     "pending",
			CreatedAt: 20242024,
			UpdatedAt: 20242025,
		},
	}

	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		State:      "active",
		Bids:       mockBids,
		ExpiresAt:  20252024,
		CreatedAt:  20242024,
		UpdatedAt:  20242025,
	}

	// Mock the FindActiveAuction behavior
	mockRepo.On("FindActiveAuction").Return(mockAuction, nil)

	// Execute the use case
	output, err := findActiveAuction.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockAuction.Id, output.Id)
	assert.Equal(t, mockAuction.Credits, output.Credits)
	assert.Equal(t, mockAuction.PriceLimit, output.PriceLimit)
	assert.Equal(t, mockAuction.State, output.State)
	assert.Equal(t, mockAuction.ExpiresAt, output.ExpiresAt)
	assert.Equal(t, mockAuction.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockAuction.UpdatedAt, output.UpdatedAt)

	assert.Len(t, output.Bids, len(mockBids))
	for i, bid := range output.Bids {
		assert.Equal(t, mockBids[i].Id, bid.Id)
		assert.Equal(t, mockBids[i].AuctionId, bid.AuctionId)
		assert.Equal(t, mockBids[i].Bidder, bid.Bidder)
		assert.Equal(t, mockBids[i].Credits, bid.Credits)
		assert.Equal(t, mockBids[i].Price, bid.Price)
		assert.Equal(t, mockBids[i].State, bid.State)
		assert.Equal(t, mockBids[i].CreatedAt, bid.CreatedAt)
		assert.Equal(t, mockBids[i].UpdatedAt, bid.UpdatedAt)
	}
	mockRepo.AssertExpectations(t)
}
