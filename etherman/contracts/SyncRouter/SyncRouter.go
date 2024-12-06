// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SyncRouter

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

// SyncRouterMetaData contains all meta data concerning the SyncRouter contract.
var SyncRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vizingPad\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LandingPadAccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_SEND_MODE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LandingPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LaunchPad\",\"outputs\":[{\"internalType\":\"contractIMessageChannel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"additionParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultBridgeMode\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGasPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGaslimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"destChainid\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"additionParams\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"estimateVizingGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vizingGasFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"}],\"name\":\"fetchOmniMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minArrivalTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGasLimit\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"mirrorEntryPoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveExtraMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcContract\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"receiveStandardMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedRelayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"destContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainUsedFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"batchsMessage\",\"type\":\"bytes\"}],\"name\":\"sendOmniMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"entryPoint\",\"type\":\"address\"}],\"name\":\"setMirrorEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610100604052600160f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660e0907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505f67ffffffffffffffff81111562000071576200007062000333565b5b6040519080825280601f01601f191660200182016040528015620000a45781602001600182028036833780820191505090505b5060049081620000b5919062000597565b5061c35060055f6101000a81548162ffffff021916908362ffffff160217905550633b9aca00600560036101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503480156200010f575f80fd5b5060405162002a9038038062002a908339818101604052810190620001359190620006e0565b338280816200014a81620001eb60201b60201c565b506200015c816200022d60201b60201c565b50505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620001d1575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620001c8919062000736565b60405180910390fd5b620001e2816200027060201b60201c565b50505062000751565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620003af57607f821691505b602082108103620003c557620003c46200036a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004297fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003ec565b620004358683620003ec565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200047f6200047962000473846200044d565b62000456565b6200044d565b9050919050565b5f819050919050565b6200049a836200045f565b620004b2620004a98262000486565b848454620003f8565b825550505050565b5f90565b620004c8620004ba565b620004d58184846200048f565b505050565b5b81811015620004fc57620004f05f82620004be565b600181019050620004db565b5050565b601f8211156200054b576200051581620003cb565b6200052084620003dd565b8101602085101562000530578190505b620005486200053f85620003dd565b830182620004da565b50505b505050565b5f82821c905092915050565b5f6200056d5f198460080262000550565b1980831691505092915050565b5f6200058783836200055c565b9150826002028217905092915050565b620005a28262000360565b67ffffffffffffffff811115620005be57620005bd62000333565b5b620005ca825462000397565b620005d782828562000500565b5f60209050601f8311600181146200060d575f8415620005f8578287015190505b6200060485826200057a565b86555062000673565b601f1984166200061d86620003cb565b5f5b8281101562000646578489015182556001820191506020850194506020810190506200061f565b8683101562000666578489015162000662601f8916826200055c565b8355505b6001600288020188555050505b505050505050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620006aa826200067f565b9050919050565b620006bc816200069e565b8114620006c7575f80fd5b50565b5f81519050620006da81620006b1565b92915050565b5f8060408385031215620006f957620006f86200067b565b5b5f6200070885828601620006ca565b92505060206200071b85828601620006ca565b9150509250929050565b62000730816200069e565b82525050565b5f6020820190506200074b5f83018462000725565b92915050565b60805160a05160c05160e0516122e8620007a85f395f818161058601528181610684015261091001525f818161077a0152610afa01525f81816107590152610b8d01525f8181610738015261088c01526122e85ff3fe608060405260043610610133575f3560e01c8063715018a6116100aa578063d5f3e0081161006e578063d5f3e008146103dd578063de8aeda014610405578063e0b838e91461042f578063e7b4294c14610459578063f2fde38b14610483578063f3148925146104ab57610133565b8063715018a61461030d57806376c81312146103235780638da5cb5b1461035f57806399e581fa14610389578063b0cfd4d2146103b357610133565b80634d297fd3116100fc5780634d297fd3146101ed5780635ad3ad06146102175780635aeb4d77146102415780635bc06ce71461026b5780635e45da23146102a7578063711abb75146102d157610133565b806273b555146101375780630e82845d14610153578063129e72961461017d5780633db063a7146101a757806345636279146101c3575b5f80fd5b610151600480360381019061014c91906111e8565b6104c7565b005b34801561015e575f80fd5b5061016761055f565b60405161017491906112d3565b60405180910390f35b348015610188575f80fd5b50610191610584565b60405161019e9190611326565b60405180910390f35b6101c160048036038101906101bc91906114b2565b6105a8565b005b3480156101ce575f80fd5b506101d76107fa565b6040516101e49190611326565b60405180910390f35b3480156101f8575f80fd5b506102016107fe565b60405161020e91906115ac565b60405180910390f35b348015610222575f80fd5b5061022b61088a565b60405161023891906115db565b60405180910390f35b34801561024c575f80fd5b506102556108ae565b6040516102629190611611565b60405180910390f35b348015610276575f80fd5b50610291600480360381019061028c919061162a565b6108b2565b60405161029e9190611664565b60405180910390f35b3480156102b2575f80fd5b506102bb6108e2565b6040516102c89190611611565b60405180910390f35b3480156102dc575f80fd5b506102f760048036038101906102f291906114b2565b6108e6565b604051610304919061168c565b60405180910390f35b348015610318575f80fd5b50610321610a0a565b005b34801561032e575f80fd5b50610349600480360381019061034491906116a5565b610a1d565b604051610356919061168c565b60405180910390f35b34801561036a575f80fd5b50610373610abc565b6040516103809190611664565b60405180910390f35b348015610394575f80fd5b5061039d610ae4565b6040516103aa9190611611565b60405180910390f35b3480156103be575f80fd5b506103c7610af8565b6040516103d49190611664565b60405180910390f35b3480156103e8575f80fd5b5061040360048036038101906103fe9190611748565b610b1c565b005b348015610410575f80fd5b50610419610b8b565b60405161042691906115db565b60405180910390f35b34801561043a575f80fd5b50610443610baf565b60405161045091906112d3565b60405180910390f35b348015610464575f80fd5b5061046d610bd2565b60405161047a91906115db565b60405180910390f35b34801561048e575f80fd5b506104a960048036038101906104a49190611786565b610bec565b005b6104c560048036038101906104c091906117e4565b610c70565b005b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461054d576040517ffb2541ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61055984848484610d0a565b50505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b4660035f8267ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610652906118c2565b60405180910390fd5b5f8260405160200161066d91906115ac565b60405160208183030381529060405290505f6106d27f00000000000000000000000000000000000000000000000000000000000000008760055f9054906101000a900462ffffff16600560039054906101000a900467ffffffffffffffff1686610d3c565b90505f6106e1888888886108e6565b905085816106ef919061190d565b3410156106fa575f80fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663209afe56347f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000338c8f60048b6040518a63ffffffff1660e01b81526004016107c2989796959493929190611a30565b5f604051808303818588803b1580156107d9575f80fd5b505af11580156107eb573d5f803e3d5ffd5b50505050505050505050505050565b5f90565b6004805461080b9061196d565b80601f01602080910402602001604051908101604052809291908181526020018280546108379061196d565b80156108825780601f1061085957610100808354040283529160200191610882565b820191905f5260205f20905b81548152906001019060200180831161086557829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f90565b6003602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f90565b5f80826040516020016108f991906115ac565b60405160208183030381529060405290505f61095e7f00000000000000000000000000000000000000000000000000000000000000008760055f9054906101000a900462ffffff16600560039054906101000a900467ffffffffffffffff1686610d3c565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd54286896004856040518563ffffffff1660e01b81526004016109bf9493929190611aba565b602060405180830381865afa1580156109da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109fe9190611b1f565b92505050949350505050565b610a12610d87565b610a1b5f610e0e565b565b5f610ab0878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505086868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050610ed1565b90509695505050505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055f9054906101000a900462ffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b610b24610d87565b8060035f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560039054906101000a900467ffffffffffffffff1681565b610bf4610d87565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c64575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610c5b9190611664565b60405180910390fd5b610c6d81610e0e565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cf6576040517ffb2541ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d038585858585610f7a565b5050505050565b6040517fff94113200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060858573ffffffffffffffffffffffffffffffffffffffff16858585604051602001610d6d959493929190611c2c565b604051602081830303815290604052905095945050505050565b610d8f6110ff565b73ffffffffffffffffffffffffffffffffffffffff16610dad610abc565b73ffffffffffffffffffffffffffffffffffffffff1614610e0c57610dd06110ff565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610e039190611664565b60405180910390fd5b565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166385fdd542868686866040518563ffffffff1660e01b8152600401610f319493929190611c86565b602060405180830381865afa158015610f4c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f709190611b1f565b9050949350505050565b8273ffffffffffffffffffffffffffffffffffffffff1660035f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611023576040517f5cb045db00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82828101906110339190611cd7565b90505f8180602001905181019061104a919061200b565b905060035f4667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166384b2b1b1826040518263ffffffff1660e01b81526004016110c99190612292565b5f604051808303815f87803b1580156110e0575f80fd5b505af11580156110f2573d5f803e3d5ffd5b5050505050505050505050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61113381611117565b811461113d575f80fd5b50565b5f8135905061114e8161112a565b92915050565b5f819050919050565b61116681611154565b8114611170575f80fd5b50565b5f813590506111818161115d565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126111a8576111a7611187565b5b8235905067ffffffffffffffff8111156111c5576111c461118b565b5b6020830191508360018202830111156111e1576111e061118f565b5b9250929050565b5f805f8060608587031215611200576111ff61110f565b5b5f61120d87828801611140565b945050602061121e87828801611173565b935050604085013567ffffffffffffffff81111561123f5761123e611113565b5b61124b87828801611193565b925092505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61129b61129661129184611259565b611278565b611259565b9050919050565b5f6112ac82611281565b9050919050565b5f6112bd826112a2565b9050919050565b6112cd816112b3565b82525050565b5f6020820190506112e65f8301846112c4565b92915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b611320816112ec565b82525050565b5f6020820190506113395f830184611317565b92915050565b5f61134982611259565b9050919050565b6113598161133f565b8114611363575f80fd5b50565b5f8135905061137481611350565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113c48261137e565b810181811067ffffffffffffffff821117156113e3576113e261138e565b5b80604052505050565b5f6113f5611106565b905061140182826113bb565b919050565b5f67ffffffffffffffff8211156114205761141f61138e565b5b6114298261137e565b9050602081019050919050565b828183375f83830152505050565b5f61145661145184611406565b6113ec565b9050828152602081018484840111156114725761147161137a565b5b61147d848285611436565b509392505050565b5f82601f83011261149957611498611187565b5b81356114a9848260208601611444565b91505092915050565b5f805f80608085870312156114ca576114c961110f565b5b5f6114d787828801611140565b94505060206114e887828801611366565b93505060406114f987828801611173565b925050606085013567ffffffffffffffff81111561151a57611519611113565b5b61152687828801611485565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561156957808201518184015260208101905061154e565b5f8484015250505050565b5f61157e82611532565b611588818561153c565b935061159881856020860161154c565b6115a18161137e565b840191505092915050565b5f6020820190508181035f8301526115c48184611574565b905092915050565b6115d581611117565b82525050565b5f6020820190506115ee5f8301846115cc565b92915050565b5f62ffffff82169050919050565b61160b816115f4565b82525050565b5f6020820190506116245f830184611602565b92915050565b5f6020828403121561163f5761163e61110f565b5b5f61164c84828501611140565b91505092915050565b61165e8161133f565b82525050565b5f6020820190506116775f830184611655565b92915050565b61168681611154565b82525050565b5f60208201905061169f5f83018461167d565b92915050565b5f805f805f80608087890312156116bf576116be61110f565b5b5f6116cc89828a01611173565b96505060206116dd89828a01611140565b955050604087013567ffffffffffffffff8111156116fe576116fd611113565b5b61170a89828a01611193565b9450945050606087013567ffffffffffffffff81111561172d5761172c611113565b5b61173989828a01611193565b92509250509295509295509295565b5f806040838503121561175e5761175d61110f565b5b5f61176b85828601611140565b925050602061177c85828601611366565b9150509250929050565b5f6020828403121561179b5761179a61110f565b5b5f6117a884828501611366565b91505092915050565b5f819050919050565b6117c3816117b1565b81146117cd575f80fd5b50565b5f813590506117de816117ba565b92915050565b5f805f805f608086880312156117fd576117fc61110f565b5b5f61180a888289016117d0565b955050602061181b88828901611140565b945050604061182c88828901611173565b935050606086013567ffffffffffffffff81111561184d5761184c611113565b5b61185988828901611193565b92509250509295509295909350565b5f82825260208201905092915050565b7f4d455000000000000000000000000000000000000000000000000000000000005f82015250565b5f6118ac600383611868565b91506118b782611878565b602082019050919050565b5f6020820190508181035f8301526118d9816118a0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61191782611154565b915061192283611154565b925082820190508082111561193a576119396118e0565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061198457607f821691505b60208210810361199757611996611940565b5b50919050565b5f819050815f5260205f209050919050565b5f81546119bb8161196d565b6119c5818661153c565b9450600182165f81146119df57600181146119f557611a27565b60ff198316865281151560200286019350611a27565b6119fe8561199d565b5f5b83811015611a1f57815481890152600182019150602081019050611a00565b808801955050505b50505092915050565b5f61010082019050611a445f83018b6115cc565b611a51602083018a6115cc565b611a5e6040830189611655565b611a6b6060830188611655565b611a78608083018761167d565b611a8560a08301866115cc565b81810360c0830152611a9781856119af565b905081810360e0830152611aab8184611574565b90509998505050505050505050565b5f608082019050611acd5f83018761167d565b611ada60208301866115cc565b8181036040830152611aec81856119af565b90508181036060830152611b008184611574565b905095945050505050565b5f81519050611b198161115d565b92915050565b5f60208284031215611b3457611b3361110f565b5b5f611b4184828501611b0b565b91505092915050565b5f819050919050565b611b64611b5f826112ec565b611b4a565b82525050565b5f819050919050565b611b84611b7f82611154565b611b6a565b82525050565b5f8160e81b9050919050565b5f611ba082611b8a565b9050919050565b611bb8611bb3826115f4565b611b96565b82525050565b5f8160c01b9050919050565b5f611bd482611bbe565b9050919050565b611bec611be782611117565b611bca565b82525050565b5f81905092915050565b5f611c0682611532565b611c108185611bf2565b9350611c2081856020860161154c565b80840191505092915050565b5f611c378288611b53565b600182019150611c478287611b73565b602082019150611c578286611ba7565b600382019150611c678285611bdb565b600882019150611c778284611bfc565b91508190509695505050505050565b5f608082019050611c995f83018761167d565b611ca660208301866115cc565b8181036040830152611cb88185611574565b90508181036060830152611ccc8184611574565b905095945050505050565b5f60208284031215611cec57611ceb61110f565b5b5f82013567ffffffffffffffff811115611d0957611d08611113565b5b611d1584828501611485565b91505092915050565b5f67ffffffffffffffff821115611d3857611d3761138e565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f60ff82169050919050565b611d6681611d51565b8114611d70575f80fd5b50565b5f81519050611d8181611d5d565b92915050565b5f81519050611d9581611350565b92915050565b5f81519050611da98161112a565b92915050565b5f611dc1611dbc84611406565b6113ec565b905082815260208101848484011115611ddd57611ddc61137a565b5b611de884828561154c565b509392505050565b5f82601f830112611e0457611e03611187565b5b8151611e14848260208601611daf565b91505092915050565b5f6101808284031215611e3357611e32611d49565b5b611e3e6101806113ec565b90505f611e4d84828501611d73565b5f830152506020611e6084828501611b0b565b6020830152506040611e7484828501611d87565b6040830152506060611e8884828501611d9b565b6060830152506080611e9c84828501611d9b565b60808301525060a082015167ffffffffffffffff811115611ec057611ebf611d4d565b5b611ecc84828501611df0565b60a08301525060c0611ee084828501611b0b565b60c08301525060e0611ef484828501611b0b565b60e083015250610100611f0984828501611b0b565b61010083015250610120611f1f84828501611b0b565b61012083015250610140611f3584828501611b0b565b61014083015250610160611f4b84828501611d87565b6101608301525092915050565b5f611f6a611f6584611d1e565b6113ec565b90508083825260208201905060208402830185811115611f8d57611f8c61118f565b5b835b81811015611fd457805167ffffffffffffffff811115611fb257611fb1611187565b5b808601611fbf8982611e1d565b85526020850194505050602081019050611f8f565b5050509392505050565b5f82601f830112611ff257611ff1611187565b5b8151612002848260208601611f58565b91505092915050565b5f602082840312156120205761201f61110f565b5b5f82015167ffffffffffffffff81111561203d5761203c611113565b5b61204984828501611fde565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61208481611d51565b82525050565b61209381611154565b82525050565b6120a28161133f565b82525050565b6120b181611117565b82525050565b5f82825260208201905092915050565b5f6120d182611532565b6120db81856120b7565b93506120eb81856020860161154c565b6120f48161137e565b840191505092915050565b5f61018083015f8301516121155f86018261207b565b506020830151612128602086018261208a565b50604083015161213b6040860182612099565b50606083015161214e60608601826120a8565b50608083015161216160808601826120a8565b5060a083015184820360a086015261217982826120c7565b91505060c083015161218e60c086018261208a565b5060e08301516121a160e086018261208a565b506101008301516121b661010086018261208a565b506101208301516121cb61012086018261208a565b506101408301516121e061014086018261208a565b506101608301516121f5610160860182612099565b508091505092915050565b5f61220b83836120ff565b905092915050565b5f602082019050919050565b5f61222982612052565b612233818561205c565b9350836020820285016122458561206c565b805f5b8581101561228057848403895281516122618582612200565b945061226c83612213565b925060208a01995050600181019050612248565b50829750879550505050505092915050565b5f6020820190508181035f8301526122aa818461221f565b90509291505056fea26469706673582212201b95d51590f80db0073672dbad4ec2b7ca479e6563a2ba112484dfe8f7d3662064736f6c63430008180033",
}

// SyncRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SyncRouterMetaData.ABI instead.
var SyncRouterABI = SyncRouterMetaData.ABI

// SyncRouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SyncRouterMetaData.Bin instead.
var SyncRouterBin = SyncRouterMetaData.Bin

// DeploySyncRouter deploys a new Ethereum contract, binding an instance of SyncRouter to it.
func DeploySyncRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _vizingPad common.Address, _owner common.Address) (common.Address, *types.Transaction, *SyncRouter, error) {
	parsed, err := SyncRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SyncRouterBin), backend, _vizingPad, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SyncRouter{SyncRouterCaller: SyncRouterCaller{contract: contract}, SyncRouterTransactor: SyncRouterTransactor{contract: contract}, SyncRouterFilterer: SyncRouterFilterer{contract: contract}}, nil
}

// SyncRouter is an auto generated Go binding around an Ethereum contract.
type SyncRouter struct {
	SyncRouterCaller     // Read-only binding to the contract
	SyncRouterTransactor // Write-only binding to the contract
	SyncRouterFilterer   // Log filterer for contract events
}

// SyncRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyncRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyncRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyncRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyncRouterSession struct {
	Contract     *SyncRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SyncRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyncRouterCallerSession struct {
	Contract *SyncRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SyncRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyncRouterTransactorSession struct {
	Contract     *SyncRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SyncRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyncRouterRaw struct {
	Contract *SyncRouter // Generic contract binding to access the raw methods on
}

// SyncRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyncRouterCallerRaw struct {
	Contract *SyncRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SyncRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyncRouterTransactorRaw struct {
	Contract *SyncRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyncRouter creates a new instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouter(address common.Address, backend bind.ContractBackend) (*SyncRouter, error) {
	contract, err := bindSyncRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyncRouter{SyncRouterCaller: SyncRouterCaller{contract: contract}, SyncRouterTransactor: SyncRouterTransactor{contract: contract}, SyncRouterFilterer: SyncRouterFilterer{contract: contract}}, nil
}

// NewSyncRouterCaller creates a new read-only instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterCaller(address common.Address, caller bind.ContractCaller) (*SyncRouterCaller, error) {
	contract, err := bindSyncRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyncRouterCaller{contract: contract}, nil
}

// NewSyncRouterTransactor creates a new write-only instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SyncRouterTransactor, error) {
	contract, err := bindSyncRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyncRouterTransactor{contract: contract}, nil
}

// NewSyncRouterFilterer creates a new log filterer instance of SyncRouter, bound to a specific deployed contract.
func NewSyncRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SyncRouterFilterer, error) {
	contract, err := bindSyncRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyncRouterFilterer{contract: contract}, nil
}

