package api

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_QueryAvailablePhoneNumberCount(t *testing.T) {
	type fields struct {
		Client  http.Client
		Context context.Context
	}
	type args struct {
		apiKey   string
		country  string
		operator string
		serverId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test QueryAvailablePhoneNumberCount",
			fields: fields{
				Client:  http.Client{},
				Context: context.Background(),
			},
			args: args{
				apiKey:   "30A4b6ce0b5Ae964eec6f0758f8e723e",
				country:  "4",
				operator: "any",
				serverId: "ds",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  tt.fields.Client,
				Context: tt.fields.Context,
			}
			got, err := c.QueryAvailablePhoneNumberCount(tt.args.apiKey, tt.args.country, tt.args.operator, tt.args.serverId)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAvailablePhoneNumberCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("QueryAvailablePhoneNumberCount() got is empty")
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		newClient http.Client
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.newClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientWithContext(t *testing.T) {
	type args struct {
		newClient http.Client
		ctx       context.Context
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientWithContext(tt.args.newClient, tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_QueryTopCountriesByService(t *testing.T) {
	type fields struct {
		Client  http.Client
		Context context.Context
	}
	type args struct {
		apiKey    string
		server    string
		freePrice bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test QueryTopCountriesByService",
			fields: fields{
				Client:  http.Client{},
				Context: context.Background(),
			},
			args: args{
				apiKey:    "30A4b6ce0b5Ae964eec6f0758f8e723e",
				server:    "ds",
				freePrice: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  tt.fields.Client,
				Context: tt.fields.Context,
			}
			got, err := c.QueryTopCountriesByService(tt.args.apiKey, tt.args.server, tt.args.freePrice)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryTopCountriesByService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryTopCountriesByService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_QueryAccountBalance(t *testing.T) {
	type fields struct {
		Client  http.Client
		Context context.Context
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Test QueryAccountBalance",
			fields: fields{
				Client:  http.Client{},
				Context: context.Background(),
			},
			args: args{
				apiKey: "30A4b6ce0b5Ae964eec6f0758f8e723e",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  tt.fields.Client,
				Context: tt.fields.Context,
			}
			got, err := c.QueryAccountBalance(tt.args.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryAccountBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryAccountBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_QueryServiceByCountry(t *testing.T) {
	type fields struct {
		Client  http.Client
		Context context.Context
	}
	type args struct {
		apiKey  string
		country string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CountryServiceRes
		wantErr bool
	}{
		{
			name: "Test QueryServiceByCountry",
			fields: fields{
				Client:  http.Client{},
				Context: context.Background(),
			},
			args: args{
				apiKey:  "30A4b6ce0b5Ae964eec6f0758f8e723e",
				country: "4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  tt.fields.Client,
				Context: tt.fields.Context,
			}
			got, err := c.QueryServiceByCountry(tt.args.apiKey, tt.args.country)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryServiceByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryServiceByCountry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_QueryActiveActivations(t *testing.T) {
	type fields struct {
		Client  http.Client
		Context context.Context
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Activations
		wantErr bool
	}{
		{
			name: "Test QueryActiveActivations",
			fields: fields{
				Client:  http.Client{},
				Context: context.Background(),
			},
			args: args{
				apiKey: "30A4b6ce0b5Ae964eec6f0758f8e723e",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Client:  tt.fields.Client,
				Context: tt.fields.Context,
			}
			got, err := c.QueryActiveActivations(tt.args.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryActiveActivations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryActiveActivations() got = %v, want %v", got, tt.want)
			}
		})
	}
}
