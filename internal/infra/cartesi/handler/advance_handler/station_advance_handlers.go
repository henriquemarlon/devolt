package advance_handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/internal/usecase/contract_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/station_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/user_usecase"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type StationAdvanceHandlers struct {
	UserRepository     entity.UserRepository
	StationRepository  entity.StationRepository
	ContractRepository entity.ContractRepository
}

func NewStationAdvanceHandlers(
	userRepository entity.UserRepository,
	stationRepository entity.StationRepository,
	contractRepository entity.ContractRepository,
) *StationAdvanceHandlers {
	return &StationAdvanceHandlers{
		UserRepository:     userRepository,
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
	res, err := createStation.Execute(&input, metadata)
	if err != nil {
		return err
	}
	station, err := json.Marshal(res)
	if err != nil {
		return err
	}
	env.Notice(append([]byte("created station - "), station...))
	return nil
}

func (h *StationAdvanceHandlers) UpdateStationHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input station_usecase.UpdateStationInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return fmt.Errorf("failed to unmarshal input: %w", err)
	}
	updateStation := station_usecase.NewUpdateStationUseCase(h.StationRepository)
	res, err := updateStation.Execute(&input, metadata)
	if err != nil {
		return err
	}
	station, err := json.Marshal(res)
	if err != nil {
		return err
	}
	env.Notice(append([]byte("updated station - "), station...))
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
	station, err := json.Marshal(input)
	if err != nil {
		return err
	}
	env.Notice(append([]byte("deleted station with - "), station...))
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

	findUserByRole := user_usecase.NewFindUserByRoleUseCase(h.UserRepository)
	auctioneer, err := findUserByRole.Execute(&user_usecase.FindUserByRoleInputDTO{Role: "auctioneer"})
	if err != nil {
		return err
	}

	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
	if err != nil {
		return err
	}

	if err := env.ERC20Transfer(volt.Address.Address, auctioneer.Address.Address, common.HexToAddress("0x1"), input.CreditsToBeOffSet.Int); err != nil {
		return err
	}

	abiJson := `[{
		"type":"function",
		"name":"burn",
		"outputs":[],
		"stateMutability":"nonpayable",
		"inputs":[{
			"internalType":"uint256",
			"name":"value",
			"type": "uint256"
		}]
	}]`
	abiInterface, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return err
	}
	payload, err = abiInterface.Pack("burn", input.CreditsToBeOffSet.Int)
	if err != nil {
		return err
	}

	env.Voucher(volt.Address.Address, payload)
	env.Notice([]byte(fmt.Sprintf("offSet Credits from station: %v by msg_sender: %v", res.Id, metadata.MsgSender)))
	return nil
}