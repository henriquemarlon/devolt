package station_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindAllStationsOutputDTO []*FindStationOutputDTO

type FindAllStationsUseCase struct {
	StationReposiory entity.StationRepository
}

func NewFindAllStationsUseCase(stationRepository entity.StationRepository) *FindAllStationsUseCase {
	return &FindAllStationsUseCase{
		StationReposiory: stationRepository,
	}
}

func (c *FindAllStationsUseCase) Execute() ([]*FindStationOutputDTO, error) {
	res, err := c.StationReposiory.FindAllStations()
	if err != nil {
		return nil, err
	}
	output := make([]*FindStationOutputDTO, len(res))
	for i, station := range res {
		orders := make([]*FindStationOutputSubDTO, len(station.Orders))
		for j, order := range station.Orders {
			orders[j] = &FindStationOutputSubDTO{
				Id:             order.Id,
				Buyer:          order.Buyer,
				StationId:      order.StationId,
				Credits:        order.Credits,
				PricePerCredit: order.PricePerCredit,
				CreatedAt:      order.CreatedAt,
				UpdatedAt:      order.UpdatedAt,
			}
		}
		output[i] = &FindStationOutputDTO{
			Id:             station.Id,
			Consumption:    station.Consumption,
			Owner:          station.Owner,
			PricePerCredit: station.PricePerCredit,
			State:          station.State,
			Orders:         orders,
			Latitude:       station.Latitude,
			Longitude:      station.Longitude,
			CreatedAt:      station.CreatedAt,
			UpdatedAt:      station.UpdatedAt,
		}
	}
	return output, nil
}
