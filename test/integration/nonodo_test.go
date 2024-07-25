package integration

import (
	"context"
	"encoding/json"
	"github.com/Khan/genqlient/graphql"
	devolt "github.com/devolthq/devolt/internal/infra/cartesi/router"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
	"math/big"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

const TestTimeout = 5 * time.Second

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
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	err = AdvanceInputBox(s.ctx, "http://127.0.0.1:8545", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", input)
	s.Require().Nil(err)

	client := graphql.NewClient("http://127.0.0.1:8080/graphql", nil)
	err = WaitForInput(s.ctx, client)
	s.Require().Nil(err)

	result, err := getNodeState(s.ctx, client)
	s.Require().Nil(err)
	s.Require().Len(result.Inputs.Edges, 1)

	expected := []byte(`created user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin`)
	output := result.Inputs.Edges[0].Node
	s.Require().Equal(expected, common.Hex2Bytes(output.Notices.Edges[0].Node.Payload[2:]))
}

func (s *NonodoSuite) TestSystemFlowAndAuctionWithoutParcialSelling() {
	token, err := NewTokenInterface("http://127.0.0.1:8545", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	if err != nil {
		s.T().Fatal(err)
	}
	volt, err := token.Deploy("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "Volt", "VOLT")
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoin, err := token.Deploy("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "USDC", "USDC")
	if err != nil {
		s.T().Fatal(err)
	}

	accounts := []string{
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

	err = token.Mint(*volt, accounts[2], big.NewInt(10000))
	s.Require().Nil(err)

	err = token.Mint(*stablecoin, accounts[0], big.NewInt(10000))
	s.Require().Nil(err)

	err = IncreaseTime("http://127.0.0.1:8545", 10)
	if err != nil {
		s.T().Fatal(err)
	}
	s.Require().Nil(err)
}
