package main

import (
	"github.com/seencxy/lsr/common"
	"log"
)

// 生成随机数示例代码
func main() {
	// 假设我想要生成12个字符的随机数
	randomString := common.GenerateRandomString(12)
	log.Println(randomString)
}
