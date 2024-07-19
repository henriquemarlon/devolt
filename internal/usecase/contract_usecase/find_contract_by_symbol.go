package contract_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
)

type FindContractBySymbolInputDTO struct {
	Symbol string
}

type FindContractBySymbolOutputDTO struct {
	Id        uint           `json:"id"`
	Symbol    string         `json:"symbol"`
	Address   common.Address `json:"address"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}

type FindContractBySymbolUseCase struct {
	ContractRepository entity.ContractRepository
}

func NewFindContractBySymbolUseCase(contractRepository entity.ContractRepository) *FindContractBySymbolUseCase {
	return &FindContractBySymbolUseCase{
		ContractRepository: contractRepository,
	}
}

func (s *FindContractBySymbolUseCase) Execute(input *FindContractBySymbolInputDTO) (*FindContractBySymbolOutputDTO, error) {
	contract, err := s.ContractRepository.FindContractBySymbol(input.Symbol)
	if err != nil {
		return nil, err
	}
	return &FindContractBySymbolOutputDTO{
		Id:        contract.Id,
		Symbol:    contract.Symbol,
		Address:   contract.Address,
		CreatedAt: contract.CreatedAt,
		UpdatedAt: contract.UpdatedAt,
	}, nil
}
