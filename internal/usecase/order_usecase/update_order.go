package order_usecase

import (
	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type UpdateOrderInputDTO struct {
	Id             uint               `json:"id"`
	Buyer          common.Address     `json:"buyer"`
	Credits        custom_type.BigInt `json:"credits"`
	StationId      uint               `json:"station_id"`
	PricePerCredit custom_type.BigInt `json:"price_per_credit"`
}

type UpdateOrderOutputDTO struct {
	Id             uint                `json:"id"`
	Buyer          custom_type.Address `json:"buyer"`
	Credits        custom_type.BigInt  `json:"credits"`
	StationId      uint                `json:"station_id"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	UpdatedAt      int64               `json:"updated_at"`
}

type UpdateOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewUpdateOrderUseCase(orderRepository entity.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *UpdateOrderUseCase) Execute(input *UpdateOrderInputDTO, metadata rollmelette.Metadata) (*UpdateOrderOutputDTO, error) {
	order, err := u.OrderRepository.UpdateOrder(&entity.Order{
		Id:             input.Id,
		Buyer:          custom_type.NewAddress(input.Buyer),
		Credits:        input.Credits,
		StationId:      input.StationId,
		PricePerCredit: input.PricePerCredit,
		UpdatedAt:      metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateOrderOutputDTO{
		Id:             order.Id,
		Buyer:          order.Buyer,
		Credits:        order.Credits,
		StationId:      order.StationId,
		PricePerCredit: order.PricePerCredit,
		UpdatedAt:      order.UpdatedAt,
	}, nil
}
