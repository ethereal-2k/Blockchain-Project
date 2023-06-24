// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fairswap

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

// FairswapMetaData contains all meta data concerning the Fairswap contract.
var FairswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitmentOfKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiverEntryKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"circuitGatesOperationArray\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"complaint\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Mverify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"circuitGatesOperation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"complaint\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"complain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recEntryKey\",\"type\":\"bytes32\"}],\"name\":\"initializeRecieverAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"enumfairswap.stage\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverEntryKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600a60006101000a81548160ff021916908360068111156200002457fe5b02179055503480156200003657600080fd5b5060405162001e2238038062001e2283398181016040526200005c91908101906200023e565b33600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600181905550836002819055508260008190555060008090505b8151811015620000fc57818181518110620000cf57fe5b602002602001015160066000838152602001908152602001600020819055508080600101915050620000b8565b5081600481905550620001146200011f60201b60201c565b505050505062000378565b6001600a60009054906101000a900460ff1660068111156200013d57fe5b0160068111156200014a57fe5b600a60006101000a81548160ff021916908360068111156200016857fe5b0217905550600a8042816200017957fe5b0401600981905550565b600082601f8301126200019557600080fd5b8151620001ac620001a68262000307565b620002d9565b91508181835260208401935060208101905083856020840282011115620001d257600080fd5b60005b83811015620002065781620001eb888262000227565b845260208401935060208301925050600181019050620001d5565b5050505092915050565b600081519050620002218162000344565b92915050565b60008151905062000238816200035e565b92915050565b600080600080600060a086880312156200025757600080fd5b6000620002678882890162000227565b95505060206200027a8882890162000210565b94505060406200028d8882890162000210565b9350506060620002a08882890162000210565b925050608086015167ffffffffffffffff811115620002be57600080fd5b620002cc8882890162000183565b9150509295509295909350565b6000604051905081810181811067ffffffffffffffff82111715620002fd57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200031f57600080fd5b602082029050602081019050919050565b6000819050919050565b6000819050919050565b6200034f8162000330565b81146200035b57600080fd5b50565b62000369816200033a565b81146200037557600080fd5b50565b611a9a80620003886000396000f3fe6080604052600436106101145760003560e01c80639ab164e5116100a0578063d654723611610064578063d654723614610360578063de30b90714610389578063eaa94830146103c6578063ee3743ab146103f1578063f6fa345f1461040857610114565b80639ab164e514610279578063a035b1fe146102a2578063a2630372146102cd578063b1c9fe6e1461030a578063ca641f861461033557610114565b80633943380c116100e75780633943380c146101b65780633e37aeb6146101e157806344d4fd19146101f85780634fc847911461022357806370dea79a1461024e57610114565b8063062a1398146101195780630b90457d1461014457806316fed3e2146101815780632852b71c146101ac575b600080fd5b34801561012557600080fd5b5061012e61041f565b60405161013b9190611887565b60405180910390f35b34801561015057600080fd5b5061016b600480360361016691908101906117dd565b610425565b60405161017891906118bd565b60405180910390f35b34801561018d57600080fd5b5061019661043d565b6040516101a39190611851565b60405180910390f35b6101b4610463565b005b3480156101c257600080fd5b506101cb610590565b6040516101d89190611887565b60405180910390f35b3480156101ed57600080fd5b506101f6610596565b005b34801561020457600080fd5b5061020d6106e7565b60405161021a91906118bd565b60405180910390f35b34801561022f57600080fd5b506102386106ed565b6040516102459190611851565b60405180910390f35b34801561025a57600080fd5b50610263610713565b60405161027091906118bd565b60405180910390f35b34801561028557600080fd5b506102a0600480360361029b91908101906117b4565b610719565b005b3480156102ae57600080fd5b506102b761076b565b6040516102c491906118bd565b60405180910390f35b3480156102d957600080fd5b506102f460048036036102ef9190810190611760565b610771565b604051610301919061186c565b60405180910390f35b34801561031657600080fd5b5061031f6109ff565b60405161032c91906118a2565b60405180910390f35b34801561034157600080fd5b5061034a610a12565b6040516103579190611887565b60405180910390f35b34801561036c57600080fd5b50610387600480360361038291908101906117b4565b610a18565b005b34801561039557600080fd5b506103b060048036036103ab91908101906116f4565b610c1f565b6040516103bd919061186c565b60405180910390f35b3480156103d257600080fd5b506103db6110da565b6040516103e89190611887565b60405180910390f35b3480156103fd57600080fd5b506104066110e0565b005b34801561041457600080fd5b5061041d611140565b005b60005481565b60066020528060005260406000206000915090505481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600180600681111561049457fe5b600a60009054906101000a900460ff1660068111156104af57fe5b146104b957600080fd5b600954600a42816104c657fe5b04106104d157600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050957600080fd5b600154341015610583573373ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610559573d6000803e3d6000fd5b506006600a60006101000a81548160ff0219169083600681111561057957fe5b021790555061058c565b61058b6110e0565b5b5050565b60035481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660028060068111156105c757fe5b600a60009054906101000a900460ff1660068111156105e257fe5b146105ec57600080fd5b600954600a42816105f957fe5b041061060457600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063c57600080fd5b600954600a428161064957fe5b041161065457600080fd5b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f193505050501580156106be573d6000803e3d6000fd5b506005600a60006101000a81548160ff021916908360068111156106de57fe5b02179055505050565b60055481565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b806004541461072757600080fd5b33600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60015481565b6000808360008151811061078157fe5b602002602001015190506000806000600590506060604080519080825280601f01601f1916602001820160405280156107c95781602001600182028038833980820191505090505b509050600092505b600182038310156109ed57600183846001901b8916901c14156108e95760008090505b602081101561086b5788600185018151811061080c57fe5b6020026020010151816020811061081f57fe5b1a60f81b82828151811061082f57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506107f4565b5060008090505b60208110156108d65785816020811061088757fe5b1a60f81b82602083018151811061089a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610872565b50808051906020012093508394506109e0565b60008090505b60208110156109505785816020811061090457fe5b1a60f81b82828151811061091457fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506108ef565b5060008090505b60208110156109d15788600185018151811061096f57fe5b6020026020010151816020811061098257fe5b1a60f81b82602083018151811061099557fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610957565b50808051906020012093508394505b82806001019350506107d1565b60005485149550505050505092915050565b600a60009054906101000a900460ff1681565b60025481565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002806006811115610a4957fe5b600a60009054906101000a900460ff166006811115610a6457fe5b14610a6e57600080fd5b600954600a4281610a7b57fe5b0410610a8657600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610abe57600080fd5b6000606060206040519080825280601f01601f191660200182016040528015610af65781602001600182028038833980820191505090505b50905060008090505b6020811015610b6057858160208110610b1457fe5b1a60f81b828281518110610b2457fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610aff565b50808051906020012091506002548214610c0857600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015610bde573d6000803e3d6000fd5b506005600a60006101000a81548160ff02191690836006811115610bfe57fe5b0217905550610c18565b84600381905550610c176110e0565b5b5050505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003806006811115610c5257fe5b600a60009054906101000a900460ff166006811115610c6d57fe5b14610c7757600080fd5b600954600a4281610c8457fe5b0410610c8f57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cc757600080fd5b600080606060008088600081518110610cdc57fe5b6020026020010151905060018a5103604051908082528060200260200182016040528015610d195781602001602082028038833980820191505090505b509250610d4e8a600081518110610d2c57fe5b60200260200101518a600081518110610d4157fe5b6020026020010151610771565b610def576006600a60006101000a81548160ff02191690836006811115610d7157fe5b0217905550600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015610de0573d6000803e3d6000fd5b506000975050505050506110d2565b600660008a600081518110610e0057fe5b60200260200101518152602001908152602001600020549350610e4a8a600081518110610e2957fe5b6020026020010151600081518110610e3d57fe5b6020026020010151611275565b915060008090505b60018b5103811015610f8857610e8e8b8281518110610e6d57fe5b60200260200101518b8381518110610e8157fe5b6020026020010151610771565b610f30576006600a60006101000a81548160ff02191690836006811115610eb157fe5b0217905550600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015610f20573d6000803e3d6000fd5b50600098505050505050506110d2565b610f638b6001830181518110610f4257fe5b6020026020010151600081518110610f5657fe5b6020026020010151611275565b848281518110610f6f57fe5b6020026020010181815250508080600101915050610e52565b50610f938484611378565b9450848214611039576005600a60006101000a81548160ff02191690836006811115610fbb57fe5b0217905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f1935050505015801561102a573d6000803e3d6000fd5b506001975050505050506110d2565b6006600a60006101000a81548160ff0219169083600681111561105857fe5b0217905550600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f193505050501580156110c7573d6000803e3d6000fd5b506000975050505050505b505092915050565b60045481565b6001600a60009054906101000a900460ff1660068111156110fd57fe5b01600681111561110957fe5b600a60006101000a81548160ff0219169083600681111561112657fe5b0217905550600a80428161113657fe5b0401600981905550565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600380600681111561117157fe5b600a60009054906101000a900460ff16600681111561118c57fe5b1461119657600080fd5b600954600a42816111a357fe5b04106111ae57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111e657600080fd5b600954600a42816111f357fe5b04116111fe57600080fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015611268573d6000803e3d6000fd5b506112716110e0565b5050565b6000606060008060216040519080825280601f01601f1916602001820160405280156112b05781602001600182028038833980820191505090505b50925060008090505b602081101561131c5760035481602081106112d057fe5b1a60f81b8482815181106112e057fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506112b9565b50600060f81b8360208151811061132f57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350828051906020012091508482189050809350505050919050565b60008060028414156114c3576060604080519080825280601f01601f1916602001820160405280156113b95781602001600182028038833980820191505090505b50905060008090505b602081101561143757846000815181106113d857fe5b602002602001015181602081106113eb57fe5b1a60f81b8282815181106113fb57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506113c2565b5060008090505b60208110156114b6578460018151811061145457fe5b6020026020010151816020811061146757fe5b1a60f81b82602083018151811061147a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061143e565b5080805190602001209150505b8091505092915050565b600082601f8301126114de57600080fd5b81356114f16114ec82611905565b6118d8565b9150818183526020840193506020810190508360005b83811015611537578135860161151d8882611541565b845260208401935060208301925050600181019050611507565b5050505092915050565b600082601f83011261155257600080fd5b81356115656115608261192d565b6118d8565b9150818183526020840193506020810190508385602084028201111561158a57600080fd5b60005b838110156115ba57816115a088826116ca565b84526020840193506020830192505060018101905061158d565b5050505092915050565b600082601f8301126115d557600080fd5b81356115e86115e382611955565b6118d8565b9150818183526020840193506020810190508385602084028201111561160d57600080fd5b60005b8381101561163d578161162388826116ca565b845260208401935060208301925050600181019050611610565b5050505092915050565b600082601f83011261165857600080fd5b813561166b6116668261197d565b6118d8565b9150818183526020840193506020810190508385602084028201111561169057600080fd5b60005b838110156116c057816116a688826116df565b845260208401935060208301925050600181019050611693565b5050505092915050565b6000813590506116d981611a29565b92915050565b6000813590506116ee81611a40565b92915050565b6000806040838503121561170757600080fd5b600083013567ffffffffffffffff81111561172157600080fd5b61172d858286016114cd565b925050602083013567ffffffffffffffff81111561174a57600080fd5b61175685828601611647565b9150509250929050565b6000806040838503121561177357600080fd5b600083013567ffffffffffffffff81111561178d57600080fd5b611799858286016115c4565b92505060206117aa858286016116df565b9150509250929050565b6000602082840312156117c657600080fd5b60006117d4848285016116ca565b91505092915050565b6000602082840312156117ef57600080fd5b60006117fd848285016116df565b91505092915050565b61180f816119a5565b82525050565b61181e816119b7565b82525050565b61182d816119c3565b82525050565b61183c81611a0a565b82525050565b61184b81611a00565b82525050565b60006020820190506118666000830184611806565b92915050565b60006020820190506118816000830184611815565b92915050565b600060208201905061189c6000830184611824565b92915050565b60006020820190506118b76000830184611833565b92915050565b60006020820190506118d26000830184611842565b92915050565b6000604051905081810181811067ffffffffffffffff821117156118fb57600080fd5b8060405250919050565b600067ffffffffffffffff82111561191c57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561194457600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561196c57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff82111561199457600080fd5b602082029050602081019050919050565b60006119b0826119e0565b9050919050565b60008115159050919050565b6000819050919050565b60008190506119db82611a1c565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611a15826119cd565b9050919050565b60078110611a2657fe5b50565b611a32816119c3565b8114611a3d57600080fd5b50565b611a4981611a00565b8114611a5457600080fd5b5056fea365627a7a7231582056bc33e7f3748abfebadaf3017c637a1d8524f54b14365b2c53c18057b8e61456c6578706572696d656e74616cf564736f6c63430005100040",
}

