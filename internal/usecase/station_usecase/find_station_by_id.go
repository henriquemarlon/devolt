package station_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindStationByIdInputDTO struct {
	Id uint `json:"id"`
}

type FindStationByIdUseCase struct {
	StationRepository entity.StationRepository
}

func NewFindStationByIdUseCase(stationRepository entity.StationRepository) *FindStationByIdUseCase {
	return &FindStationByIdUseCase{
		StationRepository: stationRepository,
	}
}

func (u *FindStationByIdUseCase) Execute(input *FindStationByIdInputDTO) (*FindStationOutputDTO, error) {
	res, err := u.StationRepository.FindStationById(input.Id)
	if err != nil {
		return nil, err
	}
	var orders []*FindStationOutputSubDTO
	for _, order := range res.Orders {
		orders = append(orders, &FindStationOutputSubDTO{
			Id:             order.Id,
			Buyer:          order.Buyer,
			StationId:      order.StationId,
			Credits:        order.Credits,
			PricePerCredit: order.PricePerCredit,
			CreatedAt:      order.CreatedAt,
			UpdatedAt:      order.UpdatedAt,
		})
	}
	return &FindStationOutputDTO{
		Id:             res.Id,
		Consumption:    res.Consumption,
		Owner:          res.Owner,
		PricePerCredit: res.PricePerCredit,
		State:          string(res.State),
		Orders:         orders,
		Latitude:       res.Latitude,
		Longitude:      res.Longitude,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}
