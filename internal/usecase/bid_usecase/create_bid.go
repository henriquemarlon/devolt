package bid_usecase

import (
	"fmt"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type CreateBidInputDTO struct {
	Price  custom_type.BigInt  `json:"price"`
}

type CreateBidOutputDTO struct {
	Id        uint                `json:"id"`
	AuctionId uint                `json:"auction_id"`
	Bidder    custom_type.Address `json:"bidder"`
	Credits   custom_type.BigInt  `json:"credits"`
	Price     custom_type.BigInt  `json:"price"`
	State     string              `json:"state"`
	CreatedAt int64               `json:"created_at"`
}

type CreateBidUseCase struct {
	BidRepository      entity.BidRepository
	ContractRepository entity.ContractRepository
	AuctionRepository  entity.AuctionRepository
}

func NewCreateBidUseCase(bidRepository entity.BidRepository, contractRepository entity.ContractRepository, auctionRepository entity.AuctionRepository) *CreateBidUseCase {
	return &CreateBidUseCase{
		BidRepository:      bidRepository,
		ContractRepository: contractRepository,
		AuctionRepository:  auctionRepository,
	}
}

func (c *CreateBidUseCase) Execute(input *CreateBidInputDTO, deposit rollmelette.Deposit, metadata rollmelette.Metadata) (*CreateBidOutputDTO, error) {
	activeAuction, err := c.AuctionRepository.FindActiveAuction()
	if err != nil {
		return nil, err
	}
	if activeAuction == nil {
		return nil, fmt.Errorf("no active auction found, cannot create bid")
	}

	if metadata.BlockTimestamp > activeAuction.ExpiresAt {
		return nil, fmt.Errorf("active auction expired, cannot create bid")
	}

	bidDeposit, ok := deposit.(*rollmelette.ERC20Deposit)
	if bidDeposit == nil || !ok {
		return nil, fmt.Errorf("unsupported deposit type for bid creation: %T", deposit)
	}

	volt, err := c.ContractRepository.FindContractBySymbol("VOLT")
	if err != nil {
		return nil, err
	}
	if bidDeposit.Token != volt.Address.Address {
		return nil, fmt.Errorf("invalid contract address provided for bid creation: %v", bidDeposit.Token)
	}

	if input.Price.Cmp(activeAuction.PriceLimit.Int) == 1 {
		return nil, fmt.Errorf("bid price exceeds active auction price limit")
	}

	bid, err := entity.NewBid(activeAuction.Id, custom_type.NewAddress(bidDeposit.Sender), custom_type.NewBigInt(bidDeposit.Amount), input.Price, metadata.BlockTimestamp)
	if err != nil {
		return nil, err
	}
	res, err := c.BidRepository.CreateBid(bid)
	if err != nil {
		return nil, err
	}

	return &CreateBidOutputDTO{
		Id:        res.Id,
		AuctionId: res.AuctionId,
		Bidder:    res.Bidder,
		Credits:   res.Credits,
		Price:     res.Price,
		State:     string(res.State),
		CreatedAt: res.CreatedAt,
	}, nil
}
