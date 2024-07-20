package order_usecase

import (
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"math/big"
)

type CreateOrderInputDTO struct {
	Buyer     common.Address `json:"buyer"`
	Credits   *big.Int       `json:"credits"`
	StationId string         `json:"station_id"`
}

type CreateOrderOutputDTO struct {
	Id             uint           `json:"id"`
	Buyer          common.Address `json:"buyer"`
	Credits        *big.Int       `json:"credits"`
	StationId      string         `json:"station_id"`
	StationOwner   common.Address `json:"station_address"`
	PricePerCredit *big.Int       `json:"price_per_credit"`
	CreatedAt      int64          `json:"created_at"`
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
	orderDeposit, ok := deposit.(*rollmelette.ERC20Deposit)
	if orderDeposit == nil || !ok {
		return nil, fmt.Errorf("unsupported deposit type for bid creation: %T", deposit)
	}

	stablecoin, err := u.ContractRepository.FindContractBySymbol("USDC")
	if err != nil {
		return nil, err
	}
	if stablecoin.Address != orderDeposit.Token {
		return nil, fmt.Errorf("invalid contract address provided for bid creation: %v", orderDeposit.Token)
	}

	station, err := u.StationRepository.FindStationById(input.StationId)
	if err != nil {
		return nil, err
	}

	paid := new(big.Int).Mul(station.PricePerCredit, input.Credits)
	if paid.Cmp(orderDeposit.Amount) == -1 {
		return nil, fmt.Errorf("order payment is less than station price")
	}

	order, err := entity.NewOrder(input.Buyer, input.Credits, input.StationId, station.PricePerCredit, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	order, err = u.OrderRepository.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return &CreateOrderOutputDTO{
		Id:             order.Id,
		Buyer:          order.Buyer,
		Credits:        order.Credits,
		StationId:      order.StationId,
		StationOwner:   station.Owner,
		PricePerCredit: order.PricePerCredit,
		CreatedAt:      order.CreatedAt,
	}, nil
}
