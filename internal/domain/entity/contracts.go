package entity

import (
	"github.com/ethereum/go-ethereum/common"
)

type ContractRepository interface {
	CreateContract(contract *Contract) (*Contract, error)
	FindAllContracts() ([]*Contract, error)
	FindContractBySymbol(symbol string) (*Contract, error)
	UpdateContract(contract *Contract) (*Contract, error)
	DeleteContract(symbol string) error
}

type Contract struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Symbol    string         `json:"symbol" gorm:"uniqueIndex;not null"`
	Address   common.Address `json:"address" gorm:"type:text;not null"`
	CreatedAt int64          `json:"created_at" gorm:"not null"`
	UpdatedAt int64          `json:"updated_at" gorm:"default:0"`
}

func NewContract(symbol string, address common.Address, createdAt int64) *Contract {
	return &Contract{
		Symbol:    symbol,
		Address:   address,
		CreatedAt: createdAt,
	}
}
