package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestFindAllAuctionsUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	findAllAuctionsUseCase := NewFindAllAuctionsUseCase(mockRepo)

	credits1 := big.NewInt(1000)
	priceLimit1 := big.NewInt(500)
	expiresAt1 := time.Now().Add(24 * time.Hour).Unix()
	createdAt1 := time.Now().Unix()
	updatedAt1 := time.Now().Unix()

	credits2 := big.NewInt(2000)
	priceLimit2 := big.NewInt(1000)
	expiresAt2 := time.Now().Add(48 * time.Hour).Unix()
	createdAt2 := time.Now().Unix()
	updatedAt2 := time.Now().Unix()

	mockBids1 := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1"),
			Credits:   big.NewInt(100),
			Price:     big.NewInt(50),
			State:     entity.BidStatePending,
			CreatedAt: createdAt1,
			UpdatedAt: updatedAt1,
		},
	}

	mockBids2 := []*entity.Bid{
		{
			Id:        2,
			AuctionId: 2,
			Bidder:    common.HexToAddress("0x2"),
			Credits:   big.NewInt(200),
			Price:     big.NewInt(150),
			State:     entity.BidStateAccepted,
			CreatedAt: createdAt2,
			UpdatedAt: updatedAt2,
		},
	}

	mockAuctions := []*entity.Auction{
		{
			Id:         1,
			Credits:    credits1,
			PriceLimit: priceLimit1,
			State:      entity.AuctionOngoing,
			Bids:       mockBids1,
			ExpiresAt:  expiresAt1,
			CreatedAt:  createdAt1,
			UpdatedAt:  updatedAt1,
		},
		{
			Id:         2,
			Credits:    credits2,
			PriceLimit: priceLimit2,
			State:      entity.AuctionFinished,
			Bids:       mockBids2,
			ExpiresAt:  expiresAt2,
			CreatedAt:  createdAt2,
			UpdatedAt:  updatedAt2,
		},
	}

	mockRepo.On("FindAllAuctions").Return(mockAuctions, nil)

	output, err := findAllAuctionsUseCase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockAuctions), len(*output))

	for i, auction := range mockAuctions {
		assert.Equal(t, auction.Id, (*output)[i].Id)
		assert.Equal(t, auction.Credits, (*output)[i].Credits)
		assert.Equal(t, auction.PriceLimit, (*output)[i].PriceLimit)
		assert.Equal(t, string(auction.State), (*output)[i].State)
		assert.Equal(t, auction.ExpiresAt, (*output)[i].ExpiresAt)
		assert.Equal(t, auction.CreatedAt, (*output)[i].CreatedAt)
		assert.Equal(t, auction.UpdatedAt, (*output)[i].UpdatedAt)

		assert.Equal(t, len(auction.Bids), len((*output)[i].Bids))
		for j, bid := range auction.Bids {
			assert.Equal(t, bid.Id, (*output)[i].Bids[j].Id)
			assert.Equal(t, bid.AuctionId, (*output)[i].Bids[j].AuctionId)
			assert.Equal(t, bid.Bidder, (*output)[i].Bids[j].Bidder)
			assert.Equal(t, bid.Credits, (*output)[i].Bids[j].Credits)
			assert.Equal(t, bid.Price, (*output)[i].Bids[j].Price)
			assert.Equal(t, string(bid.State), (*output)[i].Bids[j].State)
			assert.Equal(t, bid.CreatedAt, (*output)[i].Bids[j].CreatedAt)
			assert.Equal(t, bid.UpdatedAt, (*output)[i].Bids[j].UpdatedAt)
		}
	}

	mockRepo.AssertExpectations(t)
}
