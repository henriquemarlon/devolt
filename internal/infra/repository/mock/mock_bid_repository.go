package mock

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockBidRepository struct {
	mock.Mock
}

func (m *MockBidRepository) CreateBid(input *entity.Bid) (*entity.Bid, error) {
	args := m.Called(input)
	return args.Get(0).(*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) FindBidsByState(auctionId uint, state string) ([]*entity.Bid, error) {
	args := m.Called(auctionId, state)
	return args.Get(0).([]*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) FindBidById(id uint) (*entity.Bid, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) FindBidsByAuctionId(id uint) ([]*entity.Bid, error) {
	args := m.Called(id)
	return args.Get(0).([]*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) FindAllBids() ([]*entity.Bid, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) UpdateBid(input *entity.Bid) (*entity.Bid, error) {
	args := m.Called(input)
	return args.Get(0).(*entity.Bid), args.Error(1)
}

func (m *MockBidRepository) DeleteBid(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
