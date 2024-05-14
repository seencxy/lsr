package main

import (
	"context"
	"github.com/seencxy/lsr/common"
	"log"
)

// common/create_email.go 示例代码
func main() {
	// 调用该函数需要一个上下文 这里使用空白上下文
	ctx := context.Background()
	// 调用函数 返回一个邮箱地址 以及这个邮箱信息channel
	address, messageChan, err := common.CreateEmail(ctx)
	if err != nil {
		log.Println(err)
	}

	// 打印邮箱
	log.Println(address)

	// 从邮箱信息channel中获取信息
	message := <-messageChan
	// 打印信息
	log.Println("Received message:", message)
}
