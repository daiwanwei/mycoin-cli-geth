package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type AccountService interface {
	Balance(ctx context.Context, accountAddress string) (balance *big.Int, err error)
}

type accountService struct {
	provider *ethclient.Client
}

func NewAccountService(provider *ethclient.Client) (AccountService, error) {
	return &accountService{
		provider: provider,
	}, nil
}

func (service *accountService) Balance(ctx context.Context, accountAddress string) (balance *big.Int, err error) {
	account := common.HexToAddress(accountAddress)
	balance, err = service.provider.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
