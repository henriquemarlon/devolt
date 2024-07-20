package contract_usecase

import "github.com/devolthq/devolt/internal/domain/entity"

type FindAllContractsOutputDTO []*FindContractOutputDTO

type FindAllContractsUsecase struct {
	ContractRepository entity.ContractRepository
}

func NewFindAllContractsUseCase(contractRepository entity.ContractRepository) *FindAllContractsUsecase {
	return &FindAllContractsUsecase{
		ContractRepository: contractRepository,
	}
}

func (s *FindAllContractsUsecase) Execute() (FindAllContractsOutputDTO, error) {
	res, err := s.ContractRepository.FindAllContracts()
	if err != nil {
		return nil, err
	}
	output := make([]*FindContractOutputDTO, 0, len(res))
	for i, contract := range res {
		output[i] = &FindContractOutputDTO{
			Id:        contract.Id,
			Symbol:    contract.Symbol,
			Address:   contract.Address,
			CreatedAt: contract.CreatedAt,
			UpdatedAt: contract.UpdatedAt,
		}
	}
	return output, nil
}
