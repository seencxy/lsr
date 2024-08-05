package twitter

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

// CreateRetweet 转推
func CreateRetweet(client http.Client, auth_token string, referer string) int {
	// 获取twitter信息
	cookies, gid, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return 302
	}
	re := regexp.MustCompile(`status/(\d+)`)
	matches := re.FindStringSubmatch(referer)
	data := CreateRetweetRequest{
		Variables: struct {
			TweetID     string `json:"tweet_id"`
			DarkRequest bool   `json:"dark_request"`
		}{
			TweetID:     matches[1],
			DarkRequest: false,
		},
		QueryID: "ojPdsZsimiJrUGLR1sjUtA",
	}

	//将请求的参数json化
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to marshal data:", err.Error())
		return 500
	}

	//创建一个请求体
	request, err := http.NewRequest("POST", "https://x.com/i/api/graphql/ojPdsZsimiJrUGLR1sjUtA/CreateRetweet", bytes.NewReader(marshal))
	if err != nil {
		log.Println("failed to create request:", err.Error())
		return 500
	}

	//设置请求头部
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("no-cache", "no-cache")
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("x-twitter-active-user", "yes")
	request.Header.Set("x-twitter-auth-type", "OAuth2Session")
	request.Header.Set("x-twitter-client-language", "en")
	request.Header.Set("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	request.Header.Set("x-csrf-token", gid)
	request.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	request.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	request.Header.Set("referer", referer)
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("origin", "https://x.com")
	request.Header.Set("content-type", "application/json")

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})

	//发送请求
	response, err := client.Do(request)
	if err != nil {
		return 500
	}

	//判断状态码
	if response.Status != "200 OK" {
		return 500
	}
	return 200
}
