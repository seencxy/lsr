package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 签名示例代码
func main() {
	// 下面只是展示 必须要真实私钥 不然会报错
	prv := "b97b7a8179fe58c951889052fb5d244f940579bbf8f25a090d1d6db193ee86d5"
	// 下面是需要签名的信息
	message := "hello world"

	// 调用验证方法
	signMessage, err := common.EcdsaSignMessage(prv, message)
	if err != nil { // 代表签名报错
		log.Println(err)
		return
	}

	log.Println(signMessage)
}
