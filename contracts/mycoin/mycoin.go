// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycoin

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MycoinMetaData contains all meta data concerning the Mycoin contracts.
var MycoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600781526020017f576f6f436f696e000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f574300000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000256565b508060049080519060200190620000af92919062000256565b505050620000cd33683635c9adc5dea00000620000d360201b60201c565b620004b2565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000146576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200013d906200033e565b60405180910390fd5b6200015a600083836200024c60201b60201c565b80600260008282546200016e91906200038e565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001c591906200038e565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200022c919062000360565b60405180910390a362000248600083836200025160201b60201c565b5050565b505050565b505050565b8280546200026490620003f5565b90600052602060002090601f016020900481019282620002885760008555620002d4565b82601f10620002a357805160ff1916838001178555620002d4565b82800160010185558215620002d4579182015b82811115620002d3578251825591602001919060010190620002b6565b5b509050620002e39190620002e7565b5090565b5b8082111562000302576000816000905550600101620002e8565b5090565b600062000315601f836200037d565b9150620003228262000489565b602082019050919050565b6200033881620003eb565b82525050565b60006020820190508181036000830152620003598162000306565b9050919050565b60006020820190506200037760008301846200032d565b92915050565b600082825260208201905092915050565b60006200039b82620003eb565b9150620003a883620003eb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620003e057620003df6200042b565b5b828201905092915050565b6000819050919050565b600060028204905060018216806200040e57607f821691505b602082108114156200042557620004246200045a565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b61139580620004c26000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610e35565b60405180910390f35b6100e660048036038101906100e19190610c83565b610308565b6040516100f39190610e1a565b60405180910390f35b610104610326565b6040516101119190610f37565b60405180910390f35b610134600480360381019061012f9190610c34565b610330565b6040516101419190610e1a565b60405180910390f35b610152610428565b60405161015f9190610f52565b60405180910390f35b610182600480360381019061017d9190610c83565b610431565b60405161018f9190610e1a565b60405180910390f35b6101b260048036038101906101ad9190610bcf565b6104dd565b6040516101bf9190610f37565b60405180910390f35b6101d0610525565b6040516101dd9190610e35565b60405180910390f35b61020060048036038101906101fb9190610c83565b6105b7565b60405161020d9190610e1a565b60405180910390f35b610230600480360381019061022b9190610c83565b6106a2565b60405161023d9190610e1a565b60405180910390f35b610260600480360381019061025b9190610bf8565b6106c0565b60405161026d9190610f37565b60405180910390f35b60606003805461028590611067565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190611067565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b600061031c610315610747565b848461074f565b6001905092915050565b6000600254905090565b600061033d84848461091a565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610388610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610408576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ff90610eb7565b60405180910390fd5b61041c85610414610747565b85840361074f565b60019150509392505050565b60006012905090565b60006104d361043e610747565b84846001600061044c610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104ce9190610f89565b61074f565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60606004805461053490611067565b80601f016020809104026020016040519081016040528092919081815260200182805461056090611067565b80156105ad5780601f10610582576101008083540402835291602001916105ad565b820191906000526020600020905b81548152906001019060200180831161059057829003601f168201915b5050505050905090565b600080600160006105c6610747565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067a90610f17565b60405180910390fd5b61069761068e610747565b8585840361074f565b600191505092915050565b60006106b66106af610747565b848461091a565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156107bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b690610ef7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561082f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082690610e77565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161090d9190610f37565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561098a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098190610ed7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156109fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f190610e57565b60405180910390fd5b610a05838383610b9b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610a8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8290610e97565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b1e9190610f89565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b829190610f37565b60405180910390a3610b95848484610ba0565b50505050565b505050565b505050565b600081359050610bb481611331565b92915050565b600081359050610bc981611348565b92915050565b600060208284031215610be157600080fd5b6000610bef84828501610ba5565b91505092915050565b60008060408385031215610c0b57600080fd5b6000610c1985828601610ba5565b9250506020610c2a85828601610ba5565b9150509250929050565b600080600060608486031215610c4957600080fd5b6000610c5786828701610ba5565b9350506020610c6886828701610ba5565b9250506040610c7986828701610bba565b9150509250925092565b60008060408385031215610c9657600080fd5b6000610ca485828601610ba5565b9250506020610cb585828601610bba565b9150509250929050565b610cc881610ff1565b82525050565b6000610cd982610f6d565b610ce38185610f78565b9350610cf3818560208601611034565b610cfc816110f7565b840191505092915050565b6000610d14602383610f78565b9150610d1f82611108565b604082019050919050565b6000610d37602283610f78565b9150610d4282611157565b604082019050919050565b6000610d5a602683610f78565b9150610d65826111a6565b604082019050919050565b6000610d7d602883610f78565b9150610d88826111f5565b604082019050919050565b6000610da0602583610f78565b9150610dab82611244565b604082019050919050565b6000610dc3602483610f78565b9150610dce82611293565b604082019050919050565b6000610de6602583610f78565b9150610df1826112e2565b604082019050919050565b610e058161101d565b82525050565b610e1481611027565b82525050565b6000602082019050610e2f6000830184610cbf565b92915050565b60006020820190508181036000830152610e4f8184610cce565b905092915050565b60006020820190508181036000830152610e7081610d07565b9050919050565b60006020820190508181036000830152610e9081610d2a565b9050919050565b60006020820190508181036000830152610eb081610d4d565b9050919050565b60006020820190508181036000830152610ed081610d70565b9050919050565b60006020820190508181036000830152610ef081610d93565b9050919050565b60006020820190508181036000830152610f1081610db6565b9050919050565b60006020820190508181036000830152610f3081610dd9565b9050919050565b6000602082019050610f4c6000830184610dfc565b92915050565b6000602082019050610f676000830184610e0b565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610f948261101d565b9150610f9f8361101d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610fd457610fd3611099565b5b828201905092915050565b6000610fea82610ffd565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611052578082015181840152602081019050611037565b83811115611061576000848401525b50505050565b6000600282049050600182168061107f57607f821691505b60208210811415611093576110926110c8565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b61133a81610fdf565b811461134557600080fd5b50565b6113518161101d565b811461135c57600080fd5b5056fea264697066735822122081d488bfd6a116872c744cc7d981793b37ea5fb8f04d00ffcbc01afec9b06e7864736f6c63430008040033",
}

