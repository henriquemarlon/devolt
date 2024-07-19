package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindOrderByIdInputDTO struct {
	Id uint
}

type FindOrderByIdUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewFindOrderByIdUseCase(orderRepository entity.OrderRepository) *FindOrderByIdUseCase {
	return &FindOrderByIdUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *FindOrderByIdUseCase) Execute(input *FindOrderByIdInputDTO) (*FindOrderOutputDTO, error) {
	order, err := u.OrderRepository.FindOrderById(input.Id)
	if err != nil {
		return nil, err
	}
	return &FindOrderOutputDTO{
		Id:             order.Id,
		Buyer:          order.Buyer,
		Credits:        order.Credits,
		StationId:      order.StationId,
		PricePerCredit: order.PricePerCredit,
		CreatedAt:      order.CreatedAt,
		UpdatedAt:      order.UpdatedAt,
	}, nil
}
