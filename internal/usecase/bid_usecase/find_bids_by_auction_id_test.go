package bid_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestFindBidsByAuctionIdUseCase(t *testing.T) {
	mockBidRepo := new(repository.MockBidRepository)
	findBidsByAuctionId := NewFindBidsByAuctionIdUseCase(mockBidRepo)

	mockBids := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x0"),
			Credits:   big.NewInt(500),
			Price:     big.NewInt(1000),
			State:     "pending",
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1"),
			Credits:   big.NewInt(600),
			Price:     big.NewInt(1200),
			State:     "accepted",
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	mockBidRepo.On("FindBidsByAuctionId", uint(1)).Return(mockBids, nil)

	input := &FindBidsByAuctionIdInputDTO{
		AuctionId: 1,
	}

	output, err := findBidsByAuctionId.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, *output, 2)

	expectedOutput := FindBidsByAuctionIdOutputDTO{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x0"),
			Credits:   big.NewInt(500),
			Price:     big.NewInt(1000),
			State:     "pending",
			CreatedAt: 1600,
			UpdatedAt: 1600,
		},
		{
			Id:        2,
			AuctionId: 1,
			Bidder:    common.HexToAddress("0x1"),
			Credits:   big.NewInt(600),
			Price:     big.NewInt(1200),
			State:     "accepted",
			CreatedAt: 1700,
			UpdatedAt: 1700,
		},
	}

	assert.Equal(t, expectedOutput, *output)

	mockBidRepo.AssertExpectations(t)
}
