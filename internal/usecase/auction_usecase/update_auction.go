package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type UpdateAuctionInputDTO struct {
	Id                  uint               `json:"id"`
	Credits             custom_type.BigInt `json:"credits"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit"`
	State               string             `json:"state"`
	ExpiresAt           int64              `json:"expires_at"`
}

type UpdateAuctionOutputDTO struct {
	Id                  uint               `json:"id"`
	Credits             custom_type.BigInt `json:"credits"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit"`
	State               string             `json:"state"`
	ExpiresAt           int64              `json:"expires_at"`
	UpdatedAt           int64              `json:"updated_at"`
}

type UpdateAuctionUseCase struct {
	AuctionRepository entity.AuctionRepository
}

func NewUpdateAuctionUseCase(auctionRepository entity.AuctionRepository) *UpdateAuctionUseCase {
	return &UpdateAuctionUseCase{AuctionRepository: auctionRepository}
}

func (u *UpdateAuctionUseCase) Execute(input *UpdateAuctionInputDTO, metadata rollmelette.Metadata) (*UpdateAuctionOutputDTO, error) {
	res, err := u.AuctionRepository.UpdateAuction(&entity.Auction{
		Id:                  input.Id,
		Credits:             input.Credits,
		PriceLimitPerCredit: input.PriceLimitPerCredit,
		State:               entity.AuctionState(input.State),
		ExpiresAt:           input.ExpiresAt,
		UpdatedAt:           metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateAuctionOutputDTO{
		Id:                  res.Id,
		Credits:             res.Credits,
		PriceLimitPerCredit: res.PriceLimitPerCredit,
		State:               string(res.State),
		ExpiresAt:           res.ExpiresAt,
		UpdatedAt:           res.UpdatedAt,
	}, nil
}
