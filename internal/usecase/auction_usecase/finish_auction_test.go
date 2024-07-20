package auction_usecase

// import (
// 	"math/big"
// 	"testing"

// 	"github.com/devolthq/devolt/internal/domain/entity"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/rollmelette/rollmelette"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
// )

// func TestFinishAuctionUseCase(t *testing.T) {
// 	mockAuctionRepo := new(repository.MockAuctionRepository)
// 	mockBidRepo := new(repository.MockBidRepository)
// 	finishAuction := NewFinishAuctionUseCase(mockAuctionRepo, mockBidRepo)

// 	mockAuction := &entity.Auction{
// 		Id:         1,
// 		Credits:    big.NewInt(1000),
// 		PriceLimit: big.NewInt(2000),
// 		State:      "active",
// 		ExpiresAt:  1000,
// 		CreatedAt:  900,
// 		UpdatedAt:  900,
// 	}

// 	mockBids := []*entity.Bid{
// 		{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x0"),
// 			Credits:   big.NewInt(500),
// 			Price:     big.NewInt(1000),
// 			CreatedAt: 950,
// 		},
// 		{
// 			Id:        2,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1"),
// 			Credits:   big.NewInt(600),
// 			Price:     big.NewInt(1200),
// 			CreatedAt: 960,
// 		},
// 	}

// 	updatedBids := []*entity.Bid{
// 		{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x0"),
// 			Credits:   big.NewInt(500),
// 			Price:     big.NewInt(1000),
// 			State:     "accepted",
// 			UpdatedAt: 1100,
// 		},
// 		{
// 			Id:        2,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1"),
// 			Credits:   big.NewInt(600),
// 			Price:     big.NewInt(1200),
// 			State:     "rejected",
// 			UpdatedAt: 1100,
// 		},
// 	}

// 	mockAuctionRepo.On("FindActiveAuction").Return(mockAuction, nil)
// 	mockBidRepo.On("FindBidsByAuctionId", uint(1)).Return(mockBids, nil)
// 	mockBidRepo.On("UpdateBid", mock.AnythingOfType("*entity.Bid")).Return(func(bid *entity.Bid) *entity.Bid {
// 		for _, updatedBid := range updatedBids {
// 			if updatedBid.Id == bid.Id {
// 				return updatedBid
// 			}
// 		}
// 		return bid
// 	}, nil)
// 	mockBidRepo.On("CreateBid", mock.AnythingOfType("*entity.Bid")).Return(func(bid *entity.Bid) *entity.Bid { return bid }, nil)
// 	mockBidRepo.On("DeleteBid", uint(1)).Return(nil)
// 	mockBidRepo.On("DeleteBid", uint(2)).Return(nil)
// 	mockAuctionRepo.On("UpdateAuction", mock.AnythingOfType("*entity.Auction")).Return(func(auction *entity.Auction) *entity.Auction {
// 		auction.State = "finished"
// 		auction.UpdatedAt = 1100
// 		return auction
// 	}, nil)

// 	metadata := rollmelette.Metadata{BlockTimestamp: 1100}
// 	output, err := finishAuction.Execute(metadata)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, output)
// 	assert.Equal(t, uint(1), output.Id)

// 	mockAuctionRepo.AssertExpectations(t)
// 	mockBidRepo.AssertExpectations(t)
// }
