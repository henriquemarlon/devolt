package auction_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAuctionUseCase(t *testing.T) {
	mockRepo := new(repository.MockAuctionRepository)
	deleteAuctionUseCase := NewDeleteAuctionUseCase(mockRepo)

	input := &DeleteAuctionInputDTO{
		Id: 1,
	}

	mockRepo.On("DeleteAuction", input.Id).Return(nil)

	err := deleteAuctionUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
