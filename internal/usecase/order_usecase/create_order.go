package order_usecase

import (
	"fmt"
	"math/big"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type CreateOrderInputDTO struct {
	StationId uint `json:"station_id"`
}

type CreateOrderOutputDTO struct {
	Id             uint                `json:"id"`
	Buyer          custom_type.Address `json:"buyer"`
	Credits        custom_type.BigInt  `json:"credits"`
	StationId      uint                `json:"station_id"`
	StationOwner   custom_type.Address `json:"station_address"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	CreatedAt      int64               `json:"created_at"`
}

type CreateOrderUseCase struct {
	OrderRepository    entity.OrderRepository
	StationRepository  entity.StationRepository
	ContractRepository entity.ContractRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository, stationRepository entity.StationRepository, contractRepository entity.ContractRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:    orderRepository,
		StationRepository:  stationRepository,
		ContractRepository: contractRepository,
	}
}

func (u *CreateOrderUseCase) Execute(input *CreateOrderInputDTO, deposit rollmelette.Deposit, metadata rollmelette.Metadata) (*CreateOrderOutputDTO, error) {
	stablecoin, err := u.ContractRepository.FindContractBySymbol("STABLECOIN")
	if err != nil {
		return nil, err
	}
	if stablecoin.Address.Address != deposit.(*rollmelette.ERC20Deposit).Token {
		return nil, fmt.Errorf("invalid contract address provided for order creation: %v", deposit.(*rollmelette.ERC20Deposit).Token)
	}

	station, err := u.StationRepository.FindStationById(input.StationId)
	if err != nil {
		return nil, err
	}

	orderConsumption := new(big.Int).Div(deposit.(*rollmelette.ERC20Deposit).Amount, station.PricePerCredit.Int)

	order, err := entity.NewOrder(custom_type.NewAddress(deposit.(*rollmelette.ERC20Deposit).Sender), custom_type.NewBigInt(orderConsumption), input.StationId, station.PricePerCredit.Int, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	order, err = u.OrderRepository.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	station, err = u.StationRepository.FindStationById(input.StationId)
	if err != nil {
		return nil, err
	}

	consumption := custom_type.NewBigInt(new(big.Int).Add(station.Consumption.Int, orderConsumption))
	_, err = u.StationRepository.UpdateStation(&entity.Station{
		Id:          input.StationId,
		Consumption: consumption,
		UpdatedAt:   metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutputDTO{
		Id:             order.Id,
		Buyer:          custom_type.NewAddress(deposit.(*rollmelette.ERC20Deposit).Sender),
		Credits:        order.Credits,
		StationId:      order.StationId,
		StationOwner:   station.Owner,
		PricePerCredit: order.PricePerCredit,
		CreatedAt:      order.CreatedAt,
	}, nil
}
