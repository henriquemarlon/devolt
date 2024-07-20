package order_usecase

import (
	"testing"

	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteOrderUseCase(t *testing.T) {
	mockRepo := new(repository.MockOrderRepository)
	deleteOrderUseCase := NewDeleteOrderUseCase(mockRepo)

	input := &DeleteOrderInputDTO{
		Id: 1,
	}

	mockRepo.On("DeleteOrder", input.Id).Return(nil)

	err := deleteOrderUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
