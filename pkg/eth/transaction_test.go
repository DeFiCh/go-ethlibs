package eth_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/ethereum-interaction/pkg/eth"
)

func TestTransactionFailedContractCreation(t *testing.T) {
	// This is a transaction that tries to create a contract but failed (creates is null)
	payload := `{"blockHash":"0x774c34a19ff36b1d3669c54fe385502ad32862ed728da53b72b771f82c08b474","blockNumber":"0x6e2e24","from":"0x6b59210ade46b62b25e82e95ab390a7ccadd4c3a","gas":"0x2dc6c0","gasPrice":"0x1dcd65000","hash":"0xa6c1da26576bc41a722a3058bb1fa36e9c58188bd960bc1f40a976b6cb25e15e","input":"0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160146101000a81548160ff0219169083151502179055506000600160156101000a81548160ff021916908315150217905550611193806100d76000396000f3006080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680636ea056a9146100dd5780638da5cb5b146101425780638fc1038e1461019957806397dc97cb146101fe578063b9b8af0b14610255578063c1756a2c14610284578063c18cfe86146102dc578063ce46e04614610337578063d996c57a14610366578063dcc279c814610395578063ed6fc1bc146103c4578063f2fde38b14610407575b60001515600160159054906101000a900460ff16151514156100db57600080fd5b005b3480156100e957600080fd5b50610128600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061044a565b604051808215151515815260200191505060405180910390f35b34801561014e57600080fd5b50610157610712565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a557600080fd5b506101e4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610737565b604051808215151515815260200191505060405180910390f35b34801561020a57600080fd5b50610213610a33565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561026157600080fd5b5061026a610a59565b604051808215151515815260200191505060405180910390f35b6102c2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a6c565b604051808215151515815260200191505060405180910390f35b3480156102e857600080fd5b5061031d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b7f565b604051808215151515815260200191505060405180910390f35b34801561034357600080fd5b5061034c610e34565b604051808215151515815260200191505060405180910390f35b34801561037257600080fd5b50610393600480360381019080803515159060200190929190505050610e47565b005b3480156103a157600080fd5b506103c2600480360381019080803515159060200190929190505050610ebf565b005b3480156103d057600080fd5b50610405600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f37565b005b34801561041357600080fd5b50610448600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611012565b005b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806104f75750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561050257600080fd5b60009150600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614158015610556575060001515600160149054906101000a900460ff161515145b151561056157600080fd5b849050838173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561060057600080fd5b505af1158015610614573d6000803e3d6000fd5b505050506040513d602081101561062a57600080fd5b8101908080519060200190929190505050101561064a576000925061070a565b8073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33866040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156106ed57600080fd5b505af1158015610701573d6000803e3d6000fd5b50505050600192505b505092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806107e45750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156107ef57600080fd5b60009150600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614158015610843575060001515600160149054906101000a900460ff161515145b151561084e57600080fd5b849050838173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156108ed57600080fd5b505af1158015610901573d6000803e3d6000fd5b505050506040513d602081101561091757600080fd5b810190808051906020019092919050505010156109375760009250610a2b565b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd3033876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b158015610a0e57600080fd5b505af1158015610a22573d6000803e3d6000fd5b50505050600192505b505092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610b165750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610b2157600080fd5b60001515600160159054906101000a900460ff1615151415610b4257600080fd5b8273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050905092915050565b60008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610c2c5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610c3757600080fd5b60009150600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614158015610c8b575060001515600160149054906101000a900460ff161515145b1515610c9657600080fd5b8390508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb338373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610d5157600080fd5b505af1158015610d65573d6000803e3d6000fd5b505050506040513d6020811015610d7b57600080fd5b81019080805190602001909291905050506040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610e1157600080fd5b505af1158015610e25573d6000803e3d6000fd5b50505050600192505050919050565b600160159054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610ea257600080fd5b80600160156101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f1a57600080fd5b80600160146101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f9257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610fce57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561106d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156110a957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820bde856abe6819d4a475671ad534de55797d60abb945f5ea2ee037d554d02cf270029","nonce":"0x5f1f","r":"0xbe9872010301ebd71ea659b54b5d8b9f117249c9a5c23ef65de21c7807c7ddc8","s":"0x883971fb706e2336928a21be320c98c457238c2ef3466ba437f2f0bbfe32ad7","to":null,"transactionIndex":"0x6","v":"0x26","value":"0x0"}`

	tx := eth.Transaction{}
	err := json.Unmarshal([]byte(payload), &tx)
	require.NoError(t, err)
	require.Nil(t, tx.To)

	j, err := json.Marshal(&tx)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(payload), j)
}

