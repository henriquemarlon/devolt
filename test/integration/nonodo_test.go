package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

	devolt "github.com/devolthq/devolt/internal/infra/cartesi/router"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

const TestTimeout = 5 * time.Second

var Accounts = []string{
	"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	"0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
	"0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
	"0x90F79bf6EB2c4f870365E785982E1f101E93b906",
	"0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
	"0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc",
	"0x976EA74026E726554dB657fA54763abd0C3a0aa9",
	"0x14dC79964da2C08b23698B3D3cc7Ca32193d9955",
	"0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f",
	"0xa0Ee7A142d267C1f36714E4a8F75612F20a79720",
}

func TestNonodoSuite(t *testing.T) {
	suite.Run(t, new(NonodoSuite))
}

type NonodoSuite struct {
	suite.Suite
	group  *errgroup.Group
	ctx    context.Context
	cancel context.CancelFunc
}

// Setup ///////////////////////////////////////////////////////////////////////////////////////////

func (s *NonodoSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), TestTimeout)
	s.group, s.ctx = errgroup.WithContext(s.ctx)

	// start nonodo
	nonodo := exec.CommandContext(s.ctx, "nonodo")
	nonodo.Cancel = func() error {
		return nonodo.Process.Signal(syscall.SIGTERM)
	}
	out := NewNotifyWriter(os.Stdout, "nonodo: ready")
	nonodo.Stdout = out
	s.group.Go(nonodo.Run)
	select {
	case <-out.ready:
	case <-s.ctx.Done():
		s.T().Error(s.ctx.Err())
	}

	// start test app
	s.group.Go(func() error {
		opts := rollmelette.NewRunOpts()
		app := devolt.NewTestApp()
		return rollmelette.Run(s.ctx, opts, app)
	})
}

func (s *NonodoSuite) TearDownTest() {
	s.cancel()
	err := s.group.Wait()
	s.ErrorIs(err, context.Canceled)
}

// Test Cases //////////////////////////////////////////////////////////////////////////////////////

func (s *NonodoSuite) TestSystemFlowAndAuctionWithParcialSelling() {
	// -> Create VOLT Contract
	volt, err := NewTestToken("http://127.0.0.1:8545", Accounts[0])
	if err != nil {
		s.T().Fatal(err)
	}
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: []byte(fmt.Sprintf(`{"symbol":"VOLT","address":"%s"}`, volt.Address.String())),
	})
	s.Require().Nil(err)

	err = AddInput(s.ctx, "http://127.0.0.1:8545", Accounts[0], voltInput)
	s.Require().Nil(err)

	// -> Create USDC Contract
	usdc, err := NewTestToken("http://127.0.0.1:8545", Accounts[0])
	if err != nil {
		s.T().Fatal(err)
	}
	usdcInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: []byte(fmt.Sprintf(`{"symbol":"USDC","address":"%s"}`, usdc.Address.String())),
	})
	s.Require().Nil(err)

	err = AddInput(s.ctx, "http://127.0.0.1:8545", Accounts[0], usdcInput)
	s.Require().Nil(err)

	// -> Send AppAddress through RelayContract
	err = RelayAppAddress(s.ctx, "http://127.0.0.1:8545", Accounts[0])
	s.Require().Nil(err)

	// -> Create Station 1
	createStation1Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: []byte(fmt.Sprintf(`{"id":"station-1", "owner": "%v", "consumption": 1600, "price_per_credit": 3, "latitude": 40.7128, "longitude": -74.0060}`, Accounts[5])),
	})
	s.Require().Nil(err)

	err = AddInput(s.ctx, "http://127.0.0.1:8545", Accounts[0], createStation1Input)
	s.Require().Nil(err)

	// -> Create Station 2
	createStation2Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: []byte(fmt.Sprintf(`{"id":"station-2", "owner": "%v", "consumption": 1600, "price_per_credit": 4, "latitude": 43.7128, "longitude": -72.0060}`, Accounts[6])),
	})
	s.Require().Nil(err)

	err = AddInput(s.ctx, "http://127.0.0.1:8545", Accounts[0], createStation2Input)
	s.Require().Nil(err)

	// // -> Create Order to Station 1
	// _, err = json.Marshal(&router.AdvanceRequest{
	// 	Path:    "createOrder",
	// 	Payload: []byte(`{"station_id":"station-2"}`),
	// })
	// s.Require().Nil(err)

	// err = usdc.Mint(Accounts[0], Accounts[3], big.NewInt(1000))
	// s.Require().Nil(err)
	// -> Increase Time (5 days)
	// -> Create Order to Station 2
	// -> Increase Time (5 days)
	// -> Create Order to Station 1
	// -> Increase Time (5 days)
	// -> Create Order to Station 2
	// -> Increase Time (5 days)
	// -> Create Order to Station 1
	// -> Increase Time (5 days)
	// -> Verify if the balance of the station owner is equal to the sum of all orders
	// -> Withdraw funds as stations owners
	// -> Initiate auction with duration of 6 days
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Bid
	// -> Increase Time (half day)
	// -> Finish Auction
	// -> Verify number of outputs
}

func (s *NonodoSuite) TestSystemFlowAndAuctionWithoutParcialSelling() {}
