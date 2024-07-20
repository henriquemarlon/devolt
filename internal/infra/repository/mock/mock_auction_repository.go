package mock

import (
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockAuctionRepository struct {
	mock.Mock
}

func (m *MockAuctionRepository) CreateAuction(input *entity.Auction) (*entity.Auction, error) {
	args := m.Called(input)
	return args.Get(0).(*entity.Auction), args.Error(1)
}

func (m *MockAuctionRepository) FindActiveAuction() (*entity.Auction, error) {
	args := m.Called()
	if auction := args.Get(0); auction != nil {
		return auction.(*entity.Auction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAuctionRepository) FindAuctionById(id uint) (*entity.Auction, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Auction), args.Error(1)
}

func (m *MockAuctionRepository) FindAllAuctions() ([]*entity.Auction, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Auction), args.Error(1)
}

func (m *MockAuctionRepository) UpdateAuction(input *entity.Auction) (*entity.Auction, error) {
	args := m.Called(input)
	return args.Get(0).(*entity.Auction), args.Error(1)
}

func (m *MockAuctionRepository) DeleteAuction(id uint) error {
	args := m.Called(id)
	if !args.Bool(0) {
		return fmt.Errorf("auction with ID %d does not exist", id)
	}
	return args.Error(1)
}