package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/devolthq/devolt/pkg/router"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
	"github.com/stretchr/testify/suite"
)

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

type AppSuite struct {
	suite.Suite
	tester *rollmelette.Tester
}

func (s *AppSuite) SetupTest() {
	app := newTestApp()
	s.tester = rollmelette.NewTester(app)
}

//==> Unit Tests <==////

//////////////// User ///////////////////

func (s *AppSuite) TestItCreateUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x70997970C51812dc3A010C7d01b50e0d17dc79C8","role":"admin"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createUser",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedAdvanceOutput := `created user with address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 and role: admin`
	advanceResult := s.tester.Advance(sender, input)
	s.Len(advanceResult.Notices, 1)
	s.Equal(expectedAdvanceOutput, string(advanceResult.Notices[0].Payload))
}

func (s *AppSuite) TestItCreateUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItCreateUserWithInvalidData() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"","role":""}`)
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

func (s *AppSuite) TestItUpdateUser() {
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

func (s *AppSuite) TestItUpdateUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItUpdateNonExistentUser() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"address":"0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65","role":"admin"}`)
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

func (s *AppSuite) TestItDeleteUser() {
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

func (s *AppSuite) TestItDeleteUserWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItDeleteNonExistentUser() {
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

func (s *AppSuite) TestItWithdrawVolt() {
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

	input := []byte(`{"path":"withdrawVolt","payload":{}}`)
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

func (s *AppSuite) TestItWithdrawVoltWithInsuficientBalance() {
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

	input := []byte(`{"path":"withdrawVolt","payload":{}}`)

	expectedOutput := `no balance of VOLT to withdraw`
	withdrawResult := s.tester.Advance(sender, input)
	s.ErrorContains(withdrawResult.Err, expectedOutput)
}

func (s *AppSuite) TestItWithdrawStablecoin() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(10000), []byte(""))

	input := []byte(`{"path":"withdrawStablecoin","payload":{}}`)
	// input, err := json.Marshal(&router.AdvanceRequest{
	// 	Path:    "withdrawStablecoin",
	// 	Payload: []byte(`{}`),
	// })
	// if err != nil {
	// 	s.T().Fatal(err)
	// }

	expectedNoticePayload := `withdrawn STABLECOIN and 10000 from 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc with voucher index: 1`

	expectedSTABLECOINVoucherPayload := make([]byte, 0, 4+32+32)
	expectedSTABLECOINVoucherPayload = append(expectedSTABLECOINVoucherPayload, 0xa9, 0x05, 0x9c, 0xbb)
	expectedSTABLECOINVoucherPayload = append(expectedSTABLECOINVoucherPayload, make([]byte, 12)...)
	expectedSTABLECOINVoucherPayload = append(expectedSTABLECOINVoucherPayload, sender[:]...)
	expectedSTABLECOINVoucherPayload = append(expectedSTABLECOINVoucherPayload, big.NewInt(10000).FillBytes(make([]byte, 32))...)
	withdrawResult := s.tester.Advance(sender, input)
	s.Len(withdrawResult.Notices, 1)
	s.Len(withdrawResult.Vouchers, 1)

	s.Equal(expectedSTABLECOINVoucherPayload, withdrawResult.Vouchers[0].Payload)
	s.Equal(common.HexToAddress("0x0000000000000000000000000000000000000001"), withdrawResult.Vouchers[0].Destination)
	s.Equal(expectedNoticePayload, string(withdrawResult.Notices[0].Payload))
}

func (s *AppSuite) TestItWithdrawStablecoinWithInsuficientBalance() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	sender := common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	input := []byte(`{"path":"withdrawStablecoin","payload":{}}`)

	expectedOutput := `no balance of STABLECOIN to withdraw`
	withdrawResult := s.tester.Advance(sender, input)
	s.ErrorContains(withdrawResult.Err, expectedOutput)
}

///////////////// Contract ///////////////////

