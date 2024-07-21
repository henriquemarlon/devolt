package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindOrderByUserInputDTO struct {
	User string
}

type FindOrderByUserOutputDTO []*FindOrderOutputDTO

type FindOrdersByUserUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewFindOrdersByUserUseCase(orderRepository entity.OrderRepository) *FindOrdersByUserUseCase {
	return &FindOrdersByUserUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *FindOrdersByUserUseCase) Execute(input *FindOrderByUserInputDTO) (FindOrderByUserOutputDTO, error) {
	res, err := u.OrderRepository.FindOrdersByUser(input.User)
	if err != nil {
		return nil, err
	}
	output := make(FindOrderByUserOutputDTO, len(res))
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
