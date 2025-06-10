package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"sync/atomic"
	"time"
)

type TransactionClient struct {
	rpcURL  string
	Client  atomic.Value
	ChainId *big.Int
}

// NewTransactionClient 创建客户端，并启动自动重连
func NewTransactionClient(rpc string) (*TransactionClient, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	tc := &TransactionClient{
		rpcURL:  rpc,
		ChainId: chainId,
	}
	tc.Client.Store(client)

	// 启动自动重连协程
	go tc.StartAutoReconnect()

	return tc, nil
}

// StartAutoReconnect 定时检查连接状态，并在连接断开时自动重连
func (t *TransactionClient) StartAutoReconnect() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c := t.Client.Load()
		if c == nil {
			t.reconnect()
			continue
		}

		// 发送一个简单的请求来检查 RPC 是否可用
		_, err := c.(*ethclient.Client).ChainID(context.Background())
		if err != nil {
			t.Client.Store(&ethclient.Client{})
		}
	}
}

// 重新连接
func (tc *TransactionClient) reconnect() {
	client, err := ethclient.Dial(tc.rpcURL)
	if err != nil {
		log.Println("重连失败:", err)
		return
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("获取 ChainID 失败:", err)
		return
	}

	tc.Client.Store(client)
	tc.ChainId = chainId
	log.Println("RPC 重新连接成功")
}

// GetClient 获取当前的 ethclient
func (tc *TransactionClient) GetClient() *ethclient.Client {
	c := tc.Client.Load()
	if c == nil {
		return nil
	}
	return c.(*ethclient.Client)
}

type TransferOption struct {
	GasLimit               uint64
	GasPrice               *big.Int
	Nonce                  uint64
	Ctx                    context.Context
	WaitTransactionReceipt bool
}

type TransferOptionFunc func(*TransferOption)

func WithTransferGasLimit(gasLimit uint64) TransferOptionFunc {
	return func(option *TransferOption) {
		option.GasLimit = gasLimit
	}
}

func WithTransferGasPrice(gasPrice *big.Int) TransferOptionFunc {
	return func(option *TransferOption) {
		option.GasPrice = gasPrice
	}
}

func WithTransferNonce(nonce uint64) TransferOptionFunc {
	return func(option *TransferOption) {
		option.Nonce = nonce
	}
}

func WithTransferCtx(ctx context.Context) TransferOptionFunc {
	return func(option *TransferOption) {
		option.Ctx = ctx
	}
}

func WithTransferWaitTransactionReceipt(wait bool) TransferOptionFunc {
	return func(option *TransferOption) {
		option.WaitTransactionReceipt = wait
	}
}

func (t *TransactionClient) Transfer(prv string, toAddress string, value *big.Int, opts ...TransferOptionFunc) (string, error) {
	param := TransferOption{
		GasLimit:               21000,
		GasPrice:               big.NewInt(0),
		Nonce:                  0,
		Ctx:                    context.Background(),
		WaitTransactionReceipt: false,
	}
	for _, opt := range opts {
		opt(&param)
	}

	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return "", err
	}

	if param.Nonce == 0 {
		nonce, err := t.GetClient().PendingNonceAt(param.Ctx, crypto.PubkeyToAddress(privateKey.PublicKey))
		if err != nil {
			return "", err
		}
		param.Nonce = nonce
	}

	if param.GasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err := t.GetClient().SuggestGasPrice(param.Ctx)
		if err != nil {
			return "", err
		}
		param.GasPrice = gasPrice
	}

	toAddressHex := common.HexToAddress(toAddress)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    param.Nonce,
		To:       &toAddressHex,
		Value:    value,
		Gas:      param.GasLimit,
		GasPrice: param.GasPrice,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(t.ChainId), privateKey)
	if err != nil {
		return "", err
	}

	if err = t.GetClient().SendTransaction(param.Ctx, signedTx); err != nil {
		return "", err
	}

	txHash := signedTx.Hash()

	if !param.WaitTransactionReceipt {
		return txHash.Hex(), nil
	}

	maxRetries := 300
	retryCount := 0
	for retryCount < maxRetries {
		receipt, err := t.GetClient().TransactionReceipt(param.Ctx, txHash)
		if err != nil {
			continue
		} else if receipt != nil {
			switch receipt.Status {
			case types.ReceiptStatusSuccessful:
				return txHash.Hex(), nil
			case types.ReceiptStatusFailed:
				return "", errors.New("transaction execution failed")
			}
		}
		retryCount++
	}

	return txHash.Hex(), nil
}

type SendTransactionByAbiOptionFunc func(*SendTransactionByAbiOption)
type SendTransactionByAbiOption struct {
	GasLimit               uint64
	GasPrice               *big.Int
	Nonce                  uint64
	Ctx                    context.Context
	Value                  *big.Int
	WaitTransactionReceipt bool
	Data                   []byte
}

func WithSendTransactionByAbiGasLimit(gasLimit uint64) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.GasLimit = gasLimit
	}
}

func WithSendTransactionByAbiGasPrice(gasPrice *big.Int) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.GasPrice = gasPrice
	}
}

func WithSendTransactionByAbiNonce(nonce uint64) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.Nonce = nonce
	}
}