// FairswapABI is the input ABI used to generate the binding from.
// Deprecated: Use FairswapMetaData.ABI instead.
var FairswapABI = FairswapMetaData.ABI

// FairswapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FairswapMetaData.Bin instead.
var FairswapBin = FairswapMetaData.Bin

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentOfKey [32]byte, _EncInputRoot [32]byte, _receiverEntryKey [32]byte, circuitGatesOperationArray []*big.Int) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := FairswapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FairswapBin), backend, _price, _commitmentOfKey, _EncInputRoot, _receiverEntryKey, circuitGatesOperationArray)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// Fairswap is an auto generated Go binding around an Ethereum contract.
type Fairswap struct {
	FairswapCaller     // Read-only binding to the contract
	FairswapTransactor // Write-only binding to the contract
	FairswapFilterer   // Log filterer for contract events
}

// FairswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairswapSession struct {
	Contract     *Fairswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairswapCallerSession struct {
	Contract *FairswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FairswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairswapTransactorSession struct {
	Contract     *FairswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FairswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairswapRaw struct {
	Contract *Fairswap // Generic contract binding to access the raw methods on
}

// FairswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairswapCallerRaw struct {
	Contract *FairswapCaller // Generic read-only contract binding to access the raw methods on
}

// FairswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairswapTransactorRaw struct {
	Contract *FairswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairswap creates a new instance of Fairswap, bound to a specific deployed contract.
func NewFairswap(address common.Address, backend bind.ContractBackend) (*Fairswap, error) {
	contract, err := bindFairswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// NewFairswapCaller creates a new read-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapCaller(address common.Address, caller bind.ContractCaller) (*FairswapCaller, error) {
	contract, err := bindFairswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapCaller{contract: contract}, nil
}

// NewFairswapTransactor creates a new write-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapTransactor(address common.Address, transactor bind.ContractTransactor) (*FairswapTransactor, error) {
	contract, err := bindFairswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapTransactor{contract: contract}, nil
}