func TestTransactionSuccessfulContractCreation(t *testing.T) {
	payload := `{"blockHash":"0x00e6ace83980aa65ccde1ad2887671ca99d7a83d490204d3f8efc4d194bf8bcd","blockNumber":"0x9e346f","chainId":null,"condition":null,"creates":"0xac2475e9325b586f0b7e1928b813c9098d9ba262","from":"0x2cf222b75ffbf44cac4c5ac706dc0a756dfce9cd","gas":"0x124f80","gasPrice":"0x4a817c800","hash":"0x0cef0f5a1e7258d464646f919daaec612d27b7c2a25b11f26a090525a2feeac3","input":"0x60606040526001805460ff19169055341561001957600080fd5b604051610ef0380380610ef0833981016040528080519091019050805160031461004257600080fd5b600081805161005592916020019061005c565b50506100ea565b8280548282559060005260206000209081019282156100b3579160200282015b828111156100b35782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061007c565b506100bf9291506100c3565b5090565b6100e791905b808211156100bf578054600160a060020a03191681556001016100c9565b90565b610df7806100f96000396000f30060606040526004361061007f5763ffffffff60e060020a6000350416630dcd7a6c81146100f45780632079fb9a146101685780632da034091461019a57806339125215146101bf5780637df73e271461026f578063a0b7967b146102a2578063a68a76cc146102c7578063abe3219c146102da578063fc0f392d146102ed575b60003411156100f2577f6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f03334600036604051600160a060020a0385168152602081018490526060604082018181529082018390526080820184848082843782019150509550505050505060405180910390a15b005b34156100ff57600080fd5b6100f260048035600160a060020a03908116916024803592604435169160643591608435919060c49060a43590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061030095505050505050565b341561017357600080fd5b61017e600435610431565b604051600160a060020a03909116815260200160405180910390f35b34156101a557600080fd5b6100f2600160a060020a0360043581169060243516610459565b34156101ca57600080fd5b6100f260048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635966020808201359750919550606081019450604090810135860180830194503592508291601f8301819004810201905190810160405281815292919060208401838380828437509496506104d895505050505050565b341561027a57600080fd5b61028e600160a060020a0360043516610718565b604051901515815260200160405180910390f35b34156102ad57600080fd5b6102b5610776565b60405190815260200160405180910390f35b34156102d257600080fd5b61017e6107bd565b34156102e557600080fd5b61028e6107e2565b34156102f857600080fd5b6100f26107eb565b60008061030c33610718565b151561031757600080fd5b87878787876040517f455243323000000000000000000000000000000000000000000000000000000081526c01000000000000000000000000600160a060020a03968716810260058301526019820195909552929094169092026039820152604d810191909152606d810191909152608d01604051809103902091506103a0888385888861084a565b5085905080600160a060020a031663a9059cbb898960006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561040157600080fd5b6102c65a03f1151561041257600080fd5b50505060405180519050151561042757600080fd5b5050505050505050565b600080548290811061043f57fe5b600091825260209091200154600160a060020a0316905081565b600061046433610718565b151561046f57600080fd5b5081600160a060020a038116633ef133678360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b15156104bf57600080fd5b6102c65a03f115156104d057600080fd5b505050505050565b6000806104e433610718565b15156104ef57600080fd5b87878787876040517f455448455200000000000000000000000000000000000000000000000000000081526c01000000000000000000000000600160a060020a038716026005820152601981018590526039810184805190602001908083835b6020831061056e5780518252601f19909201916020918201910161054f565b6001836020036101000a038019825116818451161790925250505091909101938452505060208201526040908101935091505051809103902091506105b6888385888861084a565b905087600160a060020a0316878760405180828051906020019080838360005b838110156105ee5780820151838201526020016105d6565b50505050905090810190601f16801561061b5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185876187965a03f192505050151561063f57600080fd5b7f59bed9ab5d78073465dd642a9e3e76dfdb7d53bcae9d09df7d0b8f5234d5a8063382848b8b8b604051600160a060020a038088168252868116602083015260408201869052841660608201526080810183905260c060a0820181815290820183818151815260200191508051906020019080838360005b838110156106cf5780820151838201526020016106b7565b50505050905090810190601f1680156106fc5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15050505050505050565b6000805b60005481101561076b5782600160a060020a031660008281548110151561073f57fe5b600091825260209091200154600160a060020a031614156107635760019150610770565b60010161071c565b600091505b50919050565b600080805b600a8110156107b45781600282600a811061079257fe5b015411156107ac57600281600a81106107a757fe5b015491505b60010161077b565b50600101919050565b60006107c7610a44565b604051809103906000f08015156107dd57600080fd5b905090565b60015460ff1681565b6107f433610718565b15156107ff57600080fd5b6001805460ff1916811790557f0909e8f76a4fd3e970f2eaef56c0ee6dfaf8b87c5b8d3f56ffce78e825a9115733604051600160a060020a03909116815260200160405180910390a1565b60008061085786866108cf565b60015490915060ff168015610872575061087087610718565b155b1561087c57600080fd5b4284101561088957600080fd5b61089283610983565b61089b81610718565b15156108a657600080fd5b33600160a060020a031681600160a060020a031614156108c557600080fd5b9695505050505050565b60008060008084516041146108e357600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff16101561090b57601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561097057600080fd5b5050602060405103519695505050505050565b60008061098f33610718565b151561099a57600080fd5b5060009050805b600a8110156109f55782600282600a81106109b857fe5b015414156109c557600080fd5b600282600a81106109d257fe5b0154600282600a81106109e157fe5b015410156109ed578091505b6001016109a1565b600282600a8110610a0257fe5b0154831015610a1057600080fd5b600282600a8110610a1d57fe5b015461271001831115610a2f57600080fd5b82600283600a8110610a3d57fe5b0155505050565b60405161037780610a558339019056006060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561033c8061003b6000396000f30060606040526004361061003c5763ffffffff60e060020a600035041662821de381146100eb5780633ef133671461011a5780636b9f96ea1461013b575b60008054600160a060020a0316903490366040518083838082843782019150509250505060006040518083038185876187965a03f192505050151561008057600080fd5b7f69b31548dea9b3b707b4dff357d326e3e9348b24e7a6080a218a6edeeec48f9b3334600036604051600160a060020a0385168152602081018490526060604082018181529082018390526080820184848082843782019150509550505050505060405180910390a1005b34156100f657600080fd5b6100fe61014e565b604051600160a060020a03909116815260200160405180910390f35b341561012557600080fd5b610139600160a060020a036004351661015d565b005b341561014657600080fd5b6101396102d9565b600054600160a060020a031681565b600080548190819033600160a060020a0390811691161461017d57600080fd5b83925030915082600160a060020a03166370a082318360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156101da57600080fd5b6102c65a03f115156101eb57600080fd5b5050506040518051915050801515610202576102d3565b60008054600160a060020a038086169263a9059cbb929091169084906040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561026857600080fd5b6102c65a03f1151561027957600080fd5b50505060405180519050151561028e57600080fd5b7f9401e4e79c19cbe2bd774cb70a94ba660e6718be1bac1298ab3b07f454a608218482604051600160a060020a03909216825260208201526040908101905180910390a15b50505050565b600054600160a060020a039081169030163160405160006040518083038185876187965a03f192505050151561030e57600080fd5b5600a165627a7a72305820a6b61178cc9f27c0f16522b572583d67a89a7cea9f0d74293c1771a8260c38650029a165627a7a72305820498b920a8626a935d3fd3ce78cdf41089ceef2ac9798b09f7f22fa74c887a0320029000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000030000000000000000000000004cb7584d31ecfaf0c16ac5f6faa5004ddd870059000000000000000000000000aece3feb3c4c2888782778c93baefebb8a9142a40000000000000000000000002cf222b75ffbf44cac4c5ac706dc0a756dfce9cd","nonce":"0xdf","publicKey":"0x48f6c6c63206c73980e90a0f5f39cfd60352606ec0347819a6dc30cf85853588cff72945bb14e22234d0ca470e705cb033849910ec65e6c03f61ca88120e1cf0","r":"0x889510df0e466a858d0ef38e7ce4b95f3fdd43b02ca771d07cb8976a1de15679","raw":"0xf90fe481df8504a817c80083124f808080b90f9060606040526001805460ff19169055341561001957600080fd5b604051610ef0380380610ef0833981016040528080519091019050805160031461004257600080fd5b600081805161005592916020019061005c565b50506100ea565b8280548282559060005260206000209081019282156100b3579160200282015b828111156100b35782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019061007c565b506100bf9291506100c3565b5090565b6100e791905b808211156100bf578054600160a060020a03191681556001016100c9565b90565b610df7806100f96000396000f30060606040526004361061007f5763ffffffff60e060020a6000350416630dcd7a6c81146100f45780632079fb9a146101685780632da034091461019a57806339125215146101bf5780637df73e271461026f578063a0b7967b146102a2578063a68a76cc146102c7578063abe3219c146102da578063fc0f392d146102ed575b60003411156100f2577f6e89d517057028190560dd200cf6bf792842861353d1173761dfa362e1c133f03334600036604051600160a060020a0385168152602081018490526060604082018181529082018390526080820184848082843782019150509550505050505060405180910390a15b005b34156100ff57600080fd5b6100f260048035600160a060020a03908116916024803592604435169160643591608435919060c49060a43590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061030095505050505050565b341561017357600080fd5b61017e600435610431565b604051600160a060020a03909116815260200160405180910390f35b34156101a557600080fd5b6100f2600160a060020a0360043581169060243516610459565b34156101ca57600080fd5b6100f260048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094968635966020808201359750919550606081019450604090810135860180830194503592508291601f8301819004810201905190810160405281815292919060208401838380828437509496506104d895505050505050565b341561027a57600080fd5b61028e600160a060020a0360043516610718565b604051901515815260200160405180910390f35b34156102ad57600080fd5b6102b5610776565b60405190815260200160405180910390f35b34156102d257600080fd5b61017e6107bd565b34156102e557600080fd5b61028e6107e2565b34156102f857600080fd5b6100f26107eb565b60008061030c33610718565b151561031757600080fd5b87878787876040517f455243323000000000000000000000000000000000000000000000000000000081526c01000000000000000000000000600160a060020a03968716810260058301526019820195909552929094169092026039820152604d810191909152606d810191909152608d01604051809103902091506103a0888385888861084a565b5085905080600160a060020a031663a9059cbb898960006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561040157600080fd5b6102c65a03f1151561041257600080fd5b50505060405180519050151561042757600080fd5b5050505050505050565b600080548290811061043f57fe5b600091825260209091200154600160a060020a0316905081565b600061046433610718565b151561046f57600080fd5b5081600160a060020a038116633ef133678360405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b15156104bf57600080fd5b6102c65a03f115156104d057600080fd5b505050505050565b6000806104e433610718565b15156104ef57600080fd5b87878787876040517f455448455200000000000000000000000000000000000000000000000000000081526c01000000000000000000000000600160a060020a038716026005820152601981018590526039810184805190602001908083835b6020831061056e5780518252601f19909201916020918201910161054f565b6001836020036101000a038019825116818451161790925250505091909101938452505060208201526040908101935091505051809103902091506105b6888385888861084a565b905087600160a060020a0316878760405180828051906020019080838360005b838110156105ee5780820151838201526020016105d6565b50505050905090810190601f16801561061b5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185876187965a03f192505050151561063f57600080fd5b7f59bed9ab5d78073465dd642a9e3e76dfdb7d53bcae9d09df7d0b8f5234d5a8063382848b8b8b604051600160a060020a038088168252868116602083015260408201869052841660608201526080810183905260c060a0820181815290820183818151815260200191508051906020019080838360005b838110156106cf5780820151838201526020016106b7565b50505050905090810190601f1680156106fc5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15050505050505050565b6000805b60005481101561076b5782600160a060020a031660008281548110151561073f57fe5b600091825260209091200154600160a060020a031614156107635760019150610770565b60010161071c565b600091505b50919050565b600080805b600a8110156107b45781600282600a811061079257fe5b015411156107ac57600281600a81106107a757fe5b015491505b60010161077b565b50600101919050565b60006107c7610a44565b604051809103906000f08015156107dd57600080fd5b905090565b60015460ff1681565b6107f433610718565b15156107ff57600080fd5b6001805460ff1916811790557f0909e8f76a4fd3e970f2eaef56c0ee6dfaf8b87c5b8d3f56ffce78e825a9115733604051600160a060020a03909116815260200160405180910390a1565b60008061085786866108cf565b60015490915060ff168015610872575061087087610718565b155b1561087c57600080fd5b4284101561088957600080fd5b61089283610983565b61089b81610718565b15156108a657600080fd5b33600160a060020a031681600160a060020a031614156108c557600080fd5b9695505050505050565b60008060008084516041146108e357600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff16101561090b57601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561097057600080fd5b5050602060405103519695505050505050565b60008061098f33610718565b151561099a57600080fd5b5060009050805b600a8110156109f55782600282600a81106109b857fe5b015414156109c557600080fd5b600282600a81106109d257fe5b0154600282600a81106109e157fe5b015410156109ed578091505b6001016109a1565b600282600a8110610a0257fe5b0154831015610a1057600080fd5b600282600a8110610a1d57fe5b015461271001831115610a2f57600080fd5b82600283600a8110610a3d57fe5b0155505050565b60405161037780610a558339019056006060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561033c8061003b6000396000f30060606040526004361061003c5763ffffffff60e060020a600035041662821de381146100eb5780633ef133671461011a5780636b9f96ea1461013b575b60008054600160a060020a0316903490366040518083838082843782019150509250505060006040518083038185876187965a03f192505050151561008057600080fd5b7f69b31548dea9b3b707b4dff357d326e3e9348b24e7a6080a218a6edeeec48f9b3334600036604051600160a060020a0385168152602081018490526060604082018181529082018390526080820184848082843782019150509550505050505060405180910390a1005b34156100f657600080fd5b6100fe61014e565b604051600160a060020a03909116815260200160405180910390f35b341561012557600080fd5b610139600160a060020a036004351661015d565b005b341561014657600080fd5b6101396102d9565b600054600160a060020a031681565b600080548190819033600160a060020a0390811691161461017d57600080fd5b83925030915082600160a060020a03166370a082318360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156101da57600080fd5b6102c65a03f115156101eb57600080fd5b5050506040518051915050801515610202576102d3565b60008054600160a060020a038086169263a9059cbb929091169084906040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561026857600080fd5b6102c65a03f1151561027957600080fd5b50505060405180519050151561028e57600080fd5b7f9401e4e79c19cbe2bd774cb70a94ba660e6718be1bac1298ab3b07f454a608218482604051600160a060020a03909216825260208201526040908101905180910390a15b50505050565b600054600160a060020a039081169030163160405160006040518083038185876187965a03f192505050151561030e57600080fd5b5600a165627a7a72305820a6b61178cc9f27c0f16522b572583d67a89a7cea9f0d74293c1771a8260c38650029a165627a7a72305820498b920a8626a935d3fd3ce78cdf41089ceef2ac9798b09f7f22fa74c887a0320029000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000030000000000000000000000004cb7584d31ecfaf0c16ac5f6faa5004ddd870059000000000000000000000000aece3feb3c4c2888782778c93baefebb8a9142a40000000000000000000000002cf222b75ffbf44cac4c5ac706dc0a756dfce9cd1ba0889510df0e466a858d0ef38e7ce4b95f3fdd43b02ca771d07cb8976a1de15679a00106250ebdf433dff9bef988555a6c6a0f12193b52aef63c0af8cc1b39e387cd","s":"0x106250ebdf433dff9bef988555a6c6a0f12193b52aef63c0af8cc1b39e387cd","standardV":"0x0","to":null,"transactionIndex":"0x0","v":"0x1b","value":"0x0"}`

	tx := eth.Transaction{}
	err := json.Unmarshal([]byte(payload), &tx)
	require.NoError(t, err)
	require.Nil(t, tx.To)
	require.Equal(t, eth.MustAddress("0xac2475e9325b586f0b7e1928b813c9098d9ba262"), tx.Creates)

	j, err := json.Marshal(&tx)
	require.NoError(t, err)

	RequireEqualJSON(t, []byte(payload), j)
}

