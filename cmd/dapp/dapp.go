package main

import (
	"log"

	"github.com/Mugen-Builders/devolt/configs"
	"github.com/Mugen-Builders/devolt/pkg/router"
)

func NewDApp() *router.Router {
	//////////////////////// Setup Database //////////////////////////
	db, err := configs.SetupSQlite()
	if err != nil {
		log.Fatalf("Failed to setup sqlite database: %v", err)
	}

	//////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlers(db)
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlers(db)
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewares(db)
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
	}

	//////////////////////// Router //////////////////////////
	app := router.NewRouter()

	//////////////////////// Advance //////////////////////////
	app.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)

	app.HandleAdvance("createContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	app.HandleAdvance("updateContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	app.HandleAdvance("deleteContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	app.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	app.HandleAdvance("createStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	app.HandleAdvance("updateStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	app.HandleAdvance("deleteStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	app.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	app.HandleAdvance("createAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	app.HandleAdvance("finishAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	app.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawStablecoinHandler, "admin"))
	app.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	app.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)

	app.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	app.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	//////////////////////// Inspect //////////////////////////
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

func NewDAppMemory() *router.Router {
	//////////////////////// Setup Database //////////////////////////
	db, err := configs.SetupSQliteMemory()
	if err != nil {
		log.Fatalf("Failed to setup sqlite database: %v", err)
	}

	//////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlersMemory(db)
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlersMemory(db)
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewaresMemory(db)
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
	}

	//////////////////////// Router //////////////////////////
	app := router.NewRouter()

	//////////////////////// Advance //////////////////////////
	app.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)

	app.HandleAdvance("createContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	app.HandleAdvance("updateContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	app.HandleAdvance("deleteContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	app.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	app.HandleAdvance("createStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	app.HandleAdvance("updateStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	app.HandleAdvance("deleteStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	app.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	app.HandleAdvance("createAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	app.HandleAdvance("finishAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	app.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	app.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)
	app.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	app.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawStablecoinHandler, "admin"))
	app.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	//////////////////////// Inspect //////////////////////////
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
