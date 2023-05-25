package node_test

import (
	"context"
	ob64 "encoding/base64"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/go-ethlibs/eth"
	"github.com/INFURA/go-ethlibs/node"
)

const COUNTER_BYTECODE = "60806040526040518060400160405280600781526020017f436f756e74657200000000000000000000000000000000000000000000000000815250600090816200004a919062000356565b5060006002553480156200005d57600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ff15da729ec5b36e9bda8b3f71979cdac5d0f3169f8590778ac0cd82cc5cc1d4a604051620000ce906200049e565b60405180910390a1620004c0565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200015e57607f821691505b60208210810362000174576200017362000116565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620001de7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200019f565b620001ea86836200019f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062000237620002316200022b8462000202565b6200020c565b62000202565b9050919050565b6000819050919050565b620002538362000216565b6200026b62000262826200023e565b848454620001ac565b825550505050565b600090565b6200028262000273565b6200028f81848462000248565b505050565b5b81811015620002b757620002ab60008262000278565b60018101905062000295565b5050565b601f8211156200030657620002d0816200017a565b620002db846200018f565b81016020851015620002eb578190505b62000303620002fa856200018f565b83018262000294565b50505b505050565b600082821c905092915050565b60006200032b600019846008026200030b565b1980831691505092915050565b600062000346838362000318565b9150826002028217905092915050565b6200036182620000dc565b67ffffffffffffffff8111156200037d576200037c620000e7565b5b62000389825462000145565b62000396828285620002bb565b600060209050601f831160018114620003ce5760008415620003b9578287015190505b620003c5858262000338565b86555062000435565b601f198416620003de866200017a565b60005b828110156200040857848901518255600182019150602085019450602081019050620003e1565b8683101562000428578489015162000424601f89168262000318565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f48656c6c6f2c20436f756e746572000000000000000000000000000000000000600082015250565b600062000486600e836200043d565b915062000493826200044e565b602082019050919050565b60006020820190508181036000830152620004b98162000477565b9050919050565b61078d80620004d06000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b14610137578063a87d942c14610155578063c8a4ac9c14610173578063d14e62b8146101a3578063ee82ac5e146101bf5761009e565b806306fdde03146100a3578063119fbbd4146100c15780631a93d1c3146100cb578063672d5d3b146100e95780638361ff9c14610107575b600080fd5b6100ab6101ef565b6040516100b8919061043c565b60405180910390f35b6100c961027d565b005b6100d3610299565b6040516100e09190610477565b60405180910390f35b6100f16102a1565b6040516100fe9190610477565b60405180910390f35b610121600480360381019061011c91906104c3565b6102a9565b60405161012e9190610477565b60405180910390f35b61013f6102f7565b60405161014c9190610531565b60405180910390f35b61015d61031d565b60405161016a9190610477565b60405180910390f35b61018d6004803603810190610188919061054c565b610327565b60405161019a9190610477565b60405180910390f35b6101bd60048036038101906101b891906104c3565b61033d565b005b6101d960048036038101906101d491906104c3565b6103a1565b6040516101e691906105a5565b60405180910390f35b600080546101fc906105ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610228906105ef565b80156102755780601f1061024a57610100808354040283529160200191610275565b820191906000526020600020905b81548152906001019060200180831161025857829003601f168201915b505050505081565b600160026000828254610290919061064f565b92505081905550565b600045905090565b600043905090565b6000600a8211156102ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e6906106f5565b60405180910390fd5b819050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600254905090565b600081836103359190610715565b905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461039757600080fd5b8060028190555050565b600081409050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156103e65780820151818401526020810190506103cb565b60008484015250505050565b6000601f19601f8301169050919050565b600061040e826103ac565b61041881856103b7565b93506104288185602086016103c8565b610431816103f2565b840191505092915050565b600060208201905081810360008301526104568184610403565b905092915050565b6000819050919050565b6104718161045e565b82525050565b600060208201905061048c6000830184610468565b92915050565b600080fd5b6104a08161045e565b81146104ab57600080fd5b50565b6000813590506104bd81610497565b92915050565b6000602082840312156104d9576104d8610492565b5b60006104e7848285016104ae565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061051b826104f0565b9050919050565b61052b81610510565b82525050565b60006020820190506105466000830184610522565b92915050565b6000806040838503121561056357610562610492565b5b6000610571858286016104ae565b9250506020610582858286016104ae565b9150509250929050565b6000819050919050565b61059f8161058c565b82525050565b60006020820190506105ba6000830184610596565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061060757607f821691505b60208210810361061a576106196105c0565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061065a8261045e565b91506106658361045e565b925082820190508082111561067d5761067c610620565b5b92915050565b7f56616c7565206d757374206e6f742062652067726561746572207468616e203160008201527f302e000000000000000000000000000000000000000000000000000000000000602082015250565b60006106df6022836103b7565b91506106ea82610683565b604082019050919050565b6000602082019050818103600083015261070e816106d2565b9050919050565b60006107208261045e565b915061072b8361045e565b92508282026107398161045e565b915082820484148315176107505761074f610620565b5b509291505056fea264697066735822122060d14a11e936a646739497a8e93e302a1651192429e8063317c44847ec1eae7564736f6c63430008130033"

