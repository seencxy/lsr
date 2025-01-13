# lsr

[![Go Report Card](https://goreportcard.com/badge/github.com/username/projectname)](https://github.com/seencxy/lsr)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/seencxy/lsr)
[![GoDoc](https://godoc.org/github.com/username/projectname?status.svg)](https://github.com/seencxy/lsr)

Go self-use auxiliary library

## Features

- X
- sms_activate
- galxe
- yescaptcha
- mailtm
- inject([fork](https://github.com/facebookarchive/inject))
- more...

## Table of Contents

- [Installation](#installation)
- [Usage](#using)
- [License](#license)

## Installation

Make sure you have [Go](https://golang.org/dl/) installed (version 1.23.x or later).

## Using `go get`

```bash
go get https://github.com/seencxy/lsr@latest
```

##### as a library:

```bash
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

```

### License

This project is licensed under the MIT License
