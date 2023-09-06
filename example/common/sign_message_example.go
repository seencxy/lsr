package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 签名示例代码
func main() {
	// 下面只是展示 必须要真实私钥 不然会报错
	prv := "xxxxxxxxxxxxxxx"
	// 下面是需要签名的信息
	message := "this is only test..."

	signMessage, err := common.SignMessage(prv, message)
	if err != nil { // 代表签名报错
		log.Println(err)
		return
	}

	log.Println(signMessage)
}
