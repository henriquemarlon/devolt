package user_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type UpdateUserInputDTO struct {
	Id      uint           `json:"id"`
	Role    string         `json:"role"`
	Address common.Address `json:"address"`
}

type UpdateUserOutputDTO struct {
	Id        uint                `json:"id"`
	Role      string              `json:"role"`
	Address   custom_type.Address `json:"address"`
	UpdatedAt int64               `json:"update_at"`
}

type UpdateUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewUpdateUserUseCase(userRepository entity.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UpdateUserUseCase) Execute(input *UpdateUserInputDTO, metadata rollmelette.Metadata) (*UpdateUserOutputDTO, error) {
	res, err := u.UserRepository.UpdateUser(&entity.User{
		Id:        input.Id,
		Role:      input.Role,
		Address:   custom_type.NewAddress(input.Address),
		UpdatedAt: metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateUserOutputDTO{
		Id:        res.Id,
		Role:      res.Role,
		Address:   res.Address,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
