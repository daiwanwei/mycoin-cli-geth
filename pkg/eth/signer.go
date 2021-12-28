package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewSinger(provider *ethclient.Client, privateKey string) (signer *bind.TransactOpts, err error) {
	privateKeyHex, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKeyHex.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := provider.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := provider.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println(gasPrice)
	auth := bind.NewKeyedTransactor(privateKeyHex)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}