func TestNewPendingTxNotificationParams(t *testing.T) {

	{
		params := eth.NewPendingTxNotificationParams{}

		pTrx := `{"subscription":"0x70b8a3979a9df0ae3c2e006a192bbd27","result":"0xa2d258a5f2b0ff053a7bb7daa9fba98834dbb5aa6815dbad1c5c7837bbb436a3"}`

		err := json.Unmarshal([]byte(pTrx), &params)
		require.NoError(t, err)
		require.Equal(t, *eth.MustHash("0xa2d258a5f2b0ff053a7bb7daa9fba98834dbb5aa6815dbad1c5c7837bbb436a3"), params.Result)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(pTrx), j)
	}

	{
		params := eth.NewPendingTxNotificationParams{}

		parity := `{"result":"0x1c40d1cface680124774279a5806014c9d1768e46b5520e5ab1f6239b9ee2e76","subscription":"0x5beac702152c0d1f"}`

		err := json.Unmarshal([]byte(parity), &params)
		require.NoError(t, err)
		require.Equal(t, *eth.MustHash("0x1c40d1cface680124774279a5806014c9d1768e46b5520e5ab1f6239b9ee2e76"), params.Result)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(parity), j)
	}

	{
		params := eth.NewPendingTxNotificationParams{}


		kovan := `{"result":"0xc1baf4fd2b6a4f26f53ee6da7b82a69e055ea97b03a1b023afe028643dc12dc0","subscription":"0x5beac702152c0d1f"}`

		err := json.Unmarshal([]byte(kovan), &params)
		require.NoError(t, err)
		require.Equal(t, *eth.MustHash("0xc1baf4fd2b6a4f26f53ee6da7b82a69e055ea97b03a1b023afe028643dc12dc0"), params.Result)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(kovan), j)
	}
}