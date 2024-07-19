package sqlite

import (
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type UserRepositorySqlite struct {
	Db *gorm.DB
}

func NewUserRepositorySqlite(db *gorm.DB) *UserRepositorySqlite {
	return &UserRepositorySqlite{
		Db: db,
	}
}

func (r *UserRepositorySqlite) CreateUser(input *entity.User) (*entity.User, error) {
	err := r.Db.Create(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return input, nil
}

func (r *UserRepositorySqlite) FindUserByRole(role string) (*entity.User, error) {
	var user entity.User
	err := r.Db.Where("role = ?", role).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user by role: %w", err)
	}
	return &user, nil
}

func (r *UserRepositorySqlite) FindUserByAddress(address common.Address) (*entity.User, error) {
	var user entity.User
	err := r.Db.Where("address = ?", address).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find user by address: %w", err)
	}
	return &user, nil
}

func (r *UserRepositorySqlite) FindAllUsers() ([]*entity.User, error) {
	var users []*entity.User
	err := r.Db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	return users, nil
}

func (r *UserRepositorySqlite) UpdateUser(input *entity.User) (*entity.User, error) {
	err := r.Db.Save(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return input, nil
}

func (r *UserRepositorySqlite) DeleteUserByAddress(address common.Address) error {
	err := r.Db.Where("address = ?", address).Delete(&entity.User{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete user by address: %w", err)
	}
	return nil
}
