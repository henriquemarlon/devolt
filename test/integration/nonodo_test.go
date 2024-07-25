package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/Khan/genqlient/graphql"
	devolt "github.com/devolthq/devolt/internal/infra/cartesi/router"
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

func (s *NonodoSuite) TestAuctionWithParcialSelling() {
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	err = AdvanceInputBox(s.ctx, "http://127.0.0.1:8545", input)
	s.Require().Nil(err)

	client := graphql.NewClient("http://127.0.0.1:8080/graphql", nil)
	err = waitForInput(s.ctx, client)
	s.Require().Nil(err)

	result, err := getNodeState(s.ctx, client)
	s.Require().Nil(err)
	s.Require().Len(result.Inputs.Edges, 1)

	expected := []byte(`created user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin`)
	output := result.Inputs.Edges[0].Node
	s.Require().Equal(expected, common.Hex2Bytes(output.Notices.Edges[0].Node.Payload[2:]))
}

func (s *NonodoSuite) TestAuctionWithoutParcialSelling() {
	// TODO
}

// Helper functions ////////////////////////////////////////////////////////////////////////////////

func waitForInput(ctx context.Context, client graphql.Client) error {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		result, err := getInputStatus(ctx, client, 0)
		if err != nil && !strings.Contains(err.Error(), "input not found") {
			return fmt.Errorf("failed to get input status: %w", err)
		}
		if result.Input.Status == CompletionStatusAccepted {
			return nil
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
