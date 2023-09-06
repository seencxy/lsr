package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// 验证推特请求参数
type VerifyDiscordRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			Address   string `json:"address"`
			Parameter string `json:"parameter"`
			Token     string `json:"token"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

func VerifyDiscord(client http.Client, auth string, address string, token string) (bool, error) {
	// Parse the URL
	parsedURL, err := url.Parse(token)
	if err != nil {
		return false, err
	}
	queryParams := parsedURL.Query()
	code := queryParams.Get("code")
	// 实例化请求数据
	data := VerifyDiscordRequest{
		OperationName: "VerifyDiscord",
		Variables: struct {
			Input struct {
				Address   string `json:"address"`
				Parameter string `json:"parameter"`
				Token     string `json:"token"`
			} `json:"input"`
		}{Input: struct {
			Address   string `json:"address"`
			Parameter string `json:"parameter"`
			Token     string `json:"token"`
		}{Address: address, Parameter: "", Token: code},
		},
		Query: "mutation VerifyDiscord($input: VerifyDiscordAccountInput!) {\n  verifyDiscordAccount(input: $input) {\n    address\n    discordUserID\n    discordUserName\n    __typename\n  }\n}\n",
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// 创建请求
	request, err := http.NewRequest("POST", "https://graphigo.prd.galaxy.eco/query", bytes.NewReader(marshal))
	if err != nil {
		return false, err
	}
	// 设置头部
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://galxe.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	request.Header.Add("Authorization", auth)

	do, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if do.StatusCode == 200 {
		return true, nil
	}
	return false, errors.New("验证dc失败: auth_token失效")
}
