package main

import (
	"github.com/seencxy/lsr/inject"
	"log"
	"net/http"
)

type Application struct {
	Client *http.Client `inject:"http_client"`
	Name   string       `inject:"application_name"`
	// and more...
}

func (a *Application) Do(request *http.Request) error {
	_, err := a.Client.Do(request)
	return err
}

func (a *Application) Get() string {
	return a.Name
}

func main() {
	// Typically an application will have exactly one object Container, and
	// you will create it and use it within a main function:
	var g inject.Container
	var application Application
	if err := g.Provide(
		&inject.Object{Value: &application},
		&inject.Object{Value: http.DefaultClient, Name: "http_client"},
		&inject.Object{Value: "inject_example", Name: "application_name"},
	); err != nil {
		panic("Failed to provide:" + err.Error())
	}

	if err := g.Populate(); err != nil {
		panic("Failed to populate:" + err.Error())
	}

	// test request
	request, _ := http.NewRequest("GET", "https://www.baidu.com", nil)

	if err := application.Do(request); err != nil {
		panic(err)
	}

	log.Println(application.Get())
}
