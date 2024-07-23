package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *IntegrationTestSuite) SetupTest() {
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
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	createUserPayload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	createUserInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: createUserPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createUserExpectedOutput := `created user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin`
	createUserResult := s.tester.Advance(admin, createUserInput)
	s.Len(createUserResult.Notices, 1)
	s.Equal(createUserExpectedOutput, string(createUserResult.Notices[0].Payload))

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
	address := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65").String()
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

func (s *IntegrationTestSuite) TestItWithdrawVolt() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")
	
	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000001"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000001`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(10000), []byte(""))

	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawVolt",
		Payload: []byte(``),
	})
	if err != nil {
		s.T().Fatal(err)
	}

	expectedNoticePayload := `withdrawn VOLT and 10000 from 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc with voucher index: 1`

	expectedVOLTVoucherPayload := make([]byte, 0, 4+32+32)
	expectedVOLTVoucherPayload = append(expectedVOLTVoucherPayload, 0xa9, 0x05, 0x9c, 0xbb)
	expectedVOLTVoucherPayload = append(expectedVOLTVoucherPayload, make([]byte, 12)...)
	expectedVOLTVoucherPayload = append(expectedVOLTVoucherPayload, sender[:]...)
	expectedVOLTVoucherPayload = append(expectedVOLTVoucherPayload, big.NewInt(10000).FillBytes(make([]byte, 32))...)
	withdrawResult := s.tester.Advance(sender, input)
	s.Len(withdrawResult.Notices, 1)
	s.Len(withdrawResult.Vouchers, 1)

	s.Equal(expectedVOLTVoucherPayload, withdrawResult.Vouchers[0].Payload)
	s.Equal(common.HexToAddress("0x0000000000000000000000000000000000000001"), withdrawResult.Vouchers[0].Destination)
	s.Equal(expectedNoticePayload, string(withdrawResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItWithdrawVoltWithInsuficientBalance() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000001"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000001`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawVolt",
		Payload: []byte(``),
	})
	if err != nil {
		s.T().Fatal(err)
	}

	expectedOutput := `no balance of VOLT to withdraw`
	withdrawResult := s.tester.Advance(sender, input)
	s.ErrorContains(withdrawResult.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItWithdrawStablecoin() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(10000), []byte(""))

	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawStablecoin",
		Payload: []byte(``),
	})
	if err != nil {
		s.T().Fatal(err)
	}

	expectedNoticePayload := `withdrawn USDC and 10000 from 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc with voucher index: 1`

	expectedUSDCVoucherPayload := make([]byte, 0, 4+32+32)
	expectedUSDCVoucherPayload = append(expectedUSDCVoucherPayload, 0xa9, 0x05, 0x9c, 0xbb)
	expectedUSDCVoucherPayload = append(expectedUSDCVoucherPayload, make([]byte, 12)...)
	expectedUSDCVoucherPayload = append(expectedUSDCVoucherPayload, sender[:]...)
	expectedUSDCVoucherPayload = append(expectedUSDCVoucherPayload, big.NewInt(10000).FillBytes(make([]byte, 32))...)
	withdrawResult := s.tester.Advance(sender, input)
	s.Len(withdrawResult.Notices, 1)
	s.Len(withdrawResult.Vouchers, 1)

	s.Equal(expectedUSDCVoucherPayload, withdrawResult.Vouchers[0].Payload)
	s.Equal(common.HexToAddress("0x0000000000000000000000000000000000000001"), withdrawResult.Vouchers[0].Destination)
	s.Equal(expectedNoticePayload, string(withdrawResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItWithdrawStablecoinWithInsuficientBalance() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "withdrawStablecoin",
		Payload: []byte(``),
	})
	if err != nil {
		s.T().Fatal(err)
	}

	expectedOutput := `no balance of USDC to withdraw`
	withdrawResult := s.tester.Advance(sender, input)
	s.ErrorContains(withdrawResult.Err, expectedOutput)
}


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
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	createContractPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000001"}`)
	createContractInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: createContractPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createContractExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000001`
	createContractResult := s.tester.Advance(admin, createContractInput)
	s.Len(createContractResult.Notices, 1)
	s.Equal(createContractExpectedOutput, string(createContractResult.Notices[0].Payload))

	payload := []byte(`{"symbol":"VOLT"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteContract",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `deleted contract with symbol: VOLT`
	result := s.tester.Advance(admin, input)
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

///////////////// Order ///////////////////

func (s *IntegrationTestSuite) TestItCreateOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"id":"station-2", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: station-2 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, input)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)
	createOrderExpectedOutput := "created order 1 and paid 4000 as station fee and 6000 as application fee"
	s.Equal(createOrderExpectedOutput, string(createOrderResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItCreateOrderWithInvalidData() {	
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"id":"station-2", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: station-2 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(0), createOrderPayload)
	createOrderExpectedOutput := "invalid order"
	s.ErrorContains(createOrderResult.Err, createOrderExpectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"id":"station-2", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: station-2 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)
	createOrderExpectedOutput := "created order 1 and paid 4000 as station fee and 6000 as application fee"
	s.Equal(createOrderExpectedOutput, string(createOrderResult.Notices[0].Payload))

	updateOrderPayload := []byte(`{"id":1, "station_id":"station-2", "credits": 20000}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateOrder",
		Payload: updateOrderPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `updated order with id: 1 and credits: 20000`
	updateOrderResult := s.tester.Advance(admin, input)
	s.Len(updateOrderResult.Notices, 1)
	s.Equal(expectedOutput, string(updateOrderResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItUpdateOrderWithoutPermissions() {
	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")

	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)

	updateOrderPayload := []byte(`{"id":1, "credits": 20000}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateOrder",
		Payload: updateOrderPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65: record not found`
	updateOrderResult := s.tester.Advance(sender, input)
	s.ErrorContains(updateOrderResult.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"USDC","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: USDC and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"id":"station-2", "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: station-2 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)

	deleteOrderPayload := []byte(`{"id":1}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteOrder",
		Payload: deleteOrderPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `deleted order with id: 1`
	deleteOrderResult := s.tester.Advance(admin, input)
	s.Len(deleteOrderResult.Notices, 1)
	s.Equal(expectedOutput, string(deleteOrderResult.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItDeleteOrderWithoutPermissions() {
	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")

	// CREATE ORDER
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":"station-2"}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)

	deleteOrderPayload := []byte(`{"id":1}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteOrder",
		Payload: deleteOrderPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65: record not found`
	deleteOrderResult := s.tester.Advance(sender, input)
	s.ErrorContains(deleteOrderResult.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItDeleteNonExistentOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	deleteOrderPayload := []byte(`{"id":999}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "deleteOrder",
		Payload: deleteOrderPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `order not found`
	deleteOrderResult := s.tester.Advance(admin, input)
	s.ErrorContains(deleteOrderResult.Err, expectedOutput)
}

/////////////////// Bids //////////////////

func (s *IntegrationTestSuite) TestItCreateBidWhenAuctionIsNotOngoing() {
	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	payload := []byte(`{"price":"1000"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createBid",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `auction not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

//////////////// Auction //////////////////

func (s *IntegrationTestSuite) TestItCreateAuction() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit":"1000", "expires_at": %v}`, time.Now().Add(time.Hour).Unix()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created auction with id: 1`
	result := s.tester.Advance(admin, input)
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *IntegrationTestSuite) TestItCreateAuctionWithoutPermissions() {
	sender := common.HexToAddress("0x0000000000000000000000000000000000000001")
	payload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit":"1000", "expires_at": %v}`, time.Now().Add(time.Hour).Unix()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x0000000000000000000000000000000000000001: record not found`
	result := s.tester.Advance(sender, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItCreateAuctionWithInvalidData() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"credits":"0", "price_limit":"1000", "expires_at": 500}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `invalid auction`
	result := s.tester.Advance(admin, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateAuctionWithoutPermissions() {
	user := common.HexToAddress("0x0000000000000000000000000000000000000001")
	payload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit":"1000", "expires_at": %v}`, time.Now().Add(time.Hour).Unix()))
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateAuction",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `failed to find user by address 0x0000000000000000000000000000000000000001: record not found`
	result := s.tester.Advance(user, input)
	s.ErrorContains(result.Err, expectedOutput)
}

func (s *IntegrationTestSuite) TestItUpdateNonExistentAuction() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":999, "credits":"150000", "price_limit":"1200", "expires_at": 1625097600}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "updateAuction",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `auction not found`
	result := s.tester.Advance(admin, input)
	s.ErrorContains(result.Err, expectedOutput)
}