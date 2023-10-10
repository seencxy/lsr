package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 验证签名示例代码
func main() {
	verify, err := common.EcdsaVerify("0x894130767d699786d034a13c6485fe64f0ef697c", "hello world", "0xe5e5202816c9769b3413d609be7508b78776560b07d90786fa1b2a1cf11f2d805d7d033625e93282d187911817c78449630da47503f09123bd2c7d1f4489b1ae1b")
	if err != nil {
		log.Println("failed to verify ecdsa sign:", err.Error())
		return
	}

	log.Println(verify)
}
