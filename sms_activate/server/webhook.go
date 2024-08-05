package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	channelLength = 10000
)

// CreateWebhook Path: server/webhook.go
// Return sms code to the server
func CreateWebhook(ctx context.Context) (chan WebHookResponse, *gin.Engine) {
	route := gin.Default()
	messageChannel := make(chan WebHookResponse, channelLength)

	route.POST("/webhook", func(c *gin.Context) {
		var response WebHookResponse
		if err := c.ShouldBindJSON(&response); err != nil {
			return
		}
		messageChannel <- response
	})

	return messageChannel, route
}

type WebHookResponse struct {
	ActivationId int    `json:"activationId"`
	Service      string `json:"service"`
	Text         string `json:"text"`
	Code         string `json:"code"`
	Country      int    `json:"country"`
	ReceivedAt   string `json:"receivedAt"`
}

func TestSend() {
	url := "http://dd788c6.r17.cpolar.top/webhook"
	data := WebHookResponse{
		ActivationId: 1,
		Service:      "service",
		Text:         "text",
		Code:         "code",
		Country:      1,
		ReceivedAt:   "receivedAt",
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshal))
	if err != nil {
		return
	}

	http.DefaultClient.Do(request)
}
