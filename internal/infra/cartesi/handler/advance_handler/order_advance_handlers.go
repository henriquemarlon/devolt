package advance_handler

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/internal/usecase/contract_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/order_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/user_usecase"
	"github.com/rollmelette/rollmelette"
)

type OrderAdvanceHandlers struct {
	UserRepository     entity.UserRepository
	OrderRepository    entity.OrderRepository
	StationRepository  entity.StationRepository
	ContractRepository entity.ContractRepository
}

func NewOrderAdvanceHandlers(
	userRepository entity.UserRepository,
	orderRepository entity.OrderRepository,
	stationRepository entity.StationRepository,
	contractRepository entity.ContractRepository,
) *OrderAdvanceHandlers {
	return &OrderAdvanceHandlers{
		UserRepository:     userRepository,
		OrderRepository:    orderRepository,
		StationRepository:  stationRepository,
		ContractRepository: contractRepository,
	}
}

func (h *OrderAdvanceHandlers) CreateOrderHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	switch deposit := deposit.(type) {
	case *rollmelette.ERC20Deposit:
		var input order_usecase.CreateOrderInputDTO
		if err := json.Unmarshal(payload, &input); err != nil {
			return fmt.Errorf("failed to unmarshal input: %w", err)
		}
		createOrder := order_usecase.NewCreateOrderUseCase(h.OrderRepository, h.StationRepository, h.ContractRepository)
		res, err := createOrder.Execute(&input, deposit, metadata)
		if err != nil {
			return err
		}

		findUserByRole := user_usecase.NewFindUserByRoleUseCase(h.UserRepository)
		auctioneer, err := findUserByRole.Execute(&user_usecase.FindUserByRoleInputDTO{Role: "auctioneer"})
		if err != nil {
			return err
		}

		findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
		stablecoin, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "STABLECOIN"})
		if err != nil {
			return err
		}

		// The station owner gets 40% of the order amount
		stationFee := new(big.Int).Div(new(big.Int).Mul(deposit.Amount, big.NewInt(40)), big.NewInt(100))
		if err := env.ERC20Transfer(stablecoin.Address.Address, res.Buyer.Address, res.StationOwner.Address, stationFee); err != nil {
			return err
		}

		// The application gets the remainder which would be split between the cost of the energy and DeVolt fees
		remainderValue := new(big.Int).Sub(deposit.Amount, stationFee)
		if err := env.ERC20Transfer(stablecoin.Address.Address, deposit.Sender, auctioneer.Address.Address, remainderValue); err != nil {
			return err
		}

		order, err := json.Marshal(res)
		if err != nil {
			return err
		}
		env.Notice(append([]byte("created order - "), order...))
		return nil
	default:
		return fmt.Errorf("unsupported deposit type: %T", deposit)
	}
}
