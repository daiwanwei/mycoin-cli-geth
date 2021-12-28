package eth

import "github.com/ethereum/go-ethereum/ethclient"

func NewProvider(providerUrl string) (provider *ethclient.Client, err error) {
	provider, err = ethclient.Dial(providerUrl)
	if err != nil {
		return nil, err
	}
	return provider, nil
}
