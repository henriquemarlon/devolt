package advance_handler

import (
	"encoding/json"
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/contract_usecase"
	"github.com/devolthq/devolt/internal/usecase/order_usecase"
	"github.com/rollmelette/rollmelette"
	"math/big"
)

type OrderAdvanceHandlers struct {
	OrderRepository    entity.OrderRepository
	StationRepository  entity.StationRepository
	ContractRepository entity.ContractRepository
}

func NewOrderAdvanceHandlers(
	orderRepository entity.OrderRepository,
	stationRepository entity.StationRepository,
	contractRepository entity.ContractRepository,
) *OrderAdvanceHandlers {
	return &OrderAdvanceHandlers{
		OrderRepository:    orderRepository,
		StationRepository:  stationRepository,
		ContractRepository: contractRepository,
	}
}

func (h *OrderAdvanceHandlers) CreateOrderHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input order_usecase.CreateOrderInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	createOrder := order_usecase.NewCreateOrderUseCase(h.OrderRepository, h.StationRepository, h.ContractRepository)
	res, err := createOrder.Execute(&input, deposit, metadata)
	if err != nil {
		return err
	}

	application, isDefined := env.AppAddress()
	if !isDefined {
		return fmt.Errorf("no application address defined yet")
	}

	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
	if err != nil {
		return err
	}

	// The station owner gets 40% of the order amount
	stationFee := new(big.Int).Div(new(big.Int).Mul(deposit.(*rollmelette.ERC20Deposit).Amount, big.NewInt(40)), big.NewInt(100))
	if err := env.ERC20Transfer(volt.Address.Address, metadata.MsgSender, res.StationOwner.Address, stationFee); err != nil {
		return err
	}

	// The application gets the remainder which would be split between the cost of the energy and DeVolt fees
	remainderValue := new(big.Int).Sub(deposit.(*rollmelette.ERC20Deposit).Amount, stationFee)
	if err := env.ERC20Transfer(volt.Address.Address, metadata.MsgSender, application, remainderValue); err != nil {
		return err
	}

	env.Notice([]byte(fmt.Sprintf("created order %v and paied %v as station fee and %v as application fee", res, stationFee, remainderValue)))
	return nil
}

func (h *OrderAdvanceHandlers) UpdateOrderHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input order_usecase.UpdateOrderInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	updateOrder := order_usecase.NewUpdateOrderUseCase(h.OrderRepository)
	res, err := updateOrder.Execute(&input, metadata)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("updated order %v", res)))
	return nil
}

func (h *OrderAdvanceHandlers) DeleteOrderHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input order_usecase.DeleteOrderInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	deleteOrder := order_usecase.NewDeleteOrderUseCase(h.OrderRepository)
	err := deleteOrder.Execute(&order_usecase.DeleteOrderInputDTO{
		Id: input.Id,
	})
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("deleted order %v", input.Id)))
	return nil
}
