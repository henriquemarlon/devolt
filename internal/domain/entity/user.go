package entity

import "github.com/ethereum/go-ethereum/common"

type UserRepository interface {
	CreateUser(User *User) (*User, error)
	FindUserByRole(role string) (*User, error)
	FindUserByAddress(address common.Address) (*User, error)
	UpdateUser(User *User) (*User, error)
	FindAllUsers() ([]*User, error)
	DeleteUserByAddress(address common.Address) error
}

type User struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Role      string         `json:"role" gorm:"not null"`
	Address   common.Address `json:"address" gorm:"type:text;uniqueIndex;not null"`
	CreatedAt int64          `json:"created_at" gorm:"not null"`
	UpdatedAt int64          `json:"updated_at" gorm:"default:0"`
}

func NewUser(role string, address common.Address, created_at int64) *User {
	return &User{
		Role:      role,
		Address:   address,
		CreatedAt: created_at,
	}
}
