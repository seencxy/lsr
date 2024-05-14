package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 创建任务请求结构体
type CreateTaskRequest struct {
	ClientKey string `json:"clientKey"`
	Task      struct {
		WebsiteURL string `json:"websiteURL"`
		WebsiteKey string `json:"websiteKey"`
		Type       string `json:"type"`
	} `json:"task"`
}

// 创建任务响应结构体
type CreateTaskResponse struct {
	TaskId  string `json:"taskId"`
	ErrorId int    `json:"errorId"`
}

// CreateTask 创建验证码任务
func CreateTask(client http.Client, clientKey string, websiteURL string, websiteKey string, kind string) (string, error) {
	// 请求节点
	url := "https://cn.yescaptcha.com/createTask"

	//创建请求参数
	data := CreateTaskRequest{
		ClientKey: clientKey,
		Task: struct {
			WebsiteURL string `json:"websiteURL"`
			WebsiteKey string `json:"websiteKey"`
			Type       string `json:"type"`
		}{
			WebsiteURL: websiteURL,
			WebsiteKey: websiteKey,
			Type:       kind,
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 创建请求
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}

	// 发送请求
	do, err := client.Do(request)
	if err != nil {
		return "", err
	}

	if do.StatusCode != 200 {
		return "", err
	}

	//存储返回的结构体
	var response CreateTaskResponse

	// 读取请求内容
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(all, &response)
	if err != nil {
		return "", err
	}

	return response.TaskId, nil
}

// 查看任务状态请求体
type PollingTaskRequest struct {
	ClientKey string `json:"clientKey"`
	TaskId    string `json:"taskId"`
}

// 查看任务状态响应体
type PollingTaskResponse struct {
	Solution struct {
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	} `json:"solution"`
	Status  string `json:"status"`
	ErrorId int    `json:"errorId"`
}

// 查看任务状态
func GetResponse(client http.Client, taskID string, clientKey string) (string, error) {
	// 请求节点
	for times := 0; times < 120; times += 3 {
		time.Sleep(time.Duration(3) * time.Second)

		url := "https://cn.yescaptcha.com/getTaskResult"
		data := map[string]interface{}{
			"clientKey": clientKey,
			"taskId":    taskID,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
		if err != nil {
			return "", err
		}
		request.Header.Set("Content-Type", "application/json")

		response, err := client.Do(request)
		if err != nil {
			return "", err
		}
		defer response.Body.Close()

		var responseBody PollingTaskResponse
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			return "", err
		}

		if responseBody.Status == "ready" {
			return responseBody.Solution.GRecaptchaResponse, nil
		}
	}

	return "", fmt.Errorf("Polling task failed")
}
