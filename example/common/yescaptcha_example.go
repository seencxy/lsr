package main

import (
	"fmt"
	"github.com/seencxy/lsr/common"
	"log"
	"net/http"
)

// yescapcha实例代码
func main() {

	clientKey := ""
	taskType := "HCaptchaTaskProxyless"
	websiteURL := ""
	websiteKey := ""

	client := http.Client{}
	taskID, err := common.CreateTask(client, clientKey, websiteURL, websiteKey, taskType)
	if err != nil || len(taskID) == 0 {
		fmt.Println("Failed to create task:", err)
		return
	}
	fmt.Println("Created task:", taskID)

	response, err := common.GetResponse(client, taskID, clientKey)
	if err != nil || len(response) == 0 {
		fmt.Println("Failed to get response:", err)
		return
	}
	log.Println(response)
}
