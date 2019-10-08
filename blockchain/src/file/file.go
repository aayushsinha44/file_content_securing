// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package file

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FileABI is the input ABI used to generate the binding from.
const FileABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"file_name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_owner_list\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"change_owner_power\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"add_key\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"get_owner_power\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_ipfs_hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_key\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_power\",\"type\":\"uint256\"}],\"name\":\"add_owner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file_name\",\"type\":\"string\"}],\"name\":\"updateFileName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"update_hash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"update_key\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FileBin is the compiled bytecode used for deploying new contracts.
var FileBin = "0x60806040523480156200001157600080fd5b5060405162001f4738038062001f47833981810160405260208110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052505050806000908051906020019062000110929190620002c2565b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600460036000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160046000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505062000371565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030557805160ff191683800117855562000336565b8280016001018555821562000336579182015b828111156200033557825182559160200191906001019062000318565b5b50905062000345919062000349565b5090565b6200036e91905b808211156200036a57600081600090555060010162000350565b5090565b90565b611bc680620003816000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638d0651801161008c578063eac7381311610066578063eac73813146104d5578063f73eb6f314610523578063fc65e5a7146105de578063fcb6c39a14610699576100cf565b80638d06518014610377578063d131f1f5146103cf578063e072b5f614610452576100cf565b8063025e7c27146100d457806302d05d3f1461014257806324e4778d1461018c57806383ced6271461020f578063866309fa1461026e5780638c5cd3e1146102bc575b600080fd5b610100600480360360208110156100ea57600080fd5b8101908080359060200190929190505050610754565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61014a610790565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101946107b6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101d45780820151818401526020810190506101b9565b50505050905090810190601f1680156102015780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610217610854565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561025a57808201518184015260208101905061023f565b505050509050019250505060405180910390f35b6102ba6004803603604081101561028457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a59565b005b610375600480360360208110156102d257600080fd5b81019080803590602001906401000000008111156102ef57600080fd5b82018360208201111561030157600080fd5b8035906020019184600183028401116401000000008311171561032357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d82565b005b6103b96004803603602081101561038d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d9c565b6040518082815260200191505060405180910390f35b6103d7610f5c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104175780820151818401526020810190506103fc565b50505050905090810190601f1680156104445780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61045a611175565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561049a57808201518184015260208101905061047f565b50505050905090810190601f1680156104c75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610521600480360360408110156104eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061138e565b005b6105dc6004803603602081101561053957600080fd5b810190808035906020019064010000000081111561055657600080fd5b82018360208201111561056857600080fd5b8035906020019184600183028401116401000000008311171561058a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506116cd565b005b610697600480360360208110156105f457600080fd5b810190808035906020019064010000000081111561061157600080fd5b82018360208201111561062357600080fd5b8035906020019184600183028401116401000000008311171561064557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061185d565b005b610752600480360360208110156106af57600080fd5b81019080803590602001906401000000008111156106cc57600080fd5b8201836020820111156106de57600080fd5b8035906020019184600183028401116401000000008311171561070057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506119ee565b005b6002818154811061076157fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561084c5780601f106108215761010080835404028352916020019161084c565b820191906000526020600020905b81548152906001019060200180831161082f57829003601f168201915b505050505081565b60606001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610917576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156109cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b6002805480602002602001604051908101604052809291908181526020018280548015610a4e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a04575b505050505091505090565b81600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611b716021913960400191505060405180910390fd5b82600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610bc1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b6003600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c82576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610d37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b83600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050565b8060059080519060200190610d98929190611acb565b5050565b60006001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e5f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610f14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054915050919050565b60606001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661101f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156110d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60068054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561116a5780601f1061113f5761010080835404028352916020019161116a565b820191906000526020600020905b81548152906001019060200180831161114d57829003601f168201915b505050505091505090565b60606001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611238576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156112ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113835780601f1061135857610100808354040283529160200191611383565b820191906000526020600020905b81548152906001019060200180831161136657829003601f168201915b505050505091505090565b81600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561144f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220616c72656164792065786973747300000000000000000000000081525060200191505060405180910390fd5b6003600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16611510576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156115c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b60028490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b60048060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661178d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b8160009080519060200190611858929190611acb565b505050565b6002600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661191e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6f776e657220646f65736e6f742065786973747300000000000000000000000081525060200191505060405180910390fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156119d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f646f65736e6f7420686176652073756666696369656e7420706f77657200000081525060200191505060405180910390fd5b81600690805190602001906119e9929190611acb565b505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ab1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6f6e6c792063726561746f722063616e2061636365737300000000000000000081525060200191505060405180910390fd5b8060059080519060200190611ac7929190611acb565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b0c57805160ff1916838001178555611b3a565b82800160010185558215611b3a579182015b82811115611b39578251825591602001919060010190611b1e565b5b509050611b479190611b4b565b5090565b611b6d91905b80821115611b69576000816000905550600101611b51565b5090565b9056fe63616e6e6f7420706572666f726d206f7065726174696f6e206f6e206f776e6572a265627a7a72315820a828b36995879a2d60ad8998c094625fc67380bd303f194e171c46eb11b72ea064736f6c634300050b0032"

