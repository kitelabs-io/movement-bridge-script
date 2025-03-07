// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// InboundPacket is an auto generated low-level Go binding around an user-defined struct.
type InboundPacket struct {
	Origin    Origin
	DstEid    uint32
	Receiver  common.Address
	Guid      [32]byte
	Value     *big.Int
	Executor  common.Address
	Message   []byte
	ExtraData []byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// AdapterMetaData contains all meta data concerning the Adapter contract.
var AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lzEndpoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEndpointCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLocalDecimals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"InvalidOptions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"NoPeer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"}],\"name\":\"NotEnoughNative\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"OnlyEndpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"OnlyPeer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"SimulationResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"}],\"name\":\"SlippageExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"EnforcedOptionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inspector\",\"type\":\"address\"}],\"name\":\"MsgInspectorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"name\":\"OFTSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"name\":\"PeerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preCrimeAddress\",\"type\":\"address\"}],\"name\":\"PreCrimeSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SEND\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEND_AND_CALL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"}],\"name\":\"allowInitializePath\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_extraOptions\",\"type\":\"bytes\"}],\"name\":\"combineOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalConversionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"}],\"name\":\"enforcedOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"enforcedOption\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isComposeMsgSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"isPeer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structInboundPacket[]\",\"name\":\"_packets\",\"type\":\"tuple[]\"}],\"name\":\"lzReceiveAndRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceiveSimulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgInspector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nextNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oAppVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"senderVersion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"receiverVersion\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oftVersion\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"}],\"name\":\"peers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCrime\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"}],\"name\":\"quoteOFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTLimit\",\"name\":\"oftLimit\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"feeAmountLD\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structOFTFeeDetail[]\",\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_payInLzToken\",\"type\":\"bool\"}],\"name\":\"quoteSend\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"msgFee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraOptions\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"composeMsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"oftCmd\",\"type\":\"bytes\"}],\"internalType\":\"structSendParam\",\"name\":\"_sendParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"_fee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"msgReceipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountSentLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountReceivedLD\",\"type\":\"uint256\"}],\"internalType\":\"structOFTReceipt\",\"name\":\"oftReceipt\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"msgType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"internalType\":\"structEnforcedOptionParam[]\",\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\"}],\"name\":\"setEnforcedOptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgInspector\",\"type\":\"address\"}],\"name\":\"setMsgInspector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_peer\",\"type\":\"bytes32\"}],\"name\":\"setPeer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_preCrime\",\"type\":\"address\"}],\"name\":\"setPreCrime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use AdapterMetaData.ABI instead.
var AdapterABI = AdapterMetaData.ABI

// Adapter is an auto generated Go binding around an Ethereum contract.
type Adapter struct {
	AdapterCaller     // Read-only binding to the contract
	AdapterTransactor // Write-only binding to the contract
	AdapterFilterer   // Log filterer for contract events
}

// AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdapterSession struct {
	Contract     *Adapter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdapterCallerSession struct {
	Contract *AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdapterTransactorSession struct {
	Contract     *AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdapterRaw struct {
	Contract *Adapter // Generic contract binding to access the raw methods on
}

// AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdapterCallerRaw struct {
	Contract *AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdapterTransactorRaw struct {
	Contract *AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdapter creates a new instance of Adapter, bound to a specific deployed contract.
func NewAdapter(address common.Address, backend bind.ContractBackend) (*Adapter, error) {
	contract, err := bindAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adapter{AdapterCaller: AdapterCaller{contract: contract}, AdapterTransactor: AdapterTransactor{contract: contract}, AdapterFilterer: AdapterFilterer{contract: contract}}, nil
}

// NewAdapterCaller creates a new read-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterCaller(address common.Address, caller bind.ContractCaller) (*AdapterCaller, error) {
	contract, err := bindAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterCaller{contract: contract}, nil
}

// NewAdapterTransactor creates a new write-only instance of Adapter, bound to a specific deployed contract.
func NewAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*AdapterTransactor, error) {
	contract, err := bindAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterTransactor{contract: contract}, nil
}

// NewAdapterFilterer creates a new log filterer instance of Adapter, bound to a specific deployed contract.
func NewAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*AdapterFilterer, error) {
	contract, err := bindAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdapterFilterer{contract: contract}, nil
}

