package common

import (
	"context"
	"github.com/felixstrobel/mailtm"
	"time"
)

var (
	err error
)

// 创建邮箱 使用的是mail.tm API
func CreateEmail(ctx context.Context) (string, <-chan mailtm.Message, error) {
	// 创建客户端
	client, err := mailtm.NewMailClient()
	if err != nil {
		return "", nil, err
	}

	// 获取域名
	client.GetDomains()

	// 使用客户端创建账号
	account, err := client.CreateAccount()
	if err != nil {
		return "", nil, err
	}

	// 获取client为账号生成的auth_token 在接收短信的时候需要
	_, err = client.GetAuthTokenCredentials(account.Address, account.Password)
	if err != nil {
		return "", nil, err
	}

	//  创建一个信息channel
	messageChan := make(chan mailtm.Message)

	// 使用协程获取信息
	go waitForMessage(client, messageChan)

	// 返回账号和信息channel
	return account.Address, messageChan, nil
}

// 从邮箱中获取信息
func waitForMessage(client *mailtm.MailClient, messageChan chan<- mailtm.Message) {
	//  定义信息结构体
	var messages []mailtm.Message
	for {
		// 获取一条邮件就返回给channel
		messages, err = client.GetMessages(1)
		if err != nil {
			continue
		}
		if len(messages) != 0 {
			messageChan <- messages[0]
			close(messageChan)
			break
		}
		time.Sleep(1 * time.Second)
	}
}