// DeployFile deploys a new Ethereum contract, binding an instance of File to it.
func DeployFile(auth *bind.TransactOpts, backend bind.ContractBackend, _file_name string) (common.Address, *types.Transaction, *File, error) {
	parsed, err := abi.JSON(strings.NewReader(FileABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileBin), backend, _file_name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &File{FileCaller: FileCaller{contract: contract}, FileTransactor: FileTransactor{contract: contract}, FileFilterer: FileFilterer{contract: contract}}, nil
}

// File is an auto generated Go binding around an Ethereum contract.
type File struct {
	FileCaller     // Read-only binding to the contract
	FileTransactor // Write-only binding to the contract
	FileFilterer   // Log filterer for contract events
}

// FileCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileSession struct {
	Contract     *File             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileCallerSession struct {
	Contract *FileCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileTransactorSession struct {
	Contract     *FileTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileRaw struct {
	Contract *File // Generic contract binding to access the raw methods on
}

// FileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileCallerRaw struct {
	Contract *FileCaller // Generic read-only contract binding to access the raw methods on
}

// FileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileTransactorRaw struct {
	Contract *FileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFile creates a new instance of File, bound to a specific deployed contract.
func NewFile(address common.Address, backend bind.ContractBackend) (*File, error) {
	contract, err := bindFile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &File{FileCaller: FileCaller{contract: contract}, FileTransactor: FileTransactor{contract: contract}, FileFilterer: FileFilterer{contract: contract}}, nil
}

// NewFileCaller creates a new read-only instance of File, bound to a specific deployed contract.
func NewFileCaller(address common.Address, caller bind.ContractCaller) (*FileCaller, error) {
	contract, err := bindFile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileCaller{contract: contract}, nil
}

// NewFileTransactor creates a new write-only instance of File, bound to a specific deployed contract.
func NewFileTransactor(address common.Address, transactor bind.ContractTransactor) (*FileTransactor, error) {
	contract, err := bindFile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileTransactor{contract: contract}, nil
}

// NewFileFilterer creates a new log filterer instance of File, bound to a specific deployed contract.
func NewFileFilterer(address common.Address, filterer bind.ContractFilterer) (*FileFilterer, error) {
	contract, err := bindFile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileFilterer{contract: contract}, nil
}

// bindFile binds a generic wrapper to an already deployed contract.
func bindFile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_File *FileRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _File.Contract.FileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_File *FileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _File.Contract.FileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_File *FileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _File.Contract.FileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_File *FileCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _File.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_File *FileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _File.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_File *FileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _File.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_File *FileCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_File *FileSession) Creator() (common.Address, error) {
	return _File.Contract.Creator(&_File.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_File *FileCallerSession) Creator() (common.Address, error) {
	return _File.Contract.Creator(&_File.CallOpts)
}

// FileName is a free data retrieval call binding the contract method 0x24e4778d.
//
// Solidity: function file_name() constant returns(string)
func (_File *FileCaller) FileName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "file_name")
	return *ret0, err
}

// FileName is a free data retrieval call binding the contract method 0x24e4778d.
//
// Solidity: function file_name() constant returns(string)
func (_File *FileSession) FileName() (string, error) {
	return _File.Contract.FileName(&_File.CallOpts)
}

// FileName is a free data retrieval call binding the contract method 0x24e4778d.
//
// Solidity: function file_name() constant returns(string)
func (_File *FileCallerSession) FileName() (string, error) {
	return _File.Contract.FileName(&_File.CallOpts)
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_File *FileCaller) GetIpfsHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "get_ipfs_hash")
	return *ret0, err
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_File *FileSession) GetIpfsHash() (string, error) {
	return _File.Contract.GetIpfsHash(&_File.CallOpts)
}

// GetIpfsHash is a free data retrieval call binding the contract method 0xd131f1f5.
//
// Solidity: function get_ipfs_hash() constant returns(string)
func (_File *FileCallerSession) GetIpfsHash() (string, error) {
	return _File.Contract.GetIpfsHash(&_File.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_File *FileCaller) GetKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "get_key")
	return *ret0, err
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_File *FileSession) GetKey() (string, error) {
	return _File.Contract.GetKey(&_File.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0xe072b5f6.
//
// Solidity: function get_key() constant returns(string)
func (_File *FileCallerSession) GetKey() (string, error) {
	return _File.Contract.GetKey(&_File.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_File *FileCaller) GetOwnerList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "get_owner_list")
	return *ret0, err
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_File *FileSession) GetOwnerList() ([]common.Address, error) {
	return _File.Contract.GetOwnerList(&_File.CallOpts)
}

// GetOwnerList is a free data retrieval call binding the contract method 0x83ced627.
//
// Solidity: function get_owner_list() constant returns(address[])
func (_File *FileCallerSession) GetOwnerList() ([]common.Address, error) {
	return _File.Contract.GetOwnerList(&_File.CallOpts)
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_File *FileCaller) GetOwnerPower(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "get_owner_power", _owner)
	return *ret0, err
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_File *FileSession) GetOwnerPower(_owner common.Address) (*big.Int, error) {
	return _File.Contract.GetOwnerPower(&_File.CallOpts, _owner)
}

