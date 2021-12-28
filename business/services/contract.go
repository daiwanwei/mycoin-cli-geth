package services

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"mycoin-cli-geth/contracts/mycoin"
)

type ContractService interface {
	DeployMyCoin(ctx context.Context) (contractAddress string, err error)
}

type contractService struct {
	provider *ethclient.Client
	signer   *bind.TransactOpts
}

func NewContractService(provider *ethclient.Client, signer *bind.TransactOpts) (ContractService, error) {
	return &contractService{
		provider: provider,
		signer:   signer,
	}, nil
}

func (service *contractService) DeployMyCoin(ctx context.Context) (contractAddress string, err error) {
	address, _, _, err := mycoin.DeployMycoin(service.signer, service.provider)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}
