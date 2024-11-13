package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// VerifyTwitterAccountRequest 验证推特账号请求参数
type VerifyTwitterAccountRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			Address  interface{} `json:"address"`
			TweetURL interface{} `json:"tweetURL"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

// VerifyTwitterAccountResponse 验证推特账号响应参数
type VerifyTwitterAccountResponse struct {
	Data struct {
		VerifyTwitterAccount struct {
			Address         string `json:"address"`
			TwitterUserID   string `json:"twitterUserID"`
			TwitterUserName string `json:"twitterUserName"`
			Typename        string `json:"__typename"`
		} `json:"verifyTwitterAccount"`
	} `json:"data"`
}

// VerifyTwitterAccount 实现twitter账号的绑定验证
func VerifyTwitterAccount(client http.Client, address string, url string, auth string) (bool, error) {
	data := VerifyTwitterAccountRequest{
		OperationName: "VerifyTwitterAccount",
		Variables: struct {
			Input struct {
				Address  interface{} `json:"address"`
				TweetURL interface{} `json:"tweetURL"`
			} `json:"input"`
		}{
			Input: struct {
				Address  interface{} `json:"address"`
				TweetURL interface{} `json:"tweetURL"`
			}{
				Address:  address,
				TweetURL: url,
			},
		},
		Query: "mutation VerifyTwitterAccount($input: VerifyTwitterAccountInput!) {\n  verifyTwitterAccount(input: $input) {\n    address\n    twitterUserID\n    twitterUserName\n    __typename\n  }\n}\n",
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
	request.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	request.Header.Add("Authorization", auth)

	do, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if do.StatusCode == 200 {
		all, err := io.ReadAll(do.Body)
		if err != nil {
			return false, err
		}

		var result VerifyTwitterAccountResponse
		err = json.Unmarshal(all, &result)
		if err != nil {
			return false, err
		}

		// 代表twitter绑定成功
		if len(result.Data.VerifyTwitterAccount.TwitterUserName) != 0 {
			return true, nil
		}
		return false, errors.New("验证推特失败: 请重试")
	}
	return false, errors.New("验证推特失败: 请重试")
}