// GetOwnerPower is a free data retrieval call binding the contract method 0x8d065180.
//
// Solidity: function get_owner_power(address _owner) constant returns(uint256)
func (_File *FileCallerSession) GetOwnerPower(_owner common.Address) (*big.Int, error) {
	return _File.Contract.GetOwnerPower(&_File.CallOpts, _owner)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_File *FileCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_File *FileSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _File.Contract.Owners(&_File.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_File *FileCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _File.Contract.Owners(&_File.CallOpts, arg0)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_File *FileTransactor) AddKey(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "add_key", _key)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_File *FileSession) AddKey(_key string) (*types.Transaction, error) {
	return _File.Contract.AddKey(&_File.TransactOpts, _key)
}

// AddKey is a paid mutator transaction binding the contract method 0x8c5cd3e1.
//
// Solidity: function add_key(string _key) returns()
func (_File *FileTransactorSession) AddKey(_key string) (*types.Transaction, error) {
	return _File.Contract.AddKey(&_File.TransactOpts, _key)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_File *FileTransactor) AddOwner(opts *bind.TransactOpts, _owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "add_owner", _owner, _power)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_File *FileSession) AddOwner(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.Contract.AddOwner(&_File.TransactOpts, _owner, _power)
}

// AddOwner is a paid mutator transaction binding the contract method 0xeac73813.
//
// Solidity: function add_owner(address _owner, uint256 _power) returns()
func (_File *FileTransactorSession) AddOwner(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.Contract.AddOwner(&_File.TransactOpts, _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_File *FileTransactor) ChangeOwnerPower(opts *bind.TransactOpts, _owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "change_owner_power", _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_File *FileSession) ChangeOwnerPower(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.Contract.ChangeOwnerPower(&_File.TransactOpts, _owner, _power)
}

// ChangeOwnerPower is a paid mutator transaction binding the contract method 0x866309fa.
//
// Solidity: function change_owner_power(address _owner, uint256 _power) returns()
func (_File *FileTransactorSession) ChangeOwnerPower(_owner common.Address, _power *big.Int) (*types.Transaction, error) {
	return _File.Contract.ChangeOwnerPower(&_File.TransactOpts, _owner, _power)
}

// UpdateFileName is a paid mutator transaction binding the contract method 0xf73eb6f3.
//
// Solidity: function updateFileName(string _file_name) returns()
func (_File *FileTransactor) UpdateFileName(opts *bind.TransactOpts, _file_name string) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "updateFileName", _file_name)
}

// UpdateFileName is a paid mutator transaction binding the contract method 0xf73eb6f3.
//
// Solidity: function updateFileName(string _file_name) returns()
func (_File *FileSession) UpdateFileName(_file_name string) (*types.Transaction, error) {
	return _File.Contract.UpdateFileName(&_File.TransactOpts, _file_name)
}

// UpdateFileName is a paid mutator transaction binding the contract method 0xf73eb6f3.
//
// Solidity: function updateFileName(string _file_name) returns()
func (_File *FileTransactorSession) UpdateFileName(_file_name string) (*types.Transaction, error) {
	return _File.Contract.UpdateFileName(&_File.TransactOpts, _file_name)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_File *FileTransactor) UpdateHash(opts *bind.TransactOpts, _hash string) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "update_hash", _hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_File *FileSession) UpdateHash(_hash string) (*types.Transaction, error) {
	return _File.Contract.UpdateHash(&_File.TransactOpts, _hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0xfc65e5a7.
//
// Solidity: function update_hash(string _hash) returns()
func (_File *FileTransactorSession) UpdateHash(_hash string) (*types.Transaction, error) {
	return _File.Contract.UpdateHash(&_File.TransactOpts, _hash)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_File *FileTransactor) UpdateKey(opts *bind.TransactOpts, _key string) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "update_key", _key)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_File *FileSession) UpdateKey(_key string) (*types.Transaction, error) {
	return _File.Contract.UpdateKey(&_File.TransactOpts, _key)
}

// UpdateKey is a paid mutator transaction binding the contract method 0xfcb6c39a.
//
// Solidity: function update_key(string _key) returns()
func (_File *FileTransactorSession) UpdateKey(_key string) (*types.Transaction, error) {
	return _File.Contract.UpdateKey(&_File.TransactOpts, _key)
}
