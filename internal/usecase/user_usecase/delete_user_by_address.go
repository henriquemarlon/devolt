package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	// "github.com/ethereum/go-ethereum/common"
)

type DeleteUserByAddressInputDTO struct {
	Address custom_type.Address `json:"address"`
}

type DeleteUserByAddressUseCase struct {
	UserRepository entity.UserRepository
}

func NewDeleteUserByAddressUseCase(userRepository entity.UserRepository) *DeleteUserByAddressUseCase {
	return &DeleteUserByAddressUseCase{
		UserRepository: userRepository,
	}
}

func (u *DeleteUserByAddressUseCase) Execute(input *DeleteUserByAddressInputDTO) error {
	return u.UserRepository.DeleteUserByAddress(input.Address)
}
