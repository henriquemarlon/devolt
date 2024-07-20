package mock

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(input *entity.User) (*entity.User, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindUserByRole(role string) (*entity.User, error) {
	args := m.Called(role)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindUserByAddress(address common.Address) (*entity.User, error) {
	args := m.Called(address)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindAllUsers() ([]*entity.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) UpdateUser(input *entity.User) (*entity.User, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) DeleteUserByAddress(address common.Address) error {
	args := m.Called(address)
	return args.Error(0)
}
