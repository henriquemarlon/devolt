package main

import (
	"log"
	"github.com/devolthq/devolt/pkg/router"
)

func SetupApplicationPersistent() *router.Router {
	//////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlersPersistent()
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlersPersistent()
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewearsPersistent()
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
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

	dapp.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	dapp.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)
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

func SetupApplicationMemory() *router.Router {
	//////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlersMemory()
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlersMemory()
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewearsMemory()
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
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

	dapp.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	dapp.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)
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
