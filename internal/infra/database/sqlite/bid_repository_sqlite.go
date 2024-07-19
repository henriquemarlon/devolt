package sqlite

import (
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"gorm.io/gorm"
)

type BidRepositorySqlite struct {
	Db *gorm.DB
}

func NewBidRepositorySqlite(db *gorm.DB) *BidRepositorySqlite {
	return &BidRepositorySqlite{
		Db: db,
	}
}

func (r *BidRepositorySqlite) CreateBid(input *entity.Bid) (*entity.Bid, error) {
	err := r.Db.Create(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create bid: %w", err)
	}
	return input, nil
}

func (r *BidRepositorySqlite) FindBidsByState(auctionId uint, state string) ([]*entity.Bid, error) {
	var bids []*entity.Bid
	err := r.Db.Where("auction_id = ? AND state = ?", auctionId, state).Find(&bids).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find bids by state: %w", err)
	}
	return bids, nil
}

func (r *BidRepositorySqlite) FindBidById(id uint) (*entity.Bid, error) {
	var bid entity.Bid
	err := r.Db.First(&bid, id).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find bid by ID: %w", err)
	}
	return &bid, nil
}

func (r *BidRepositorySqlite) FindBidsByAuctionId(id uint) ([]*entity.Bid, error) {
	var bids []*entity.Bid
	err := r.Db.Where("auction_id = ?", id).Find(&bids).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find bids by auction ID: %w", err)
	}
	return bids, nil
}

func (r *BidRepositorySqlite) FindAllBids() ([]*entity.Bid, error) {
	var bids []*entity.Bid
	err := r.Db.Find(&bids).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find all bids: %w", err)
	}
	return bids, nil
}

func (r *BidRepositorySqlite) UpdateBid(input *entity.Bid) (*entity.Bid, error) {
	err := r.Db.Save(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update bid: %w", err)
	}
	return input, nil
}

func (r *BidRepositorySqlite) DeleteBid(id uint) error {
	err := r.Db.Delete(&entity.Bid{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete bid: %w", err)
	}
	return nil
}