// NewFairswapFilterer creates a new log filterer instance of Fairswap, bound to a specific deployed contract.
func NewFairswapFilterer(address common.Address, filterer bind.ContractFilterer) (*FairswapFilterer, error) {
	contract, err := bindFairswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairswapFilterer{contract: contract}, nil
}

// bindFairswap binds a generic wrapper to an already deployed contract.
func bindFairswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.FairswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transact(opts, method, params...)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() view returns(bytes32)
func (_Fairswap *FairswapCaller) EncInputRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "EncInputRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() view returns(bytes32)
func (_Fairswap *FairswapSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() view returns(bytes32)
func (_Fairswap *FairswapCallerSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) view returns(bool)
func (_Fairswap *FairswapCaller) Mverify(opts *bind.CallOpts, complaint [][32]byte, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "Mverify", complaint, _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) view returns(bool)
func (_Fairswap *FairswapSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Mverify is a free data retrieval call binding the contract method 0xa2630372.
//
// Solidity: function Mverify(bytes32[] complaint, uint256 _index) view returns(bool)
func (_Fairswap *FairswapCallerSession) Mverify(complaint [][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() view returns(uint256)
func (_Fairswap *FairswapCaller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "Now")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() view returns(uint256)
func (_Fairswap *FairswapSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() view returns(uint256)
func (_Fairswap *FairswapCallerSession) Now() (*big.Int, error) {
	return _Fairswap.Contract.Now(&_Fairswap.CallOpts)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) view returns(uint256)
func (_Fairswap *FairswapCaller) CircuitGatesOperation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "circuitGatesOperation", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) view returns(uint256)
func (_Fairswap *FairswapSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) view returns(uint256)
func (_Fairswap *FairswapCallerSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() view returns(bytes32)
func (_Fairswap *FairswapCaller) CommitmentOfKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "commitmentOfKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() view returns(bytes32)
func (_Fairswap *FairswapSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// CommitmentOfKey is a free data retrieval call binding the contract method 0xca641f86.
//
// Solidity: function commitmentOfKey() view returns(bytes32)
func (_Fairswap *FairswapCallerSession) CommitmentOfKey() ([32]byte, error) {
	return _Fairswap.Contract.CommitmentOfKey(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(bytes32)
func (_Fairswap *FairswapCaller) Key(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "key")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(bytes32)
func (_Fairswap *FairswapSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() view returns(bytes32)
func (_Fairswap *FairswapCallerSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Fairswap *FairswapCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "phase")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Fairswap *FairswapSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Fairswap *FairswapCallerSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Fairswap *FairswapCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Fairswap *FairswapSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Fairswap *FairswapCallerSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() view returns(address)
func (_Fairswap *FairswapCaller) ReceiverAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "receiverAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() view returns(address)
func (_Fairswap *FairswapSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() view returns(address)
func (_Fairswap *FairswapCallerSession) ReceiverAddress() (common.Address, error) {
	return _Fairswap.Contract.ReceiverAddress(&_Fairswap.CallOpts)
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() view returns(bytes32)
func (_Fairswap *FairswapCaller) ReceiverEntryKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "receiverEntryKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() view returns(bytes32)
func (_Fairswap *FairswapSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
}

// ReceiverEntryKey is a free data retrieval call binding the contract method 0xeaa94830.
//
// Solidity: function receiverEntryKey() view returns(bytes32)
func (_Fairswap *FairswapCallerSession) ReceiverEntryKey() ([32]byte, error) {
	return _Fairswap.Contract.ReceiverEntryKey(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() view returns(address)
func (_Fairswap *FairswapCaller) SenderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "senderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() view returns(address)
func (_Fairswap *FairswapSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() view returns(address)
func (_Fairswap *FairswapCallerSession) SenderAddress() (common.Address, error) {
	return _Fairswap.Contract.SenderAddress(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Fairswap *FairswapCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fairswap.contract.Call(opts, &out, "timeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Fairswap *FairswapSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint256)
func (_Fairswap *FairswapCallerSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() payable returns()
func (_Fairswap *FairswapTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() payable returns()
func (_Fairswap *FairswapSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() payable returns()
func (_Fairswap *FairswapTransactorSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactor) Complain(opts *bind.TransactOpts, complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "complain", complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapSession) Complain(complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0xde30b907.
//
// Solidity: function complain(bytes32[][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactorSession) Complain(complaint [][][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactor) InitializeRecieverAddress(opts *bind.TransactOpts, recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "initializeRecieverAddress", recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
}

// InitializeRecieverAddress is a paid mutator transaction binding the contract method 0x9ab164e5.
//
// Solidity: function initializeRecieverAddress(bytes32 recEntryKey) returns()
func (_Fairswap *FairswapTransactorSession) InitializeRecieverAddress(recEntryKey [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.InitializeRecieverAddress(&_Fairswap.TransactOpts, recEntryKey)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactor) NextStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "nextStage")
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactorSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactor) ReceiverGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "receiverGetEther")
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Fairswap *FairswapTransactorSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.ReceiverGetEther(&_Fairswap.TransactOpts)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactor) RevealKey(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "revealKey", _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactorSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactor) SenderGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "senderGetEther")
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Fairswap *FairswapTransactorSession) SenderGetEther() (*types.Transaction, error) {
	return _Fairswap.Contract.SenderGetEther(&_Fairswap.TransactOpts)
}
