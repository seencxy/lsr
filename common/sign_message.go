package common

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	sign "github.com/etaaa/Golang-Ethereum-Personal-Sign"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// EcdsaSignMessage 使用私钥对信息进行签名 返回哈希 椭圆曲线签名
func EcdsaSignMessage(prv string, message string) (string, error) {
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

// Eip712SignMessage 实现eip712签名
func Eip712SignMessage(typedData apitypes.TypedData, prv string) (string, error) {
	// 获取私钥
	privateKey, err := crypto.HexToECDSA(prv)
	if err != nil {
		return "", err
	}

	signature, err := signEIP712(typedData, privateKey)
	if err != nil {
		return "", err
	}

	sign := fmt.Sprintf("0x%x", signature)

	return sign, nil
}

func eip712Hash(typedData apitypes.TypedData) []byte {
	domainSeparator, _ := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	messageHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)

	return crypto.Keccak256(bytes.Join([][]byte{
		[]byte("\x19\x01"),
		domainSeparator[:],
		messageHash[:],
	}, nil))
}

func signEIP712(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := eip712Hash(typedData)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
