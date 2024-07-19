package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindAuctionByIdInputDTO struct {
	Id uint `json:"id"`
}

type FindAuctionByIdUseCase struct {
	AuctionRepository entity.AuctionRepository
}

func NewFindAuctionByIdUseCase(auctionRepository entity.AuctionRepository) *FindAuctionByIdUseCase {
	return &FindAuctionByIdUseCase{AuctionRepository: auctionRepository}
}

func (f *FindAuctionByIdUseCase) Execute(input *FindAuctionByIdInputDTO) (*FindAuctionOutputDTO, error) {
	res, err := f.AuctionRepository.FindAuctionById(input.Id)
	if err != nil {
		return nil, err
	}
	var bids []*FindAuctionOutputSubDTO
	for _, bid := range res.Bids {
		bids = append(bids, &FindAuctionOutputSubDTO{
			Id:        bid.Id,
			AuctionId: bid.AuctionId,
			Bidder:    bid.Bidder,
			Credits:   bid.Credits,
			Price:     bid.Price,
			State:     bid.State,
			CreatedAt: bid.CreatedAt,
			UpdatedAt: bid.UpdatedAt,
		})
	}
	return &FindAuctionOutputDTO{
		Id:         res.Id,
		Credits:    res.Credits,
		PriceLimit: res.PriceLimit,
		State:      res.State,
		Bids:       bids,
		ExpiresAt:  res.ExpiresAt,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}, nil
}
