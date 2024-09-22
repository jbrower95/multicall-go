package multicall

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

const ethNodeURL = "https://rpc.ankr.com/eth"

func setupClient(t *testing.T) *ethclient.Client {
	client, err := ethclient.Dial(ethNodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}

func TestMulticallClientCreation(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	// Test creating a new multicall client
	multicallClient, err := NewMulticallClient(context.Background(), client, nil)
	assert.NoError(t, err)
	assert.NotNil(t, multicallClient)
	assert.Equal(t, multicallClient.Address.Hex(), "0xcA11bde05977b3631167028862bE2a173976CA11") // Default multicall address
}

func TestMulticallGetBalance(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	// Create multicall client
	multicallClient, err := NewMulticallClient(context.Background(), client, nil)
	assert.NoError(t, err)
	assert.NotNil(t, multicallClient)

	// Sample Ethereum address (replace with a valid one for testing)
	testAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")

	// Get balance
	balanceCall := multicallClient.GetBalance(testAddress)
	assert.NotNil(t, balanceCall)

	// Execute the call
	balances, err := DoMany(*multicallClient, balanceCall)
	assert.NoError(t, err)
	assert.NotNil(t, balances)

	// Validate the balance (just check if it's non-nil for now)
	assert.NotNil(t, (*balances)[0])
}

func TestMulticallGetBlockNumber(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	// Create multicall client
	multicallClient, err := NewMulticallClient(context.Background(), client, nil)
	assert.NoError(t, err)
	assert.NotNil(t, multicallClient)

	// Get block number
	blockNumberCall := multicallClient.GetBlockNumber()
	assert.NotNil(t, blockNumberCall)

	// Execute the call
	blockNumbers, err := DoMany(*multicallClient, blockNumberCall, blockNumberCall)
	assert.NoError(t, err)
	assert.NotNil(t, blockNumbers)
	assert.Equal(t, (*blockNumbers)[0], (*blockNumbers)[1])

	// Validate the block number (should be greater than 0)
	blockNumber := (*blockNumbers)[0]
	assert.NotNil(t, blockNumber)
	assert.True(t, blockNumber.Cmp(big.NewInt(0)) > 0)
}

func TestMulticallCustomCall(t *testing.T) {
	client := setupClient(t)
	defer client.Close()

	// Create multicall client
	multicallClient, err := NewMulticallClient(context.Background(), client, nil)
	assert.NoError(t, err)
	assert.NotNil(t, multicallClient)

	// ERC721 abi
	contractAbi, _ := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"mint","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"constant":true,"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"}]`))
	contractAddress := common.HexToAddress("0x60E4d786628Fea6478F785A6d7e704777c86a7c6") // MAYC
	methodName := "balanceOf"

	// Create a custom call
	customCall, err := Describe[big.Int](contractAddress, contractAbi, methodName, common.HexToAddress("0x040881CA16C00358E9f02E2373310C2fDaC1a5b8"))
	assert.NoError(t, err)
	assert.NotNil(t, customCall)

	// Execute the custom call
	result, result2, err := Do(*multicallClient, customCall, customCall)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, result2)

	// Validate the result (for now, just check that it's non-nil)
	assert.NotNil(t, result)
}
