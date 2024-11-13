package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// BindEmailRequest 绑定邮箱请求参数
type BindEmailRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			Address          string `json:"address"`
			Email            string `json:"email"`
			VerificationCode string `json:"verificationCode"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

// BindEmailResponse 绑定邮箱请求参数
type BindEmailResponse struct {
	Data struct {
		UpdateEmail interface{} `json:"updateEmail"`
	} `json:"data"`
}

// BindEmail 绑定galxe邮箱
func BindEmail(client http.Client, address string, email string, code string, auth string) (bool, error) {
	// 实例化请求参数
	data := BindEmailRequest{
		OperationName: "UpdateEmail",
		Variables: struct {
			Input struct {
				Address          string `json:"address"`
				Email            string `json:"email"`
				VerificationCode string `json:"verificationCode"`
			} `json:"input"`
		}{
			Input: struct {
				Address          string `json:"address"`
				Email            string `json:"email"`
				VerificationCode string `json:"verificationCode"`
			}{
				Address:          address,
				Email:            email,
				VerificationCode: code,
			},
		},
		Query: "mutation UpdateEmail($input: UpdateEmailInput!) {\n  updateEmail(input: $input) {\n    code\n    message\n    __typename\n  }\n}\n",
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

	var body BindEmailResponse

	err = json.Unmarshal(readAll, &body)
	if err != nil {
		return false, err
	}

	if body.Data.UpdateEmail == nil {
		return true, nil
	}

	return false, errors.New("failed to send verify code...")

}
