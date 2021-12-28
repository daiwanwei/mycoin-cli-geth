package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"mycoin-cli-geth/contracts/mycoin"
	"strings"
)

type EventService interface {
	FindAllEvent(ctx context.Context, contractAddress string) (err error)
}

type eventService struct {
	provider *ethclient.Client
}

func NewEventService(provider *ethclient.Client) (EventService, error) {
	return &eventService{
		provider: provider,
	}, nil
}

func (service *eventService) FindAllEvent(ctx context.Context, contractAddress string) (err error) {
	address := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(9),
		Addresses: []common.Address{
			address,
		},
	}

	logs, err := service.provider.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(mycoin.MycoinABI))
	if err != nil {
		return err
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				return err
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				panic(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Value.String())
		}

		fmt.Printf("\n\n")
	}
	return nil
}

// LogTransfer ..
type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Value      *big.Int
}
