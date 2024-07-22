package auction_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type CreateAuctionInputDTO struct {
	Credits    custom_type.BigInt `json:"credits"`
	PriceLimit custom_type.BigInt `json:"price_limit"`
	ExpiresAt  int64              `json:"expires_at"`
	CreatedAt  int64              `json:"created_at"`
}

type CreateAuctionOutputDTO struct {
	Id         uint               `json:"id"`
	Credits    custom_type.BigInt `json:"credits"`
	PriceLimit custom_type.BigInt `json:"price_limit"`
	State      string             `json:"state"`
	ExpiresAt  int64              `json:"expires_at"`
	CreatedAt  int64              `json:"created_at"`
}

type CreateAuctionUseCase struct {
	DeviceRepository entity.AuctionRepository
}

func NewCreateAuctionUseCase(deviceRepository entity.AuctionRepository) *CreateAuctionUseCase {
	return &CreateAuctionUseCase{DeviceRepository: deviceRepository}
}

func (c *CreateAuctionUseCase) Execute(input *CreateAuctionInputDTO, metadata rollmelette.Metadata) (*CreateAuctionOutputDTO, error) {
	auction, err := entity.NewAuction(input.Credits, input.PriceLimit, input.ExpiresAt, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	res, err := c.DeviceRepository.CreateAuction(auction)
	if err != nil {
		return nil, err
	}
	return &CreateAuctionOutputDTO{
		Id:         res.Id,
		Credits:    res.Credits,
		PriceLimit: res.PriceLimit,
		State:      string(res.State),
		ExpiresAt:  res.ExpiresAt,
		CreatedAt:  res.CreatedAt,
	}, nil
}
