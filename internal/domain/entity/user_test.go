package entity

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	role := "admin"
	address := common.HexToAddress("0x123")
	createdAt := time.Now().Unix()

	user, err := NewUser(role, address, createdAt)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, address, user.Address)
	assert.NotZero(t, user.CreatedAt)
}

func TestUser_Validate(t *testing.T) {
	address := common.HexToAddress("0x123")
	createdAt := time.Now().Unix()

	// Invalid role
	user := &User{
		Role:      "",
		Address:   address,
		CreatedAt: createdAt,
	}
	err := user.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidUser, err)

	// Invalid address
	user.Role = "admin"
	user.Address = common.Address{}
	err = user.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidUser, err)

	// Valid user
	user.Address = address
	err = user.Validate()
	assert.Nil(t, err)
}
