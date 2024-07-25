package integration

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	devolt "github.com/devolthq/devolt/internal/infra/cartesi/router"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
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
	// err := DeleteFileIfExists("./devolt.db")
	// if err != nil {
	// 	s.T().Fatal(err)
	// }

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
		app := devolt.SetupTestApplication()
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
	err := IncreaseTime("http://127.0.0.1:8545", 10)
	s.Require().Nil(err)
}
