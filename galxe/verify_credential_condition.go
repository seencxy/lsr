package galxe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// VerifyCredentialConditionRequest 验证完成请求参数
type VerifyCredentialConditionRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Input struct {
			CampaignId        string `json:"campaignId"`
			CredentialGroupId string `json:"credentialGroupId"`
			Address           string `json:"address"`
			ConditionIndex    int    `json:"conditionIndex"`
		} `json:"input"`
	} `json:"variables"`
	Query string `json:"query"`
}

// VerifyCredentialConditionResponse 验证完成响应参数
type VerifyCredentialConditionResponse struct {
	Data struct {
		VerifyCondition bool `json:"verifyCondition"`
	} `json:"data"`
}

// VerifyCredentialCondition 验证该活动的该列表是否完成
func VerifyCredentialCondition(client http.Client, auth string, CampaignId string, address string, ConditionIndex int, CredentialGroupId string) (bool, error) {
	data := VerifyCredentialConditionRequest{
		OperationName: "VerifyCredentialCondition",
		Query:         "mutation VerifyCredentialCondition($input: VerifyCredentialGroupConditionInput!) {\n  verifyCondition(input: $input)\n}\n",
		Variables: struct {
			Input struct {
				CampaignId        string `json:"campaignId"`
				CredentialGroupId string `json:"credentialGroupId"`
				Address           string `json:"address"`
				ConditionIndex    int    `json:"conditionIndex"`
			} `json:"input"`
		}{
			Input: struct {
				CampaignId        string `json:"campaignId"`
				CredentialGroupId string `json:"credentialGroupId"`
				Address           string `json:"address"`
				ConditionIndex    int    `json:"conditionIndex"`
			}{
				CampaignId:        CampaignId,
				CredentialGroupId: CredentialGroupId,
				Address:           address,
				ConditionIndex:    ConditionIndex,
			},
		},
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
		return false, err
	}

	readAll, err := io.ReadAll(all.Body)
	if err != nil {
		return false, err
	}

	var body VerifyCredentialConditionResponse

	err = json.Unmarshal(readAll, &body)
	if err != nil {
		return false, err
	}

	return body.Data.VerifyCondition, nil
}
