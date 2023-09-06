package twitter

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

// 点赞推文
type FavoriteTwitter struct {
	Variables struct {
		TweetID string `json:"tweet_id"`
	} `json:"variables"`
	QueryID string `json:"queryId"`
}

// 转推
type CreateRetweetRequest struct {
	Variables struct {
		TweetID     string `json:"tweet_id"`
		DarkRequest bool   `json:"dark_request"`
	} `json:"variables"`
	QueryID string `json:"queryId"`
}

// 实现推文点赞
func FavoriteTweet(client http.Client, auth_token string, referer string) int {
	re := regexp.MustCompile(`status/(\d+)`)
	matches := re.FindStringSubmatch(referer)

	cookies, x_csrf_token, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return 500
	}

	data := FavoriteTwitter{
		Variables: struct {
			TweetID string `json:"tweet_id"`
		}{
			TweetID: matches[1],
		},
		QueryID: "lI07N6Otwv1PhnEgXILM7A", //该ID为twitter独有点赞标识
	}

	//将data json化
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to json data:", err.Error())
		return 500
	}

	newRequest, err := http.NewRequest("POST", "https://twitter.com/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet", bytes.NewReader(jsonData))
	if err != nil {
		log.Println("failed to create request:", err.Error())
		return 500
	}

	//设置请求头部
	newRequest.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	newRequest.Header.Set("accept-encoding", "gzip, deflate, br")
	newRequest.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	newRequest.Header.Set("cache-control", "no-cache")
	newRequest.Header.Set("no-cache", "no-cache")
	newRequest.Header.Set("pragma", "no-cache")
	newRequest.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	newRequest.Header.Set("upgrade-insecure-requests", "1")
	newRequest.Header.Set("x-twitter-active-user", "yes")
	newRequest.Header.Set("x-twitter-auth-type", "OAuth2Session")
	newRequest.Header.Set("x-twitter-client-language", "en")
	newRequest.Header.Set("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	newRequest.Header.Set("x-csrf-token", x_csrf_token)
	newRequest.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	newRequest.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	newRequest.Header.Set("referer", referer)
	newRequest.Header.Set("sec-fetch-site", "same-origin")
	newRequest.Header.Set("origin", "https://twitter.com")
	newRequest.Header.Set("content-type", "application/json")

	for _, cookie := range cookies {
		newRequest.AddCookie(cookie)
	}

	newRequest.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})
	response, err := client.Do(newRequest)
	if err != nil {
		return 500
	}

	//判断状态码
	if response.StatusCode != 200 {
		return 500
	}
	return 200
}
