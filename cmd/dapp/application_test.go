package main

import (
	"encoding/json"
	"fmt"
	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *IntegrationTestSuite) SetupSuite() {
	app := SetupApplication()
	s.tester = rollmelette.NewTester(app)
}

////////////////// User ///////////////////

func (s *IntegrationTestSuite) TestItCreateUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateUser() {
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

func (s *IntegrationTestSuite) TestItDeleteUser() {
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

// TODO: withdraw
// TODO: withdraw app

///////////////// Contract ///////////////////

func (s *IntegrationTestSuite) TestItCreateContract() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000001"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000001`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateContract() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000003"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `updated contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000003`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItDeleteContract() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"VOLT"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `deleted contract with symbol: VOLT`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

///////////////// Station ///////////////////

func (s *IntegrationTestSuite) TestItCreateStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"station-1", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created station with id: station-1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"station-1", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `updated station with id: station-1, address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and consumption: 100`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItDeleteStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"station-1"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `deleted station with id: station-1`
	result := s.tester.Advance(sender, input)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

// TODO: OffSet Station Consumption