func (s *AppSuite) TestItCreateContract() {
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

func (s *AppSuite) TestItDeleteContract() {
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

func (s *AppSuite) TestItUpdateContract() {
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

func (s *AppSuite) TestItCreateContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItCreateContractWithInvalidData() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"symbol":"","address":""}`)
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

func (s *AppSuite) TestItUpdateContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItUpdateNonExistentContract() {
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

func (s *AppSuite) TestItDeleteContractWithoutPermissions() {
	sender := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
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

func (s *AppSuite) TestItDeleteNonExistentContract() {
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

func (s *AppSuite) TestItCreateStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	result := s.tester.Advance(sender, input)
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *AppSuite) TestItCreateStationWithoutPermissions() {
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

func (s *AppSuite) TestItCreateStationWithInvalidData() {
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

func (s *AppSuite) TestItUpdateStationWithoutPermissions() {
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

func (s *AppSuite) TestItUpdateNonExistentStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":12, "owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 150, "price_per_credit": 75, "latitude": 34.0522, "longitude": -118.2437}`)
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

func (s *AppSuite) TestItDeleteStationWithoutPermissions() {
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

func (s *AppSuite) TestItDeleteNonExistentStation() {
	sender := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":1000}`)
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

func (s *AppSuite) TestItCreateOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, input)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":1}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)
	createOrderExpectedOutput := "created order 1 and paid 4000 as station fee and 6000 as application fee"
	s.Equal(createOrderExpectedOutput, string(createOrderResult.Notices[0].Payload))
}

func (s *AppSuite) TestItCreateOrderWithInvalidData() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":1}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(0), createOrderPayload)
	createOrderExpectedOutput := "invalid order"
	s.ErrorContains(createOrderResult.Err, createOrderExpectedOutput)
}

func (s *AppSuite) TestItUpdateOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":1}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createOrderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), createOrderPayload)
	createOrderExpectedOutput := "created order 1 and paid 4000 as station fee and 6000 as application fee"
	s.Equal(createOrderExpectedOutput, string(createOrderResult.Notices[0].Payload))

	updateOrderPayload := []byte(`{"id":1, "station_id":1, "credits": 20000}`)
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

func (s *AppSuite) TestItUpdateOrderWithoutPermissions() {
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

func (s *AppSuite) TestItDeleteOrder() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000002"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000002`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	createStationPayload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	createOrderPayload, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createOrder",
		Payload: []byte(`{"station_id":1}`),
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

func (s *AppSuite) TestItDeleteOrderWithoutPermissions() {
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

func (s *AppSuite) TestItDeleteNonExistentOrder() {
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

func (s *AppSuite) TestItCreateBid() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// Create STABLECOIN contract
	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000002"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000002`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	// Create Station 1
	station1Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	stationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station1Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	stationResult := s.tester.Advance(admin, stationInput)
	s.Len(stationResult.Notices, 1)
	s.Equal(stationExpectedOutput, string(stationResult.Notices[0].Payload))

	// Relay app address
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	// Create Orders for Stations
	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 10000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 15000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 20000)

	createAuctionPayload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit_per_credit":"1000", "expires_at": %v, "orders_time_range": 10}`, time.Now().Add(time.Hour).Unix()))
	createAuctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: createAuctionPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createAuctionExpectedOutput := `created auction with id: 1`
	createAuctionResult := s.tester.Advance(admin, createAuctionInput)
	s.Len(createAuctionResult.Notices, 1)
	s.Equal(createAuctionExpectedOutput, string(createAuctionResult.Notices[0].Payload))

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	payload := []byte(`{"price_per_credit":"1000"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createBid",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `created bid with id: 1 and amount of credits: 10000 and price per credit: 1000`
	result := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), input)
	s.Len(result.Notices, 1)
	s.Equal(expectedOutput, string(result.Notices[0].Payload))
}

func (s *AppSuite) TestItCreateBidWhenAuctionIsNotOngoing() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000002"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000002`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	createStationPayload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	createStationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: createStationPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createStationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	createStationResult := s.tester.Advance(admin, createStationInput)
	s.Len(createStationResult.Notices, 1)
	s.Equal(createStationExpectedOutput, string(createStationResult.Notices[0].Payload))

	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	// Simulate orders for stations with incremented time periods
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 10000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 15000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 20000)

	createAuctionPayload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit_per_credit":"1000", "expires_at": %v, "orders_time_range": 2}`, time.Now().Add(5*time.Second).Unix()))
	createAuctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: createAuctionPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	createAuctionExpectedOutput := `created auction with id: 1`
	createAuctionResult := s.tester.Advance(admin, createAuctionInput)
	s.Len(createAuctionResult.Notices, 1)
	s.Equal(createAuctionExpectedOutput, string(createAuctionResult.Notices[0].Payload))

	time.Sleep(6 * time.Second) // wait for auction to expire

	sender := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	payload := []byte(`{"price_per_credit":"1000"}`)
	input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createBid",
		Payload: payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	expectedOutput := `active auction expired, cannot create bid`
	result := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000002"), sender, big.NewInt(10000), input)
	s.ErrorContains(result.Err, expectedOutput)
}

//////////////// Auction //////////////////

func (s *AppSuite) TestItCreateAuction() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// Create VOLT contract
	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000022"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000022`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	// Create STABLECOIN contract
	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	// Relay app address
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	// Create Station 1
	station1Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	stationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station1Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	stationResult := s.tester.Advance(admin, stationInput)
	s.Len(stationResult.Notices, 1)
	s.Equal(stationExpectedOutput, string(stationResult.Notices[0].Payload))

	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	// Simulate orders for stations with incremented time periods
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 10000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 15000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 20000)

	// admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(fmt.Sprintf(`{"price_limit_per_credit":"1000", "expires_at": %v, "orders_time_range": 100}`, time.Now().Add(time.Hour).Unix()))
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

