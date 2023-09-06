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

	message := twitter.CreateRetweetMessage(client, "this is lsr", "cd682646f2d10b5914817fbda078894b9acb4a74", "https://twitter.com/heyibinance/status/1652232554844413956")
	if message == 200 {
		log.Println("评论成功...")
	}

	user := twitter.FollowUser(client, "cd682646f2d10b5914817fbda078894b9acb4a74", "https://twitter.com/heyibinance")
	if user == 200 {
		log.Println("关注用户成功...")
	}

	retweet := twitter.CreateRetweet(client, "cd682646f2d10b5914817fbda078894b9acb4a74", "https://twitter.com/heyibinance/status/1652232554844413956")
	if retweet == 200 {
		log.Println("转推帖子成功...")
	}

	tweet := twitter.FavoriteTweet(client, "cd682646f2d10b5914817fbda078894b9acb4a74", "https://twitter.com/heyibinance/status/1652232554844413956")
	if tweet == 200 {
		log.Println("喜欢帖子成功...")
	}
}
