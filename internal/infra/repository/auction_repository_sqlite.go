package db

import (
	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"gorm.io/gorm"
)

type AuctionRepositorySqlite struct {
	Db *gorm.DB
}

func NewAuctionRepositorySqlite(db *gorm.DB) *AuctionRepositorySqlite {
	return &AuctionRepositorySqlite{
		Db: db,
	}
}

func (r *AuctionRepositorySqlite) CreateAuction(input *entity.Auction) (*entity.Auction, error) {
	err := r.Db.Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *AuctionRepositorySqlite) FindActiveAuction() (*entity.Auction, error) {
	var auction entity.Auction
	err := r.Db.Preload("Bids").Where("state = ?", "ongoing").First(&auction).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrAuctionNotFound
		}
		return nil, err
	}
	return &auction, nil
}

func (r *AuctionRepositorySqlite) FindAuctionById(id uint) (*entity.Auction, error) {
	var auction entity.Auction
	err := r.Db.Preload("Bids").First(&auction, id).Error
	if err != nil {
		return nil, err
	}
	return &auction, nil
}

func (r *AuctionRepositorySqlite) FindAllAuctions() ([]*entity.Auction, error) {
	var auctions []*entity.Auction
	err := r.Db.Preload("Bids").Find(&auctions).Error
	if err != nil {
		return nil, err
	}
	return auctions, nil
}

func (r *AuctionRepositorySqlite) UpdateAuction(input *entity.Auction) (*entity.Auction, error) {
	var auction entity.Auction
	err := r.Db.First(&auction, "id = ?", input.Id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrAuctionNotFound
		}
		return nil, err
	}

	auction.PriceLimitPerCredit = input.PriceLimitPerCredit
	auction.State = input.State
	auction.ExpiresAt = input.ExpiresAt
	auction.UpdatedAt = input.UpdatedAt

	res := r.Db.Save(auction)
	if res.Error != nil {
		return nil, res.Error
	}
	return &auction, nil
}

func (r *AuctionRepositorySqlite) DeleteAuction(id uint) error {
	err := r.Db.Delete(&entity.Auction{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
