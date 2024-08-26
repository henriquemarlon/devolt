package auction_usecase

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type FinishAuctionOutputDTO struct {
	Id                  uint               `json:"id"`
	RequiredCredits     custom_type.BigInt `json:"required_credits,omitempty"`
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
		log.Println("No bids placed for active auction, finishing auction without bids")
	}

	sort.Slice(bids, func(i, j int) bool {
		return bids[i].PricePerCredit.Cmp(bids[j].PricePerCredit.Int) < 0
	})

	requiredCreditsRemaining := new(big.Int).Set(activeAuction.RequiredCredits.Int)

	for _, bid := range bids {
		if requiredCreditsRemaining.Sign() == 0 {
			bid.State = "rejected"
			bid.UpdatedAt = metadata.BlockTimestamp
			_, err := u.BidRepository.UpdateBid(bid)
			if err != nil {
				return nil, err
			}
			continue
		}

		if requiredCreditsRemaining.Cmp(bid.Credits.Int) >= 0 {
			bid.State = "accepted"
			bid.UpdatedAt = metadata.BlockTimestamp
			_, err := u.BidRepository.UpdateBid(bid)
			if err != nil {
				return nil, err
			}
			requiredCreditsRemaining.Sub(requiredCreditsRemaining, bid.Credits.Int)
		} else {
			// Partially accept the bid
			partiallyAcceptedCredits := new(big.Int).Set(requiredCreditsRemaining)
			_, err := u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(partiallyAcceptedCredits),
				PricePerCredit: bid.PricePerCredit,
				State:          "partially_accepted",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			// Reject the remaining credits
			rejectedCredits := new(big.Int).Sub(bid.Credits.Int, partiallyAcceptedCredits)
			_, err = u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(rejectedCredits),
				PricePerCredit: bid.PricePerCredit,
				State:          "rejected",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			// Delete original bid
			err = u.BidRepository.DeleteBid(bid.Id)
			if err != nil {
				return nil, err
			}

			requiredCreditsRemaining.SetInt64(0)
		}
	}

	state := "finished"
	if requiredCreditsRemaining.Sign() > 0 {
		state = "partially_awarded"
	}

	activeAuction.State = entity.AuctionState(state)
	activeAuction.UpdatedAt = metadata.BlockTimestamp
	res, err := u.AuctionRepository.UpdateAuction(activeAuction)
	if err != nil {
		return nil, err
	}

	return &FinishAuctionOutputDTO{
		Id:                  res.Id,
		RequiredCredits:     res.RequiredCredits,
		PriceLimitPerCredit: res.PriceLimitPerCredit,
		State:               string(res.State),
		Bids:                bids,
		ExpiresAt:           res.ExpiresAt,
		CreatedAt:           res.CreatedAt,
		UpdatedAt:           res.UpdatedAt,
	}, nil
}
