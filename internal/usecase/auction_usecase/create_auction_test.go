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

func TestCreateAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	createAuction := NewCreateAuctionUseCase(mockRepo)

	mockAuction := &entity.Auction{
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		CreatedAt:  20242024,
		ExpiresAt:  20252024,
	}

	input := &CreateAuctionInputDTO{
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		CreatedAt:  20242024,
		ExpiresAt:  20252024,
	}

	mockRepo.On("CreateAuction", mock.AnythingOfType("*entity.Auction")).Return(mockAuction, nil)

	output, err := createAuction.Execute(input, rollmelette.Metadata{BlockTimestamp: 1000})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.Credits, output.Credits)
	assert.Equal(t, input.PriceLimit, output.PriceLimit)
	assert.Equal(t, input.CreatedAt, output.CreatedAt)
	assert.Equal(t, input.ExpiresAt, output.ExpiresAt)

	mockRepo.AssertExpectations(t)
}
