package main

import (
	"github.com/devolthq/devolt/internal/usecase/user_usecase"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"log"
	"log/slog"
	"strings"
)

func SetupApplication() *router.Router {
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
	initialOwner := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	createUser := user_usecase.NewCreateUserUseCase(ah.UserAdvanceHandlers.UserRepository)
	if _, err = createUser.Execute(&user_usecase.CreateUserInputDTO{
		Address: strings.ToLower(initialOwner.String()),
		Role:    "admin",
	}, rollmelette.Metadata{}); err != nil {
		slog.Error("failed to setup initial onwer", "error", err)
	}

	//////////////////////// Router //////////////////////////
	dapp := router.NewRouter()

	//////////////////////// Advance //////////////////////////
	dapp.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)
	dapp.HandleAdvance("updateOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.UpdateOrderHandler, "admin"))
	dapp.HandleAdvance("deleteOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.DeleteOrderHandler, "admin"))

	dapp.HandleAdvance("createContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	dapp.HandleAdvance("updateContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	dapp.HandleAdvance("deleteContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	dapp.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	dapp.HandleAdvance("createStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	dapp.HandleAdvance("updateStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	dapp.HandleAdvance("deleteStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	dapp.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	dapp.HandleAdvance("createAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	dapp.HandleAdvance("updateAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.UpdateAuctionHandler, "admin"))
	dapp.HandleAdvance("finishAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	dapp.HandleAdvance("withdraw", ah.UserAdvanceHandlers.WithdrawHandler)
	dapp.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	dapp.HandleAdvance("updateUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.UpdateUserHandler, "admin"))
	dapp.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawAppHandler, "admin"))
	dapp.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

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

	return dapp
}
