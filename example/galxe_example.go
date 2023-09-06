package main

import (
	"context"
	"github.com/seencxy/lsr/common"
	"github.com/seencxy/lsr/galxe"
	"github.com/seencxy/lsr/twitter"
	"log"
	"net/http"
)

func main() {
	address := ""
	prv := ""
	auth_token := ""
	followUser := "https://twitter.com/LLshuorong"
	post := ""
	dc_token := ""

	// 生成客户端
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	// 获取网站登录签名
	sign, err := galxe.GalxeLoginSign(client, address, prv)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 获取账号信息gid has_email has_twitter
	gid, has_email, has_twitter, has_discord, err := galxe.BasicUserInfo(client, address, sign)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(gid, has_email, has_twitter, has_discord)

	// 在这里进行判断如果未绑定邮箱则绑定邮箱
	if has_email == false {
		load, err := galxe.GetGeetestV4CaptchaLoad(client)
		if err != nil {
			log.Println(err)
			return
		}

		verify, err := galxe.GetGeetestV4CaptchaVerify(client, load, "w")
		if err != nil {
			log.Println(err)
			return
		}

		// 在这里创建一个邮箱
		email, messageChan, err := common.CreateEmail(context.TODO())
		if err != nil {
			log.Println(err)
			return
		}

		code, err := galxe.SendVerifyCode(client, verify, sign, email, address)
		if code == true && err == nil {
			log.Println("邮件发送成功...")
		}

		messages := <-messageChan
		message := messages.Intro[:6]
		log.Println("Received message:", message)

		bindEmail, err := galxe.BindEmail(client, address, email, message, sign)
		if bindEmail == true && err == nil {
			log.Println("邮件绑定成功...")
		}
	}

	// 如果推特绑定未绑定则绑定
	if has_twitter == false {
		//  绑定推特
		twitters, err := galxe.BindTwitter(client, auth_token, gid, address, sign)
		if err == nil && twitters == true {
			log.Println("推特绑定成功...")
		} else {
			log.Println(err)
			return
		}

		user := twitter.FollowUser(client, auth_token, followUser)
		if user == 200 {
			log.Println("关注用户成功...")
		}

		retweet := twitter.CreateRetweet(client, auth_token, post)
		if retweet == 200 {
			log.Println("转推帖子成功...")
		}

		tweet := twitter.FavoriteTweet(client, auth_token, post)
		if tweet == 200 {
			log.Println("喜欢帖子成功...")
		}
	}

	// 绑定discord
	if has_discord == false {
		discord, err := galxe.BindDiscord(client, address, dc_token, "cf_clearance", sign)
		if err == nil && discord == true {
			log.Println("discord绑定成功...")
		}
	}

	//log.Println(twitter, err)

}
