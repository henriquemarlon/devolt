package auction_usecase

import (
	"fmt"
	"math/big"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type CreateAuctionInputDTO struct {
	OrdersTimeRange     int64              `json:"orders_time_range"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit"`
	ExpiresAt           int64              `json:"expires_at"`
	CreatedAt           int64              `json:"created_at"`
}

type CreateAuctionOutputDTO struct {
	Id                  uint               `json:"id"`
	Credits             custom_type.BigInt `json:"credits"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit"`
	State               string             `json:"state"`
	ExpiresAt           int64              `json:"expires_at"`
	CreatedAt           int64              `json:"created_at"`
}

type CreateAuctionUseCase struct {
	OrderRepository  entity.OrderRepository
	DeviceRepository entity.AuctionRepository
}

func NewCreateAuctionUseCase(orderRepository entity.OrderRepository, deviceRepository entity.AuctionRepository) *CreateAuctionUseCase {
	return &CreateAuctionUseCase{
		OrderRepository:  orderRepository,
		DeviceRepository: deviceRepository,
	}
}

func (c *CreateAuctionUseCase) Execute(input *CreateAuctionInputDTO, metadata rollmelette.Metadata) (*CreateAuctionOutputDTO, error) {
	orders, err := c.OrderRepository.FindOrdersByTimeRange(metadata.BlockTimestamp-input.OrdersTimeRange, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("no orders found in time range")
	}

	credits := new(big.Int)
	for _, order := range orders {
		credits.Add(credits, order.Credits.Int)
	}

	if credits.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("no credits found in orders in time range to create auction")
	}

	auction, err := entity.NewAuction(custom_type.NewBigInt(credits), input.PriceLimitPerCredit, input.ExpiresAt, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	res, err := c.DeviceRepository.CreateAuction(auction)
	if err != nil {
		return nil, err
	}
	return &CreateAuctionOutputDTO{
		Id:                  res.Id,
		Credits:             res.Credits,
		PriceLimitPerCredit: res.PriceLimitPerCredit,
		State:               string(res.State),
		ExpiresAt:           res.ExpiresAt,
		CreatedAt:           res.CreatedAt,
	}, nil
}
