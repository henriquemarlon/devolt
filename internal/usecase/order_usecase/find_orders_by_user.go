package order_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
)

type FindOrderByUserInputDTO struct {
	User common.Address `json:"user"`
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
	res, err := u.OrderRepository.FindOrdersByUser(custom_type.NewAddress(input.User))
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
