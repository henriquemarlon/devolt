package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
)

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, new(UserIntegrationTestSuite))
}

type UserIntegrationTestSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *UserIntegrationTestSuite) SetupSuite() {
	dapp := SetupApplication()
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
	address := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8").String()
	payload := []byte(`{"address":"` + address + `"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := fmt.Sprintf(`deleted user with address: %v`, address)
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}
