package user_usecase

type FindUserOutputDTO struct {
	Id        uint   `json:"id"`
	Role      string `json:"role"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
