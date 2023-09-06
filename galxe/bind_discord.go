package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/andybalholm/brotli"
	"io/ioutil"
	"log"
	"net/http"
)

// 绑定discord账号请求体
type BindDiscordRequest struct {
	Permissions string `json:"permissions"`
	Authorize   bool   `json:"authorize"`
}

// BindDiscordErrorResponse
type BindDiscordErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// BindDiscordSuccessResponse
type BindDiscordSuccessResponse struct {
	Location string `json:"location"`
}

// 测试绑定discord 通过token
func BindDiscord(client http.Client, address string, dc_token string, cf_clearance string, auth string) (bool, error) {
	cookies := GetDiscord(client, dc_token, cf_clearance)

	// 创建请求
	data := BindDiscordRequest{
		Permissions: "0",
		Authorize:   true,
	}

	// 将数据序列化
	marshal, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// 初始URL
	url := "https://discord.com/api/v9/oauth2/authorize?client_id=947863296789323776&response_type=code&redirect_uri=https%3A%2F%2Fgalxe.com&scope=identify%20guilds%20guilds.members.read&state=Discord_Auth%3" + address
	referer := "https://discord.com/oauth2/authorize?client_id=947863296789323776&redirect_uri=https://galxe.com&response_type=code&scope=identify%20guilds%20guilds.members.read&prompt=consent&state=Discord_Auth;" + address

	// 创建请求
	request, err := http.NewRequest("POST", url, bytes.NewReader(marshal))
	if err != nil {
		return false, err
	}

	// 设置请求头部
	request.Header.Add("authorization", dc_token)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	request.Header.Add("accept", "*/*")
	request.Header.Add("content-Type", "application/json")
	request.Header.Add("accept-Encoding", "gzip, deflate, br")
	request.Header.Add("origin", "https://discord.com")
	request.Header.Add("referer", referer)
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("pragma", "no-cache")
	request.Header.Add("sec-ch-ua-mobile", "?0")

	request.Header.Add("x-Discord-Timezone", "Asia/Shanghai")
	request.Header.Add("x-Debug-Options", "bugReporterEnabled")
	request.Header.Add("x-Discord-Locale", "zh-CN")

	for _, v := range cookies {
		request.AddCookie(v)
	}

	request.AddCookie(&http.Cookie{Name: "cf_clearance", Value: cf_clearance})
	//PcM9dTLdK_sart_rsNQmES36wpN0pd0TPhALCa0XfrQ-1693492568-0-1-9f54b960.ad2fa180.c855e680-0.2.1693492568;
	//PcM9dTLdK_sart_rsNQmES36wpN0pd0TPhALCa0XfrQ-1693492568-0-1-9f54b960.ad2fa180.c855e680-0.2.1693492568;

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	reader := brotli.NewReader(response.Body)
	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return false, err
	}
	// 在这里做一个判断 如果status code为200则代表成功
	if response.StatusCode == 200 {
		var resp BindDiscordSuccessResponse

		if err := json.Unmarshal(bodyBytes, &resp); err != nil {
			return false, err
		}
		discord, err := VerifyDiscord(client, auth, address, resp.Location)
		if err == nil && discord == true {
			return true, nil
		}
		return false, errors.New("绑定dc失败: 绑定失败请重试")
	} else { // 代表失败获取失败原因
		return false, errors.New("绑定dc失败: auth_token失效")
	}
}

func GetDiscord(client http.Client, token string, cf_clearance string) []*http.Cookie {
	// 测试一下
	request, err := http.NewRequest("POST", "https://discord.com/channels/@me", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	// 设置请求头部
	request.Header.Add("authorization", token)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	request.Header.Add("accept", "*/*")
	request.Header.Add("content-Type", "application/json")
	request.Header.Add("accept-Encoding", "gzip, deflate, br")
	request.Header.Add("origin", "https://discord.com")
	request.Header.Add("referer", "https://discord.com")
	request.Header.Add("x-Discord-Timezone", "Asia/Shanghai")
	request.Header.Add("x-Debug-Options", "bugReporterEnabled")
	request.Header.Add("x-Discord-Locale", "zh-CN")
	request.Header.Add("sec-Fetch-Site", "same-origin")
	request.Header.Add("sec-Fetch-Mode", "cors")
	request.Header.Add("sec-Ch-Ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\"")

	request.AddCookie(&http.Cookie{Name: "cf_clearance", Value: cf_clearance})
	do, err := client.Do(request)
	if err != nil {
		return nil
	}

	return do.Cookies()

}
