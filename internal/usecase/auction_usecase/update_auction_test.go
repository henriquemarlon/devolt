package auction_usecase

import (
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	updateAuctionUseCase := NewUpdateAuctionUseCase(mockRepo)

	credits := custom_type.NewBigInt(big.NewInt(1000))
	priceLimit := custom_type.NewBigInt(big.NewInt(500))
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	updatedAt := time.Now().Unix()

	input := &UpdateAuctionInputDTO{
		Id:         1,
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      "ongoing",
		ExpiresAt:  expiresAt,
	}

	mockAuction := &entity.Auction{
		Id:         input.Id,
		Credits:    input.Credits,
		PriceLimit: input.PriceLimit,
		State:      entity.AuctionState(input.State),
		ExpiresAt:  input.ExpiresAt,
		UpdatedAt:  updatedAt,
	}

	mockRepo.On("UpdateAuction", mock.AnythingOfType("*entity.Auction")).Return(mockAuction, nil)

	metadata := rollmelette.Metadata{
		BlockTimestamp: updatedAt,
	}

	output, err := updateAuctionUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockAuction.Id, output.Id)
	assert.Equal(t, mockAuction.Credits, output.Credits)
	assert.Equal(t, mockAuction.PriceLimit, output.PriceLimit)
	assert.Equal(t, string(mockAuction.State), output.State)
	assert.Equal(t, mockAuction.ExpiresAt, output.ExpiresAt)
	assert.Equal(t, mockAuction.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