func (s *AppSuite) TestItCreateAuctionWithoutPermissions() {
	sender := common.HexToAddress("0x0000000000000000000000000000000000000001")
	payload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit_per_credit":"1000", "expires_at": %v}`, time.Now().Add(time.Hour).Unix()))
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

func (s *AppSuite) TestItCreateAuctionWithInvalidData() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// Create STABLECOIN contract
	stablecoinPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000001"}`)
	stablecoinInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: stablecoinPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stablecoinExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000001`
	stablecoinResult := s.tester.Advance(admin, stablecoinInput)
	s.Len(stablecoinResult.Notices, 1)
	s.Equal(stablecoinExpectedOutput, string(stablecoinResult.Notices[0].Payload))

	// Create Station 1
	station1Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "consumption": 100, "price_per_credit": 50, "latitude": 40.7128, "longitude": -74.0060}`)
	stationInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station1Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	stationExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	stationResult := s.tester.Advance(admin, stationInput)
	s.Len(stationResult.Notices, 1)
	s.Equal(stationExpectedOutput, string(stationResult.Notices[0].Payload))

	// Relay app address
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	// Create Orders for Stations
	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000001"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 10000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 15000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 20000)

	payload := []byte(fmt.Sprintf(`{"price_limit_per_credit":"1000", "expires_at": %v, "orders_time_range": 10}`, time.Now().Unix()))
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

