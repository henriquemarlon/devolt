package integration

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rollmelette/rollmelette"
)

const ApplicationAddress = "0x70ac08179605AF2D9e75782b8DEcDD3c22aA4D0C"
const SenderAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const TestMnemonic = "test test test test test test test test test test test junk"

// Advance sends an input using the devnet contract addresses.
func AdvanceInputBox(ctx context.Context, url string, payload []byte) error {
	if len(payload) == 0 {
		return fmt.Errorf("cannot send empty payload")
	}
	book := rollmelette.NewAddressBook()
	input := hexutil.Encode(payload)
	cmd := exec.CommandContext(ctx,
		"cast", "send",
		"--mnemonic", TestMnemonic,
		"--rpc-url", url,
		book.InputBox.String(),             // TO
		"addInput(address,bytes)(bytes32)", // SIG
		ApplicationAddress, input,          // ARGS
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cast: %w: %v", err, string(output))
	}
	return nil
}

func AdvanceDepositERC20Tokens(ctx context.Context, url string, tokenAddress common.Address, value uint64, payload []byte) error {
	if len(payload) == 0 {
		return fmt.Errorf("cannot send empty payload")
	}
	book := rollmelette.NewAddressBook()
	input := hexutil.Encode(payload)

	approve := exec.CommandContext(ctx,
		"cast", "send",
		"--mnemonic", TestMnemonic,
		"--rpc-url", url,
		tokenAddress.String(),                                    // TO
		"approve(address,uint256)",                               // SIG
		book.ERC20Portal.String(), strconv.FormatUint(value, 10), // ARGS
	)
	approveOutput, err := approve.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cast: %w: %v", err, string(approveOutput))
	}

	deposit := exec.CommandContext(ctx,
		"cast", "send",
		"--mnemonic", TestMnemonic,
		"--rpc-url", url,
		book.ERC20Portal.String(),                                                       // TO
		"depositERC20Tokens(IERC20,address,uint256,bytes)",                              // SIG
		tokenAddress.String(), ApplicationAddress, strconv.FormatUint(value, 10), input, // ARGS
	)
	depositOutput, err := deposit.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cast: %w: %v", err, string(depositOutput))
	}
	return nil
}

func IncreaseTime(ctx context.Context, url string, seconds uint64) error {
	cmd := exec.CommandContext(ctx,
		"cast", "send",
		"--mnemonic", TestMnemonic,
		"--rpc-url", url,
		"clock", "increaseTime(uint256)", strconv.FormatUint(seconds, 10),
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cast: %w: %v", err, string(output))
	}
	return nil
}
