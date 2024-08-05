package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"sync"
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	type args struct {
		ctx  context.Context
		port string
	}
	tests := []struct {
		name string
		args args
		want chan WebHookResponse
	}{
		{
			name: "Test CreateWebhook",
			args: args{
				ctx:  context.Background(),
				port: "8080",
			},
			want: make(chan WebHookResponse, channelLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channel, router := CreateWebhook(tt.args.ctx)
			// get data form channel
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				err := router.Run(":" + tt.args.port)
				if err != nil {
					wg.Done()
				}
			}()

			go func() {
				for {
					select {
					case data := <-channel:
						log.Println(data)
					}
				}
			}()

			wg.Wait()
			log.Println("Test CreateWebhook Done")
		})
	}
}

func TestCreateWebhook1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name  string
		args  args
		want  chan WebHookResponse
		want1 *gin.Engine
	}{
		{
			name: "Test CreateWebhook1",
			args: args{
				ctx: context.Background(),
			},
			want: make(chan WebHookResponse, channelLength),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateWebhook(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateWebhook() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateWebhook() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSendS(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Send",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TestSend()
		})
	}
}