// bindSyncRouter binds a generic wrapper to an already deployed contract.
func bindSyncRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SyncRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncRouter *SyncRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncRouter.Contract.SyncRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncRouter *SyncRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.Contract.SyncRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncRouter *SyncRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncRouter.Contract.SyncRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncRouter *SyncRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncRouter *SyncRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncRouter *SyncRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncRouter.Contract.contract.Transact(opts, method, params...)
}

// BRIDGESENDMODE is a free data retrieval call binding the contract method 0x129e7296.
//
// Solidity: function BRIDGE_SEND_MODE() view returns(bytes1)
func (_SyncRouter *SyncRouterCaller) BRIDGESENDMODE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "BRIDGE_SEND_MODE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// BRIDGESENDMODE is a free data retrieval call binding the contract method 0x129e7296.
//
// Solidity: function BRIDGE_SEND_MODE() view returns(bytes1)
func (_SyncRouter *SyncRouterSession) BRIDGESENDMODE() ([1]byte, error) {
	return _SyncRouter.Contract.BRIDGESENDMODE(&_SyncRouter.CallOpts)
}

// BRIDGESENDMODE is a free data retrieval call binding the contract method 0x129e7296.
//
// Solidity: function BRIDGE_SEND_MODE() view returns(bytes1)
func (_SyncRouter *SyncRouterCallerSession) BRIDGESENDMODE() ([1]byte, error) {
	return _SyncRouter.Contract.BRIDGESENDMODE(&_SyncRouter.CallOpts)
}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterCaller) LandingPad(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "LandingPad")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterSession) LandingPad() (common.Address, error) {
	return _SyncRouter.Contract.LandingPad(&_SyncRouter.CallOpts)
}