func (s *AppSuite) TestItUpdateAuctionWithoutPermissions() {
	user := common.HexToAddress("0x0000000000000000000000000000000000000001")
	payload := []byte(fmt.Sprintf(`{"credits":"100000", "price_limit_per_credit":"1000", "expires_at": %v}`, time.Now().Add(time.Hour).Unix()))
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

func (s *AppSuite) TestItUpdateNonExistentAuction() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	payload := []byte(`{"id":999, "credits":"150000", "price_limit_per_credit":"1200", "expires_at": 1625097600}`)
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

func (s *AppSuite) TestItFinishAuctionWithoutPartialSelling() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// Create VOLT contract
	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000022"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000022`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	// Create STABLECOIN contract
	usdcPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000033"}`)
	usdcInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: usdcPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	usdcExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000033`
	usdcResult := s.tester.Advance(admin, usdcInput)
	s.Len(usdcResult.Notices, 1)
	s.Equal(usdcExpectedOutput, string(usdcResult.Notices[0].Payload))

	// Relay app address
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	// Create Station 1
	station1Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "price_per_credit": 10, "latitude": 40.7128, "longitude": -74.0060}`)
	station1Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station1Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	station1ExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	station1Result := s.tester.Advance(admin, station1Input)
	s.Len(station1Result.Notices, 1)
	s.Equal(station1ExpectedOutput, string(station1Result.Notices[0].Payload))

	// Create Station 2
	station2Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C9", "price_per_credit": 10, "latitude": 34.0522, "longitude": -118.2437}`)
	station2Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station2Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	station2ExpectedOutput := `created station with id: 2 and owner: 0x70997970c51812dc3a010c7D01B50E0D17Dc79c9`
	station2Result := s.tester.Advance(admin, station2Input)
	s.Len(station2Result.Notices, 1)
	s.Equal(station2ExpectedOutput, string(station2Result.Notices[0].Payload))

	// Create Orders for Stations
	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000033"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	// Simulate orders for stations with incremented time periods
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 10000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C2"), 2, 20000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 15000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C4"), 2, 25000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 20000)

	// Initiate auction with mock duration
	auctionPayload := []byte(fmt.Sprintf(`{"price_limit_per_credit":"10", "expires_at": %v, "orders_time_range": 1000}`, time.Now().Add(5*time.Second).Unix()))
	auctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: auctionPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	auctionExpectedOutput := `created auction with id: 1`
	auctionResult := s.tester.Advance(admin, auctionInput)
	s.Len(auctionResult.Notices, 1)
	s.Equal(auctionExpectedOutput, string(auctionResult.Notices[0].Payload))

	// Simulate bids with a counter for expected bid ID
	bidCounter := 1
	placeBid := func(sender common.Address, pricePeCredit string, amount *big.Int) {
		bidPayload := []byte(fmt.Sprintf(`{"price_per_credit":"%s"}`, pricePeCredit))
		bidInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createBid",
			Payload: bidPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		bidResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000022"), sender, amount, bidInput)
		expectedBidOutput := fmt.Sprintf("created bid with id: %d and amount of credits: %v and price per credit: %v", bidCounter, amount, pricePeCredit)
		s.Len(bidResult.Notices, 1)
		s.Equal(expectedBidOutput, string(bidResult.Notices[0].Payload))
		bidCounter++
	}

	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000001"), "7", big.NewInt(500))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000002"), "8", big.NewInt(2000))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000003"), "10", big.NewInt(3000))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000004"), "9", big.NewInt(500))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000005"), "4", big.NewInt(3000))

	time.Sleep(5 * time.Second)

	// Finish auction and verify results
	finishAuctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path: "finishAuction",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	time.Sleep(5 * time.Second)
	finishAuctionResult := s.tester.Advance(admin, finishAuctionInput)
	finishAuctionExpectedOutput := fmt.Sprintf("finished auction with id: 1 at: %v", time.Now().Unix())
	s.Len(finishAuctionResult.Notices, 1)
	s.Equal(finishAuctionExpectedOutput, string(finishAuctionResult.Notices[0].Payload))

	// OffSetConsumption as station owner
	offSetStationConsumptionAsOwnerInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "offSetStationConsumption",
		Payload: []byte(`{"id":1, "credits_to_be_offSet": 4500}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	offSetStationConsumptionAsOwnerExpectedOutput := `offSet Credits from station: 1 by msg_sender: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	offSetStationConsumptionAsOwnerResult := s.tester.Advance(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"), offSetStationConsumptionAsOwnerInput)
	s.Len(offSetStationConsumptionAsOwnerResult.Notices, 1)
	s.Equal(offSetStationConsumptionAsOwnerExpectedOutput, string(offSetStationConsumptionAsOwnerResult.Notices[0].Payload))
}

func (s *AppSuite) TestItFinishAuctionWithPartialSelling() {
	admin := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	// Create VOLT contract
	voltPayload := []byte(`{"symbol":"VOLT","address":"0x0000000000000000000000000000000000000022"}`)
	voltInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: voltPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	voltExpectedOutput := `created contract with symbol: VOLT and address: 0x0000000000000000000000000000000000000022`
	voltResult := s.tester.Advance(admin, voltInput)
	s.Len(voltResult.Notices, 1)
	s.Equal(voltExpectedOutput, string(voltResult.Notices[0].Payload))

	// Create STABLECOIN contract
	usdcPayload := []byte(`{"symbol":"STABLECOIN","address":"0x0000000000000000000000000000000000000033"}`)
	usdcInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createContract",
		Payload: usdcPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	usdcExpectedOutput := `created contract with symbol: STABLECOIN and address: 0x0000000000000000000000000000000000000033`
	usdcResult := s.tester.Advance(admin, usdcInput)
	s.Len(usdcResult.Notices, 1)
	s.Equal(usdcExpectedOutput, string(usdcResult.Notices[0].Payload))

	// Relay app address
	appAddressResult := s.tester.RelayAppAddress(common.HexToAddress("0xdadadadadadadadadadadadadadadadadadadada"))
	s.Nil(appAddressResult.Err)

	// Create Station 1
	station1Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "price_per_credit": 10, "latitude": 40.7128, "longitude": -74.0060}`)
	station1Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station1Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	station1ExpectedOutput := `created station with id: 1 and owner: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	station1Result := s.tester.Advance(admin, station1Input)
	s.Len(station1Result.Notices, 1)
	s.Equal(station1ExpectedOutput, string(station1Result.Notices[0].Payload))

	// Create Station 2
	station2Payload := []byte(`{"owner": "0x70997970C51812dc3A010C7d01b50e0d17dc79C9", "price_per_credit": 10, "latitude": 34.0522, "longitude": -118.2437}`)
	station2Input, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createStation",
		Payload: station2Payload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	station2ExpectedOutput := `created station with id: 2 and owner: 0x70997970c51812dc3a010c7D01B50E0D17Dc79c9`
	station2Result := s.tester.Advance(admin, station2Input)
	s.Len(station2Result.Notices, 1)
	s.Equal(station2ExpectedOutput, string(station2Result.Notices[0].Payload))

	// Create Orders for Stations
	orderCounter := 1
	createOrder := func(sender common.Address, stationID uint, credits int64) {
		orderPayload := []byte(fmt.Sprintf(`{"station_id":%v}`, stationID))
		orderInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createOrder",
			Payload: orderPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		orderResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000033"), sender, big.NewInt(credits), orderInput)
		expectedOutput := fmt.Sprintf("created order %v and paid %d as station fee and %d as application fee", orderCounter, credits*40/100, credits*60/100)
		s.Equal(expectedOutput, string(orderResult.Notices[0].Payload))
		orderCounter++
	}

	// Simulate orders for stations with incremented time periods
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C1"), 1, 1000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C2"), 2, 2000)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C3"), 1, 1500)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C4"), 2, 2500)
	createOrder(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C5"), 1, 2000)

	// Initiate auction with mock duration
	auctionPayload := []byte(fmt.Sprintf(`{"price_limit_per_credit":"10", "expires_at": %v, "orders_time_range": 1000}`, time.Now().Add(5*time.Second).Unix()))
	auctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "createAuction",
		Payload: auctionPayload,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	auctionExpectedOutput := `created auction with id: 1`
	auctionResult := s.tester.Advance(admin, auctionInput)
	s.Len(auctionResult.Notices, 1)
	s.Equal(auctionExpectedOutput, string(auctionResult.Notices[0].Payload))

	// Simulate bids with a counter for expected bid ID
	bidCounter := 1
	placeBid := func(sender common.Address, pricePeCredit string, amount *big.Int) {
		bidPayload := []byte(fmt.Sprintf(`{"price_per_credit":"%s"}`, pricePeCredit))
		bidInput, err := json.Marshal(&router.AdvanceRequest{
			Path:    "createBid",
			Payload: bidPayload,
		})
		if err != nil {
			s.T().Fatal(err)
		}
		bidResult := s.tester.DepositERC20(common.HexToAddress("0x0000000000000000000000000000000000000022"), sender, amount, bidInput)
		expectedBidOutput := fmt.Sprintf("created bid with id: %d and amount of credits: %v and price per credit: %v", bidCounter, amount, pricePeCredit)
		s.Len(bidResult.Notices, 1)
		s.Equal(expectedBidOutput, string(bidResult.Notices[0].Payload))
		bidCounter++
	}

	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000002"), "9", big.NewInt(2000))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000003"), "6", big.NewInt(3000))
	placeBid(common.HexToAddress("0x0000000000000000000000000000000000000004"), "7", big.NewInt(2000))

	time.Sleep(5 * time.Second)

	// Finish auction and verify results
	finishAuctionInput, err := json.Marshal(&router.AdvanceRequest{
		Path: "finishAuction",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	time.Sleep(5 * time.Second)
	finishAuctionResult := s.tester.Advance(admin, finishAuctionInput)
	finishAuctionExpectedOutput := fmt.Sprintf("finished auction with id: 1 at: %v", time.Now().Unix())
	s.Len(finishAuctionResult.Notices, 1)
	s.Equal(finishAuctionExpectedOutput, string(finishAuctionResult.Notices[0].Payload))

	// offSetConsumption as station owner
	offSetStationConsumptionAsOwnerInput, err := json.Marshal(&router.AdvanceRequest{
		Path:    "offSetStationConsumption",
		Payload: []byte(`{"id":1, "credits_to_be_offSet": 450}`),
	})
	if err != nil {
		s.T().Fatal(err)
	}
	offSetStationConsumptionAsOwnerExpectedOutput := `offSet Credits from station: 1 by msg_sender: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
	offSetStationConsumptionAsOwnerResult := s.tester.Advance(common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"), offSetStationConsumptionAsOwnerInput)
	s.Len(offSetStationConsumptionAsOwnerResult.Notices, 1)
	s.Equal(offSetStationConsumptionAsOwnerExpectedOutput, string(offSetStationConsumptionAsOwnerResult.Notices[0].Payload))
}

func newTestApp() *router.Router {
	////////////////////// Setup Handlers //////////////////////////
	ah, err := NewAdvanceHandlersMemory()
	if err != nil {
		log.Fatalf("Failed to initialize advance handlers from wire: %v", err)
	}

	ih, err := NewInspectHandlersMemory()
	if err != nil {
		log.Fatalf("Failed to initialize inspect handlers from wire: %v", err)
	}

	ms, err := NewMiddlewaresMemory()
	if err != nil {
		log.Fatalf("Failed to initialize middlewares from wire: %v", err)
	}

	////////////////////// Router //////////////////////////
	app := router.NewRouter()

	////////////////////// Advance //////////////////////////
	app.HandleAdvance("createOrder", ah.OrderAdvanceHandlers.CreateOrderHandler)
	app.HandleAdvance("updateOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.UpdateOrderHandler, "admin"))
	app.HandleAdvance("deleteOrder", ms.RBAC.Middleware(ah.OrderAdvanceHandlers.DeleteOrderHandler, "admin"))

	app.HandleAdvance("createContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.CreateContractHandler, "admin"))
	app.HandleAdvance("updateContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.UpdateContractHandler, "admin"))
	app.HandleAdvance("deleteContract", ms.RBAC.Middleware(ah.ContractAdvanceHandlers.DeleteContractHandler, "admin"))

	app.HandleAdvance("createBid", ah.BidAdvanceHandlers.CreateBidHandler)

	app.HandleAdvance("createStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.CreateStationHandler, "admin"))
	app.HandleAdvance("updateStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.UpdateStationHandler, "admin"))
	app.HandleAdvance("deleteStation", ms.RBAC.Middleware(ah.StationAdvanceHandlers.DeleteStationHandler, "admin"))
	app.HandleAdvance("offSetStationConsumption", ah.StationAdvanceHandlers.OffSetStationConsumptionHandler)

	app.HandleAdvance("createAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.CreateAuctionHandler, "admin"))
	app.HandleAdvance("updateAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.UpdateAuctionHandler, "admin"))
	app.HandleAdvance("finishAuction", ms.RBAC.Middleware(ah.AuctionAdvanceHandlers.FinishAuctionHandler, "admin"))

	app.HandleAdvance("withdrawVolt", ah.UserAdvanceHandlers.WithdrawVoltHandler)
	app.HandleAdvance("withdrawStablecoin", ah.UserAdvanceHandlers.WithdrawStablecoinHandler)
	app.HandleAdvance("createUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.CreateUserHandler, "admin"))
	app.HandleAdvance("updateUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.UpdateUserHandler, "admin"))
	app.HandleAdvance("withdrawApp", ms.RBAC.Middleware(ah.UserAdvanceHandlers.WithdrawAppHandler, "admin"))
	app.HandleAdvance("deleteUser", ms.RBAC.Middleware(ah.UserAdvanceHandlers.DeleteUserByAddressHandler, "admin"))

	////////////////////// Inspect //////////////////////////
	app.HandleInspect("order", ih.OrderInspectHandlers.FindAllOrdersHandler)
	app.HandleInspect("order/{id}", ih.OrderInspectHandlers.FindOrderByIdHandler)
	app.HandleInspect("order/user/{address}", ih.OrderInspectHandlers.FindOrdersByUserHandler)

	app.HandleInspect("auction", ih.AuctionInspectHandlers.FindAllAuctionsHandler)
	app.HandleInspect("auction/{id}", ih.AuctionInspectHandlers.FindAuctionByIdHandler)
	app.HandleInspect("auction/active", ih.AuctionInspectHandlers.FindActiveAuctionHandler)

	app.HandleInspect("station", ih.StationInspectHandlers.FindAllStationsHandler)
	app.HandleInspect("station/{id}", ih.StationInspectHandlers.FindStationByIdHandler)

	app.HandleInspect("bid", ih.BidInspectHandlers.FindAllBidsHandler)
	app.HandleInspect("bid/{id}", ih.BidInspectHandlers.FindBidByIdHandler)
	app.HandleInspect("bid/auction/{id}", ih.BidInspectHandlers.FindBisdByAuctionIdHandler)

	app.HandleInspect("contract", ih.ContractInspectHandlers.FindAllContractsHandler)
	app.HandleInspect("contract/{symbol}", ih.ContractInspectHandlers.FindContractBySymbolHandler)

	app.HandleInspect("user", ih.UserInspectHandlers.FindAllUsersHandler)
	app.HandleInspect("user/{address}", ih.UserInspectHandlers.FindUserByAddressHandler)
	app.HandleInspect("balance/{symbol}/{address}", ih.UserInspectHandlers.BalanceHandler)

	return app
}
