package auction_usecase

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/devolthq/devolt/internal/domain/entity"
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteAuction(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)

	createAuction := NewCreateAuctionUseCase(mockRepo)
	mockAuction := &entity.Auction{
		Id:         1,
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		CreatedAt:  20242024,
		ExpiresAt:  20252024,
	}

	auctionInput := &CreateAuctionInputDTO{
		Credits:    big.NewInt(1000),
		PriceLimit: big.NewInt(1000),
		CreatedAt:  20242024,
		ExpiresAt:  20252024,
	}

	mockRepo.On("CreateAuction", mock.AnythingOfType("*entity.Auction")).Return(mockAuction, nil)
	_, err := createAuction.Execute(auctionInput, rollmelette.Metadata{BlockTimestamp: 1000})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	deleteAuction := NewDeleteAuctionUseCase(mockRepo)
	input := &DeleteAuctionInputDTO{
		Id: 1,
	}

	mockRepo.On("DeleteAuction", input.Id).Return(true, nil)
	err = deleteAuction.Execute(input)
	assert.Nil(t, err)

	mockRepo.On("DeleteAuction", uint(2)).Return(false, fmt.Errorf("auction with ID 2 does not exist"))
	err = deleteAuction.Execute(&DeleteAuctionInputDTO{Id: 2})
	assert.Error(t, err)
	assert.EqualError(t, err, "auction with ID 2 does not exist")

	mockRepo.AssertExpectations(t)
}