// LandingPad is a free data retrieval call binding the contract method 0x0e82845d.
//
// Solidity: function LandingPad() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) LandingPad() (common.Address, error) {
	return _SyncRouter.Contract.LandingPad(&_SyncRouter.CallOpts)
}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterCaller) LaunchPad(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "LaunchPad")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterSession) LaunchPad() (common.Address, error) {
	return _SyncRouter.Contract.LaunchPad(&_SyncRouter.CallOpts)
}

// LaunchPad is a free data retrieval call binding the contract method 0xe0b838e9.
//
// Solidity: function LaunchPad() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) LaunchPad() (common.Address, error) {
	return _SyncRouter.Contract.LaunchPad(&_SyncRouter.CallOpts)
}

// AdditionParams is a free data retrieval call binding the contract method 0x4d297fd3.
//
// Solidity: function additionParams() view returns(bytes)
func (_SyncRouter *SyncRouterCaller) AdditionParams(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "additionParams")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AdditionParams is a free data retrieval call binding the contract method 0x4d297fd3.
//
// Solidity: function additionParams() view returns(bytes)
func (_SyncRouter *SyncRouterSession) AdditionParams() ([]byte, error) {
	return _SyncRouter.Contract.AdditionParams(&_SyncRouter.CallOpts)
}