func WithSendTransactionByAbiCtx(ctx context.Context) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.Ctx = ctx
	}
}

func WithSendTransactionByAbiValue(value *big.Int) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.Value = value
	}
}

func WithSendTransactionByAbiWaitTransactionReceipt(wait bool) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.WaitTransactionReceipt = wait
	}
}

func WithSendTransactionByAbiData(data []byte) SendTransactionByAbiOptionFunc {
	return func(option *SendTransactionByAbiOption) {
		option.Data = data
	}
}

func (t *TransactionClient) SendTransactionByAbi(prv string, contractAddr string, contractAbi string, method string, args []interface{}, opts ...SendTransactionByAbiOptionFunc) (string, error) {
	param := SendTransactionByAbiOption{
		GasLimit: 0,
		GasPrice: big.NewInt(0),
		Nonce:    0,
		Ctx:      context.Background(),
		Value:    big.NewInt(0),
		Data:     nil,
	}
	for _, opt := range opts {
		opt(&param)
	}

	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return "", err
	}

	if param.Nonce == 0 {
		nonce, err := t.GetClient().PendingNonceAt(param.Ctx, crypto.PubkeyToAddress(privateKey.PublicKey))
		if err != nil {
			return "", err
		}
		param.Nonce = nonce
	}

	if param.GasPrice.Cmp(big.NewInt(0)) == 0 {
		gasPrice, err := t.GetClient().SuggestGasPrice(param.Ctx)
		if err != nil {
			return "", err
		}
		param.GasPrice = gasPrice
	}
	data := make([]byte, 0)
	if param.Data != nil {
		data = param.Data
	} else {
		parsedAbi, err := abi.JSON(strings.NewReader(contractAbi))
		if err != nil {
			return "", err
		}
		data, err = parsedAbi.Pack(method, args...)
		if err != nil {
			return "", err
		}
	}

	// 转换合约地址
	toAddress := common.HexToAddress(contractAddr)

	// 如果未指定 gasLimit，则预估
	if param.GasLimit == 0 {
		// 构造预估调用的 CallMsg
		callMsg := ethereum.CallMsg{
			From:     crypto.PubkeyToAddress(privateKey.PublicKey),
			To:       &toAddress,
			Gas:      0, // 不设置具体值，由节点预估
			GasPrice: param.GasPrice,
			Value:    param.Value,
			Data:     data,
		}
		gasLimit, err := t.GetClient().EstimateGas(param.Ctx, callMsg)
		if err != nil {
			return "", err
		}
		param.GasLimit = gasLimit
	}

	// 构造交易对象
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    param.Nonce,
		To:       &toAddress,
		Value:    param.Value,
		Gas:      param.GasLimit,
		GasPrice: param.GasPrice,
		Data:     data,
	})
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(t.ChainId), privateKey)
	if err != nil {
		return "", err
	}

	// 发送交易
	if err = t.GetClient().SendTransaction(param.Ctx, signedTx); err != nil {
		return "", err
	}

	txHash := signedTx.Hash()

	// 如果不等待交易回执，直接返回 txHash
	if !param.WaitTransactionReceipt {
		return txHash.Hex(), nil
	}

	// 等待交易回执
	maxRetries := 300
	retryCount := 0
	for retryCount < maxRetries {
		receipt, err := t.GetClient().TransactionReceipt(param.Ctx, txHash)
		if err != nil {
			// 若出现错误则跳过，继续等待
			retryCount++
			continue
		} else if receipt != nil {
			switch receipt.Status {
			case types.ReceiptStatusSuccessful:
				return txHash.Hex(), nil
			case types.ReceiptStatusFailed:
				return "", errors.New("transaction execution failed")
			}
		}
		retryCount++
	}

	return txHash.Hex(), nil
}

type GetMaxTransferableAmountOption struct {
	GasLimit      uint64
	ReserveAmount uint64
}

type GetMaxTransferableAmountOptionFunc func(*GetMaxTransferableAmountOption)

func WithGetMaxTransferableAmountGasLimit(gasLimit uint64) GetMaxTransferableAmountOptionFunc {
	return func(option *GetMaxTransferableAmountOption) {
		option.GasLimit = gasLimit
	}
}

func WithGetMaxTransferableAmountReserve(reserve uint64) GetMaxTransferableAmountOptionFunc {
	return func(option *GetMaxTransferableAmountOption) {
		option.ReserveAmount = reserve
	}
}

func (t *TransactionClient) GetMaxTransferableAmount(address common.Address, opts ...GetMaxTransferableAmountOptionFunc) (*big.Int, error) {
	// 默认配置：GasLimit 为 21000，ReserveAmount 默认为 0（单位：wei）
	param := GetMaxTransferableAmountOption{
		GasLimit:      21000,
		ReserveAmount: 0,
	}

	for _, opt := range opts {
		opt(&param)
	}

	balance, err := t.GetClient().BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	gasPrice, err := t.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(param.GasLimit)))

	available := new(big.Int).Sub(balance, gasCost)

	reserve := big.NewInt(int64(param.ReserveAmount))
	if available.Cmp(reserve) <= 0 {
		return nil, fmt.Errorf("余额不足以支付 Gas 费用和保留金额")
	}

	transferable := new(big.Int).Sub(available, reserve)
	return transferable, nil
}
