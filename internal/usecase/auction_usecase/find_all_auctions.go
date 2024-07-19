package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindAllAuctionsOutputDTO []*FindAuctionOutputDTO

type FindAllAuctionsUseCase struct {
	AuctionRepository entity.AuctionRepository
}

func NewFindAllAuctionsUseCase(auctionRepository entity.AuctionRepository) *FindAllAuctionsUseCase {
	return &FindAllAuctionsUseCase{AuctionRepository: auctionRepository}
}

func (f *FindAllAuctionsUseCase) Execute() (*FindAllAuctionsOutputDTO, error) {
	res, err := f.AuctionRepository.FindAllAuctions()
	if err != nil {
		return nil, err
	}
	output := make(FindAllAuctionsOutputDTO, len(res))
	for i, auction := range res {
		bids := make([]*FindAuctionOutputSubDTO, len(auction.Bids))
		for j, bid := range auction.Bids {
			bids[j] = &FindAuctionOutputSubDTO{
				Id:        bid.Id,
				AuctionId: bid.AuctionId,
				Bidder:    bid.Bidder,
				Credits:   bid.Credits,
				Price:     bid.Price,
				State:     bid.State,
				CreatedAt: bid.CreatedAt,
				UpdatedAt: bid.UpdatedAt,
			}
		}
		output[i] = &FindAuctionOutputDTO{
			Id:         auction.Id,
			Credits:    auction.Credits,
			PriceLimit: auction.PriceLimit,
			State:      auction.State,
			Bids:       bids,
			ExpiresAt:  auction.ExpiresAt,
			CreatedAt:  auction.CreatedAt,
			UpdatedAt:  auction.UpdatedAt,
		}
	}
	return &output, nil
}