// AdditionParams is a free data retrieval call binding the contract method 0x4d297fd3.
//
// Solidity: function additionParams() view returns(bytes)
func (_SyncRouter *SyncRouterCallerSession) AdditionParams() ([]byte, error) {
	return _SyncRouter.Contract.AdditionParams(&_SyncRouter.CallOpts)
}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterCaller) DefaultBridgeMode(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultBridgeMode")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterSession) DefaultBridgeMode() ([1]byte, error) {
	return _SyncRouter.Contract.DefaultBridgeMode(&_SyncRouter.CallOpts)
}

// DefaultBridgeMode is a free data retrieval call binding the contract method 0x45636279.
//
// Solidity: function defaultBridgeMode() view returns(bytes1)
func (_SyncRouter *SyncRouterCallerSession) DefaultBridgeMode() ([1]byte, error) {
	return _SyncRouter.Contract.DefaultBridgeMode(&_SyncRouter.CallOpts)
}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) DefaultGasPrice(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultGasPrice")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterSession) DefaultGasPrice() (uint64, error) {
	return _SyncRouter.Contract.DefaultGasPrice(&_SyncRouter.CallOpts)
}

// DefaultGasPrice is a free data retrieval call binding the contract method 0xe7b4294c.
//
// Solidity: function defaultGasPrice() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) DefaultGasPrice() (uint64, error) {
	return _SyncRouter.Contract.DefaultGasPrice(&_SyncRouter.CallOpts)
}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) DefaultGaslimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "defaultGaslimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) DefaultGaslimit() (*big.Int, error) {
	return _SyncRouter.Contract.DefaultGaslimit(&_SyncRouter.CallOpts)
}

