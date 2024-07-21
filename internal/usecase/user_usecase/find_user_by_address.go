package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type FindUserByAddressInputDTO struct {
	Address string `json:"address"`
}

type FindUserByAddressUseCase struct {
	UserRepository entity.UserRepository
}

func NewFindUserByAddressUseCase(userRepository entity.UserRepository) *FindUserByAddressUseCase {
	return &FindUserByAddressUseCase{
		UserRepository: userRepository,
	}
}

func (u *FindUserByAddressUseCase) Execute(input *FindUserByAddressInputDTO) (*FindUserOutputDTO, error) {
	res, err := u.UserRepository.FindUserByAddress(input.Address)
	if err != nil {
		return nil, err
	}
	return &FindUserOutputDTO{
		Id:        res.Id,
		Role:      res.Role,
		Address:   res.Address,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
