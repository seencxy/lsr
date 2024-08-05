package api

import (
	"context"
	"fmt"
	"net/http"
	"sms-activate/pkg"
	"strings"
)

const (
	baseUrl = "https://api.sms-activate.org/stubs/handler_api.php"
)

type ServicePriceStruct struct {
	RetailPrice  float64                `json:"retail_price"`
	Country      int64                  `json:"country"`
	FreePriceMap map[string]interface{} `json:"freePriceMap"`
	Price        interface{}            `json:"price"`
	Count        int64                  `json:"count"`
}

type Client struct {
	http.Client
	context.Context
}

// NewClient creates a new client with the provided http.Client
func NewClient(newClient http.Client) *Client {
	return &Client{
		Client:  newClient,
		Context: context.Background(),
	}
}

// NewClientWithContext creates a new client with the provided http.Client and context
func NewClientWithContext(newClient http.Client, ctx context.Context) *Client {
	return &Client{
		Client:  newClient,
		Context: ctx,
	}
}

// QueryAvailablePhoneNumberCount returns the number of available phone numbers for the specified country and operator.
// apiKey: SMS Activate API key.
// country: Country code.
// operator: Operator code.
func (c *Client) QueryAvailablePhoneNumberCount(apiKey string, country string, operator string, serverId string) (map[string]interface{}, error) {
	// If the operator is not specified, set it to "any".
	if len(operator) == 0 {
		operator = "any"
	}

	req, err := http.NewRequest(http.MethodGet, baseUrl+fmt.Sprintf("?api_key=%s&action=getNumbersStatus&country=%s&operator=%s", apiKey, country, operator), nil)
	if err != nil {
		return nil, err
	}

	serverMap := make(map[string]interface{})
	err = pkg.HandRequestUnmarshalData(c.Client, c.Context, req, &serverMap)
	if err != nil {
		return nil, err
	}

	// Check if the serverId is in the map.
	returnMap := make(map[string]interface{})
	for key, value := range serverMap {
		if strings.Contains(key, serverId) {
			returnMap[key] = value
		}
	}

	return returnMap, nil
}

// QueryTopCountriesByService returns the top countries by service.
// apiKey: SMS Activate API key.
// server: Server code.
// freePrice: Show free price.
func (c *Client) QueryTopCountriesByService(apiKey string, server string, freePrice bool) (map[string]*ServicePriceStruct, error) {
	req, err := http.NewRequest(http.MethodGet, baseUrl+fmt.Sprintf("?api_key=%s&action=getTopCountriesByService&service=%s&freePrice=%v", apiKey, server, freePrice), nil)
	if err != nil {
		return nil, err
	}

	serverMap := make(map[string]*ServicePriceStruct)
	err = pkg.HandRequestUnmarshalData(c.Client, c.Context, req, &serverMap)
	if err != nil {
		return nil, err
	}

	return serverMap, nil
}

// QueryAccountBalance returns the account balance.
// apiKey: SMS Activate API key.
func (c *Client) QueryAccountBalance(apiKey string) (interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, baseUrl+fmt.Sprintf("?api_key=%s&action=getBalance", apiKey), nil)
	if err != nil {
		return nil, err
	}

	resBodyString, err := pkg.HandRequestData(c.Client, c.Context, req)
	if err != nil {
		return nil, err
	}

	// Check if the response body string contains "ACCESS_BALANCE".
	if !strings.Contains(resBodyString, "ACCESS_BALANCE") {
		return nil, fmt.Errorf("response body string not contains ACCESS_NUMBER")
	}

	split := strings.Split(resBodyString, "ACCESS_BALANCE:")

	return split[1], nil
}

// QueryServiceByCountry returns the service by country.
// apiKey: SMS Activate API key.
// country: Country code.
func (c *Client) QueryServiceByCountry(apiKey string, country string) (*CountryServiceRes, error) {
	req, err := http.NewRequest(http.MethodGet, baseUrl+fmt.Sprintf("?api_key=%s&action=getOperators&country=%s", apiKey, country), nil)
	if err != nil {
		return nil, err
	}

	var res CountryServiceRes
	err = pkg.HandRequestUnmarshalData(c.Client, c.Context, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type CountryServiceRes struct {
	Status           string              `json:"status"`
	CountryOperators map[string][]string `json:"countryOperators"`
}

// QueryActiveActivations returns the active activations.
// apiKey: SMS Activate API key.
func (c *Client) QueryActiveActivations(apiKey string) (*Activations, error) {
	req, err := http.NewRequest(http.MethodGet, baseUrl+fmt.Sprintf("?api_key=%s&action=getActiveActivations", apiKey), nil)
	if err != nil {
		return nil, err
	}

	var callBackData Activations
	err = pkg.HandRequestUnmarshalData(c.Client, c.Context, req, &callBackData)
	if err != nil {
		return nil, err
	}

	return &callBackData, nil
}

type Activations struct {
	Status            string `json:"status"`
	ActiveActivations []struct {
		ActivationId     string   `json:"activationId"`
		ServiceCode      string   `json:"serviceCode"`
		PhoneNumber      string   `json:"phoneNumber"`
		ActivationCost   string   `json:"activationCost"`
		ActivationStatus string   `json:"activationStatus"`
		SmsCode          []string `json:"smsCode"`
		SmsText          string   `json:"smsText"`
		ActivationTime   string   `json:"activationTime"`
		Discount         string   `json:"discount"`
		Repeated         string   `json:"repeated"`
		CountryCode      string   `json:"countryCode"`
		CountryName      string   `json:"countryName"`
		CanGetAnotherSms string   `json:"canGetAnotherSms"`
	} `json:"activeActivations"`
}

// GetPhoneNumber returns the phone number.
// apiKey: SMS Activate API key.
// country: Country code.
// operator: Operator code.
// service: Service code.
func (c *Client) GetPhoneNumber(apiKey string, country string, operator string, service string) (*GetPhoneNumberResponse, error) {
	// request url
	url := baseUrl + fmt.Sprintf("?api_key=%s&action=getNumberV2&service=%s&operator=%s&country=%s", apiKey, service, operator, country)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var res GetPhoneNumberResponse
	err = pkg.HandRequestUnmarshalData(c.Client, c.Context, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetPhoneNumberResponse struct {
	ActivationId       string `json:"activationId"`
	PhoneNumber        string `json:"phoneNumber"`
	ActivationCost     string `json:"activationCost"`
	CountryCode        string `json:"countryCode"`
	CanGetAnotherSms   bool   `json:"canGetAnotherSms"`
	ActivationTime     string `json:"activationTime"`
	ActivationOperator string `json:"activationOperator"`
}
