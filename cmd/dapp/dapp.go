package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/devolthq/devolt/pkg/router"
	"github.com/rollmelette/rollmelette"
)

func SetupApplication() *router.Router {
	//////////////////////// Setup Dependencies //////////////////////////
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

	//////////////////////// Setup Router //////////////////////////
	app := router.NewRouter()

	//////////////////////// Advance Handlers //////////////////////////
	app.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)
	app.HandleAdvance("updateOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.UpdateOrderHandler, "admin"))
	app.HandleAdvance("deleteOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.DeleteOrderHandler, "admin"))

	app.HandleAdvance("createContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	app.HandleAdvance("updateContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	app.HandleAdvance("deleteContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	app.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	app.HandleAdvance("createStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	app.HandleAdvance("updateStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	app.HandleAdvance("deleteStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	app.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	app.HandleAdvance("createAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	app.HandleAdvance("updateAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.UpdateAuctionHandler, "admin"))
	app.HandleAdvance("finishAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	app.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	app.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)
	app.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	app.HandleAdvance("updateUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.UpdateUserHandler, "admin"))
	app.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawAppHandler, "admin"))
	app.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	//////////////////////// Inspect Handlers //////////////////////////
	app.HandleInspect("order", ih.OrderInspectHandlers.FindAllOrdersHandler)
	app.HandleInspect("order/{id}", ih.OrderInspectHandlers.FindOrderByIdHandler)
	app.HandleInspect("order/user/{address}", ih.OrderInspectHandlers.FindOrdersByUserHandler)

	app.HandleInspect("auction", ih.AuctionInspectHandlers.FindAllAuctionsHandler)
	app.HandleInspect("auction/{id}", ih.AuctionInspectHandlers.FindAuctionByIdHandler)
	app.HandleInspect("auction/active", ih.AuctionInspectHandlers.FindActiveAuctionHandler)

	app.HandleInspect("station", ih.StationInspectHandlers.FindAllStationsHandler)
	app.HandleInspect("station/{id}", ih.StationInspectHandlers.FindStationByIdHandler)

	app.HandleInspect("bid", ih.BidInspectHandlers.FindAllBidsHandler)
	app.HandleInspect("bid/{id}", ih.BidInspectHandlers.FindBidByIdHandler)
	app.HandleInspect("bid/auction/{id}", ih.BidInspectHandlers.FindBisdByAuctionIdHandler)

	app.HandleInspect("contract", ih.ContractInspectHandlers.FindAllContractsHandler)
	app.HandleInspect("contract/{symbol}", ih.ContractInspectHandlers.FindContractBySymbolHandler)

	app.HandleInspect("user", ih.UserInspectHandlers.FindAllUsersHandler)
	app.HandleInspect("user/{address}", ih.UserInspectHandlers.FindUserByAddressHandler)
	app.HandleInspect("balance/{symbol}/{address}", ih.UserInspectHandlers.BalanceHandler)

	return app
}

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := SetupApplication()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
