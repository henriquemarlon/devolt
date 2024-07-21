package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/devolthq/devolt/internal/usecase/user_usecase"
	"github.com/devolthq/devolt/pkg/router"
	di "github.com/devolthq/devolt/tools/dependency_injection"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
)

func setupUserIntegrationTestApplication() *router.Router {
	//////////////////////// Setup Handlers //////////////////////////
	ah, err := di.NewAdvanceHandlers()
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := di.NewInspectHandlers()
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := di.NewMiddlewares()
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
	dapp.HandleAdvance("withdraw", ah.UserAdvanceHandlers.WithdrawHandler)
	dapp.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	dapp.HandleAdvance("updateUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.UpdateUserHandler, "admin"))
	dapp.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawAppHandler, "admin"))
	dapp.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	//////////////////////// Inspect //////////////////////////
	dapp.HandleInspect("user", ih.UserInspectHandlers.FindAllUsersHandler)
	dapp.HandleInspect("user/{address}", ih.UserInspectHandlers.FindUserByAddressHandler)
	dapp.HandleInspect("balance/{symbol}/{address}", ih.UserInspectHandlers.BalanceHandler)

	return dapp
}

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, new(UserIntegrationTestSuite))
}

type UserIntegrationTestSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *UserIntegrationTestSuite) SetupSuite() {
	dapp := setupUserIntegrationTestApplication()
	s.tester = rollmelette.NewTester(dapp)
}

func (s *UserIntegrationTestSuite) TestItCreateUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := fmt.Sprintf(`{"id":2,"role":"admin","address":"0x70997970c51812dc3a010c7d01b50e0d17dc79c8","created_at":%d}`, time.Now().Unix())
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *UserIntegrationTestSuite) TestItUpdateUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := "updated user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin"
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *UserIntegrationTestSuite) TestItDeleteUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := fmt.Sprintf(`deleted user with address: %v`, "0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}
