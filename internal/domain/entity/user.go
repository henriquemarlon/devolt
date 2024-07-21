package entity

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidUser  = errors.New("invalid user")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(User *User) (*User, error)
	FindUserByRole(role string) (*User, error)
	FindUserByAddress(address string) (*User, error)
	UpdateUser(User *User) (*User, error)
	FindAllUsers() ([]*User, error)
	DeleteUserByAddress(address string) error
}

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Role      string `json:"role" gorm:"not null"`
	Address   string `json:"address" gorm:"type:text;uniqueIndex;not null"`
	CreatedAt int64  `json:"created_at" gorm:"not null"`
	UpdatedAt int64  `json:"updated_at" gorm:"default:0"`
}

func NewUser(role string, address string, created_at int64) (*User, error) {
	user := &User{
		Role:      role,
		Address:   address,
		CreatedAt: created_at,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Role == "" || u.Address == (common.Address{}.String()) {
		return ErrInvalidUser
	}
	return nil
}
