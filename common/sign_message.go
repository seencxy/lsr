package common

import (
	sign "github.com/etaaa/Golang-Ethereum-Personal-Sign"
	"github.com/ethereum/go-ethereum/crypto"
)

// 使用私钥对信息进行签名 返回哈希
func SignMessage(prv string, message string) (string, error) {
	// 将prv转化为ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return "", err
	}
	// 进行签名
	personalSign, err := sign.PersonalSign(message, privateKey)
	if err != nil {
		return "", err
	}

	return personalSign, nil
}
