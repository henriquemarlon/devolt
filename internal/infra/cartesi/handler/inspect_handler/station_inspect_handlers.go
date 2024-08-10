package inspect_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/station_usecase"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/rollmelette/rollmelette"
	"strconv"
)

type StationInspectHandlers struct {
	StationRepository entity.StationRepository
}

func NewStationInspectHandlers(stationRepository entity.StationRepository) *StationInspectHandlers {
	return &StationInspectHandlers{
		StationRepository: stationRepository,
	}
}

func (h *StationInspectHandlers) FindStationByIdHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	id, err := strconv.Atoi(router.PathValue(ctx, "id"))
	if err != nil {
		return fmt.Errorf("failed to parse id into int: %v", router.PathValue(ctx, "id"))
	}
	findStationById := station_usecase.NewFindStationByIdUseCase(h.StationRepository)
	res, err := findStationById.Execute(&station_usecase.FindStationByIdInputDTO{
		Id: uint(id),
	})
	if err != nil {
		return err
	}
	station, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal station: %w", err)
	}
	env.Report(station)
	return nil
}

func (h *StationInspectHandlers) FindAllStationsHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	findAllStationsUseCase := station_usecase.NewFindAllStationsUseCase(h.StationRepository)
	res, err := findAllStationsUseCase.Execute()
	if err != nil {
		return fmt.Errorf("failed to find all stations: %w", err)
	}
	allStations, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal all stations: %w", err)
	}
	env.Report(allStations)
	return nil
}
