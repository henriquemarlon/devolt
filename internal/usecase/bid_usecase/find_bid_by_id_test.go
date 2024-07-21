package bid_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestFindBidByIdUseCase(t *testing.T) {
	mockRepo := new(repository.MockBidRepository)
	findBidByIdUseCase := NewFindBidByIdUseCase(mockRepo)

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	mockBid := &entity.Bid{
		Id:        1,
		AuctionId: 1,
		Bidder:    common.HexToAddress("0x1").String(),
		Credits:   big.NewInt(100),
		Price:     big.NewInt(50),
		State:     entity.BidStatePending,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	mockRepo.On("FindBidById", uint(1)).Return(mockBid, nil)

	input := &FindBidByIdInputDTO{
		Id: 1,
	}

	output, err := findBidByIdUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockBid.Id, output.Id)
	assert.Equal(t, mockBid.AuctionId, output.AuctionId)
	assert.Equal(t, mockBid.Bidder, output.Bidder)
	assert.Equal(t, mockBid.Credits, output.Credits)
	assert.Equal(t, mockBid.Price, output.Price)
	assert.Equal(t, string(mockBid.State), output.State)
	assert.Equal(t, mockBid.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockBid.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
