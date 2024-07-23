package db

import (
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"gorm.io/gorm"
)

type ContractRepositorySqlite struct {
	Db *gorm.DB
}

func NewContractRepositorySqlite(db *gorm.DB) *ContractRepositorySqlite {
	return &ContractRepositorySqlite{
		Db: db,
	}
}

func (r *ContractRepositorySqlite) CreateContract(contract *entity.Contract) (*entity.Contract, error) {
	err := r.Db.Create(contract).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create contract: %w", err)
	}
	return contract, nil
}

func (r *ContractRepositorySqlite) FindAllContracts() ([]*entity.Contract, error) {
	var contracts []*entity.Contract
	err := r.Db.Find(&contracts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch contracts: %w", err)
	}
	return contracts, nil
}

func (r *ContractRepositorySqlite) FindContractBySymbol(symbol string) (*entity.Contract, error) {
	var contract entity.Contract
	err := r.Db.Where("symbol = ?", symbol).First(&contract).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find contract by symbol: %w", err)
	}
	return &contract, nil
}

func (r *ContractRepositorySqlite) UpdateContract(contract *entity.Contract) (*entity.Contract, error) {
	res := r.Db.Model(&entity.Contract{}).Where("symbol = ?", contract.Symbol).Omit("created_at").Updates(contract)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to update contract: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, entity.ErrContractNotFound
	}
	return contract, nil
}

func (r *ContractRepositorySqlite) DeleteContract(symbol string) error {
	res := r.Db.Where("symbol = ?", symbol).Delete(&entity.Contract{})
	if res.Error != nil {
		return fmt.Errorf("failed to delete contract: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.ErrContractNotFound
	}
	return nil
}
