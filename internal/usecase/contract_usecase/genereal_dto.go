package contract_usecase

type FindContractOutputDTO struct {
	Id        uint   `json:"id"`
	Symbol    string `json:"symbol"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
