package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/seencxy/lsr/common"
	"io"
	"net/http"
	"strings"
	"time"
)

// 实现galxe登录签名 获取token
func GalxeLoginSign(client http.Client, address string, prv string) (string, error) {
	// 下面是原始签名 需要补充自己的签名
	messageSign := "galxe.com wants you to sign in with your Ethereum account:\n0x435479fBE9B19f9f4776f8Bf4B1c7d1f341572A1\n\nSign in with Ethereum to the app.\n\nURI: https://galxe.com\nVersion: 1\nChain ID: 1\nNonce: EIs5NZLaU2D2xi2Ev\nIssued At: 2023-08-26T06:49:41.277Z\nExpiration Time: 2023-09-02T06:49:41.065Z"
	// 将地址替换
	messageSign = strings.Replace(messageSign, "0x435479fBE9B19f9f4776f8Bf4B1c7d1f341572A1", address, 1)
	// 替换nonce
	messageSign = strings.Replace(messageSign, "EIs5NZLaU2D2xi2Ev", common.GenerateRandomString(17), 1)
	// 获取当前时间  替换签名时间
	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	messageSign = strings.Replace(messageSign, "2023-08-26T06:49:41.277Z", currentTime, 1)
	// 替换过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour).UTC().Format("2006-01-02T15:04:05.999Z")
	messageSign = strings.Replace(messageSign, "2023-09-02T06:49:41.065Z", expireTime, 1)

	personalSign, err := common.SignMessage(prv, messageSign)
	if err != nil {
		return "", err
	}

	data := GalxeGetTokenRequest{
		OperationName: "SignIn",
		Query:         "mutation SignIn($input: Auth) {\n  signin(input: $input)\n}\n",
		Variables: struct {
			Input struct {
				Address   string `json:"address"`
				Message   string `json:"message"`
				Signature string `json:"signature"`
			} `json:"input"`
		}{
			Input: struct {
				Address   string `json:"address"`
				Message   string `json:"message"`
				Signature string `json:"signature"`
			}{Address: address,
				Message:   messageSign,
				Signature: personalSign},
		},
	}

	// 将请求参数序列化
	marshal, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 创建请求
	request, err := http.NewRequest("POST", "https://graphigo.prd.galaxy.eco/query", bytes.NewReader(marshal))
	if err != nil {
		return "", err
	}
	// 增加头部
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://galxe.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	request.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")

	do, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// 代表请求成功
	if do.StatusCode == 200 {
		all, err := io.ReadAll(do.Body)
		if err != nil {
			return "", err
		}
		var res GalxeGetTokenResponse
		json.Unmarshal(all, &res)
		if err != nil {
			return "", err
		}

		return res.Data.Signin, nil
	}

	return "", errors.New("获取签名失败：请重试")
}

// galxe获取token请求体
type GalxeGetTokenRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			Address   string `json:"address"`
			Message   string `json:"message"`
			Signature string `json:"signature"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

// galxe获取token响应体
type GalxeGetTokenResponse struct {
	Data struct {
		Signin string `json:"signin"`
	} `json:"data"`
}