// MycoinABI is the input ABI used to generate the binding from.
// Deprecated: Use MycoinMetaData.ABI instead.
var MycoinABI = MycoinMetaData.ABI

// MycoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MycoinMetaData.Bin instead.
var MycoinBin = MycoinMetaData.Bin

// DeployMycoin deploys a new Ethereum contracts, binding an instance of Mycoin to it.
func DeployMycoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mycoin, error) {
	parsed, err := MycoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MycoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mycoin{MycoinCaller: MycoinCaller{contract: contract}, MycoinTransactor: MycoinTransactor{contract: contract}, MycoinFilterer: MycoinFilterer{contract: contract}}, nil
}

// Mycoin is an auto generated Go binding around an Ethereum contracts.
type Mycoin struct {
	MycoinCaller     // Read-only binding to the contracts
	MycoinTransactor // Write-only binding to the contracts
	MycoinFilterer   // Log filterer for contracts events
}

// MycoinCaller is an auto generated read-only Go binding around an Ethereum contracts.
type MycoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycoinTransactor is an auto generated write-only Go binding around an Ethereum contracts.
type MycoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycoinFilterer is an auto generated log filtering Go binding around an Ethereum contracts events.
type MycoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycoinSession is an auto generated Go binding around an Ethereum contracts,
// with pre-set call and transact options.
type MycoinSession struct {
	Contract     *Mycoin           // Generic contracts binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycoinCallerSession is an auto generated read-only Go binding around an Ethereum contracts,
// with pre-set call options.
type MycoinCallerSession struct {
	Contract *MycoinCaller // Generic contracts caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MycoinTransactorSession is an auto generated write-only Go binding around an Ethereum contracts,
// with pre-set transact options.
type MycoinTransactorSession struct {
	Contract     *MycoinTransactor // Generic contracts transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycoinRaw is an auto generated low-level Go binding around an Ethereum contracts.
type MycoinRaw struct {
	Contract *Mycoin // Generic contracts binding to access the raw methods on
}

// MycoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contracts.
type MycoinCallerRaw struct {
	Contract *MycoinCaller // Generic read-only contracts binding to access the raw methods on
}

// MycoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contracts.
type MycoinTransactorRaw struct {
	Contract *MycoinTransactor // Generic write-only contracts binding to access the raw methods on
}

// NewMycoin creates a new instance of Mycoin, bound to a specific deployed contracts.
func NewMycoin(address common.Address, backend bind.ContractBackend) (*Mycoin, error) {
	contract, err := bindMycoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycoin{MycoinCaller: MycoinCaller{contract: contract}, MycoinTransactor: MycoinTransactor{contract: contract}, MycoinFilterer: MycoinFilterer{contract: contract}}, nil
}

// NewMycoinCaller creates a new read-only instance of Mycoin, bound to a specific deployed contracts.
func NewMycoinCaller(address common.Address, caller bind.ContractCaller) (*MycoinCaller, error) {
	contract, err := bindMycoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycoinCaller{contract: contract}, nil
}

// NewMycoinTransactor creates a new write-only instance of Mycoin, bound to a specific deployed contracts.
func NewMycoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MycoinTransactor, error) {
	contract, err := bindMycoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycoinTransactor{contract: contract}, nil
}

// NewMycoinFilterer creates a new log filterer instance of Mycoin, bound to a specific deployed contracts.
func NewMycoinFilterer(address common.Address, filterer bind.ContractFilterer) (*MycoinFilterer, error) {
	contract, err := bindMycoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycoinFilterer{contract: contract}, nil
}

// bindMycoin binds a generic wrapper to an already deployed contracts.
func bindMycoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MycoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contracts method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycoin *MycoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycoin.Contract.MycoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contracts, calling
// its default method if one is available.
func (_Mycoin *MycoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycoin.Contract.MycoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contracts method with params as input values.
func (_Mycoin *MycoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycoin.Contract.MycoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contracts method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycoin *MycoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contracts, calling
// its default method if one is available.
func (_Mycoin *MycoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contracts method with params as input values.
func (_Mycoin *MycoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contracts method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mycoin *MycoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contracts method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mycoin *MycoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mycoin.Contract.Allowance(&_Mycoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contracts method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mycoin *MycoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mycoin.Contract.Allowance(&_Mycoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contracts method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mycoin *MycoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contracts method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mycoin *MycoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mycoin.Contract.BalanceOf(&_Mycoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contracts method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mycoin *MycoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mycoin.Contract.BalanceOf(&_Mycoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contracts method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mycoin *MycoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contracts method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mycoin *MycoinSession) Decimals() (uint8, error) {
	return _Mycoin.Contract.Decimals(&_Mycoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contracts method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mycoin *MycoinCallerSession) Decimals() (uint8, error) {
	return _Mycoin.Contract.Decimals(&_Mycoin.CallOpts)
}

// Name is a free data retrieval call binding the contracts method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycoin *MycoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contracts method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycoin *MycoinSession) Name() (string, error) {
	return _Mycoin.Contract.Name(&_Mycoin.CallOpts)
}

// Name is a free data retrieval call binding the contracts method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycoin *MycoinCallerSession) Name() (string, error) {
	return _Mycoin.Contract.Name(&_Mycoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contracts method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycoin *MycoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contracts method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycoin *MycoinSession) Symbol() (string, error) {
	return _Mycoin.Contract.Symbol(&_Mycoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contracts method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycoin *MycoinCallerSession) Symbol() (string, error) {
	return _Mycoin.Contract.Symbol(&_Mycoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contracts method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycoin *MycoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contracts method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycoin *MycoinSession) TotalSupply() (*big.Int, error) {
	return _Mycoin.Contract.TotalSupply(&_Mycoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contracts method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mycoin *MycoinCallerSession) TotalSupply() (*big.Int, error) {
	return _Mycoin.Contract.TotalSupply(&_Mycoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contracts method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contracts method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mycoin *MycoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.Approve(&_Mycoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contracts method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.Approve(&_Mycoin.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contracts method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mycoin *MycoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contracts method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mycoin *MycoinSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.DecreaseAllowance(&_Mycoin.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contracts method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mycoin *MycoinTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.DecreaseAllowance(&_Mycoin.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contracts method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mycoin *MycoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contracts method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mycoin *MycoinSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.IncreaseAllowance(&_Mycoin.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contracts method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mycoin *MycoinTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.IncreaseAllowance(&_Mycoin.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contracts method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contracts method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.Transfer(&_Mycoin.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contracts method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.Transfer(&_Mycoin.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contracts method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contracts method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.TransferFrom(&_Mycoin.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contracts method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mycoin *MycoinTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mycoin.Contract.TransferFrom(&_Mycoin.TransactOpts, sender, recipient, amount)
}

// MycoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mycoin contracts.
type MycoinApprovalIterator struct {
	Event *MycoinApproval // Event containing the contracts specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycoinApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycoinApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycoinApproval represents a Approval event raised by the Mycoin contracts.
type MycoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contracts event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mycoin *MycoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MycoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mycoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MycoinApprovalIterator{contract: _Mycoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contracts event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mycoin *MycoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MycoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mycoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycoinApproval)
				if err := _Mycoin.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contracts event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mycoin *MycoinFilterer) ParseApproval(log types.Log) (*MycoinApproval, error) {
	event := new(MycoinApproval)
	if err := _Mycoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mycoin contracts.
type MycoinTransferIterator struct {
	Event *MycoinTransfer // Event containing the contracts specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MycoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycoinTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MycoinTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MycoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycoinTransfer represents a Transfer event raised by the Mycoin contracts.
type MycoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contracts event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mycoin *MycoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MycoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MycoinTransferIterator{contract: _Mycoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contracts event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mycoin *MycoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MycoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mycoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycoinTransfer)
				if err := _Mycoin.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contracts event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mycoin *MycoinFilterer) ParseTransfer(log types.Log) (*MycoinTransfer, error) {
	event := new(MycoinTransfer)
	if err := _Mycoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
