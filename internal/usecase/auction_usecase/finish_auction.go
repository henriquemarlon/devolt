package auction_usecase

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type FinishAuctionOutputDTO struct {
	Id                  uint               `json:"id"`
	Credits             custom_type.BigInt `json:"credits,omitempty"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit,omitempty"`
	State               string             `json:"state,omitempty"`
	Bids                []*entity.Bid      `json:"bids,omitempty"`
	ExpiresAt           int64              `json:"expires_at,omitempty"`
	CreatedAt           int64              `json:"created_at,omitempty"`
	UpdatedAt           int64              `json:"updated_at,omitempty"`
}

type FinishAuctionUseCase struct {
	BidRepository     entity.BidRepository
	AuctionRepository entity.AuctionRepository
}

func NewFinishAuctionUseCase(auctionRepository entity.AuctionRepository, bidRepository entity.BidRepository) *FinishAuctionUseCase {
	return &FinishAuctionUseCase{
		AuctionRepository: auctionRepository,
		BidRepository:     bidRepository,
	}
}

func (u *FinishAuctionUseCase) Execute(metadata rollmelette.Metadata) (*FinishAuctionOutputDTO, error) {
	activeAuction, err := u.AuctionRepository.FindActiveAuction()
	if err != nil {
		return nil, err
	}

	if metadata.BlockTimestamp < activeAuction.ExpiresAt {
		return nil, fmt.Errorf("active auction not expired, you can't finish it yet")
	}

	bids, err := u.BidRepository.FindBidsByAuctionId(activeAuction.Id)
	if err != nil {
		return nil, err
	}

	if len(bids) == 0 {
		log.Println("no bids placed for active auction, finishing auction without bids")
	}

	sort.Slice(bids, func(i, j int) bool {
		return bids[i].PricePerCredit.Cmp(bids[j].PricePerCredit.Int) < 0
	})

	requireCreditsRemaining := activeAuction.Credits

	for _, bid := range bids {
		if requireCreditsRemaining.Cmp(big.NewInt(0)) == 0 {
			_, err := u.BidRepository.UpdateBid(&entity.Bid{
				Id:        bid.Id,
				State:     "rejected",
				UpdatedAt: metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
			continue
		}

		if requireCreditsRemaining.Cmp(bid.Credits.Int) >= 0 {
			_, err := u.BidRepository.UpdateBid(&entity.Bid{
				Id:        bid.Id,
				State:     "accepted",
				UpdatedAt: metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
			requireCreditsRemaining.Sub(requireCreditsRemaining.Int, bid.Credits.Int)
			continue
		}

		if bid.Credits.Int.Cmp(requireCreditsRemaining.Int) == 1 {
			remainingCredits := new(big.Int).Set(requireCreditsRemaining.Int)

			// Create the partially accepted bid
			res, err := u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(remainingCredits),
				PricePerCredit: bid.PricePerCredit,
				State:          "partial_accepted",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			// Create the rejected part of the bid
			_, err = u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(new(big.Int).Sub(bid.Credits.Int, remainingCredits)),
				PricePerCredit: bid.PricePerCredit,
				State:          "rejected",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			requireCreditsRemaining.Sub(requireCreditsRemaining.Int, res.Credits.Int)

			// Delete the original bid
			err = u.BidRepository.DeleteBid(bid.Id)
			if err != nil {
				return nil, err
			}
			continue
		}
	}

	res, err := u.AuctionRepository.UpdateAuction(&entity.Auction{
		Id:        activeAuction.Id,
		State:     "finished",
		UpdatedAt: metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &FinishAuctionOutputDTO{
		Id:                  res.Id,
		Credits:             res.Credits,
		PriceLimitPerCredit: res.PriceLimitPerCredit,
		State:               string(res.State),
		Bids:                bids,
		ExpiresAt:           res.ExpiresAt,
		CreatedAt:           res.CreatedAt,
		UpdatedAt:           res.UpdatedAt,
	}, nil
}
