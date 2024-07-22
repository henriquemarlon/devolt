package auction_usecase

import (
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
	"math/big"
	"sort"
)

type FinishAuctionSubDTO struct {
	Id        uint                `json:"id"`
	AuctionId uint                `json:"auction_id"`
	Bidder    custom_type.Address `json:"bidder"`
	Credits   custom_type.BigInt        `json:"credits"`
	Price     custom_type.BigInt        `json:"price"`
	CreatedAt int64               `json:"created_at"`
}

type FinishAuctionOutputDTO struct {
	Id uint `json:"id"`
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

type ByPriceCreditsRatioCreatedAt []*FinishAuctionSubDTO

func (a ByPriceCreditsRatioCreatedAt) Len() int {
	return len(a)
}

func (a ByPriceCreditsRatioCreatedAt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByPriceCreditsRatioCreatedAt) Less(i, j int) bool {
	pricePerCreditI := new(big.Int).Div(a[i].Price.Int, a[i].Credits.Int)
	pricePerCreditJ := new(big.Int).Div(a[j].Price.Int, a[j].Credits.Int)
	if pricePerCreditI.Cmp(pricePerCreditJ) == 0 {
		return a[i].CreatedAt < a[j].CreatedAt
	}
	return pricePerCreditI.Cmp(pricePerCreditJ) < 0
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

	totalCredits := big.NewInt(0)
	requiredCredits := activeAuction.Credits.Int

	var bidsDTO []*FinishAuctionSubDTO
	for _, bid := range bids {
		bidsDTO = append(bidsDTO, &FinishAuctionSubDTO{
			Id:        bid.Id,
			AuctionId: bid.AuctionId,
			Bidder:    bid.Bidder,
			Credits:   bid.Credits,
			Price:     bid.Price,
			CreatedAt: bid.CreatedAt,
		})
	}

	sort.Sort(ByPriceCreditsRatioCreatedAt(bidsDTO))

	for _, bid := range bidsDTO {
		if totalCredits.Cmp(requiredCredits) >= 0 {
			break
		}

		remainingCredits := new(big.Int).Sub(requiredCredits, totalCredits)
		if bid.Credits.Int.Cmp(remainingCredits) > 0 {
			acceptedBidCredits := custom_type.NewBigInt(remainingCredits)
			rejectedBidCredits := custom_type.NewBigInt(new(big.Int).Sub(bid.Credits.Int, remainingCredits))
			acceptedBidPrice := custom_type.NewBigInt(new(big.Int).Div(new(big.Int).Mul(bid.Price.Int, remainingCredits), bid.Credits.Int))
			rejectedBidPrice := custom_type.NewBigInt(new(big.Int).Div(new(big.Int).Mul(bid.Price.Int, rejectedBidCredits.Int), bid.Credits.Int))

			// Create bid with exactly required credits
			acceptedBid := &entity.Bid{
				AuctionId: bid.AuctionId,
				Bidder:    bid.Bidder,
				Credits:   acceptedBidCredits,
				Price:     acceptedBidPrice,
				State:     "partial_accepted",
				CreatedAt: bid.CreatedAt,
				UpdatedAt: metadata.BlockTimestamp,
			}

			// Create reject bid with remaining credits the exceed the required
			rejectedBid := &entity.Bid{
				AuctionId: bid.AuctionId,
				Bidder:    bid.Bidder,
				Credits:   rejectedBidCredits,
				Price:     rejectedBidPrice,
				State:     "rejected",
				CreatedAt: bid.CreatedAt,
				UpdatedAt: metadata.BlockTimestamp,
			}

			_, err = u.BidRepository.CreateBid(acceptedBid)
			if err != nil {
				return nil, err
			}

			_, err = u.BidRepository.CreateBid(rejectedBid)
			if err != nil {
				return nil, err
			}

			// Delete the original bid
			err = u.BidRepository.DeleteBid(bid.Id)
			if err != nil {
				return nil, err
			}

			totalCredits.Add(totalCredits, remainingCredits)
		} else {
			totalCredits.Add(totalCredits, bid.Credits.Int)
			_, err = u.BidRepository.UpdateBid(&entity.Bid{
				Id:        bid.Id,
				AuctionId: bid.AuctionId,
				Bidder:    bid.Bidder,
				Credits:   bid.Credits,
				Price:     bid.Price,
				State:     "accepted",
				UpdatedAt: metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	for _, bid := range bidsDTO {
		if bid.Credits.Int.Cmp(big.NewInt(0)) > 0 && bid.Credits.Int.Cmp(totalCredits) > 0 {
			_, err = u.BidRepository.UpdateBid(&entity.Bid{
				Id:        bid.Id,
				AuctionId: bid.AuctionId,
				Bidder:    bid.Bidder,
				Credits:   bid.Credits,
				Price:     bid.Price,
				State:     "rejected",
				UpdatedAt: metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	_, err = u.AuctionRepository.UpdateAuction(&entity.Auction{
		Id:        activeAuction.Id,
		State:     "finished",
		UpdatedAt: metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &FinishAuctionOutputDTO{Id: activeAuction.Id}, nil
}
