package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 转账示例代码
func main() {
	// 转账代码 数据需真实填
	transfer, err := common.Transfer("address", "prv", "rpc", "toAddress", 1, 1)
	if err != nil || transfer == true {
		log.Println("转账失败...")
	}
}
