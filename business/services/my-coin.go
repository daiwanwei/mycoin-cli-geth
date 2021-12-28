package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"mycoin-cli-geth/contracts/mycoin"
)

type MyCoinService interface {
	BalanceOf(ctx context.Context, accountAddress, contractAddress string) (balance *big.Int, err error)
	Transfer(ctx context.Context, toAddress, contractAddress string, amount int64) (err error)
}

type myCoinService struct {
	provider *ethclient.Client
	signer   *bind.TransactOpts
}

func NewMyCoinService(provider *ethclient.Client, signer *bind.TransactOpts) (MyCoinService, error) {
	return &myCoinService{
		provider: provider,
		signer:   signer,
	}, nil
}

func (service *myCoinService) BalanceOf(ctx context.Context, accountAddress, contractAddress string) (balance *big.Int, err error) {
	account := common.HexToAddress(accountAddress)
	address := common.HexToAddress(contractAddress)
	instance, err := mycoin.NewMycoin(address, service.provider)
	if err != nil {
		panic(err)
	}
	balance, err = instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		panic(err)
	}
	return
}

func (service *myCoinService) Transfer(ctx context.Context, toAddress, contractAddress string, amount int64) (err error) {
	to := common.HexToAddress(toAddress)
	address := common.HexToAddress(contractAddress)
	contract, err := mycoin.NewMycoin(address, service.provider)
	if err != nil {
		panic(err)
	}
	txn, err := contract.Transfer(service.signer, to, big.NewInt(amount))
	if err != nil {
		panic(err)
	}
	fmt.Println(txn)
	return
}
