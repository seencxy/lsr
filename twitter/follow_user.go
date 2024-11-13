package twitter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// FollowUserResponse 查找用户信息
type FollowUserResponse struct {
	Data struct {
		User struct {
			Result struct {
				Typename                   string `json:"__typename"`
				Id                         string `json:"id"`
				RestId                     string `json:"rest_id"`
				AffiliatesHighlightedLabel struct {
				} `json:"affiliates_highlighted_label"`
				HasGraduatedAccess bool   `json:"has_graduated_access"`
				IsBlueVerified     bool   `json:"is_blue_verified"`
				ProfileImageShape  string `json:"profile_image_shape"`
				Legacy             struct {
					CanDm               bool   `json:"can_dm"`
					CanMediaTag         bool   `json:"can_media_tag"`
					CreatedAt           string `json:"created_at"`
					DefaultProfile      bool   `json:"default_profile"`
					DefaultProfileImage bool   `json:"default_profile_image"`
					Description         string `json:"description"`
					Entities            struct {
						Description struct {
							Urls []interface{} `json:"urls"`
						} `json:"description"`
					} `json:"entities"`
					FastFollowersCount      int           `json:"fast_followers_count"`
					FavouritesCount         int           `json:"favourites_count"`
					FollowersCount          int           `json:"followers_count"`
					FriendsCount            int           `json:"friends_count"`
					HasCustomTimelines      bool          `json:"has_custom_timelines"`
					IsTranslator            bool          `json:"is_translator"`
					ListedCount             int           `json:"listed_count"`
					Location                string        `json:"location"`
					MediaCount              int           `json:"media_count"`
					Name                    string        `json:"name"`
					NormalFollowersCount    int           `json:"normal_followers_count"`
					PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
					PossiblySensitive       bool          `json:"possibly_sensitive"`
					ProfileBannerUrl        string        `json:"profile_banner_url"`
					ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
					ProfileInterstitialType string        `json:"profile_interstitial_type"`
					ScreenName              string        `json:"screen_name"`
					StatusesCount           int           `json:"statuses_count"`
					TranslatorType          string        `json:"translator_type"`
					Verified                bool          `json:"verified"`
					WantRetweets            bool          `json:"want_retweets"`
					WithheldInCountries     []interface{} `json:"withheld_in_countries"`
				} `json:"legacy"`
				Professional struct {
					RestId           string        `json:"rest_id"`
					ProfessionalType string        `json:"professional_type"`
					Category         []interface{} `json:"category"`
				} `json:"professional"`
				SmartBlockedBy        bool `json:"smart_blocked_by"`
				SmartBlocking         bool `json:"smart_blocking"`
				LegacyExtendedProfile struct {
					Birthdate struct {
						Day            int    `json:"day"`
						Month          int    `json:"month"`
						Year           int    `json:"year"`
						Visibility     string `json:"visibility"`
						YearVisibility string `json:"year_visibility"`
					} `json:"birthdate"`
				} `json:"legacy_extended_profile"`
				IsProfileTranslatable           bool `json:"is_profile_translatable"`
				HasHiddenSubscriptionsOnProfile bool `json:"has_hidden_subscriptions_on_profile"`
				VerificationInfo                struct {
				} `json:"verification_info"`
				HighlightsInfo struct {
					CanHighlightTweets bool   `json:"can_highlight_tweets"`
					HighlightedTweets  string `json:"highlighted_tweets"`
				} `json:"highlights_info"`
				BusinessAccount struct {
				} `json:"business_account"`
				CreatorSubscriptionsCount int `json:"creator_subscriptions_count"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

// GetUserInfoRequest 定义一个结构体来表示请求的参数
type GetUserInfoRequest struct {
	ScreenName                                                string
	WithSafetyModeUserFields                                  bool
	HiddenProfileLikesEnabled                                 bool
	HiddenProfileSubscriptions                                bool
	ResponsiveWebGraphqlExcludeDirectiveEnabled               bool
	VerifiedPhoneLabelEnabled                                 bool
	SubscriptionsVerificationInfoIsIdentityVerifiedEnabled    bool
	SubscriptionsVerificationInfoVerifiedSinceEnabled         bool
	HighlightsTweetsTabUIEnabled                              bool
	CreatorSubscriptionsTweetPreviewAPIEnabled                bool
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled bool
	ResponsiveWebGraphqlTimelineNavigationEnabled             bool
	WithAuxiliaryUserLabels                                   bool
}

// FollowUser 实现twitter关注账号
func FollowUser(client http.Client, auth_token string, referer string) int {
	cookies, gid, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return 500
	}
	uid, err := GetUserInfo(client, auth_token, referer)
	if err != nil {
		return 500
	}

	//创建数据
	data := url.Values{}
	data.Set("include_profile_interstitial_type", "1")
	data.Set("include_blocking", "1")
	data.Set("include_blocked_by", "1")
	data.Set("include_followed_by", "1")
	data.Set("include_want_retweets", "1")
	data.Set("include_mute_edge", "1")
	data.Set("include_can_dm", "1")
	data.Set("include_can_media_tag", "1")
	data.Set("include_ext_has_nft_avatar", "1")
	data.Set("include_ext_is_blue_verified", "1")
	data.Set("include_ext_verified_type", "1")
	data.Set("include_ext_profile_image_shape", "1")
	data.Set("skip_status", "1")
	data.Set("user_id", uid)

	newRequest, err := http.NewRequest("POST", "https://x.com/i/api/1.1/friendships/create.json", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return 500
	}

	//设置请求头部
	newRequest.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	newRequest.Header.Set("accept-encoding", "gzip, deflate, br")
	newRequest.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	newRequest.Header.Set("cache-control", "no-cache")
	newRequest.Header.Set("no-cache", "no-cache")
	newRequest.Header.Set("pragma", "no-cache")
	newRequest.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	newRequest.Header.Set("upgrade-insecure-requests", "1")
	newRequest.Header.Set("x-twitter-active-user", "yes")
	newRequest.Header.Set("x-twitter-auth-type", "OAuth2Session")
	newRequest.Header.Set("x-twitter-client-language", "en")
	newRequest.Header.Set("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	newRequest.Header.Set("x-csrf-token", gid)
	newRequest.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	newRequest.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	newRequest.Header.Set("referer", referer)
	newRequest.Header.Set("sec-fetch-site", "same-origin")
	newRequest.Header.Set("origin", "https://x.com")
	newRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, cookie := range cookies {
		newRequest.AddCookie(cookie)
	}

	newRequest.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})
	response, err := client.Do(newRequest)
	if err != nil {
		return 500
	}

	//判断状态码
	if response.StatusCode != 200 {
		return 500
	}
	return 200
}

// GetUserInfo 返回用户信息
func GetUserInfo(client http.Client, auth_token string, referer string) (string, error) {
	split := strings.Split(referer, "https://x.com/")
	params := GetUserInfoRequest{
		ScreenName:                                                split[1],
		WithSafetyModeUserFields:                                  true,
		HiddenProfileLikesEnabled:                                 false,
		HiddenProfileSubscriptions:                                true,
		ResponsiveWebGraphqlExcludeDirectiveEnabled:               true,
		VerifiedPhoneLabelEnabled:                                 false,
		SubscriptionsVerificationInfoIsIdentityVerifiedEnabled:    false,
		SubscriptionsVerificationInfoVerifiedSinceEnabled:         true,
		HighlightsTweetsTabUIEnabled:                              true,
		CreatorSubscriptionsTweetPreviewAPIEnabled:                true,
		ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled: false,
		ResponsiveWebGraphqlTimelineNavigationEnabled:             true,
		WithAuxiliaryUserLabels:                                   false,
	}
	baseURL := "https://x.com/i/api/graphql/G3KGOASz96M-Qu0nwmGXNg/UserByScreenName?"

	queryParams := url.Values{}
	queryParams.Set("variables", fmt.Sprintf(`{"screen_name":"%s","withSafetyModeUserFields":%t}`, params.ScreenName, params.WithSafetyModeUserFields))
	queryParams.Set("features", fmt.Sprintf(`{"hidden_profile_likes_enabled":%t,"hidden_profile_subscriptions_enabled":%t,"responsive_web_graphql_exclude_directive_enabled":%t,"verified_phone_label_enabled":%t,"subscriptions_verification_info_is_identity_verified_enabled":%t,"subscriptions_verification_info_verified_since_enabled":%t,"highlights_tweets_tab_ui_enabled":%t,"creator_subscriptions_tweet_preview_api_enabled":%t,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":%t,"responsive_web_graphql_timeline_navigation_enabled":%t}`, params.HiddenProfileLikesEnabled, params.HiddenProfileSubscriptions, params.ResponsiveWebGraphqlExcludeDirectiveEnabled, params.VerifiedPhoneLabelEnabled, params.SubscriptionsVerificationInfoIsIdentityVerifiedEnabled, params.SubscriptionsVerificationInfoVerifiedSinceEnabled, params.HighlightsTweetsTabUIEnabled, params.CreatorSubscriptionsTweetPreviewAPIEnabled, params.ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled, params.ResponsiveWebGraphqlTimelineNavigationEnabled))
	queryParams.Set("fieldToggles", fmt.Sprintf(`{"withAuxiliaryUserLabels":%t}`, params.WithAuxiliaryUserLabels))

	cookies, gid, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("GET", baseURL+queryParams.Encode(), nil)
	if err != nil {
		return "", err
	}

	//设置请求头部
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-encoding", "gzip, deflate, br")
	request.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Set("cache-control", "no-cache")
	request.Header.Set("no-cache", "no-cache")
	request.Header.Set("pragma", "no-cache")
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	request.Header.Set("upgrade-insecure-requests", "1")
	request.Header.Set("x-twitter-active-user", "yes")
	request.Header.Set("x-twitter-auth-type", "OAuth2Session")
	request.Header.Set("x-twitter-client-language", "en")
	request.Header.Set("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	request.Header.Set("x-csrf-token", gid)
	request.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	request.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	request.Header.Set("referer", referer)
	request.Header.Set("sec-fetch-site", "same-origin")
	request.Header.Set("origin", "https://x.com")
	request.Header.Set("content-type", "application/json")

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	var reader *gzip.Reader
	reader, err = gzip.NewReader(response.Body)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	var data FollowUserResponse
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return "", err
	}

	return data.Data.User.Result.RestId, nil
}

// GetTwitterInfo 获取twitter用户信息
func GetTwitterInfo(client http.Client, auth_token string) ([]*http.Cookie, string, error) {
	//首先对https://x.com/home发送一个请求获取ct0
	request, err := http.NewRequest("GET", "https://x.com/home", nil)
	if err != nil {
		return nil, "", err
	}

	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})

	do, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}

	// 获取ct0
	var ct0 string
	for _, v := range do.Cookies() {
		if v.Name == "ct0" {
			ct0 = v.Value
		}
	}

	return do.Cookies(), ct0, nil
}
