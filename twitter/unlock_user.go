package twitter

import (
	"bytes"
	"net/http"
	"net/url"
)

// UnlockUser 取关用户
// 实现twitter取关账号
func UnlockUser(client http.Client, auth_token string, referer string) int {
	cookies, gid, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return 500
	}
	uid, err := GetUserInfo(client, auth_token, referer)
	if err != nil {
		return 500
	}

	//创建数据
	data := url.Values{}
	data.Set("include_profile_interstitial_type", "1")
	data.Set("include_blocking", "1")
	data.Set("include_blocked_by", "1")
	data.Set("include_followed_by", "1")
	data.Set("include_want_retweets", "1")
	data.Set("include_mute_edge", "1")
	data.Set("include_can_dm", "1")
	data.Set("include_can_media_tag", "1")
	data.Set("include_ext_has_nft_avatar", "1")
	data.Set("include_ext_is_blue_verified", "1")
	data.Set("include_ext_verified_type", "1")
	data.Set("include_ext_profile_image_shape", "1")
	data.Set("skip_status", "1")
	data.Set("user_id", uid)

	newRequest, err := http.NewRequest("POST", "https://x.com/i/api/1.1/friendships/destroy.json", bytes.NewBufferString(data.Encode()))
	if err != nil {
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
	newRequest.Header.Set("x-csrf-token", gid)
	newRequest.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	newRequest.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	newRequest.Header.Set("referer", referer)
	newRequest.Header.Set("sec-fetch-site", "same-origin")
	newRequest.Header.Set("origin", "https://x.com")
	newRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
