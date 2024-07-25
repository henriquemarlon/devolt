package integration

import (
	"context"
	"fmt"
	"github.com/devolthq/devolt/test/integration/artifacts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type TokenInterface struct {
	Client *ethclient.Client
	Opts   *bind.TransactOpts
}

func NewTokenInterface(url string, sender string) (*TokenInterface, error) {
	client, opts, err := setup(url, sender)
	if err != nil {
		return nil, err
	}
	return &TokenInterface{Client: client, Opts: opts}, nil
}

func (dc *TokenInterface) Deploy(initialOwner string, name string, symbol string) (*common.Address, error) {
	tokenAddress, tx, _, err := artifacts.DeployToken(dc.Opts, dc.Client, common.HexToAddress(initialOwner), name, symbol)
	if err != nil {
		return nil, err
	}
	_, err = bind.WaitMined(context.Background(), dc.Client, tx)
	if err != nil {
		return nil, err
	}
	return &tokenAddress, nil
}

func (dc *TokenInterface) Mint(
	token common.Address,
	to string,
	amount *big.Int,
) error {
	instance, err := artifacts.NewToken(token, dc.Client)
	if err != nil {
		return err
	}
	tx, err := instance.Mint(dc.Opts, common.HexToAddress(to), amount)
	if err != nil {
		return err
	}
	_, err = bind.WaitMined(context.Background(), dc.Client, tx)
	if err != nil {
		return err
	}
	return nil
}

func setup(rpcurl string, sender string) (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(rpcurl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to blockchain: %v", err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get chain id: %v", err)
	}
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(AddressToPrivateKey[common.HexToAddress(sender)], "0x"))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load private key: %v", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	return client, opts, err
}
