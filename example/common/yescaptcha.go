package main

import (
	"fmt"
	"github.com/seencxy/lsr/common"
	"log"
	"net/http"
)

// yescapcha实例代码
func main() {

	clientKey := "5a1c753cc414fc4872650d2f066e88bea15d011125689"
	taskType := "HCaptchaTaskProxyless"
	websiteURL := "https://dashboard.alt.technology/flashlayers/altitude/create?usecase=generic&method=custom&tier=RESTAKING_TRIAL"
	websiteKey := "a7f3edce-9191-4d46-878e-e14fa24ab3e1"

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
