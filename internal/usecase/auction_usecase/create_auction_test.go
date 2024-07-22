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

func TestCreateAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	createAuctionUseCase := NewCreateAuctionUseCase(mockRepo)

	credits := custom_type.NewBigInt(big.NewInt(1000))
	priceLimit := custom_type.NewBigInt(big.NewInt(500))
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	createdAt := time.Now().Unix()

	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    credits,
		PriceLimit: priceLimit,
		State:      entity.AuctionOngoing,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
	}

	mockRepo.On("CreateAuction", mock.AnythingOfType("*entity.Auction")).Return(mockAuction, nil)

	input := &CreateAuctionInputDTO{
		Credits:    credits,
		PriceLimit: priceLimit,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
	}

	metadata := rollmelette.Metadata{
		BlockTimestamp: createdAt,
	}

	output, err := createAuctionUseCase.Execute(input, metadata)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockAuction.Id, output.Id)
	assert.Equal(t, mockAuction.Credits, output.Credits)
	assert.Equal(t, mockAuction.PriceLimit, output.PriceLimit)
	assert.Equal(t, string(mockAuction.State), output.State)
	assert.Equal(t, mockAuction.ExpiresAt, output.ExpiresAt)
	assert.Equal(t, mockAuction.CreatedAt, output.CreatedAt)

	mockRepo.AssertExpectations(t)
}
