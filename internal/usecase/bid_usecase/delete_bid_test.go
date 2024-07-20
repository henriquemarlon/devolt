package bid_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteBidUseCase(t *testing.T) {
	mockBidRepo := new(repository.MockBidRepository)
	deleteBid := NewDeleteBidUseCase(mockBidRepo)

	mockBidRepo.On("DeleteBid", uint(1)).Return(nil)

	input := &DeleteBidInputDTO{
		Id: 1,
	}

	err := deleteBid.Execute(input)
	assert.Nil(t, err)

	mockBidRepo.AssertExpectations(t)
}
