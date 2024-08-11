package order_usecase

import (
	"github.com/Mugen-Builders/devolt/internal/domain/entity"
)

type FindAllOrdersOutputDTO []*FindOrderOutputDTO

type FindAllOrderUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewFindAllOrderUseCase(orderRepository entity.OrderRepository) *FindAllOrderUseCase {
	return &FindAllOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *FindAllOrderUseCase) Execute() (FindAllOrdersOutputDTO, error) {
	res, err := u.OrderRepository.FindAllOrders()
	if err != nil {
		return nil, err
	}
	output := make(FindAllOrdersOutputDTO, len(res))
	for i, order := range res {
		output[i] = &FindOrderOutputDTO{
			Id:             order.Id,
			Buyer:          order.Buyer,
			Credits:        order.Credits,
			StationId:      order.StationId,
			PricePerCredit: order.PricePerCredit,
			CreatedAt:      order.CreatedAt,
			UpdatedAt:      order.UpdatedAt,
		}
	}
	return output, nil
}
