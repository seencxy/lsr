package main

import (
	"fmt"
	"github.com/seencxy/lsr/common"
	"log"
)

func main() {
	// clientKey can from env
	clientKey := ""
	taskType := "HCaptchaTaskProxyless"
	websiteURL := ""
	websiteKey := ""

	// new instance
	client := common.New(clientKey)
	// or
	// client := common.NewWithOptions(clientKey, common.WithHttpClient(&http.Client{}), common.WithBaseUrl(""))
	taskID, err := client.CreateTask(websiteURL, websiteKey, taskType)
	if err != nil || len(taskID) == 0 {
		fmt.Println("Failed to create task:", err)
		return
	}
	fmt.Println("Created task:", taskID)

	response, err := client.GetResponse(taskID)
	if err != nil || len(response) == 0 {
		fmt.Println("Failed to get response:", err)
		return
	}

	// or
	//resChan := make(chan string, 1)
	//go func() {
	//	err = client.GetResponseWithChannel(taskID, resChan)
	//	if err != nil {
	//		// handle error
	//	}
	//}()
	//log.Println(<-resChan)
	log.Println(response)
}
