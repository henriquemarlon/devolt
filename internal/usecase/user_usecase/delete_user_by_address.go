package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type DeleteUserByAddressInputDTO struct {
	Address string
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
