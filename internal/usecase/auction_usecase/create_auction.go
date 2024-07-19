package auction_usecase

import (
	"math/big"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/rollmelette/rollmelette"
)

type CreateAuctionInputDTO struct {
	Credits    *big.Int `json:"credits"`
	PriceLimit *big.Int `json:"price_limit"`
	State      string   `json:"state"`
	ExpiresAt  int64    `json:"expires_at"`
	CreatedAt  int64    `json:"created_at"`
}

type CreateAuctionOutputDTO struct {
	Id         uint     `json:"id"`
	Credits    *big.Int `json:"credits"`
	PriceLimit *big.Int `json:"price_limit"`
	State      string   `json:"state"`
	ExpiresAt  int64    `json:"expires_at"`
	CreatedAt  int64    `json:"created_at"`
}

type CreateAuctionUseCase struct {
	DeviceRepository entity.AuctionRepository
}

func NewCreateAuctionUseCase(deviceRepository entity.AuctionRepository) *CreateAuctionUseCase {
	return &CreateAuctionUseCase{DeviceRepository: deviceRepository}
}

func (c *CreateAuctionUseCase) Execute(input *CreateAuctionInputDTO, metadata rollmelette.Metadata) (*CreateAuctionOutputDTO, error) {
	auction := entity.NewAuction(input.Credits, input.PriceLimit, "ongoing", input.ExpiresAt, metadata.BlockTimestamp)
	res, err := c.DeviceRepository.CreateAuction(auction)
	if err != nil {
		return nil, err
	}
	return &CreateAuctionOutputDTO{
		Id:         res.Id,
		Credits:    res.Credits,
		PriceLimit: res.PriceLimit,
		State:      res.State,
		ExpiresAt:  res.ExpiresAt,
		CreatedAt:  res.CreatedAt,
	}, nil
}
