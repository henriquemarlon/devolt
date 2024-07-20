package entity

// import (
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"math/big"
// 	"testing"
// 	"time"
// )

// // BidRepositoryMock is a mock implementation of the BidRepository interface
// type BidRepositoryMock struct {
// 	mock.Mock
// }

// func (m *BidRepositoryMock) CreateBid(bid *Bid) (*Bid, error) {
// 	args := m.Called(bid)
// 	return args.Get(0).(*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) FindBidsByState(auctionId uint, state string) ([]*Bid, error) {
// 	args := m.Called(auctionId, state)
// 	return args.Get(0).([]*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) FindBidById(id uint) (*Bid, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) FindBidsByAuctionId(id uint) ([]*Bid, error) {
// 	args := m.Called(id)
// 	return args.Get(0).([]*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) FindAllBids() ([]*Bid, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) UpdateBid(bid *Bid) (*Bid, error) {
// 	args := m.Called(bid)
// 	return args.Get(0).(*Bid), args.Error(1)
// }

// func (m *BidRepositoryMock) DeleteBid(id uint) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func TestNewBid(t *testing.T) {
// 	auctionId := uint(1)
// 	bidder := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
// 	credits := big.NewInt(1000)
// 	price := big.NewInt(500)
// 	state := "pending"
// 	createdAt := time.Now().Unix()

// 	bid := NewBid(auctionId, bidder, credits, price, state, createdAt)

// 	assert.NotNil(t, bid)
// 	assert.Equal(t, auctionId, bid.AuctionId)
// 	assert.Equal(t, bidder, bid.Bidder)
// 	assert.Equal(t, credits, bid.Credits)
// 	assert.Equal(t, price, bid.Price)
// 	assert.Equal(t, state, bid.State)
// 	assert.Equal(t, createdAt, bid.CreatedAt)
// }

// func TestBidRepository(t *testing.T) {
// 	mockRepo := new(BidRepositoryMock)

// 	t.Run("CreateBid", func(t *testing.T) {
// 		bid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "pending",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("CreateBid", bid).Return(bid, nil)

// 		createdBid, err := mockRepo.CreateBid(bid)

// 		assert.Nil(t, err)
// 		assert.Equal(t, bid, createdBid)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindBidsByState", func(t *testing.T) {
// 		bid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "accepted",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		bids := []*Bid{bid}

// 		mockRepo.On("FindBidsByState", uint(1), "accepted").Return(bids, nil)

// 		foundBids, err := mockRepo.FindBidsByState(1, "accepted")

// 		assert.Nil(t, err)
// 		assert.Equal(t, bids, foundBids)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindBidById", func(t *testing.T) {
// 		bid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "pending",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("FindBidById", uint(1)).Return(bid, nil)

// 		foundBid, err := mockRepo.FindBidById(1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, bid, foundBid)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindBidsByAuctionId", func(t *testing.T) {
// 		bid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "pending",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		bids := []*Bid{bid}

// 		mockRepo.On("FindBidsByAuctionId", uint(1)).Return(bids, nil)

// 		foundBids, err := mockRepo.FindBidsByAuctionId(1)

// 		assert.Nil(t, err)
// 		assert.Equal(t, bids, foundBids)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindAllBids", func(t *testing.T) {
// 		bid1 := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "pending",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		bid2 := &Bid{
// 			Id:        2,
// 			AuctionId: 2,
// 			Bidder:    common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
// 			Credits:   big.NewInt(2000),
// 			Price:     big.NewInt(1000),
// 			State:     "accepted",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		bids := []*Bid{bid1, bid2}

// 		mockRepo.On("FindAllBids").Return(bids, nil)

// 		foundBids, err := mockRepo.FindAllBids()

// 		assert.Nil(t, err)
// 		assert.Equal(t, bids, foundBids)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("UpdateBid", func(t *testing.T) {
// 		bid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			Credits:   big.NewInt(1000),
// 			Price:     big.NewInt(500),
// 			State:     "pending",
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		updatedBid := &Bid{
// 			Id:        1,
// 			AuctionId: 1,
// 			Bidder:    bid.Bidder,
// 			Credits:   big.NewInt(1500),
// 			Price:     big.NewInt(750),
// 			State:     "accepted",
// 			CreatedAt: bid.CreatedAt,
// 			UpdatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("UpdateBid", updatedBid).Return(updatedBid, nil)

// 		resultBid, err := mockRepo.UpdateBid(updatedBid)

// 		assert.Nil(t, err)
// 		assert.Equal(t, updatedBid, resultBid)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("DeleteBid", func(t *testing.T) {
// 		mockRepo.On("DeleteBid", uint(1)).Return(nil)

// 		err := mockRepo.DeleteBid(1)

// 		assert.Nil(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }
