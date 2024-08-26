package station_usecase

import (
	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type CreateStationInputDTO struct {
	Owner          custom_type.Address `json:"owner"`
	Consumption    custom_type.BigInt  `json:"consumption"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	Latitude       float64             `json:"latitude"`
	Longitude      float64             `json:"longitude"`
}

type CreateStationOutputDTO struct {
	Id             uint                `json:"id"`
	Owner          custom_type.Address `json:"owner"`
	State          string              `json:"state"`
	Consumption    custom_type.BigInt  `json:"consumption"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	Latitude       float64             `json:"latitude"`
	Longitude      float64             `json:"longitude"`
	CreatedAt      int64               `json:"created_at"`
}

type CreateStationUseCase struct {
	StationRepository entity.StationRepository
}

func NewCreateStationUseCase(stationRepository entity.StationRepository) *CreateStationUseCase {
	return &CreateStationUseCase{
		StationRepository: stationRepository,
	}
}

func (u *CreateStationUseCase) Execute(input *CreateStationInputDTO, metadata rollmelette.Metadata) (*CreateStationOutputDTO, error) {
	station, err := entity.NewStation(input.Owner, input.Consumption, input.PricePerCredit, input.Latitude, input.Longitude, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	res, err := u.StationRepository.CreateStation(station)
	if err != nil {
		return nil, err
	}
	return &CreateStationOutputDTO{
		Id:             res.Id,
		Owner:          res.Owner,
		PricePerCredit: res.PricePerCredit,
		State:          string(res.State),
		Latitude:       res.Latitude,
		Longitude:      res.Longitude,
		CreatedAt:      res.CreatedAt,
	}, nil
}
