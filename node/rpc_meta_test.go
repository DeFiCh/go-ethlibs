package node_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/go-ethlibs/eth"
	"github.com/INFURA/go-ethlibs/node"
)

const ALICE = "0x9b8a4af42140d8a4c153a822f02571a1dd037e89"
const BOB = "0x6C34CBb9219d8cAa428835D2073E8ec88BA0a110"

func getMetaDevClient(t *testing.T, ctx context.Context) node.Client {
	url, isSet := os.LookupEnv("NODE_URL")
	if !isSet {
		url = "http://35.187.53.161:20551/"
	}

	// Create connection
	conn, err := node.NewClient(ctx, url, http.Header{})
	require.NoError(t, err, "creating connection should not fail")
	return conn
}

func TestMetaRpc_Block_GetBlockByNumber(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	blockNumber, err := conn.BlockNumber(ctx)
	require.NoError(t, err)

	next, err := conn.BlockByNumber(ctx, blockNumber+1000, false)
	require.Nil(t, next, "future block should be nil")
	require.Error(t, err, "requesting a future block should return an error")
	require.Equal(t, node.ErrBlockNotFound, err)

	// NOTE(canonbrother): getBlockByNumber(0) -> null
	block, err := conn.BlockByNumber(ctx, blockNumber, true)
	println("block: ", block, err)
	require.NoError(t, err, "getBlockByNumber(n) should not fail")
}

func TestMetaRpc_Block_GetBlockByHash(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

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

	blockNumber, _ := conn.BlockNumber(ctx)
	block, err := conn.BlockByNumber(ctx, blockNumber, true)
	require.NoError(t, err, "getBlockByNumber should be no error")
	hash := (*block.Hash).String()
	b, err = conn.BlockByHash(ctx, hash, true)
	require.NoError(t, err, "getBlockByHash should not fail")
	require.NotNil(t, b, "block should be retrievable by hash")
}

func TestMetaRpc_Client_Accounts(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	accountList, err := conn.GetAccounts(ctx)
	require.NoError(t, err)
	require.Equal(t, accountList[0], eth.Address(BOB))
	require.Equal(t, accountList[1], eth.Address(ALICE))
}

func TestMetaRpc_Client_NetVersion(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	netVersion, err := conn.NetVersion(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, netVersion, "net version id must not be nil")
	require.Equal(t, netVersion, "1133")
}

func TestMetaRpc_Client_ChainId(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	chainId, err := conn.ChainId(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, chainId, "chain id must not be nil")
	require.Equal(t, chainId, "0x46d") // 1133
}

func TestMetaRpc_Execute_Call(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	tx := eth.Transaction{
		From: *eth.MustAddress(ALICE),
		// pragma solidity ^0.8.2;
		// contract Counter {
		// 	string public name = 'Counter';
		// 	address public owner;
		// 	uint256 count = 0;
		// 	event echo(string message);
		// 	constructor() {
		// 		owner = msg.sender;
		// 		emit echo('Hello, Counter');
		// 	}
		// 	modifier onlyOwner() {
		// 		require(msg.sender == owner);
		// 		_;
		// 	}
		// 	function mul(uint256 a, uint256 b) public pure returns (uint256) {
		// 		return a * b;
		// 	}
		// 	function max10(uint256 a) public pure returns (uint256) {
		// 		if (a > 10) revert('Value must not be greater than 10.');
		// 		return a;
		// 	}
		// 	function getCount() public view returns (uint256) {
		// 		return count;
		// 	}
		// 	function setCount(uint256 _count) public onlyOwner {
		// 		count = _count;
		// 	}
		// 	function incr() public {
		// 		count += 1;
		// 	}
		// 	function getBlockHash(uint256 number) public view returns (bytes32) {
		// 		return blockhash(number);
		// 	}
		// 	function getCurrentBlock() public view returns (uint256) {
		// 		return block.number;
		// 	}
		// 	function getGasLimit() public view returns (uint256) {
		// 		return block.gaslimit;
		// 	}
		// }
		To:    eth.MustAddress("0xaaf2bfd7b8cb16fe6f719f66280bf7292763c121"), // contract address
		Value: *eth.MustQuantity("0x00"),
		// mul(2,3)
		Input: *eth.MustData("0xc8a4ac9c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003"),
	}

	hash, err := conn.Call(ctx, tx, *eth.MustBlockNumberOrTag("latest"))
	println("hash: ", hash, err)
	require.NoError(t, err)
	require.Equal(t, hash, "0x0000000000000000000000000000000000000000000000000000000000000006")
}

