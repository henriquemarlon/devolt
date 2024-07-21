package auction_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindAuctionByIdUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	findAuctionByIdUseCase := NewFindAuctionByIdUseCase(mockRepo)

	credits := big.NewInt(1000)
	priceLimit := big.NewInt(500)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockBids := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1").String(),
			Credits:   big.NewInt(100),
			Price:     big.NewInt(50),
			State:     entity.BidStatePending,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      entity.AuctionOngoing,
		Bids:       mockBids,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	mockRepo.On("FindAuctionById", uint(1)).Return(mockAuction, nil)

	input := &FindAuctionByIdInputDTO{
		Id: 1,
	}

	output, err := findAuctionByIdUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockAuction.Id, output.Id)
	assert.Equal(t, mockAuction.Credits, output.Credits)
	assert.Equal(t, mockAuction.PriceLimit, output.PriceLimit)
	assert.Equal(t, string(mockAuction.State), output.State)
	assert.Equal(t, mockAuction.ExpiresAt, output.ExpiresAt)
	assert.Equal(t, mockAuction.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockAuction.UpdatedAt, output.UpdatedAt)

	assert.Equal(t, len(mockAuction.Bids), len(output.Bids))
	for i, bid := range mockAuction.Bids {
		assert.Equal(t, bid.Id, output.Bids[i].Id)
		assert.Equal(t, bid.AuctionId, output.Bids[i].AuctionId)
		assert.Equal(t, bid.Bidder, output.Bids[i].Bidder.String())
		assert.Equal(t, bid.Credits, output.Bids[i].Credits)
		assert.Equal(t, bid.Price, output.Bids[i].Price)
		assert.Equal(t, string(bid.State), output.Bids[i].State)
		assert.Equal(t, bid.CreatedAt, output.Bids[i].CreatedAt)
		assert.Equal(t, bid.UpdatedAt, output.Bids[i].UpdatedAt)
	}

	mockRepo.AssertExpectations(t)
}
