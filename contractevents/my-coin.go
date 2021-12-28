package contractevents

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

type MyCoinSubManager struct {
	provider  *ethclient.Client
	isRunning bool
	subs      map[string]*MyCoinSub
	subCh     chan string
	unsubCh   chan string
	stopCh    chan struct{}
}

func NewMyCoinSubManager(provider *ethclient.Client) (*MyCoinSubManager, error) {
	return &MyCoinSubManager{
		provider:  provider,
		subs:      make(map[string]*MyCoinSub),
		subCh:     make(chan string, 20),
		unsubCh:   make(chan string, 20),
		stopCh:    make(chan struct{}),
		isRunning: false,
	}, nil
}

func (manager *MyCoinSubManager) Run() error {
	for {
		select {
		case address := <-manager.subCh:
			if _, isExisted := manager.subs[address]; !isExisted {
				sub, err := NewMyCoinSub(manager.provider, address)
				if err != nil {
					fmt.Printf("contract(%s) subscribe fail: %s", address, err)
					break
				}
				manager.subs[address] = sub
				go func() {
					err = sub.Run()
					if err != nil {
						fmt.Printf("contract(%s) subscriber start fail: %s", address, err)
					}
					delete(manager.subs, address)
				}()
				fmt.Printf("contract(%s) subscribe successfully", address)
			}
		case address := <-manager.unsubCh:
			sub, isExisted := manager.subs[address]
			if !isExisted {
				break
			}
			delete(manager.subs, address)
			err := sub.Stop()
			if err != nil {
				fmt.Printf("contract(%s) subscriber stop fail: %s", address, err)
			}
			fmt.Printf("contract(%s) unsubscribe successfully", address)
		case <-manager.stopCh:
			for address, sub := range manager.subs {
				err := sub.Stop()
				if err != nil {
					fmt.Printf("contract(%s) subscriber stop fail: %s", address, err)
				}
			}
			close(manager.subCh)
			close(manager.unsubCh)
			fmt.Printf("manager stop successfully")
			return nil
		}
	}
}

func (manager *MyCoinSubManager) Close() {
	close(manager.stopCh)
}

func (manager *MyCoinSubManager) Subscribe(myCoinAddress string) error {
	select {
	case manager.subCh <- myCoinAddress:
		return nil
	case <-time.After(time.Second * 5):
		return errors.New("timeout for subscribe")
	}
}

func (manager *MyCoinSubManager) Unsubscribe(myCoinAddress string) error {
	select {
	case manager.subCh <- myCoinAddress:
		return nil
	case <-time.After(time.Second * 5):
		return errors.New("timeout for subscribe")
	}
}

type MyCoinSub struct {
	logCh  chan types.Log
	ethSub ethereum.Subscription
	stopCh chan struct{}
}

func NewMyCoinSub(provider *ethclient.Client, myCoinAddress string) (*MyCoinSub, error) {
	address := common.HexToAddress(myCoinAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logCh := make(chan types.Log)
	sub, err := provider.SubscribeFilterLogs(context.Background(), query, logCh)
	if err != nil {
		return nil, err
	}
	return &MyCoinSub{
		logCh:  logCh,
		ethSub: sub,
		stopCh: make(chan struct{}),
	}, nil
}

func (sub *MyCoinSub) Run() error {
	for {
		select {
		case err := <-sub.ethSub.Err():
			fmt.Println(err)
		case log := <-sub.logCh:
			fmt.Println(log) // pointer to event log
		case <-sub.stopCh:
			sub.ethSub.Unsubscribe()
			close(sub.logCh)
			fmt.Println("close sub")
			return nil
		}
	}
}

func (sub *MyCoinSub) Stop() error {
	close(sub.stopCh)
	return nil
}