func TestMetaRpc_Execute_EstimateGas(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	from := eth.MustAddress(ALICE)
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

// IGNORE(): rpc is not ready yet
// func TestMetaRpc_Fee_MaxPriorityFeePerGas(t *testing.T) {
// 	ctx := context.Background()
// 	conn := getMetaDevClient(t, ctx)

// 	fee, err := conn.MaxPriorityFeePerGas(ctx)
// 	require.NoError(t, err)
// 	require.NotEqual(t, fee, 0, "fee cannot be equal to 0")
// }

func TestMetaRpc_Fee_GasPrice(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	gasPrice, err := conn.GasPrice(ctx) // 10_000_000_000
	require.NoError(t, err)
	require.NotEqual(t, gasPrice, 0, "gas price cannot be equal to 0")
}

func TestMetaRpc_State_GetBalance(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	bal, err := conn.GetBalance(ctx, *eth.MustAddress(ALICE), *eth.MustBlockNumberOrTag("latest"))
	require.NoError(t, err)
	require.NotNil(t, bal)
}

func TestMetaRpc_State_GetTransactionCount(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	blockNum := eth.MustBlockNumberOrTag("latest")
	count, err := conn.GetTransactionCount(ctx, ALICE, *blockNum)
	require.NoError(t, err)
	require.NotNil(t, count, "must not be nil")

	// NOTE(canonbrother): return latest nonce of account, no latest block checked
	// Should catch failure since it is looking for a nonce of a future block
	// blockNum = eth.MustBlockNumberOrTag("0x7654321")
	// _, err = conn.GetTransactionCount(ctx, ALICE, *blockNum)
	// require.Error(t, err)
	// require.Empty(t, count, "must not exist since it is a future block")
}

func TestMetaRpc_Submit_SendRawTransactionInValidEmpty(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	txHash, err := conn.SendRawTransaction(ctx, "0x0")
	require.Error(t, err)
	require.Empty(t, txHash, "txHash must be nil")
}

// ERROR(): expect error
// func TestMetaRpc_Submit_SendRawTransactionInValidOldNonce(t *testing.T) {
// 	ctx := context.Background()
// 	conn := getMetaDevClient(t, ctx)

// 	data := eth.MustData("0x02f9015c8080852363e7f0008522ecb25c00870aa87bee5380008080b8fe608060405234801561001057600080fd5b5060df8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063165c4a1614602d575b600080fd5b603c6038366004605f565b604e565b60405190815260200160405180910390f35b600060588284607f565b9392505050565b600080604083850312156070578182fd5b50508035926020909101359150565b600081600019048311821515161560a457634e487b7160e01b81526011600452602481fd5b50029056fea2646970667358221220223df7833fd08eb1cd3ce363a9c4cb4619c1068a5f5517ea8bb862ed45d994f764736f6c63430008020033c080a048cbbd708f4f0db5f6232f1404cea1c5b305138038866de624d37f82d7477b95a063ef66a0dc703e98f6224110ea0917f29acd384038b3b2e731445316abfbc7e8")
// 	txHash, err := conn.SendRawTransaction(ctx, data.String())
// 	println("txHash: ", txHash)
// 	require.Error(t, err)
// 	require.Equal(t, err.Error(), "{\"code\":-32000,\"message\":\"nonce too low\"}")
// 	require.Empty(t, txHash, "txHash must be nil")
// }

func TestMetaRpc_Submit_SendRawTransaction(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

	data := eth.MustData("0x02f9015c8080852363e7f0008522ecb25c00870aa87bee5380008080b8fe608060405234801561001057600080fd5b5060df8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063165c4a1614602d575b600080fd5b603c6038366004605f565b604e565b60405190815260200160405180910390f35b600060588284607f565b9392505050565b600080604083850312156070578182fd5b50508035926020909101359150565b600081600019048311821515161560a457634e487b7160e01b81526011600452602481fd5b50029056fea2646970667358221220223df7833fd08eb1cd3ce363a9c4cb4619c1068a5f5517ea8bb862ed45d994f764736f6c63430008020033c080a048cbbd708f4f0db5f6232f1404cea1c5b305138038866de624d37f82d7477b95a063ef66a0dc703e98f6224110ea0917f29acd384038b3b2e731445316abfbc7e8")
	txHash, err := conn.SendRawTransaction(ctx, data.String())
	require.NoError(t, err, "sendRawTransaction should be no error")
	require.NotNil(t, txHash, "sendRawTransaction should not be null")
}

func TestMetaRpc_Transaction_GetTransactionByHash(t *testing.T) {
	ctx := context.Background()
	conn := getMetaDevClient(t, ctx)

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

	// IGNORE(): can't pre-obtain tx hash yet
	// tx, err = conn.TransactionByHash(ctx, "0xb0d129c0d84eef2db189f268d6510a3b24e51822c48e1a810fbd367ef8c1028c")
	// require.NoError(t, err, "early tx should not return an error")
	// require.NotNil(t, tx, "early tx should be retrievable by hash")
}
