package order_usecase

import (
	repository "github.com/devolthq/devolt/internal/infra/repository/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteOrderUseCase(t *testing.T) {
	mockOrderRepo := new(repository.MockOrderRepository)
	deleteOrder := NewDeleteOrderUseCase(mockOrderRepo)

	mockOrderRepo.On("DeleteOrder", uint(1)).Return(nil)

	input := &DeleteOrderInputDTO{
		Id: 1,
	}

	err := deleteOrder.Execute(input)
	assert.Nil(t, err)

	mockOrderRepo.AssertExpectations(t)
}
