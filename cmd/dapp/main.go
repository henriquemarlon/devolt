package main

import (
	"context"
	"github.com/devolthq/devolt/internal/usecase/user_usecase"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"log"
	"log/slog"
)

func main() {
	//////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlers()
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlers()
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewares()
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
	}

	//////////////////////// Initial Owner //////////////////////////
	initialOwner := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	createUser := user_usecase.NewCreateUserUseCase(ah.UserAdvanceHandlers.UserRepository)
	if _, err = createUser.Execute(&user_usecase.CreateUserInputDTO{
		Address: initialOwner,
		Role:    "admin",
	}, rollmelette.Metadata{}); err != nil {
		slog.Error("failed to setup initial onwer", "error", err)
	}

	//////////////////////// Router //////////////////////////
	dapp := router.NewRouter()

	//////////////////////// Advance //////////////////////////
	dapp.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)
	dapp.HandleAdvance("updateOrder", ms.RBACMiddleware.Middleware(ah.OrderAdvanceHandlers.UpdateOrderHandler, "admin"))
	dapp.HandleAdvance("deleteOrder", ms.RBACMiddleware.Middleware(ah.OrderAdvanceHandlers.DeleteOrderHandler, "admin"))

	dapp.HandleAdvance("createContract", ms.RBACMiddleware.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	dapp.HandleAdvance("updateContract", ms.RBACMiddleware.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	dapp.HandleAdvance("deleteContract", ms.RBACMiddleware.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	dapp.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	dapp.HandleAdvance("createStation", ms.RBACMiddleware.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	dapp.HandleAdvance("updateStation", ms.RBACMiddleware.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	dapp.HandleAdvance("deleteStation", ms.RBACMiddleware.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	dapp.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	dapp.HandleAdvance("createAuction", ms.RBACMiddleware.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	dapp.HandleAdvance("updateAuction", ms.RBACMiddleware.Middleware(ah.AuctionAdvanceHandlers.UpdateAuctionHandler, "admin"))
	dapp.HandleAdvance("finishAuction", ms.RBACMiddleware.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	dapp.HandleAdvance("withdraw", ah.UserAdvanceHandlers.WithdrawHandler)
	dapp.HandleAdvance("createUser", ms.RBACMiddleware.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	dapp.HandleAdvance("updateUser", ms.RBACMiddleware.Middleware(ah.UserAdvanceHandlers.UpdateUserHandler, "admin"))
	dapp.HandleAdvance("withdrawApp", ms.RBACMiddleware.Middleware(ah.UserAdvanceHandlers.WithdrawAppHandler, "admin"))
	dapp.HandleAdvance("deleteUser", ms.RBACMiddleware.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	//////////////////////// Inspect //////////////////////////
	dapp.HandleInspect("order", ih.OrderInspectHandlers.FindAllOrdersHandler)
	dapp.HandleInspect("order/{id}", ih.OrderInspectHandlers.FindOrderByIdHandler)
	dapp.HandleInspect("order/user/{address}", ih.OrderInspectHandlers.FindOrdersByUserHandler)

	dapp.HandleInspect("auction", ih.AuctionInspectHandlers.FindAllAuctionsHandler)
	dapp.HandleInspect("auction/{id}", ih.AuctionInspectHandlers.FindAuctionByIdHandler)
	dapp.HandleInspect("auction/active", ih.AuctionInspectHandlers.FindActiveAuctionHandler)

	dapp.HandleInspect("station", ih.StationInspectHandlers.FindAllStationsHandler)
	dapp.HandleInspect("station/{id}", ih.StationInspectHandlers.FindStationByIdHandler)

	dapp.HandleInspect("bid", ih.BidInspectHandlers.FindAllBidsHandler)
	dapp.HandleInspect("bid/{id}", ih.BidInspectHandlers.FindBidByIdHandler)
	dapp.HandleInspect("bid/auction/{id}", ih.BidInspectHandlers.FindBisdByAuctionIdHandler)

	dapp.HandleInspect("contract", ih.ContractInspectHandlers.FindAllContractsHandler)
	dapp.HandleInspect("contract/{symbol}", ih.ContractInspectHandlers.FindContractBySymbolHandler)

	dapp.HandleInspect("user", ih.UserInspectHandlers.FindAllUsersHandler)
	dapp.HandleInspect("user/{address}", ih.UserInspectHandlers.FindUserByAddressHandler)
	dapp.HandleInspect("balance/{symbol}/{address}", ih.UserInspectHandlers.BalanceHandler)

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err = rollmelette.Run(ctx, opts, dapp)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
