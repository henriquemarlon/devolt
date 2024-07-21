package advance_handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/contract_usecase"
	"github.com/devolthq/devolt/internal/usecase/station_usecase"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type StationAdvanceHandlers struct {
	StationRepository  entity.StationRepository
	ContractRepository entity.ContractRepository
}

func NewStationAdvanceHandlers(
	stationRepository entity.StationRepository,
	contractRepository entity.ContractRepository,
) *StationAdvanceHandlers {
	return &StationAdvanceHandlers{
		StationRepository:  stationRepository,
		ContractRepository: contractRepository,
	}
}

func (h *StationAdvanceHandlers) CreateStationHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input station_usecase.CreateStationInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	createStation := station_usecase.NewCreateStationUseCase(h.StationRepository)
	input.Owner = strings.ToLower(input.Owner)
	res, err := createStation.Execute(&input, metadata)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("created station with id: %v, address: %v.", res.Id, res.Owner)))
	return nil
}

func (h *StationAdvanceHandlers) UpdateStationHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input station_usecase.UpdateStationInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	updateStation := station_usecase.NewUpdateStationUseCase(h.StationRepository)
	input.Owner = strings.ToLower(input.Owner)
	res, err := updateStation.Execute(&input, metadata)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("updated station with id: %v, address: %v and consumption: %v", res.Id, res.Owner, res.Consumption)))
	return nil
}

func (h *StationAdvanceHandlers) DeleteStationHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input station_usecase.DeleteStationInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	deleteStation := station_usecase.NewDeleteStationUseCase(h.StationRepository)
	err := deleteStation.Execute(&input)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("deleted station with id: %v", input.Id)))
	return nil
}

func (h *StationAdvanceHandlers) OffSetStationConsumptionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input station_usecase.OffSetStationConsumptionInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	offSetStationConsumption := station_usecase.NewOffSetStationConsumptionUseCase(h.StationRepository)
	res, err := offSetStationConsumption.Execute(&input, metadata)
	if err != nil {
		return err
	}
	application, isDefined := env.AppAddress()
	if !isDefined {
		return fmt.Errorf("no application address defined yet")
	}
	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
	if err != nil {
		return err
	}
	if err := env.ERC20Transfer(common.HexToAddress(volt.Address), application, metadata.MsgSender, input.CreditsToBeOffSet); err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("offSetCredits from station: %v by msg_sender: %v", res, metadata.MsgSender)))
	return nil
}
