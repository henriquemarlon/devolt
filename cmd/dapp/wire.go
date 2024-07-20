//go:build wireinject
// +build wireinject

package main

import (
	"github.com/devolthq/devolt/configs"
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/devolthq/devolt/internal/infra/cartesi/handler/advance_handler"
	"github.com/devolthq/devolt/internal/infra/cartesi/handler/inspect_handler"
	"github.com/devolthq/devolt/internal/infra/cartesi/middleware"
	"github.com/devolthq/devolt/internal/infra/repository/sqlite"
	"github.com/google/wire"
)

var setBidRepositoryDependency = wire.NewSet(
	sqlite.NewBidRepositorySqlite,
	wire.Bind(new(entity.BidRepository), new(*sqlite.BidRepositorySqlite)),
)

var setAuctionRepositoryDependency = wire.NewSet(
	sqlite.NewAuctionRepositorySqlite,
	wire.Bind(new(entity.AuctionRepository), new(*sqlite.AuctionRepositorySqlite)),
)

var setOrderRepositoryDependency = wire.NewSet(
	sqlite.NewOrderRepositorySqlite,
	wire.Bind(new(entity.OrderRepository), new(*sqlite.OrderRepositorySqlite)),
)

var setStationRepositoryDependency = wire.NewSet(
	sqlite.NewStationRepositorySqlite,
	wire.Bind(new(entity.StationRepository), new(*sqlite.StationRepositorySqlite)),
)

var setContractRepositoryDependency = wire.NewSet(
	sqlite.NewContractRepositorySqlite,
	wire.Bind(new(entity.ContractRepository), new(*sqlite.ContractRepositorySqlite)),
)

var setUserRepositoryDependency = wire.NewSet(
	sqlite.NewUserRepositorySqlite,
	wire.Bind(new(entity.UserRepository), new(*sqlite.UserRepositorySqlite)),
)

var setAdvanceHandlers = wire.NewSet(
	advance_handler.NewOrderAdvanceHandlers,
	advance_handler.NewStationAdvanceHandlers,
	advance_handler.NewContractAdvanceHandlers,
	advance_handler.NewUserAdvanceHandlers,
	advance_handler.NewAuctionAdvanceHandlers,
	advance_handler.NewBidAdvanceHandlers,
)

var setInspectHandlers = wire.NewSet(
	inspect_handler.NewBidInspectHandlers,
	inspect_handler.NewUserInspectHandlers,
	inspect_handler.NewOrderInspectHandlers,
	inspect_handler.NewStationInspectHandlers,
	inspect_handler.NewAuctionInspectHandlers,
	inspect_handler.NewContractInspectHandlers,
)

var setMiddleware = wire.NewSet(
	middleware.NewRBACMiddleware,
)

func NewMiddlewares() (*Middlewares, error) {
	wire.Build(
		configs.SetupSQlite,
		setUserRepositoryDependency,
		setMiddleware,
		wire.Struct(new(Middlewares), "*"),
	)
	return nil, nil
}

func NewAdvanceHandlers() (*AdvanceHandlers, error) {
	wire.Build(
		configs.SetupSQlite,
		setBidRepositoryDependency,
		setUserRepositoryDependency,
		setOrderRepositoryDependency,
		setStationRepositoryDependency,
		setAuctionRepositoryDependency,
		setContractRepositoryDependency,
		setAdvanceHandlers,
		wire.Struct(new(AdvanceHandlers), "*"),
	)
	return nil, nil
}

func NewInspectHandlers() (*InspectHandlers, error) {
	wire.Build(
		configs.SetupSQlite,
		setBidRepositoryDependency,
		setUserRepositoryDependency,
		setOrderRepositoryDependency,
		setStationRepositoryDependency,
		setAuctionRepositoryDependency,
		setContractRepositoryDependency,
		setInspectHandlers,
		wire.Struct(new(InspectHandlers), "*"),
	)
	return nil, nil
}

type Middlewares struct {
	RBAC *middleware.RBACMiddleware
}

type AdvanceHandlers struct {
	BidAdvanceHandlers      *advance_handler.BidAdvanceHandlers
	UserAdvanceHandlers     *advance_handler.UserAdvanceHandlers
	OrderAdvanceHandlers    *advance_handler.OrderAdvanceHandlers
	StationAdvanceHandlers  *advance_handler.StationAdvanceHandlers
	AuctionAdvanceHandlers  *advance_handler.AuctionAdvanceHandlers
	ContractAdvanceHandlers *advance_handler.ContractAdvanceHandlers
}

type InspectHandlers struct {
	BidInspectHandlers      *inspect_handler.BidInspectHandlers
	UserInspectHandlers     *inspect_handler.UserInspectHandlers
	OrderInspectHandlers    *inspect_handler.OrderInspectHandlers
	StationInspectHandlers  *inspect_handler.StationInspectHandlers
	AuctionInspectHandlers  *inspect_handler.AuctionInspectHandlers
	ContractInspectHandlers *inspect_handler.ContractInspectHandlers
}
