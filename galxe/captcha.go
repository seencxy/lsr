package galxe

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"time"
)

// 在galxe中绑定email以及领取奖励需要过captcha

// captcha load响应结构体
type CaptchaLoadResponse struct {
	Status string `json:"status"`
	Data   struct {
		LotNumber   string `json:"lot_number"`
		CaptchaType string `json:"captcha_type"`
		Js          string `json:"js"`
		Css         string `json:"css"`
		StaticPath  string `json:"static_path"`
		GctPath     string `json:"gct_path"`
		ShowVoice   bool   `json:"show_voice"`
		Feedback    string `json:"feedback"`
		Logo        bool   `json:"logo"`
		Pt          string `json:"pt"`
		CaptchaMode string `json:"captcha_mode"`
		Guard       bool   `json:"guard"`
		CheckDevice bool   `json:"check_device"`
		Language    string `json:"language"`
		CustomTheme struct {
			Style      string `json:"_style"`
			Color      string `json:"_color"`
			Gradient   string `json:"_gradient"`
			Hover      string `json:"_hover"`
			Brightness string `json:"_brightness"`
			Radius     string `json:"_radius"`
		} `json:"custom_theme"`
		PowDetail struct {
			Version  string    `json:"version"`
			Bits     int       `json:"bits"`
			Datetime time.Time `json:"datetime"`
			Hashfunc string    `json:"hashfunc"`
		} `json:"pow_detail"`
		Payload         string `json:"payload"`
		ProcessToken    string `json:"process_token"`
		PayloadProtocol int    `json:"payload_protocol"`
	} `json:"data"`
}

// captcha verify响应结构体
type CaptchaVerifyResponse struct {
	Status string `json:"status"`
	Data   struct {
		LotNumber string `json:"lot_number"`
		Result    string `json:"result"`
		Seccode   struct {
			CaptchaId     string `json:"captcha_id"`
			LotNumber     string `json:"lot_number"`
			PassToken     string `json:"pass_token"`
			GenTime       string `json:"gen_time"`
			CaptchaOutput string `json:"captcha_output"`
		} `json:"seccode"`
		Score           string `json:"score"`
		Payload         string `json:"payload"`
		ProcessToken    string `json:"process_token"`
		PayloadProtocol int    `json:"payload_protocol"`
	} `json:"data"`
}

// 获取galxe请求中所需要参数

// CaptchaLoad
func GetGeetestV4CaptchaLoad(client http.Client) (CaptchaLoadResponse, error) {
	// 生成uuid作为 challenge 参数
	challenge := uuid.New().String()
	// 生成callback
	callback := "geetest_" + fmt.Sprint(time.Now().UnixNano()/1e6)
	//初始化url
	url := fmt.Sprintf("https://gcaptcha4.geetest.com/load?captcha_id=244bcb8b9846215df5af4c624a750db4&challenge=%s&client_type=web&lang=zh-cn&callback=%s", challenge, callback)

	// 发送请求
	resp, err := client.Get(url)
	if err != nil {
		return CaptchaLoadResponse{}, err
	}

	if resp.StatusCode != 200 {
		return CaptchaLoadResponse{}, errors.New("failed to send request...")
	}

	// 读取响应参数
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return CaptchaLoadResponse{}, err
	}

	// 实例化对象
	var data CaptchaLoadResponse

	jsonStr := strings.TrimPrefix(string(all), callback+"(")
	jsonStr = strings.TrimSuffix(jsonStr, ")")

	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return CaptchaLoadResponse{}, err
	}
	return data, nil
}

// CaptchaVerify
func GetGeetestV4CaptchaVerify(client http.Client, params CaptchaLoadResponse, w string) (CaptchaVerifyResponse, error) {
	// 生成callback
	callback := "geetest_" + fmt.Sprint(time.Now().UnixNano()/1e6)
	// 初始化url
	url := fmt.Sprintf("https://gcaptcha4.geetest.com/verify?callback=%s&captcha_id=244bcb8b9846215df5af4c624a750db4&client_type=web&lot_number=%s&payload=%s&process_token=%s&payload_protocol=%d&pt=%s&w=%s", callback, params.Data.LotNumber, params.Data.Payload, params.Data.ProcessToken, params.Data.PayloadProtocol, params.Data.Pt, w)
	// 发送请求
	resp, err := client.Get(url)
	if err != nil {
		return CaptchaVerifyResponse{}, err
	}
	if resp.StatusCode != 200 {
		return CaptchaVerifyResponse{}, errors.New("failed to send request...")
	}

	// 读取响应参数
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return CaptchaVerifyResponse{}, err
	}

	// 实例化对象
	var data CaptchaVerifyResponse

	jsonStr := strings.TrimPrefix(string(all), callback+"(")
	jsonStr = strings.TrimSuffix(jsonStr, ")")

	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return CaptchaVerifyResponse{}, err
	}
	return data, nil
}
