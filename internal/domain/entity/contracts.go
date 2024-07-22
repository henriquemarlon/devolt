package entity

import (
	"errors"

	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidContract  = errors.New("invalid contract")
	ErrContractNotFound = errors.New("contract not found")
)

type ContractRepository interface {
	CreateContract(contract *Contract) (*Contract, error)
	FindAllContracts() ([]*Contract, error)
	FindContractBySymbol(symbol string) (*Contract, error)
	UpdateContract(contract *Contract) (*Contract, error)
	DeleteContract(symbol string) error
}

type Contract struct {
	Id        uint                `json:"id" gorm:"primaryKey"`
	Symbol    string              `json:"symbol" gorm:"uniqueIndex;not null"`
	Address   custom_type.Address `json:"address" gorm:"type:text;not null"`
	CreatedAt int64               `json:"created_at" gorm:"not null"`
	UpdatedAt int64               `json:"updated_at" gorm:"default:0"`
}

func NewContract(symbol string, address custom_type.Address, createdAt int64) (*Contract, error) {
	contract := &Contract{
		Symbol:    symbol,
		Address:   address,
		CreatedAt: createdAt,
	}
	if err := contract.Validate(); err != nil {
		return nil, err
	}
	return contract, nil
}

func (c *Contract) Validate() error {
	if c.Symbol == "" || c.Address.Address == (common.Address{}) {
		return ErrInvalidContract
	}
	return nil
}
