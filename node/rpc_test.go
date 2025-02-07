package node_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/go-ethlibs/eth"
	"github.com/INFURA/go-ethlibs/node"
)

func getClient(t *testing.T, ctx context.Context) node.Client {
	url := "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"

	// Create connection
	conn, err := node.NewClient(ctx, url, http.Header{})
	require.NoError(t, err, "creating websocket connection should not fail")
	return conn
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

	data := eth.MustData("0xf90150808522ecb25c008307a1208080b8fe608060405234801561001057600080fd5b5060df8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063165c4a1614602d575b600080fd5b603c6038366004605f565b604e565b60405190815260200160405180910390f35b600060588284607f565b9392505050565b600080604083850312156070578182fd5b50508035926020909101359150565b600081600019048311821515161560a457634e487b7160e01b81526011600452602481fd5b50029056fea2646970667358221220223df7833fd08eb1cd3ce363a9c4cb4619c1068a5f5517ea8bb862ed45d994f764736f6c634300080200331ca02096ae58676338a357356c57ed96c6b33573fdf6ac47a575b641a53f198d4f09a02bb1989de4389734be188af15329b3177ca9f1f766020bd11225bc46e242947d")
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