func getClient(t *testing.T, ctx context.Context) node.Client {
	base_url := os.Getenv("ETHLIBS_TEST_URL")
	if base_url == "" {
		t.Skip("ETHLIBS_TEST_URL not set, skipping test. Set to a valid http/ws URL to execute this test.")
	}
	auth_id := os.Getenv("AUTH_ID")
	if auth_id == "" {
		t.Skip("AUTH_ID not set, skipping test.")
	}
	url := fmt.Sprintf("%s%s", base_url, auth_id)

	auth_pass := os.Getenv("AUTH_PASS")
	if auth_pass == "" {
		t.Skip("AUTH_PASS not set, skipping test.")
	}

	base64Header := ob64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", auth_id, auth_pass)))

	header := http.Header{
		"Authorization": {fmt.Sprintf("Basic %s", base64Header)},
	}
	conn, err := node.NewClient(ctx, url, header)
	require.NoError(t, err, "creating websocket connection should not fail")
	return conn
}

func setup(t *testing.T, conn node.Client) {
	// create smart contract
	// tx := eth.Transaction{
	// 	From:     *eth.MustAddress("0x"),
	// 	To:       nil, // nil for contract creation
	// 	Value:    *eth.MustQuantity("0x00"),
	// 	GasPrice: eth.MustQuantity("0x01"),
	// 	Gas:      *eth.MustQuantity("0x11940"), // gas limit
	// 	Nonce:    *eth.MustQuantity("0x00"),
	// 	ChainId:  eth.MustQuantity("1337"),
	// 	Input:    *eth.MustData(COUNTER_BYTECODE),
	// }
}

func TestConnection_Call_Simple_Contract(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	tx := eth.Transaction{
		From:  *eth.MustAddress("0x148772F29058DcC772613260b078dCa8C14afF6c"),
		To:    eth.MustAddress("0x790b2DaF786774BEe346b9B8913b259eD2d354D0"), // contract address
		Value: *eth.MustQuantity("0x00"),
		// mul(2,3)
		Input: *eth.MustData("0xc8a4ac9c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003"),
	}

	hash, err := conn.Call(ctx, tx, *eth.MustBlockNumberOrTag("latest"))
	require.NoError(t, err)
	require.Equal(t, hash, "0x0000000000000000000000000000000000000000000000000000000000000006")
}

func TestConnection_Get_Accounts(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	accountList, err := conn.GetAccounts(ctx)
	require.NoError(t, err)
	require.Empty(t, accountList)
}

func TestConnection_Get_Balance(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	bal, err := conn.GetBalance(ctx, *eth.MustAddress("0x148772F29058DcC772613260b078dCa8C14afF6c"), *eth.MustBlockNumberOrTag("latest"))
	require.NoError(t, err)
	require.NotNil(t, bal)
}

func TestConnection_GetTransactionCount(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	// Checks the current pending nonce for account can be retrieved
	blockNum1 := eth.MustBlockNumberOrTag("latest")
	pendingNonce1, err := conn.GetTransactionCount(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum1)
	require.NoError(t, err)
	require.NotEmpty(t, pendingNonce1, "pending nonce must not be nil")

	// Should catch failure since it is looking for a nonce of a future block
	blockNum2 := eth.MustBlockNumberOrTag("0x7654321")
	pendingNonce2, err := conn.GetTransactionCount(ctx, "0xed28874e52A12f0D42118653B0FBCee0ACFadC00", *blockNum2)
	require.Error(t, err)
	require.Empty(t, pendingNonce2, "pending nonce must not exist since it is a future block")
}

func TestConnection_EstimateGas(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	from := eth.MustAddress("0xed28874e52A12f0D42118653B0FBCee0ACFadC00")
	tx := eth.Transaction{
		Nonce:    eth.QuantityFromUInt64(146),
		GasPrice: eth.OptionalQuantityFromInt(3000000000),
		Gas:      eth.QuantityFromUInt64(22000),
		To:       eth.MustAddress("0x43700db832E9Ac990D36d6279A846608643c904E"),
		Value:    *eth.OptionalQuantityFromInt(100),
		From:     *from,
	}

	gas, err := conn.EstimateGas(ctx, tx)
	require.NoError(t, err)
	require.NotEqual(t, gas, 0, "estimate gas cannot be equal to zero.")
}

