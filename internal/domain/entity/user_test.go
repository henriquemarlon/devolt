package entity

// import (
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"testing"
// 	"time"
// )

// // UserRepositoryMock is a mock implementation of the UserRepository interface
// type UserRepositoryMock struct {
// 	mock.Mock
// }

// func (m *UserRepositoryMock) CreateUser(user *User) (*User, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(*User), args.Error(1)
// }

// func (m *UserRepositoryMock) FindUserByRole(role string) (*User, error) {
// 	args := m.Called(role)
// 	return args.Get(0).(*User), args.Error(1)
// }

// func (m *UserRepositoryMock) FindUserByAddress(address common.Address) (*User, error) {
// 	args := m.Called(address)
// 	return args.Get(0).(*User), args.Error(1)
// }

// func (m *UserRepositoryMock) UpdateUser(user *User) (*User, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(*User), args.Error(1)
// }

// func (m *UserRepositoryMock) FindAllUsers() ([]*User, error) {
// 	args := m.Called()
// 	return args.Get(0).([]*User), args.Error(1)
// }

// func (m *UserRepositoryMock) DeleteUserByAddress(address common.Address) error {
// 	args := m.Called(address)
// 	return args.Error(0)
// }

// func TestNewUser(t *testing.T) {
// 	role := "admin"
// 	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
// 	createdAt := time.Now().Unix()

// 	user := NewUser(role, address, createdAt)

// 	assert.NotNil(t, user)
// 	assert.Equal(t, role, user.Role)
// 	assert.Equal(t, address, user.Address)
// 	assert.Equal(t, createdAt, user.CreatedAt)
// }

// func TestUserRepository(t *testing.T) {
// 	mockRepo := new(UserRepositoryMock)

// 	t.Run("CreateUser", func(t *testing.T) {
// 		user := &User{
// 			Id:        1,
// 			Role:      "admin",
// 			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("CreateUser", user).Return(user, nil)

// 		createdUser, err := mockRepo.CreateUser(user)

// 		assert.Nil(t, err)
// 		assert.Equal(t, user, createdUser)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindUserByRole", func(t *testing.T) {
// 		user := &User{
// 			Id:        1,
// 			Role:      "admin",
// 			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("FindUserByRole", "admin").Return(user, nil)

// 		foundUser, err := mockRepo.FindUserByRole("admin")

// 		assert.Nil(t, err)
// 		assert.Equal(t, user, foundUser)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindUserByAddress", func(t *testing.T) {
// 		user := &User{
// 			Id:        1,
// 			Role:      "admin",
// 			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("FindUserByAddress", user.Address).Return(user, nil)

// 		foundUser, err := mockRepo.FindUserByAddress(user.Address)

// 		assert.Nil(t, err)
// 		assert.Equal(t, user, foundUser)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("UpdateUser", func(t *testing.T) {
// 		user := &User{
// 			Id:        1,
// 			Role:      "admin",
// 			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		updatedUser := &User{
// 			Id:        1,
// 			Role:      "superadmin",
// 			Address:   user.Address,
// 			CreatedAt: user.CreatedAt,
// 			UpdatedAt: time.Now().Unix(),
// 		}

// 		mockRepo.On("UpdateUser", updatedUser).Return(updatedUser, nil)

// 		resultUser, err := mockRepo.UpdateUser(updatedUser)

// 		assert.Nil(t, err)
// 		assert.Equal(t, updatedUser, resultUser)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("FindAllUsers", func(t *testing.T) {
// 		user1 := &User{
// 			Id:        1,
// 			Role:      "admin",
// 			Address:   common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		user2 := &User{
// 			Id:        2,
// 			Role:      "user",
// 			Address:   common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
// 			CreatedAt: time.Now().Unix(),
// 		}

// 		users := []*User{user1, user2}

// 		mockRepo.On("FindAllUsers").Return(users, nil)

// 		foundUsers, err := mockRepo.FindAllUsers()

// 		assert.Nil(t, err)
// 		assert.Equal(t, users, foundUsers)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("DeleteUserByAddress", func(t *testing.T) {
// 		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

// 		mockRepo.On("DeleteUserByAddress", address).Return(nil)

// 		err := mockRepo.DeleteUserByAddress(address)

// 		assert.Nil(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }
