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
	app := SetupApplicationMemory()
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
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItCreateUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C9","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItCreateUserWithInvalidData() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"","role":""}`) // Dados inválidos
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `invalid user`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateUser() {
	createUserSender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	createUserPayload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C6","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: createUserPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createUserExpectedOutput := `created user with address: 0x70997970c51812Dc3a010c7D01b50e0d17Dc79C6 and role: admin`
	result := s.tester.Advance(createUserSender, input)
	s.Len(result.Notices, 1)
	s.Equal(createUserExpectedOutput, string(result.Notices[0].Payload))

	updateUserSender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	updateUserPayload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C6","role":"admin"}`)
	updateUserInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateUser",
		Payload: updateUserPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := "updated user with address: 0x70997970c51812Dc3a010c7D01b50e0d17Dc79C6 and role: admin"
	updateUserResult := s.tester.Advance(updateUserSender, updateUserInput)
	s.Len(updateUserResult.Notices, 1)
	s.Equal(expectedOutput, string(updateUserResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C6","role":"user"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateNonExistentUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65","role":"admin"}`) // Usuário que não existe
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `user not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
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
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItDeleteUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteNonExistentUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	address := common.HexToAddress("0xNonExistentAddress").String()
	payload := []byte(`{"address":"` + address + `"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `user not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
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
	s.Len(result.Notices, 1)
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
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateContract() {
	createContractSender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	createContractPayload := []byte(`{"symbol":"TEST","address":"0x0000000000000000000000000000000000000005"}`)
	createContractInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: createContractPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createContractExpectedOutput := `created contract with symbol: TEST and address: 0x0000000000000000000000000000000000000005`
	result := s.tester.Advance(createContractSender, createContractInput)
	s.Len(result.Notices, 1)
	s.Equal(createContractExpectedOutput, string(result.Notices[0].Payload))

	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	updateContractPayload := []byte(`{"symbol":"TEST","address":"0x0000000000000000000000000000000000000005"}`)
	updateContractInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateContract",
		Payload: updateContractPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutputUpdateContract := `updated contract with symbol: TEST and address: 0x0000000000000000000000000000000000000005`
	resultUpdateContract := s.tester.Advance(sender, updateContractInput)
	s.Len(result.Notices, 1)
	s.Equal(expectedOutputUpdateContract, string(resultUpdateContract.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItCreateContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000002"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItCreateContractWithInvalidData() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"","address":""}`) // Dados inválidos
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `invalid contract`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000003"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateNonExistentContract() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"NONEXISTENT","address":"0x0000000000000000000000000000000000000003"}`) // Contrato que não existe
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `contract not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Não é admin
	payload := []byte(`{"symbol":"VOLT"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteNonExistentContract() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"NONEXISTENT"}`) // Contrato que não existe
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `contract not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
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
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItCreateStationWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Not an admin
	payload := []byte(`{"id":"station-2", "owner": "0x1234567890abcdef1234567890abcdef12345678", "consumption": 200, "price_per_credit": 100, "latitude": 34.0522, "longitude": -118.2437}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItCreateStationWithInvalidData() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"", "owner": "", "consumption": -100, "price_per_credit": -50, "latitude": 91.0000, "longitude": 181.0000}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `invalid station`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateStationWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Not an admin
	payload := []byte(`{"id":"station-1", "owner": "0x1234567890abcdef1234567890abcdef12345678", "consumption": 150, "price_per_credit": 75, "latitude": 34.0522, "longitude": -118.2437}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateNonExistentStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"non-existent-station", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 150, "price_per_credit": 75, "latitude": 34.0522, "longitude": -118.2437}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `station not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteStationWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678") // Not an admin
	payload := []byte(`{"id":"station-1"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x1234567890AbcdEF1234567890aBcdef12345678: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteNonExistentStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":"non-existent-station"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `station not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

// TODO: OffSet Station Consumption
