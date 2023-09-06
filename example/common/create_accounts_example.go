package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 创建代码实例代码
func main() {
	// 创建单个账号
	prv, addr, err := common.CreateOneAddress()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(prv, addr)
	// 创建助记词和钱包 10个地址
	mnemonic, accounts, err := common.Mnemonic(10)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(mnemonic)
	log.Println(accounts)

	// 通过已经有的助记词获取账号信息 从第五个开始获取 获取十个账号
	createAccounts, err := common.ByMnemonicCreateAccounts(mnemonic, 5, 10)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(createAccounts)
}