// DefaultGaslimit is a free data retrieval call binding the contract method 0x99e581fa.
//
// Solidity: function defaultGaslimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) DefaultGaslimit() (*big.Int, error) {
	return _SyncRouter.Contract.DefaultGaslimit(&_SyncRouter.CallOpts)
}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterCaller) EstimateVizingGasFee(opts *bind.CallOpts, value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "estimateVizingGasFee", value, destChainid, additionParams, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterSession) EstimateVizingGasFee(value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	return _SyncRouter.Contract.EstimateVizingGasFee(&_SyncRouter.CallOpts, value, destChainid, additionParams, message)
}

// EstimateVizingGasFee is a free data retrieval call binding the contract method 0x76c81312.
//
// Solidity: function estimateVizingGasFee(uint256 value, uint64 destChainid, bytes additionParams, bytes message) view returns(uint256 vizingGasFee)
func (_SyncRouter *SyncRouterCallerSession) EstimateVizingGasFee(value *big.Int, destChainid uint64, additionParams []byte, message []byte) (*big.Int, error) {
	return _SyncRouter.Contract.EstimateVizingGasFee(&_SyncRouter.CallOpts, value, destChainid, additionParams, message)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x711abb75.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) view returns(uint256)
func (_SyncRouter *SyncRouterCaller) FetchOmniMessageFee(opts *bind.CallOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "fetchOmniMessageFee", destChainId, destContract, destChainUsedFee, batchsMessage)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x711abb75.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) view returns(uint256)
func (_SyncRouter *SyncRouterSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, batchsMessage)
}

