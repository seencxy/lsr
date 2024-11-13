package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// SendVerifyCodeRequest 发送邮箱验证码服务请求体
type SendVerifyCodeRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			Address string `json:"address"`
			Email   string `json:"email"`
			Captcha struct {
				LotNumber     string `json:"lotNumber"`
				CaptchaOutput string `json:"captchaOutput"`
				PassToken     string `json:"passToken"`
				GenTime       string `json:"genTime"`
			} `json:"captcha"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

// SendVerifyCodeResponse 发送邮箱验证码服务响应体
type SendVerifyCodeResponse struct {
	Data struct {
		SendVerificationCode interface{} `json:"sendVerificationCode"`
	} `json:"data"`
}

// SendVerifyCode 用于邮箱验证时发送邮箱
func SendVerifyCode(client http.Client, params CaptchaVerifyResponse, auth string, email string, address string) (bool, error) {
	// 实例化请求参数
	data := SendVerifyCodeRequest{
		OperationName: "SendVerifyCode",
		Variables: struct {
			Input struct {
				Address string `json:"address"`
				Email   string `json:"email"`
				Captcha struct {
					LotNumber     string `json:"lotNumber"`
					CaptchaOutput string `json:"captchaOutput"`
					PassToken     string `json:"passToken"`
					GenTime       string `json:"genTime"`
				} `json:"captcha"`
			} `json:"input"`
		}{
			Input: struct {
				Address string `json:"address"`
				Email   string `json:"email"`
				Captcha struct {
					LotNumber     string `json:"lotNumber"`
					CaptchaOutput string `json:"captchaOutput"`
					PassToken     string `json:"passToken"`
					GenTime       string `json:"genTime"`
				} `json:"captcha"`
			}{
				Address: address,
				Email:   email,
				Captcha: struct {
					LotNumber     string `json:"lotNumber"`
					CaptchaOutput string `json:"captchaOutput"`
					PassToken     string `json:"passToken"`
					GenTime       string `json:"genTime"`
				}{
					LotNumber:     params.Data.LotNumber,
					CaptchaOutput: params.Data.Seccode.CaptchaOutput,
					PassToken:     params.Data.Seccode.PassToken,
					GenTime:       params.Data.Seccode.GenTime,
				},
			},
		},
		Query: "mutation SendVerifyCode($input: SendVerificationEmailInput!) {\n  sendVerificationCode(input: $input) {\n    code\n    message\n    __typename\n  }\n}\n",
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

	all, err := client.Do(request)
	if err != nil {
		return false, err
	}

	if all.StatusCode != 200 {
		return false, nil
	}

	readAll, err := io.ReadAll(all.Body)
	if err != nil {
		return false, err
	}

	var body SendVerifyCodeResponse

	err = json.Unmarshal(readAll, &body)
	if err != nil {
		return false, err
	}

	if body.Data.SendVerificationCode == nil {
		return true, nil
	}

	return false, errors.New("failed to send verify code...")
}
