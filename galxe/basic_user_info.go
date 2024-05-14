package galxe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// 获取galxe用户信息
// 获取用户信息 主要是检测twitter和email是否绑定
// 返回四个参数 第一个是gid 第二个是是否绑定email 第三个是否绑定twitter  第四个是是否绑定dc
func BasicUserInfo(client http.Client, address string, auth string) (string, bool, bool, bool, error) {
	// 实例化请求参数
	data := BasicUserInfoRequest{
		OperationName: "BasicUserInfo",
		Variables: struct {
			Address        interface{} `json:"address"`
			ListSpaceInput struct {
				First int `json:"first"`
			} `json:"listSpaceInput"`
		}{
			Address: address,
			ListSpaceInput: struct {
				First int `json:"first"`
			}{First: 30},
		},
		Query: "query BasicUserInfo($address: String!, $listSpaceInput: ListSpaceInput!) {\n  addressInfo(address: $address) {\n    id\n    username\n    address\n    hasEmail\n    avatar\n    solanaAddress\n    aptosAddress\n    seiAddress\n    injectiveAddress\n    flowAddress\n    hasEvmAddress\n    hasSolanaAddress\n    hasAptosAddress\n    hasInjectiveAddress\n    hasFlowAddress\n    hasTwitter\n    hasGithub\n    hasDiscord\n    hasTelegram\n    displayEmail\n    displayTwitter\n    displayGithub\n    displayDiscord\n    displayTelegram\n    email\n    twitterUserID\n    twitterUserName\n    githubUserID\n    githubUserName\n    passport {\n      status\n      pendingRedactAt\n      id\n      __typename\n    }\n    isVerifiedTwitterOauth2\n    isVerifiedDiscordOauth2\n    displayNamePref\n    discordUserID\n    discordUserName\n    telegramUserID\n    telegramUserName\n    subscriptions\n    isWhitelisted\n    isInvited\n    isAdmin\n    passportPendingRedactAt\n    spaces(input: $listSpaceInput) {\n      list {\n        ...SpaceBasicFrag\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment SpaceBasicFrag on Space {\n  id\n  name\n  info\n  thumbnail\n  alias\n  links\n  isVerified\n  status\n  followersCount\n  __typename\n}\n",
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return "", false, false, false, err
	}

	// 创建请求
	request, err := http.NewRequest("POST", "https://graphigo.prd.galaxy.eco/query", bytes.NewReader(marshal))
	if err != nil {
		return "", false, false, false, err
	}

	// 设置头部
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://galxe.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	request.Header.Add("Sec-Ch-Ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"Google Chrome\";v=\"116\"")
	request.Header.Add("Authorization", auth)

	all, err := client.Do(request)
	if err != nil {
		return "", false, false, false, err
	}

	if all.StatusCode != 200 {
		return "", false, false, false, err
	}

	readAll, err := io.ReadAll(all.Body)
	if err != nil {
		return "", false, false, false, err
	}

	var body BasicUserInfoResponse

	err = json.Unmarshal(readAll, &body)
	if err != nil {
		return "", false, false, false, err
	}

	return body.Data.AddressInfo.Id, body.Data.AddressInfo.HasEmail, body.Data.AddressInfo.HasTwitter, body.Data.AddressInfo.HasDiscord, nil
}

// BasicUserInfo请求体
type BasicUserInfoRequest struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		Address        interface{} `json:"address"`
		ListSpaceInput struct {
			First int `json:"first"`
		} `json:"listSpaceInput"`
	} `json:"variables"`
	Query string `json:"query"`
}

// BasicUserInfo响应体
type BasicUserInfoResponse struct {
	Data struct {
		AddressInfo struct {
			Id                  string `json:"id"`
			Username            string `json:"username"`
			Address             string `json:"address"`
			HasEmail            bool   `json:"hasEmail"`
			Avatar              string `json:"avatar"`
			SolanaAddress       string `json:"solanaAddress"`
			AptosAddress        string `json:"aptosAddress"`
			SeiAddress          string `json:"seiAddress"`
			InjectiveAddress    string `json:"injectiveAddress"`
			FlowAddress         string `json:"flowAddress"`
			HasEvmAddress       bool   `json:"hasEvmAddress"`
			HasSolanaAddress    bool   `json:"hasSolanaAddress"`
			HasAptosAddress     bool   `json:"hasAptosAddress"`
			HasInjectiveAddress bool   `json:"hasInjectiveAddress"`
			HasFlowAddress      bool   `json:"hasFlowAddress"`
			HasTwitter          bool   `json:"hasTwitter"`
			HasGithub           bool   `json:"hasGithub"`
			HasDiscord          bool   `json:"hasDiscord"`
			HasTelegram         bool   `json:"hasTelegram"`
			DisplayEmail        bool   `json:"displayEmail"`
			DisplayTwitter      bool   `json:"displayTwitter"`
			DisplayGithub       bool   `json:"displayGithub"`
			DisplayDiscord      bool   `json:"displayDiscord"`
			DisplayTelegram     bool   `json:"displayTelegram"`
			Email               string `json:"email"`
			TwitterUserID       string `json:"twitterUserID"`
			TwitterUserName     string `json:"twitterUserName"`
			GithubUserID        string `json:"githubUserID"`
			GithubUserName      string `json:"githubUserName"`
			Passport            struct {
				Status          string      `json:"status"`
				PendingRedactAt interface{} `json:"pendingRedactAt"`
				Id              string      `json:"id"`
				Typename        string      `json:"__typename"`
			} `json:"passport"`
			IsVerifiedTwitterOauth2 bool     `json:"isVerifiedTwitterOauth2"`
			IsVerifiedDiscordOauth2 bool     `json:"isVerifiedDiscordOauth2"`
			DisplayNamePref         string   `json:"displayNamePref"`
			DiscordUserID           string   `json:"discordUserID"`
			DiscordUserName         string   `json:"discordUserName"`
			TelegramUserID          string   `json:"telegramUserID"`
			TelegramUserName        string   `json:"telegramUserName"`
			Subscriptions           []string `json:"subscriptions"`
			IsWhitelisted           bool     `json:"isWhitelisted"`
			IsInvited               bool     `json:"isInvited"`
			IsAdmin                 bool     `json:"isAdmin"`
			PassportPendingRedactAt int      `json:"passportPendingRedactAt"`
			Spaces                  struct {
				List     []interface{} `json:"list"`
				Typename string        `json:"__typename"`
			} `json:"spaces"`
			Typename string `json:"__typename"`
		} `json:"addressInfo"`
	} `json:"data"`
}
