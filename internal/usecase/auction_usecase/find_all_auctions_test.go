package auction_usecase

import (
	"math/big"
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindAllAuctionsUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	findAllAuctions := NewFindAllAuctionsUseCase(mockRepo)

	mockBids1 := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678").String(),
			Credits:   big.NewInt(500),
			Price:     big.NewInt(500),
			State:     "active",
			CreatedAt: 20242024,
			UpdatedAt: 20242025,
		},
	}

	mockBids2 := []*entity.Bid{
		{
			Id:        2,
			AuctionId: 2,
			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678").String(),
			Credits:   big.NewInt(600),
			Price:     big.NewInt(600),
			State:     "active",
			CreatedAt: 20242026,
			UpdatedAt: 20242027,
		},
	}

	mockAuction1 := &entity.Auction{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		State:      "active",
		Bids:       mockBids1,
		ExpiresAt:  20252024,
		CreatedAt:  20242024,
		UpdatedAt:  20242025,
	}

	mockAuction2 := &entity.Auction{
		Id:         2,
		Credits:    big.NewInt(2000),
		PriceLimit: big.NewInt(2000),
		State:      "active",
		Bids:       mockBids2,
		ExpiresAt:  20252026,
		CreatedAt:  20242026,
		UpdatedAt:  20242027,
	}

	mockRepo.On("FindAllAuctions").Return([]*entity.Auction{mockAuction1, mockAuction2}, nil)

	// Execute the use case
	output, err := findAllAuctions.Execute()
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, *output, 2)

	assert.Equal(t, mockAuction1.Id, (*output)[0].Id)
	assert.Equal(t, mockAuction1.Credits, (*output)[0].Credits)
	assert.Equal(t, mockAuction1.PriceLimit, (*output)[0].PriceLimit)
	assert.Equal(t, string(mockAuction1.State), (*output)[0].State)
	assert.Equal(t, mockAuction1.ExpiresAt, (*output)[0].ExpiresAt)
	assert.Equal(t, mockAuction1.CreatedAt, (*output)[0].CreatedAt)
	assert.Equal(t, mockAuction1.UpdatedAt, (*output)[0].UpdatedAt)

	assert.Len(t, (*output)[0].Bids, len(mockBids1))
	for i, bid := range (*output)[0].Bids {
		assert.Equal(t, mockBids1[i].Id, bid.Id)
		assert.Equal(t, mockBids1[i].AuctionId, bid.AuctionId)
		assert.Equal(t, mockBids1[i].Bidder, bid.Bidder.String())
		assert.Equal(t, mockBids1[i].Credits, bid.Credits)
		assert.Equal(t, mockBids1[i].Price, bid.Price)
		assert.Equal(t, string(mockBids1[i].State), bid.State)
		assert.Equal(t, mockBids1[i].CreatedAt, bid.CreatedAt)
		assert.Equal(t, mockBids1[i].UpdatedAt, bid.UpdatedAt)
	}

	assert.Equal(t, mockAuction2.Id, (*output)[1].Id)
	assert.Equal(t, mockAuction2.Credits, (*output)[1].Credits)
	assert.Equal(t, mockAuction2.PriceLimit, (*output)[1].PriceLimit)
	assert.Equal(t, string(mockAuction2.State), (*output)[1].State)
	assert.Equal(t, mockAuction2.ExpiresAt, (*output)[1].ExpiresAt)
	assert.Equal(t, mockAuction2.CreatedAt, (*output)[1].CreatedAt)
	assert.Equal(t, mockAuction2.UpdatedAt, (*output)[1].UpdatedAt)

	assert.Len(t, (*output)[1].Bids, len(mockBids2))
	for i, bid := range (*output)[1].Bids {
		assert.Equal(t, mockBids2[i].Id, bid.Id)
		assert.Equal(t, mockBids2[i].AuctionId, bid.AuctionId)
		assert.Equal(t, mockBids2[i].Bidder, bid.Bidder.String())
		assert.Equal(t, mockBids2[i].Credits, bid.Credits)
		assert.Equal(t, mockBids2[i].Price, bid.Price)
		assert.Equal(t, string(mockBids2[i].State), bid.State)
		assert.Equal(t, mockBids2[i].CreatedAt, bid.CreatedAt)
		assert.Equal(t, mockBids2[i].UpdatedAt, bid.UpdatedAt)
	}

	mockRepo.AssertExpectations(t)
}
