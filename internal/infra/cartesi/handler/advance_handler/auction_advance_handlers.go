package advance_handler

import (
	"encoding/json"
	"fmt"

	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/usecase/auction_usecase"
	"github.com/devolthq/devolt/internal/usecase/bid_usecase"
	"github.com/devolthq/devolt/internal/usecase/contract_usecase"
	"github.com/rollmelette/rollmelette"
)

type AuctionAdvanceHandlers struct {
	BidRepository      entity.BidRepository
	UserRepository     entity.UserRepository
	AuctionRepository  entity.AuctionRepository
	ContractRepository entity.ContractRepository
}

func NewAuctionAdvanceHandlers(
	bidRepository entity.BidRepository,
	userRepository entity.UserRepository,
	auctionRepository entity.AuctionRepository,
	contractRepository entity.ContractRepository,
) *AuctionAdvanceHandlers {
	return &AuctionAdvanceHandlers{
		BidRepository:      bidRepository,
		UserRepository:     userRepository,
		AuctionRepository:  auctionRepository,
		ContractRepository: contractRepository,
	}
}

func (h *AuctionAdvanceHandlers) CreateAuctionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input *auction_usecase.CreateAuctionInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return err
	}
	createAuction := auction_usecase.NewCreateAuctionUseCase(h.AuctionRepository)
	res, err := createAuction.Execute(input, metadata)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("created auction with id: %v", res.Id)))
	return nil
}

func (h *AuctionAdvanceHandlers) UpdateAuctionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input *auction_usecase.UpdateAuctionInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return err
	}
	updateAuction := auction_usecase.NewUpdateAuctionUseCase(h.AuctionRepository)
	res, err := updateAuction.Execute(input, metadata)
	if err != nil {
		return err
	}
	env.Notice([]byte(fmt.Sprintf("updated auction with id: %v and expiration: %v", res.Id, res.ExpiresAt)))
	return nil
}

func (h *AuctionAdvanceHandlers) FinishAuctionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	finishAuction := auction_usecase.NewFinishAuctionUseCase(h.AuctionRepository, h.BidRepository)
	finishedAuction, err := finishAuction.Execute(metadata)
	if err != nil {
		return err
	}

	application, isDefined := env.AppAddress()
	if !isDefined {
		return fmt.Errorf("no application address defined yet, contact the DeVolt support")
	}

	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	// Find Volt contract address
	volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
	if err != nil {
		return err
	}

	// Find Stablecoin contract address
	stablecoin, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "USDC"})
	if err != nil {
		return err
	}

	findBidsByState := bid_usecase.NewFindBidsByStateUseCase(h.BidRepository)
	acceptedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "accepted",
	})
	if err != nil {
		return err
	}
	for _, bid := range acceptedBids {
		if err := env.ERC20Transfer(stablecoin.Address.Address, application, bid.Bidder.Address, bid.Price.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	partialAcceptedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "partial_accepted",
	})
	if err != nil {
		return err
	}
	for _, bid := range partialAcceptedBids {
		if err := env.ERC20Transfer(stablecoin.Address.Address, application, bid.Bidder.Address, bid.Price.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	rejectedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "rejected",
	})
	if err != nil {
		return err
	}
	for _, bid := range rejectedBids {
		if err := env.ERC20Transfer(volt.Address.Address, application, bid.Bidder.Address, bid.Credits.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	findAuctionById := auction_usecase.NewFindAuctionByIdUseCase(h.AuctionRepository)
	res, err := findAuctionById.Execute(&auction_usecase.FindAuctionByIdInputDTO{
		Id: finishedAuction.Id,
	})
	if err != nil {
		return fmt.Errorf("failed to find auction: %w", err)
	}
	auction, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal auction: %w", err)
	}
	env.Notice([]byte(fmt.Sprintf("finished auction: %v at: %v", auction, metadata.BlockTimestamp)))
	return nil
}
