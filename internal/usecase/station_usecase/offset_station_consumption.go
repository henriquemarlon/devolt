package station_usecase

import (
	"fmt"
	"math/big"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type OffSetStationConsumptionInputDTO struct {
	Id                string   `json:"id"`
	CreditsToBeOffSet *big.Int `json:"credits_to_be_offSet"`
}

type OffSetStationConsumptionOutputDTO struct {
	Id             string         `json:"id"`
	Consumption    *big.Int       `json:"consumption"`
	Owner          common.Address `json:"owner"`
	PricePerCredit *big.Int       `json:"price_per_credit"`
	State          string         `json:"state"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	UpdatedAt      int64          `json:"updated_at"`
}

type OffSetStationConsumptionUseCase struct {
	StationRepository entity.StationRepository
}

func NewOffSetStationConsumptionUseCase(stationRepository entity.StationRepository) *OffSetStationConsumptionUseCase {
	return &OffSetStationConsumptionUseCase{
		StationRepository: stationRepository,
	}
}

func (u *OffSetStationConsumptionUseCase) Execute(input *OffSetStationConsumptionInputDTO, metadata rollmelette.Metadata) (*OffSetStationConsumptionOutputDTO, error) {
	station, err := u.StationRepository.FindStationById(input.Id)
	if err != nil {
		return nil, err
	}
	if station.Owner != metadata.MsgSender {
		return nil, fmt.Errorf("can't offSet station consumption, because the station owner is not equal to the msg_sender address, expected: %v, got: %v", station.Owner, metadata.MsgSender)
	}

	consumption := new(big.Int)
	consumption.Sub(station.Consumption, input.CreditsToBeOffSet)

	res, err := u.StationRepository.UpdateStation(&entity.Station{
		Id:          input.Id,
		Consumption: consumption,
		UpdatedAt:   metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &OffSetStationConsumptionOutputDTO{
		Id:             res.Id,
		Consumption:    res.Consumption,
		Owner:          res.Owner,
		PricePerCredit: res.PricePerCredit,
		State:          res.State,
		Latitude:       res.Latitude,
		Longitude:      res.Longitude,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}