// FetchOmniMessageFee is a free data retrieval call binding the contract method 0x711abb75.
//
// Solidity: function fetchOmniMessageFee(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) view returns(uint256)
func (_SyncRouter *SyncRouterCallerSession) FetchOmniMessageFee(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*big.Int, error) {
	return _SyncRouter.Contract.FetchOmniMessageFee(&_SyncRouter.CallOpts, destChainId, destContract, destChainUsedFee, batchsMessage)
}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) MaxArrivalTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "maxArrivalTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterSession) MaxArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MaxArrivalTime(&_SyncRouter.CallOpts)
}

// MaxArrivalTime is a free data retrieval call binding the contract method 0xde8aeda0.
//
// Solidity: function maxArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) MaxArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MaxArrivalTime(&_SyncRouter.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) MaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "maxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) MaxGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MaxGasLimit(&_SyncRouter.CallOpts)
}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) MaxGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MaxGasLimit(&_SyncRouter.CallOpts)
}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCaller) MinArrivalTime(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "minArrivalTime")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterSession) MinArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MinArrivalTime(&_SyncRouter.CallOpts)
}

// MinArrivalTime is a free data retrieval call binding the contract method 0x5ad3ad06.
//
// Solidity: function minArrivalTime() view returns(uint64)
func (_SyncRouter *SyncRouterCallerSession) MinArrivalTime() (uint64, error) {
	return _SyncRouter.Contract.MinArrivalTime(&_SyncRouter.CallOpts)
}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCaller) MinGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "minGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterSession) MinGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MinGasLimit(&_SyncRouter.CallOpts)
}

// MinGasLimit is a free data retrieval call binding the contract method 0x5aeb4d77.
//
// Solidity: function minGasLimit() view returns(uint24)
func (_SyncRouter *SyncRouterCallerSession) MinGasLimit() (*big.Int, error) {
	return _SyncRouter.Contract.MinGasLimit(&_SyncRouter.CallOpts)
}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x5bc06ce7.
//
// Solidity: function mirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterCaller) MirrorEntryPoint(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "mirrorEntryPoint", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x5bc06ce7.
//
// Solidity: function mirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterSession) MirrorEntryPoint(arg0 uint64) (common.Address, error) {
	return _SyncRouter.Contract.MirrorEntryPoint(&_SyncRouter.CallOpts, arg0)
}

// MirrorEntryPoint is a free data retrieval call binding the contract method 0x5bc06ce7.
//
// Solidity: function mirrorEntryPoint(uint64 ) view returns(address)
func (_SyncRouter *SyncRouterCallerSession) MirrorEntryPoint(arg0 uint64) (common.Address, error) {
	return _SyncRouter.Contract.MirrorEntryPoint(&_SyncRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterSession) Owner() (common.Address, error) {
	return _SyncRouter.Contract.Owner(&_SyncRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) Owner() (common.Address, error) {
	return _SyncRouter.Contract.Owner(&_SyncRouter.CallOpts)
}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterCaller) SelectedRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncRouter.contract.Call(opts, &out, "selectedRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterSession) SelectedRelayer() (common.Address, error) {
	return _SyncRouter.Contract.SelectedRelayer(&_SyncRouter.CallOpts)
}

// SelectedRelayer is a free data retrieval call binding the contract method 0xb0cfd4d2.
//
// Solidity: function selectedRelayer() view returns(address)
func (_SyncRouter *SyncRouterCallerSession) SelectedRelayer() (common.Address, error) {
	return _SyncRouter.Contract.SelectedRelayer(&_SyncRouter.CallOpts)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactor) ReceiveExtraMessage(opts *bind.TransactOpts, messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "receiveExtraMessage", messageId, srcChainId, srcContract, message)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterSession) ReceiveExtraMessage(messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveExtraMessage(&_SyncRouter.TransactOpts, messageId, srcChainId, srcContract, message)
}

