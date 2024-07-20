package mock

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(input *entity.Order) (*entity.Order, error) {
	args := m.Called(input)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockOrderRepository) FindAllOrders() ([]*entity.Order, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (m *MockOrderRepository) FindOrderById(id uint) (*entity.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockOrderRepository) FindOrdersByUser(buyer common.Address) ([]*entity.Order, error) {
	args := m.Called(buyer)
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateOrder(order *entity.Order) (*entity.Order, error) {
	args := m.Called(order)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (m *MockOrderRepository) DeleteOrder(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