func TestConnection_MaxPriorityFeePerGas(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	fee, err := conn.MaxPriorityFeePerGas(ctx)
	require.NoError(t, err)
	require.NotEqual(t, fee, 0, "fee cannot be equal to 0")
}

func TestConnection_GasPrice(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	gasPrice, err := conn.GasPrice(ctx)
	require.NoError(t, err)
	require.NotEqual(t, gasPrice, 0, "gas price cannot be equal to 0")
}

func TestConnection_NetVersion(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	netVersion, err := conn.NetVersion(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, netVersion, "net version id must not be nil")
}

func TestConnection_ChainId(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	chainId, err := conn.ChainId(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, chainId, "chain id must not be nil")
}

func TestConnection_SendRawTransactionInValidEmpty(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	txHash, err := conn.SendRawTransaction(ctx, "0x0")
	require.Error(t, err)
	require.Empty(t, txHash, "txHash must be nil")
}

func TestConnection_SendRawTransactionInValidOldNonce(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	data := eth.MustData("0xf86e0185174876e8008252089460c063d3f3b744e2d153fcbe66a068b09109cf1b865af3107a400084baadf00d2ea0b4d9e2edbd2a2d9a38cf0415f9d03849e6a6f2de8562d7cd74eda89397882030a056edb455e9ffa07ad22f8b06f9065564911f796a026e1b2177ecaad995198aaa")
	txHash, err := conn.SendRawTransaction(ctx, data.String())
	require.Error(t, err)
	require.Equal(t, err.Error(), "{\"code\":-32000,\"message\":\"nonce too low\"}")
	require.Empty(t, txHash, "txHash must be nil")
}

func TestConnection_FutureBlockByNumber(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	blockNumber, err := conn.BlockNumber(ctx)
	require.NoError(t, err)

	next, err := conn.BlockByNumber(ctx, blockNumber+1000, false)
	require.Nil(t, next, "future block should be nil")
	require.Error(t, err, "requesting a future block should return an error")
	require.Equal(t, node.ErrBlockNotFound, err)

	// get a the genesis block by number which should _not_ fail
	genesis, err := conn.BlockByNumber(ctx, 0, false)
	require.NoError(t, err, "requesting genesis block by number should not fail")
	require.NotNil(t, genesis, "genesis block must not be nil")
}

func TestConnection_InvalidBlockByHash(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	b, err := conn.BlockByHash(ctx, "invalid", false)
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, b, "block from invalid hash should be nil")

	b, err = conn.BlockByHash(ctx, "0x1234", false)
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, b, "block from invalid hash should be nil")

	b, err = conn.BlockByHash(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d", false)
	require.Error(t, err, "requesting a non-existent block should should return an error")
	require.Nil(t, b, "block from non-existent hash should be nil")
	require.Equal(t, node.ErrBlockNotFound, err)

	// get the genesis block which should _not_ fail
	// https://goerli.etherscan.io/block/8000000
	b, err = conn.BlockByHash(ctx, "0x2ae83825ac6b2a2b2509da8617cf31072a5628e9a818f177316f4f4bcdfafd06", true)
	require.NoError(t, err, "genesis block hash should not return an error")
	require.NotNil(t, b, "genesis block should be retrievable by hash")
}

func TestConnection_InvalidTransactionByHash(t *testing.T) {
	ctx := context.Background()
	conn := getClient(t, ctx)

	tx, err := conn.TransactionByHash(ctx, "invalid")
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, tx, "tx from invalid hash should be nil")

	tx, err = conn.TransactionByHash(ctx, "0x1234")
	require.Error(t, err, "requesting an invalid hash should return an error")
	require.Nil(t, tx, "tx from invalid hash should be nil")

	tx, err = conn.TransactionByHash(ctx, "0x0badf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00dbadf00d")
	require.Error(t, err, "requesting an non-existent hash should return an error")
	require.Nil(t, tx, "tx from non-existent hash should be nil")
	require.Equal(t, node.ErrTransactionNotFound, err)

	// get an early tx which should _not_ fail
	// https://goerli.etherscan.io/tx/0x752ca2e3175c0dfd8b8612abcd2dac3134445f29e764d33645726cbcd57aefd1
	tx, err = conn.TransactionByHash(ctx, "0x752ca2e3175c0dfd8b8612abcd2dac3134445f29e764d33645726cbcd57aefd1")
	require.NoError(t, err, "early tx should not return an error")
	require.NotNil(t, tx, "early tx should be retrievable by hash")
}
