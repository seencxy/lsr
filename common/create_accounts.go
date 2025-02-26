package common

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"strconv"
	"strings"
)

// Account 账号信息
type Account struct {
	Address string `json:"address"`
	Prv     string `json:"prv"`
}

// CreateOneAddress 创建单个以太坊地址
func CreateOneAddress() (prv, address string) {
	//创建私钥
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	priv := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", ""
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return priv, addr
}

// Mnemonic 创建助记词 /*
// number 创建账号的数量
func Mnemonic(number int) (mnemonic string, accounts []Account, err error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", nil, err
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return "", nil, err
	}

	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return "", nil, err
	}

	for i := 0; i < number; i++ {
		path := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		child, err := deriveChildKey(masterKey, path)
		if err != nil {
			return "", nil, err
		}

		privateKeyECDSA, err := crypto.ToECDSA(child.Key)
		if err != nil {
			return "", nil, err
		}

		address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).Hex()

		var account Account
		account.Address = address
		account.Prv = hex.EncodeToString(crypto.FromECDSA(privateKeyECDSA))

		accounts = append(accounts, account)
	}

	return mnemonic, accounts, err
}

// ByMnemonicCreateAccounts  通过助记词创建账号/*
// mnemonic 助记词
// startIndex 之前创建地址的数量
// additionalAddresses创建地址的数量
func ByMnemonicCreateAccounts(mnemonic string, startIndex, additionalAddresses int) (accounts []Account, err error) {
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	for i := startIndex; i < startIndex+additionalAddresses; i++ {
		path := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		child, err := deriveChildKey(masterKey, path)
		if err != nil {
			return nil, err
		}

		privateKeyECDSA, err := crypto.ToECDSA(child.Key)
		if err != nil {
			return nil, err
		}

		address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).Hex()

		var account Account
		account.Address = address
		account.Prv = hex.EncodeToString(crypto.FromECDSA(privateKeyECDSA))
		accounts = append(accounts, account)
	}

	return accounts, err
}

// 私有函数不供外部调用
func deriveChildKey(masterKey *bip32.Key, path string) (*bip32.Key, error) {
	segments := strings.Split(path, "/")
	key := masterKey
	for _, segment := range segments {
		if segment == "m" {
			continue
		}

		value, err := strconv.Atoi(strings.TrimRight(segment, "'"))
		if err != nil {
			return nil, err
		}

		if strings.HasSuffix(segment, "'") {
			value += int(bip32.FirstHardenedChild)
		}

		key, err = key.NewChildKey(uint32(value))
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}
