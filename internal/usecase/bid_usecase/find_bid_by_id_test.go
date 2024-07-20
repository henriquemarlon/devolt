package bid_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestFindBidByIdUseCase(t *testing.T) {
	mockBidRepo := new(repository.MockBidRepository)
	findBidById := NewFindBidByIdUseCase(mockBidRepo)

	mockBid := &entity.Bid{
		Id:        1,
		AuctionId: 1,
		Bidder:    common.HexToAddress("0x0"),
		Credits:   big.NewInt(500),
		Price:     big.NewInt(1000),
		State:     "pending",
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}

	mockBidRepo.On("FindBidById", uint(1)).Return(mockBid, nil)

	input := &FindBidByIdInputDTO{
		Id: 1,
	}

	output, err := findBidById.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, &FindBidOutputDTO{
		Id:        1,
		AuctionId: 1,
		Bidder:    common.HexToAddress("0x0"),
		Credits:   big.NewInt(500),
		Price:     big.NewInt(1000),
		State:     "pending",
		CreatedAt: 1600,
		UpdatedAt: 1600,
	}, output)

	mockBidRepo.AssertExpectations(t)
}
