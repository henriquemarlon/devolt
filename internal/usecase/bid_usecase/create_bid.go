package bid_usecase

import (
	"fmt"
	"math/big"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type CreateBidInputDTO struct {
	Bidder string   `json:"bidder"`
	Price  *big.Int `json:"price"`
}

type CreateBidOutputDTO struct {
	Id        uint     `json:"id"`
	AuctionId uint     `json:"auction_id"`
	Bidder    string   `json:"bidder"`
	Credits   *big.Int `json:"credits"`
	Price     *big.Int `json:"price"`
	State     string   `json:"state"`
	CreatedAt int64    `json:"created_at"`
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
	bidDeposit, ok := deposit.(*rollmelette.ERC20Deposit)
	if bidDeposit == nil || !ok {
		return nil, fmt.Errorf("unsupported deposit type for bid creation: %T", deposit)
	}

	volt, err := c.ContractRepository.FindContractBySymbol("VOLT")
	if err != nil {
		return nil, err
	}
	if bidDeposit.Token != common.HexToAddress(volt.Address) {
		return nil, fmt.Errorf("invalid contract address provided for bid creation: %v", bidDeposit.Token)
	}

	activeAuctionRes, err := c.AuctionRepository.FindActiveAuction()
	if err != nil {
		return nil, err
	}

	if input.Price.Cmp(activeAuctionRes.PriceLimit) == 1 {
		return nil, fmt.Errorf("bid price exceeds active auction price limit")
	}

	bid, err := entity.NewBid(activeAuctionRes.Id, input.Bidder, bidDeposit.Amount, input.Price, metadata.BlockTimestamp)
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
