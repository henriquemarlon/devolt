package inspect_handler

import (
	// "context"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/contract_usecase"
	"github.com/devolthq/devolt/internal/usecase/user_usecase"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type UserInspectHandlers struct {
	UserRepository     entity.UserRepository
	ContractRepository entity.ContractRepository
}

func NewUserInspectHandlers(userRepository entity.UserRepository, contractRepository entity.ContractRepository) *UserInspectHandlers {
	return &UserInspectHandlers{
		UserRepository:     userRepository,
		ContractRepository: contractRepository,
	}
}

func (h *UserInspectHandlers) FindUserByAddressHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	findUserByAddress := user_usecase.NewFindUserByAddressUseCase(h.UserRepository)
	res, err := findUserByAddress.Execute(&user_usecase.FindUserByAddressInputDTO{
		Address: common.BytesToAddress([]byte(router.PathValue(ctx, "address"))),
	})
	if err != nil {
		return fmt.Errorf("failed to find User: %w", err)
	}
	User, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal User: %w", err)
	}
	env.Report(User)
	return nil
}

func (h *UserInspectHandlers) FindAllUsersHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	findAllUsers := user_usecase.NewFindAllUsersUseCase(h.UserRepository)
	res, err := findAllUsers.Execute()
	if err != nil {
		return fmt.Errorf("failed to find all Users: %w", err)
	}
	allUsers, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal all Users: %w", err)
	}
	env.Report(allUsers)
	return nil
}

func (h *UserInspectHandlers) BalanceHandler(env rollmelette.EnvInspector, ctx context.Context) error {
	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	log.Printf("symbol: %s", router.PathValue(ctx, "symbol"))
	contract, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{
		Symbol: router.PathValue(ctx, "symbol"),
	})
	if err != nil {
		return fmt.Errorf("failed to find contract: %w", err)
	}
	balanceBytes, err := json.Marshal(env.ERC20BalanceOf(contract.Address, common.HexToAddress(router.PathValue(ctx, "address"))))
	if err != nil {
		return fmt.Errorf("failed to marshal balance: %w", err)
	}
	env.Report(balanceBytes)
	return nil
}
