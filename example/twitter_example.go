package main

import (
	"github.com/seencxy/lsr/twitter"
	"log"
	"net/http"
)

func main() {
	// 生成客户端
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	message := twitter.CreateRetweetMessage(client, "this is lsr", "", "")
	if message == 200 {
		log.Println("评论成功...")
	}

	user := twitter.FollowUser(client, "", "")
	if user == 200 {
		log.Println("关注用户成功...")
	}

	retweet := twitter.CreateRetweet(client, "", "")
	if retweet == 200 {
		log.Println("转推帖子成功...")
	}

	tweet := twitter.FavoriteTweet(client, "", "")
	if tweet == 200 {
		log.Println("喜欢帖子成功...")
	}
}
