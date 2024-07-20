package mock

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockContractRepository struct {
	mock.Mock
}

func (m *MockContractRepository) CreateContract(contract *entity.Contract) (*entity.Contract, error) {
	args := m.Called(contract)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Contract), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockContractRepository) FindAllContracts() ([]*entity.Contract, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entity.Contract), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockContractRepository) FindContractBySymbol(symbol string) (*entity.Contract, error) {
	args := m.Called(symbol)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Contract), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockContractRepository) UpdateContract(contract *entity.Contract) (*entity.Contract, error) {
	args := m.Called(contract)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Contract), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockContractRepository) DeleteContract(symbol string) error {
	args := m.Called(symbol)
	return args.Error(0)
}
