package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// CreateTaskRequest 创建任务请求结构体
type CreateTaskRequest struct {
	ClientKey string `json:"clientKey"`
	Task      struct {
		WebsiteURL string `json:"websiteURL"`
		WebsiteKey string `json:"websiteKey"`
		Type       string `json:"type"`
	} `json:"task"`
}

// CreateTaskResponse 创建任务响应结构体
type CreateTaskResponse struct {
	TaskId  string `json:"taskId"`
	ErrorId int    `json:"errorId"`
}

type CaptchaClient struct {
	client    *http.Client
	clientKey string
	baseUrl   string
}

func New(_clientKey string) *CaptchaClient {
	// if _clientKey invalid,try get key form env
	if _clientKey == "" {
		_clientKey = os.Getenv("CaptchaKey")
	}
	return &CaptchaClient{
		client:    &http.Client{},
		clientKey: _clientKey,
	}
}

type Options func(*CaptchaClient)

func WithHttpClient(customClient *http.Client) Options {
	return func(c *CaptchaClient) {
		c.client = customClient
	}
}

func WithBaseUrl(url string) Options {
	return func(c *CaptchaClient) {
		c.baseUrl = url
	}
}

func NewWithOptions(_clientKey string, _options ...Options) *CaptchaClient {
	newClient := &CaptchaClient{
		client:    &http.Client{},
		clientKey: _clientKey,
		baseUrl:   "https://cn.yescaptcha.com/createTask",
	}
	for _, option := range _options {
		option(newClient)
	}

	return newClient
}

// CreateTask 创建验证码任务
func (c *CaptchaClient) CreateTask(websiteURL string, websiteKey string, kind string) (string, error) {

	//创建请求参数
	data := CreateTaskRequest{
		ClientKey: c.clientKey,
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
	request, err := http.NewRequest("POST", c.baseUrl+"createTask", bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}

	// 发送请求
	do, err := c.client.Do(request)
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

// PollingTaskRequest 查看任务状态请求体
type PollingTaskRequest struct {
	ClientKey string `json:"clientKey"`
	TaskId    string `json:"taskId"`
}

// PollingTaskResponse 查看任务状态响应体
type PollingTaskResponse struct {
	Solution struct {
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	} `json:"solution"`
	Status  string `json:"status"`
	ErrorId int    `json:"errorId"`
}

// GetResponse 查看任务状态
func (c *CaptchaClient) GetResponse(taskID string) (string, error) {
	// 请求节点
	for times := 0; times < 120; times += 3 {
		time.Sleep(time.Duration(3) * time.Second)

		data := map[string]interface{}{
			"clientKey": c.clientKey,
			"taskId":    taskID,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		request, err := http.NewRequest("POST", c.baseUrl+"getTaskResult", bytes.NewReader(jsonData))
		if err != nil {
			return "", err
		}
		request.Header.Set("Content-Type", "application/json")

		response, err := c.client.Do(request)
		if err != nil {
			_ = response.Body.Close()
			return "", err
		}

		var responseBody PollingTaskResponse
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			return "", err
		}

		if responseBody.Status == "ready" {
			_ = response.Body.Close()
			return responseBody.Solution.GRecaptchaResponse, nil
		}
	}

	return "", fmt.Errorf("polling task failed")
}

func (c *CaptchaClient) GetResponseWithChannel(taskID string, res chan string) error {
	// 请求节点
	for times := 0; times < 120; times += 3 {
		time.Sleep(time.Duration(3) * time.Second)

		data := map[string]interface{}{
			"clientKey": c.clientKey,
			"taskId":    taskID,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		request, err := http.NewRequest("POST", c.baseUrl+"getTaskResult", bytes.NewReader(jsonData))
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")

		response, err := c.client.Do(request)
		if err != nil {
			_ = response.Body.Close()
			return err
		}

		var responseBody PollingTaskResponse
		if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
			return err
		}

		if responseBody.Status == "ready" {
			_ = response.Body.Close()
			res <- responseBody.Solution.GRecaptchaResponse
			return nil
		}
	}

	return fmt.Errorf("polling task failed")
}
