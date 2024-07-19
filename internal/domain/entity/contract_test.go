package entity

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// ContractRepositoryMock is a mock implementation of the ContractRepository interface
type ContractRepositoryMock struct {
	mock.Mock
}

func (m *ContractRepositoryMock) CreateContract(contract *Contract) (*Contract, error) {
	args := m.Called(contract)
	return args.Get(0).(*Contract), args.Error(1)
}

func (m *ContractRepositoryMock) FindAllContracts() ([]*Contract, error) {
	args := m.Called()
	return args.Get(0).([]*Contract), args.Error(1)
}

func (m *ContractRepositoryMock) FindContractBySymbol(symbol string) (*Contract, error) {
	args := m.Called(symbol)
	return args.Get(0).(*Contract), args.Error(1)
}

func (m *ContractRepositoryMock) UpdateContract(contract *Contract) (*Contract, error) {
	args := m.Called(contract)
	return args.Get(0).(*Contract), args.Error(1)
}

func (m *ContractRepositoryMock) DeleteContract(symbol string) error {
	args := m.Called(symbol)
	return args.Error(0)
}

func TestNewContract(t *testing.T) {
	symbol := "ETH"
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	createdAt := time.Now().Unix()

	contract := NewContract(symbol, address, createdAt)

	assert.NotNil(t, contract)
	assert.Equal(t, symbol, contract.Symbol)
	assert.Equal(t, address, contract.Address)
	assert.Equal(t, createdAt, contract.CreatedAt)
}

func TestContractRepository(t *testing.T) {
	mockRepo := new(ContractRepositoryMock)

	t.Run("CreateContract", func(t *testing.T) {
		contract := &Contract{
			Id:        1,
			Symbol:    "ETH",
			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			CreatedAt: time.Now().Unix(),
		}

		mockRepo.On("CreateContract", contract).Return(contract, nil)

		createdContract, err := mockRepo.CreateContract(contract)

		assert.Nil(t, err)
		assert.Equal(t, contract, createdContract)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindAllContracts", func(t *testing.T) {
		contract1 := &Contract{
			Id:        1,
			Symbol:    "ETH",
			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			CreatedAt: time.Now().Unix(),
		}

		contract2 := &Contract{
			Id:        2,
			Symbol:    "USDC",
			Address:   common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			CreatedAt: time.Now().Unix(),
		}

		contracts := []*Contract{contract1, contract2}

		mockRepo.On("FindAllContracts").Return(contracts, nil)

		foundContracts, err := mockRepo.FindAllContracts()

		assert.Nil(t, err)
		assert.Equal(t, contracts, foundContracts)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindContractBySymbol", func(t *testing.T) {
		contract := &Contract{
			Id:        1,
			Symbol:    "ETH",
			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			CreatedAt: time.Now().Unix(),
		}

		mockRepo.On("FindContractBySymbol", "ETH").Return(contract, nil)

		foundContract, err := mockRepo.FindContractBySymbol("ETH")

		assert.Nil(t, err)
		assert.Equal(t, contract, foundContract)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateContract", func(t *testing.T) {
		contract := &Contract{
			Id:        1,
			Symbol:    "ETH",
			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			CreatedAt: time.Now().Unix(),
		}

		updatedContract := &Contract{
			Id:        1,
			Symbol:    "ETH",
			Address:   common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			CreatedAt: contract.CreatedAt,
			UpdatedAt: time.Now().Unix(),
		}

		mockRepo.On("UpdateContract", updatedContract).Return(updatedContract, nil)

		resultContract, err := mockRepo.UpdateContract(updatedContract)

		assert.Nil(t, err)
		assert.Equal(t, updatedContract, resultContract)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteContract", func(t *testing.T) {
		mockRepo.On("DeleteContract", "ETH").Return(nil)

		err := mockRepo.DeleteContract("ETH")

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})
}
