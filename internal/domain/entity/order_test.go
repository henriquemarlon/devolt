package entity

// import (
// 	"math/big"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/stretchr/testify/mock"
// )

// type MockOrderRepository struct {
// 	mock.Mock
// }

// func (m *MockOrderRepository) CreateOrder(order *Order) (*Order, error) {
// 	args := m.Called(order)
// 	return args.Get(0).(*Order), args.Error(1)
// }

// func TestCreateOrder(t *testing.T) {
// 	mockRepo := new(MockOrderRepository)
// 	order := NewOrder(common.HexToAddress("0xFb05c72178c0b88BFB8C5cFb8301e542A21aF1b7"), big.NewInt(100), "Station42", big.NewInt(50), 1622548000)

// 	mockRepo.On("CreateOrder", mock.Anything).Return(order, nil)

// 	createdOrder, err := mockRepo.CreateOrder(order)
// 	if err != nil {
// 		t.Errorf("expected no error, got %v", err)
// 	}

// 	if createdOrder != order {
// 		t.Errorf("expected order %v, got %v", order, createdOrder)
// 	}

// 	mockRepo.AssertExpectations(t)
// }
