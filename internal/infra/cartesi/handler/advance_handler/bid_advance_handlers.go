package advance_handler

import (
	"encoding/json"
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/bid_usecase"
	"github.com/devolthq/devolt/internal/usecase/contract_usecase"
	"github.com/rollmelette/rollmelette"
)

type BidAdvanceHandlers struct {
	BidRepository      entity.BidRepository
	UserRepository     entity.UserRepository
	AuctionRepository  entity.AuctionRepository
	ContractRepository entity.ContractRepository
}

func NewBidAdvanceHandlers(bidRepository entity.BidRepository, userRepository entity.UserRepository, contractRepository entity.ContractRepository, auctionRepository entity.AuctionRepository) *BidAdvanceHandlers {
	return &BidAdvanceHandlers{
		BidRepository:      bidRepository,
		UserRepository:     userRepository,
		AuctionRepository:  auctionRepository,
		ContractRepository: contractRepository,
	}
}

func (h *BidAdvanceHandlers) CreateBidHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	switch deposit := deposit.(type) {
	case *rollmelette.ERC20Deposit:
		var input bid_usecase.CreateBidInputDTO
		if err := json.Unmarshal(payload, &input); err != nil {
			return err
		}
		createBid := bid_usecase.NewCreateBidUseCase(h.BidRepository, h.ContractRepository, h.AuctionRepository)
		res, err := createBid.Execute(&input, deposit, metadata)
		if err != nil {
			return err
		}

		findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
		volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
		if err != nil {
			return err
		}

		application, isDefined := env.AppAddress()
		if !isDefined {
			return fmt.Errorf("no application address defined yet, contact the DeVolt support")
		}

		if err := env.ERC20Transfer(volt.Address.Address, res.Bidder.Address, application, res.Credits.Int); err != nil {
			return err
		}
		env.Notice([]byte(fmt.Sprintf("created bid with id: %v and amount of credits: %v and price per credit: %v", res.Id, res.Credits, res.PricePerCredit)))
		return nil
	default:
		return fmt.Errorf("unsupported deposit type")
	}
}
