package entity

// import (
// 	"math/big"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // AuctionRepositoryMock is a mock implementation of the AuctionRepository interface
// type AuctionRepositoryMock struct {
// 	mock.Mock
// }

// func (m *AuctionRepositoryMock) CreateAuction(auction *Auction) (*Auction, error) {
// 	args := m.Called(auction)
// 	return args.Get(0).(*Auction), args.Error(1)
// }

// func (m *AuctionRepositoryMock) FindActiveAuction() (*Auction, error) {
// 	args := m.Called()
// 	return args.Get(0).(*Auction), args.Error(1)
// }

// func (m *AuctionRepositoryMock) FindAuctionById(id uint) (*Auction, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*Auction), args.Error(1)
// }

// func (m *AuctionRepositoryMock) FindAllAuctions() ([]*Auction, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*Auction), args.Error(1)
// }

// func (m *AuctionRepositoryMock) UpdateAuction(auction *Auction) (*Auction, error) {
// 	args := m.Called(auction)
// 	return args.Get(0).(*Auction), args.Error(1)
// }

// func (m *AuctionRepositoryMock) DeleteAuction(id uint) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func TestNewAuction(t *testing.T) {
// 	credits := big.NewInt(1000)
// 	priceLimit := big.NewInt(500)
// 	state := "ongoing"
// 	expiresAt := time.Now().Add(time.Hour * 24).Unix()
// 	createdAt := time.Now().Unix()

// 	auction := NewAuction(credits, priceLimit, expiresAt, createdAt)

// 	assert.NotNil(t, auction)
// 	assert.Equal(t, credits, auction.Credits)
// 	assert.Equal(t, priceLimit, auction.PriceLimit)
// 	assert.Equal(t, state, auction.State)
// 	assert.Equal(t, expiresAt, auction.ExpiresAt)
// 	assert.Equal(t, createdAt, auction.CreatedAt)
// 	assert.Empty(t, auction.Bids)
// }

// func TestAuctionRepository(t *testing.T) {
// 	mockRepo := new(AuctionRepositoryMock)

// 	t.Run("CreateAuction", func(t *testing.T) {
// 		auction := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1000),
// 			PriceLimit: big.NewInt(500),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		mockRepo.On("CreateAuction", auction).Return(auction, nil)

// 		createdAuction, err := mockRepo.CreateAuction(auction)

// 		assert.Nil(t, err)
// 		assert.Equal(t, auction, createdAuction)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindActiveAuction", func(t *testing.T) {
// 		auction := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1000),
// 			PriceLimit: big.NewInt(500),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		mockRepo.On("FindActiveAuction").Return(auction, nil)

// 		activeAuction, err := mockRepo.FindActiveAuction()

// 		assert.Nil(t, err)
// 		assert.Equal(t, auction, activeAuction)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindAuctionById", func(t *testing.T) {
// 		auction := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1000),
// 			PriceLimit: big.NewInt(500),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		mockRepo.On("FindAuctionById", uint(1)).Return(auction, nil)

// 		foundAuction, err := mockRepo.FindAuctionById(1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, auction, foundAuction)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindAllAuctions", func(t *testing.T) {
// 		auction1 := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1000),
// 			PriceLimit: big.NewInt(500),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		auction2 := &Auction{
// 			Id:         2,
// 			Credits:    big.NewInt(2000),
// 			PriceLimit: big.NewInt(1000),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 48).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		auctions := []*Auction{auction1, auction2}

// 		mockRepo.On("FindAllAuctions").Return(auctions, nil)

// 		foundAuctions, err := mockRepo.FindAllAuctions()

// 		assert.Nil(t, err)
// 		assert.Equal(t, auctions, foundAuctions)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("UpdateAuction", func(t *testing.T) {
// 		auction := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1000),
// 			PriceLimit: big.NewInt(500),
// 			State:      "ongoing",
// 			ExpiresAt:  time.Now().Add(time.Hour * 24).Unix(),
// 			CreatedAt:  time.Now().Unix(),
// 		}

// 		updatedAuction := &Auction{
// 			Id:         1,
// 			Credits:    big.NewInt(1500),
// 			PriceLimit: big.NewInt(750),
// 			State:      "completed",
// 			ExpiresAt:  auction.ExpiresAt,
// 			CreatedAt:  auction.CreatedAt,
// 			UpdatedAt:  time.Now().Unix(),
// 		}

// 		mockRepo.On("UpdateAuction", updatedAuction).Return(updatedAuction, nil)

// 		resultAuction, err := mockRepo.UpdateAuction(updatedAuction)

// 		assert.Nil(t, err)
// 		assert.Equal(t, updatedAuction, resultAuction)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("DeleteAuction", func(t *testing.T) {
// 		mockRepo.On("DeleteAuction", uint(1)).Return(nil)

// 		err := mockRepo.DeleteAuction(1)

// 		assert.Nil(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }
