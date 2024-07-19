package inspect_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/order_usecase"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type OrderInspectHandlers struct {
	OrderRepository entity.OrderRepository
}

func NewOrderInspectHandlers(orderRepository entity.OrderRepository) *OrderInspectHandlers {
	return &OrderInspectHandlers{
		OrderRepository: orderRepository,
	}
}

func (h *OrderInspectHandlers) FindAllOrdersHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	findAllOrders := order_usecase.NewFindAllOrderUseCase(h.OrderRepository)
	orders, err := findAllOrders.Execute()
	if err != nil {
		return err
	}
	ordersBytes, err := json.Marshal(orders)
	if err != nil {
		return fmt.Errorf("failed to marshal orders: %w", err)
	}
	env.Report(ordersBytes)
	return nil
}

func (h *OrderInspectHandlers) FindOrderByIdHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	id, err := strconv.Atoi(router.PathValue(ctx, "id"))
	if err != nil {
		return fmt.Errorf("failed to parse id into int: %v", router.PathValue(ctx, "id"))
	}
	findOrderById := order_usecase.NewFindOrderByIdUseCase(h.OrderRepository)
	res, err := findOrderById.Execute(&order_usecase.FindOrderByIdInputDTO{
		Id: uint(id),
	})
	if err != nil {
		return fmt.Errorf("failed to find order by id: %w from id: %s", err, router.PathValue(ctx, "id"))
	}
	orderBytes, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal order: %w", err)
	}
	env.Report(orderBytes)
	return nil
}

func (h *OrderInspectHandlers) FindOrdersByUserHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	user := common.HexToAddress(router.PathValue(ctx, "address"))
	findOrderByUser := order_usecase.NewFindOrdersByUserUseCase(h.OrderRepository)
	res, err := findOrderByUser.Execute(&order_usecase.FindOrderByUserInputDTO{
		User: user,
	})
	if err != nil {
		return fmt.Errorf("failed to find order by user: %w from user: %s", err, router.PathValue(ctx, "user"))
	}
	ordersBytes, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal orders: %w", err)
	}
	env.Report(ordersBytes)
	return nil
}
