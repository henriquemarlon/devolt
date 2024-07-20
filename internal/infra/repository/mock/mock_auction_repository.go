package mock

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockAuctionRepository struct {
	mock.Mock
}

func (m *MockAuctionRepository) CreateAuction(input *entity.Auction) (*entity.Auction, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) FindActiveAuction() (*entity.Auction, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) FindAuctionById(id uint) (*entity.Auction, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) FindAllAuctions() ([]*entity.Auction, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) UpdateAuction(input *entity.Auction) (*entity.Auction, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) DeleteAuction(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