// ReceiveExtraMessage is a paid mutator transaction binding the contract method 0xf3148925.
//
// Solidity: function receiveExtraMessage(bytes32 messageId, uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) ReceiveExtraMessage(messageId [32]byte, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveExtraMessage(&_SyncRouter.TransactOpts, messageId, srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactor) ReceiveStandardMessage(opts *bind.TransactOpts, srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "receiveStandardMessage", srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterSession) ReceiveStandardMessage(srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveStandardMessage(&_SyncRouter.TransactOpts, srcChainId, srcContract, message)
}

// ReceiveStandardMessage is a paid mutator transaction binding the contract method 0x0073b555.
//
// Solidity: function receiveStandardMessage(uint64 srcChainId, uint256 srcContract, bytes message) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) ReceiveStandardMessage(srcChainId uint64, srcContract *big.Int, message []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.ReceiveStandardMessage(&_SyncRouter.TransactOpts, srcChainId, srcContract, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyncRouter.Contract.RenounceOwnership(&_SyncRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SyncRouter *SyncRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SyncRouter.Contract.RenounceOwnership(&_SyncRouter.TransactOpts)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x3db063a7.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) payable returns()
func (_SyncRouter *SyncRouterTransactor) SendOmniMessage(opts *bind.TransactOpts, destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "sendOmniMessage", destChainId, destContract, destChainUsedFee, batchsMessage)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x3db063a7.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) payable returns()
func (_SyncRouter *SyncRouterSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainUsedFee, batchsMessage)
}

// SendOmniMessage is a paid mutator transaction binding the contract method 0x3db063a7.
//
// Solidity: function sendOmniMessage(uint64 destChainId, address destContract, uint256 destChainUsedFee, bytes batchsMessage) payable returns()
func (_SyncRouter *SyncRouterTransactorSession) SendOmniMessage(destChainId uint64, destContract common.Address, destChainUsedFee *big.Int, batchsMessage []byte) (*types.Transaction, error) {
	return _SyncRouter.Contract.SendOmniMessage(&_SyncRouter.TransactOpts, destChainId, destContract, destChainUsedFee, batchsMessage)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterTransactor) SetMirrorEntryPoint(opts *bind.TransactOpts, chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "setMirrorEntryPoint", chainId, entryPoint)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterSession) SetMirrorEntryPoint(chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetMirrorEntryPoint(&_SyncRouter.TransactOpts, chainId, entryPoint)
}

// SetMirrorEntryPoint is a paid mutator transaction binding the contract method 0xd5f3e008.
//
// Solidity: function setMirrorEntryPoint(uint64 chainId, address entryPoint) returns()
func (_SyncRouter *SyncRouterTransactorSession) SetMirrorEntryPoint(chainId uint64, entryPoint common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.SetMirrorEntryPoint(&_SyncRouter.TransactOpts, chainId, entryPoint)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.TransferOwnership(&_SyncRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SyncRouter *SyncRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SyncRouter.Contract.TransferOwnership(&_SyncRouter.TransactOpts, newOwner)
}

// SyncRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SyncRouter contract.
type SyncRouterOwnershipTransferredIterator struct {
	Event *SyncRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SyncRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncRouterOwnershipTransferred)
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
		it.Event = new(SyncRouterOwnershipTransferred)
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
func (it *SyncRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncRouterOwnershipTransferred represents a OwnershipTransferred event raised by the SyncRouter contract.
type SyncRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyncRouter *SyncRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SyncRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyncRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SyncRouterOwnershipTransferredIterator{contract: _SyncRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SyncRouter *SyncRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SyncRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SyncRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncRouterOwnershipTransferred)
				if err := _SyncRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SyncRouter *SyncRouterFilterer) ParseOwnershipTransferred(log types.Log) (*SyncRouterOwnershipTransferred, error) {
	event := new(SyncRouterOwnershipTransferred)
	if err := _SyncRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
