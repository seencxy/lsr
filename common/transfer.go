package common

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

// Transfer 转账
func Transfer(address, prv string, rpc string, toAddress string, chainId int, value float64) (bool, error) {
	// 连接到 bnb 网络
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return false, err
	}
	defer client.Close()

	// 构造交易参数
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return false, err
	}

	//gas price
	// 设置 Gas 价格为 7 Gwei
	// 获取当前的gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	//// 增加到当前gas price的3倍
	//fastGasPrice := new(big.Int).Mul(gasPrice, big.NewInt(3))

	ToAddress := common.HexToAddress(toAddress)

	//设置gas限制
	gasLimit := uint64(21000)

	// 将 ETH 转换为 wei 的整数值
	weiAmountFloat := new(big.Float).Mul(big.NewFloat(value), big.NewFloat(1e18))

	// 将 weiAmountFloat 转换为 big.Int 类型
	weiAmount := new(big.Int)
	weiAmountFloat.Int(weiAmount)

	// 创建交易
	tx := types.NewTransaction(nonce, ToAddress, weiAmount, gasLimit, gasPrice, nil)

	// 签名交易
	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return false, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainId))), privateKey)
	if err != nil {
		return false, err
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, err
	}
	log.Println("Transfer: " + signedTx.Hash().Hex())
	txHash := signedTx.Hash()
	maxRetries := 3000
	retryCount := 0

	for retryCount < maxRetries {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			log.Println("查看transfer交易状态中:", err)
		} else if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
			return true, nil
		}

		time.Sleep(5 * time.Second)
		retryCount++
	}

	log.Println("Exceeded max retries, transaction might not have succeeded.")
	return false, err
}
