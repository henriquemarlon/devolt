package bid_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFindBidsByStateUseCase(t *testing.T) {
	mockRepo := new(repository.MockBidRepository)
	findBidsByStateUseCase := NewFindBidsByStateUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockBids := []*entity.Bid{
		{
			Id:        1,
			AuctionId: 1,
			Bidder:    custom_type.NewAddress(common.HexToAddress("0x1")),
			Credits:   custom_type.NewBigInt(big.NewInt(100)),
			Price:     custom_type.NewBigInt(big.NewInt(50)),
			State:     entity.BidStatePending,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		{
			Id:        2,
			AuctionId: 1,
			Bidder:    custom_type.NewAddress(common.HexToAddress("0x2")),
			Credits:   custom_type.NewBigInt(big.NewInt(200)),
			Price:     custom_type.NewBigInt(big.NewInt(150)),
			State:     entity.BidStatePending,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	mockRepo.On("FindBidsByState", uint(1), "pending").Return(mockBids, nil)

	input := &FindBidsByStateInputDTO{
		AuctionId: 1,
		State:     "pending",
	}

	output, err := findBidsByStateUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, len(mockBids), len(output))

	for i, bid := range mockBids {
		assert.Equal(t, bid.Id, output[i].Id)
		assert.Equal(t, bid.AuctionId, output[i].AuctionId)
		assert.Equal(t, bid.Bidder, output[i].Bidder)
		assert.Equal(t, bid.Credits, output[i].Credits)
		assert.Equal(t, bid.Price, output[i].Price)
		assert.Equal(t, string(bid.State), output[i].State)
		assert.Equal(t, bid.CreatedAt, output[i].CreatedAt)
		assert.Equal(t, bid.UpdatedAt, output[i].UpdatedAt)
	}
	mockRepo.AssertExpectations(t)
}