// bindAdapter binds a generic wrapper to an already deployed contract.
func bindAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adapter *AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adapter *AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adapter *AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adapter.Contract.contract.Transact(opts, method, params...)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Adapter *AdapterCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Adapter *AdapterSession) SEND() (uint16, error) {
	return _Adapter.Contract.SEND(&_Adapter.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_Adapter *AdapterCallerSession) SEND() (uint16, error) {
	return _Adapter.Contract.SEND(&_Adapter.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Adapter *AdapterCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Adapter *AdapterSession) SENDANDCALL() (uint16, error) {
	return _Adapter.Contract.SENDANDCALL(&_Adapter.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_Adapter *AdapterCallerSession) SENDANDCALL() (uint16, error) {
	return _Adapter.Contract.SENDANDCALL(&_Adapter.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Adapter *AdapterCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Adapter *AdapterSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Adapter.Contract.AllowInitializePath(&_Adapter.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_Adapter *AdapterCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _Adapter.Contract.AllowInitializePath(&_Adapter.CallOpts, origin)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Adapter *AdapterCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Adapter *AdapterSession) ApprovalRequired() (bool, error) {
	return _Adapter.Contract.ApprovalRequired(&_Adapter.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_Adapter *AdapterCallerSession) ApprovalRequired() (bool, error) {
	return _Adapter.Contract.ApprovalRequired(&_Adapter.CallOpts)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Adapter *AdapterCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Adapter *AdapterSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Adapter.Contract.CombineOptions(&_Adapter.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_Adapter *AdapterCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _Adapter.Contract.CombineOptions(&_Adapter.CallOpts, _eid, _msgType, _extraOptions)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Adapter *AdapterCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Adapter *AdapterSession) DecimalConversionRate() (*big.Int, error) {
	return _Adapter.Contract.DecimalConversionRate(&_Adapter.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_Adapter *AdapterCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _Adapter.Contract.DecimalConversionRate(&_Adapter.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Adapter *AdapterCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Adapter *AdapterSession) Endpoint() (common.Address, error) {
	return _Adapter.Contract.Endpoint(&_Adapter.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_Adapter *AdapterCallerSession) Endpoint() (common.Address, error) {
	return _Adapter.Contract.Endpoint(&_Adapter.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Adapter *AdapterCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Adapter *AdapterSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Adapter.Contract.EnforcedOptions(&_Adapter.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_Adapter *AdapterCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _Adapter.Contract.EnforcedOptions(&_Adapter.CallOpts, eid, msgType)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Adapter *AdapterCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Adapter *AdapterSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Adapter.Contract.IsComposeMsgSender(&_Adapter.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_Adapter *AdapterCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _Adapter.Contract.IsComposeMsgSender(&_Adapter.CallOpts, arg0, arg1, _sender)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Adapter *AdapterCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Adapter *AdapterSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Adapter.Contract.IsPeer(&_Adapter.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_Adapter *AdapterCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _Adapter.Contract.IsPeer(&_Adapter.CallOpts, _eid, _peer)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Adapter *AdapterCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Adapter *AdapterSession) MsgInspector() (common.Address, error) {
	return _Adapter.Contract.MsgInspector(&_Adapter.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_Adapter *AdapterCallerSession) MsgInspector() (common.Address, error) {
	return _Adapter.Contract.MsgInspector(&_Adapter.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Adapter *AdapterCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Adapter *AdapterSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Adapter.Contract.NextNonce(&_Adapter.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_Adapter *AdapterCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _Adapter.Contract.NextNonce(&_Adapter.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Adapter *AdapterCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Adapter *AdapterSession) OApp() (common.Address, error) {
	return _Adapter.Contract.OApp(&_Adapter.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_Adapter *AdapterCallerSession) OApp() (common.Address, error) {
	return _Adapter.Contract.OApp(&_Adapter.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Adapter *AdapterCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Adapter *AdapterSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Adapter.Contract.OAppVersion(&_Adapter.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_Adapter *AdapterCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _Adapter.Contract.OAppVersion(&_Adapter.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Adapter *AdapterCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Adapter *AdapterSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Adapter.Contract.OftVersion(&_Adapter.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_Adapter *AdapterCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _Adapter.Contract.OftVersion(&_Adapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adapter *AdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adapter *AdapterSession) Owner() (common.Address, error) {
	return _Adapter.Contract.Owner(&_Adapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Adapter *AdapterCallerSession) Owner() (common.Address, error) {
	return _Adapter.Contract.Owner(&_Adapter.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Adapter *AdapterCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Adapter *AdapterSession) Peers(eid uint32) ([32]byte, error) {
	return _Adapter.Contract.Peers(&_Adapter.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_Adapter *AdapterCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _Adapter.Contract.Peers(&_Adapter.CallOpts, eid)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Adapter *AdapterCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Adapter *AdapterSession) PreCrime() (common.Address, error) {
	return _Adapter.Contract.PreCrime(&_Adapter.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_Adapter *AdapterCallerSession) PreCrime() (common.Address, error) {
	return _Adapter.Contract.PreCrime(&_Adapter.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "quoteOFT", _sendParam)

	outstruct := new(struct {
		OftLimit      OFTLimit
		OftFeeDetails []OFTFeeDetail
		OftReceipt    OFTReceipt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OftLimit = *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	outstruct.OftFeeDetails = *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	outstruct.OftReceipt = *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return *outstruct, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Adapter.Contract.QuoteOFT(&_Adapter.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _Adapter.Contract.QuoteOFT(&_Adapter.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Adapter *AdapterCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Adapter *AdapterSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Adapter.Contract.QuoteSend(&_Adapter.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_Adapter *AdapterCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _Adapter.Contract.QuoteSend(&_Adapter.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Adapter *AdapterCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Adapter *AdapterSession) SharedDecimals() (uint8, error) {
	return _Adapter.Contract.SharedDecimals(&_Adapter.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_Adapter *AdapterCallerSession) SharedDecimals() (uint8, error) {
	return _Adapter.Contract.SharedDecimals(&_Adapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Adapter *AdapterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Adapter.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Adapter *AdapterSession) Token() (common.Address, error) {
	return _Adapter.Contract.Token(&_Adapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Adapter *AdapterCallerSession) Token() (common.Address, error) {
	return _Adapter.Contract.Token(&_Adapter.CallOpts)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceive(&_Adapter.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceive(&_Adapter.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Adapter *AdapterTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Adapter *AdapterSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceiveAndRevert(&_Adapter.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_Adapter *AdapterTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceiveAndRevert(&_Adapter.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceiveSimulate(&_Adapter.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_Adapter *AdapterTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Adapter.Contract.LzReceiveSimulate(&_Adapter.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adapter *AdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adapter *AdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Adapter.Contract.RenounceOwnership(&_Adapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Adapter *AdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Adapter.Contract.RenounceOwnership(&_Adapter.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.Send(&_Adapter.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_Adapter *AdapterTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.Send(&_Adapter.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Adapter *AdapterTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Adapter *AdapterSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetDelegate(&_Adapter.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Adapter *AdapterTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetDelegate(&_Adapter.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Adapter *AdapterTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Adapter *AdapterSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Adapter.Contract.SetEnforcedOptions(&_Adapter.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_Adapter *AdapterTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _Adapter.Contract.SetEnforcedOptions(&_Adapter.TransactOpts, _enforcedOptions)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Adapter *AdapterTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Adapter *AdapterSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetMsgInspector(&_Adapter.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_Adapter *AdapterTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetMsgInspector(&_Adapter.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Adapter *AdapterTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Adapter *AdapterSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Adapter.Contract.SetPeer(&_Adapter.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_Adapter *AdapterTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _Adapter.Contract.SetPeer(&_Adapter.TransactOpts, _eid, _peer)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Adapter *AdapterTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Adapter *AdapterSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetPreCrime(&_Adapter.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_Adapter *AdapterTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.SetPreCrime(&_Adapter.TransactOpts, _preCrime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adapter *AdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Adapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adapter *AdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.TransferOwnership(&_Adapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Adapter *AdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Adapter.Contract.TransferOwnership(&_Adapter.TransactOpts, newOwner)
}

// AdapterEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the Adapter contract.
type AdapterEnforcedOptionSetIterator struct {
	Event *AdapterEnforcedOptionSet // Event containing the contract specifics and raw log

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
func (it *AdapterEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterEnforcedOptionSet)
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
		it.Event = new(AdapterEnforcedOptionSet)
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
func (it *AdapterEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterEnforcedOptionSet represents a EnforcedOptionSet event raised by the Adapter contract.
type AdapterEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Adapter *AdapterFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*AdapterEnforcedOptionSetIterator, error) {

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &AdapterEnforcedOptionSetIterator{contract: _Adapter.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Adapter *AdapterFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *AdapterEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterEnforcedOptionSet)
				if err := _Adapter.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
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

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_Adapter *AdapterFilterer) ParseEnforcedOptionSet(log types.Log) (*AdapterEnforcedOptionSet, error) {
	event := new(AdapterEnforcedOptionSet)
	if err := _Adapter.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the Adapter contract.
type AdapterMsgInspectorSetIterator struct {
	Event *AdapterMsgInspectorSet // Event containing the contract specifics and raw log

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
func (it *AdapterMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterMsgInspectorSet)
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
		it.Event = new(AdapterMsgInspectorSet)
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
func (it *AdapterMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterMsgInspectorSet represents a MsgInspectorSet event raised by the Adapter contract.
type AdapterMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Adapter *AdapterFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*AdapterMsgInspectorSetIterator, error) {

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &AdapterMsgInspectorSetIterator{contract: _Adapter.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Adapter *AdapterFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *AdapterMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterMsgInspectorSet)
				if err := _Adapter.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
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

// ParseMsgInspectorSet is a log parse operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_Adapter *AdapterFilterer) ParseMsgInspectorSet(log types.Log) (*AdapterMsgInspectorSet, error) {
	event := new(AdapterMsgInspectorSet)
	if err := _Adapter.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the Adapter contract.
type AdapterOFTReceivedIterator struct {
	Event *AdapterOFTReceived // Event containing the contract specifics and raw log

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
func (it *AdapterOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterOFTReceived)
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
		it.Event = new(AdapterOFTReceived)
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
func (it *AdapterOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterOFTReceived represents a OFTReceived event raised by the Adapter contract.
type AdapterOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*AdapterOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &AdapterOFTReceivedIterator{contract: _Adapter.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *AdapterOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterOFTReceived)
				if err := _Adapter.contract.UnpackLog(event, "OFTReceived", log); err != nil {
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

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) ParseOFTReceived(log types.Log) (*AdapterOFTReceived, error) {
	event := new(AdapterOFTReceived)
	if err := _Adapter.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the Adapter contract.
type AdapterOFTSentIterator struct {
	Event *AdapterOFTSent // Event containing the contract specifics and raw log

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
func (it *AdapterOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterOFTSent)
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
		it.Event = new(AdapterOFTSent)
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
func (it *AdapterOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterOFTSent represents a OFTSent event raised by the Adapter contract.
type AdapterOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*AdapterOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &AdapterOFTSentIterator{contract: _Adapter.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *AdapterOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterOFTSent)
				if err := _Adapter.contract.UnpackLog(event, "OFTSent", log); err != nil {
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

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_Adapter *AdapterFilterer) ParseOFTSent(log types.Log) (*AdapterOFTSent, error) {
	event := new(AdapterOFTSent)
	if err := _Adapter.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Adapter contract.
type AdapterOwnershipTransferredIterator struct {
	Event *AdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterOwnershipTransferred)
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
		it.Event = new(AdapterOwnershipTransferred)
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
func (it *AdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterOwnershipTransferred represents a OwnershipTransferred event raised by the Adapter contract.
type AdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Adapter *AdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdapterOwnershipTransferredIterator{contract: _Adapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Adapter *AdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterOwnershipTransferred)
				if err := _Adapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Adapter *AdapterFilterer) ParseOwnershipTransferred(log types.Log) (*AdapterOwnershipTransferred, error) {
	event := new(AdapterOwnershipTransferred)
	if err := _Adapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the Adapter contract.
type AdapterPeerSetIterator struct {
	Event *AdapterPeerSet // Event containing the contract specifics and raw log

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
func (it *AdapterPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterPeerSet)
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
		it.Event = new(AdapterPeerSet)
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
func (it *AdapterPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterPeerSet represents a PeerSet event raised by the Adapter contract.
type AdapterPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Adapter *AdapterFilterer) FilterPeerSet(opts *bind.FilterOpts) (*AdapterPeerSetIterator, error) {

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &AdapterPeerSetIterator{contract: _Adapter.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Adapter *AdapterFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *AdapterPeerSet) (event.Subscription, error) {

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterPeerSet)
				if err := _Adapter.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_Adapter *AdapterFilterer) ParsePeerSet(log types.Log) (*AdapterPeerSet, error) {
	event := new(AdapterPeerSet)
	if err := _Adapter.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdapterPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the Adapter contract.
type AdapterPreCrimeSetIterator struct {
	Event *AdapterPreCrimeSet // Event containing the contract specifics and raw log

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
func (it *AdapterPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterPreCrimeSet)
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
		it.Event = new(AdapterPreCrimeSet)
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
func (it *AdapterPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterPreCrimeSet represents a PreCrimeSet event raised by the Adapter contract.
type AdapterPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Adapter *AdapterFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*AdapterPreCrimeSetIterator, error) {

	logs, sub, err := _Adapter.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &AdapterPreCrimeSetIterator{contract: _Adapter.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Adapter *AdapterFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *AdapterPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _Adapter.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterPreCrimeSet)
				if err := _Adapter.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
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

// ParsePreCrimeSet is a log parse operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_Adapter *AdapterFilterer) ParsePreCrimeSet(log types.Log) (*AdapterPreCrimeSet, error) {
	event := new(AdapterPreCrimeSet)
	if err := _Adapter.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
