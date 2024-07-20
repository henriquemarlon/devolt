package bid_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteBidUseCase(t *testing.T) {
	mockRepo := new(repository.MockBidRepository)
	deleteBidUseCase := NewDeleteBidUseCase(mockRepo)

	input := &DeleteBidInputDTO{
		Id: 1,
	}

	mockRepo.On("DeleteBid", input.Id).Return(nil)

	err := deleteBidUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
