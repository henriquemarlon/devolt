package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
)

func TestUpdateAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	updateAuction := NewUpdateAuctionUseCase(mockRepo)

	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(2000),
		State:      "updated",
		ExpiresAt:  20262024,
		UpdatedAt:  1000,
	}

	input := &UpdateAuctionInputDTO{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(2000),
		State:      "updated",
		ExpiresAt:  20262024,
	}

	mockRepo.On("UpdateAuction", mock.AnythingOfType("*entity.Auction")).Return(mockAuction, nil)

	metadata := rollmelette.Metadata{BlockTimestamp: 1000}
	output, err := updateAuction.Execute(input, metadata)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockAuction.Id, output.Id)
	assert.Equal(t, mockAuction.Credits, output.Credits)
	assert.Equal(t, mockAuction.PriceLimit, output.PriceLimit)
	assert.Equal(t, mockAuction.State, output.State)
	assert.Equal(t, mockAuction.ExpiresAt, output.ExpiresAt)
	assert.Equal(t, mockAuction.UpdatedAt, output.UpdatedAt)

	mockRepo.AssertExpectations(t)
}
