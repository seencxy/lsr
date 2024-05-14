package twitter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

type CreateRetweetMessageRequest struct {
	Variables struct {
		TweetText string `json:"tweet_text"`
		Reply     struct {
			InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
			ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
		} `json:"reply"`
		DarkRequest bool `json:"dark_request"`
		Media       struct {
			MediaEntities     []interface{} `json:"media_entities"`
			PossiblySensitive bool          `json:"possibly_sensitive"`
		} `json:"media"`
		SemanticAnnotationIds []interface{} `json:"semantic_annotation_ids"`
	} `json:"variables"`
	Features struct {
		TweetypieUnmentionOptimizationEnabled                          bool `json:"tweetypie_unmention_optimization_enabled"`
		ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
		GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
		ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
		LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
		ResponsiveWebTwitterArticleTweetConsumptionEnabled             bool `json:"responsive_web_twitter_article_tweet_consumption_enabled"`
		TweetAwardsWebTippingEnabled                                   bool `json:"tweet_awards_web_tipping_enabled"`
		LongformNotetweetsRichTextReadEnabled                          bool `json:"longform_notetweets_rich_text_read_enabled"`
		LongformNotetweetsInlineMediaEnabled                           bool `json:"longform_notetweets_inline_media_enabled"`
		ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
		VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
		FreedomOfSpeechNotReachFetchEnabled                            bool `json:"freedom_of_speech_not_reach_fetch_enabled"`
		StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
		TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
		ResponsiveWebMediaDownloadVideoEnabled                         bool `json:"responsive_web_media_download_video_enabled"`
		ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
		ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
		ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
	} `json:"features"`
	QueryId string `json:"queryId"`
}

// 对某条推文发起评论
func CreateRetweetMessage(client http.Client, info string, auth_token string, referer string) int {
	// 获取twitter信息
	cookies, gid, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return 302
	}
	re := regexp.MustCompile(`status/(\d+)`)
	matches := re.FindStringSubmatch(referer)
	data := CreateRetweetMessageRequest{
		Variables: struct {
			TweetText string `json:"tweet_text"`
			Reply     struct {
				InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
				ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
			} `json:"reply"`
			DarkRequest bool `json:"dark_request"`
			Media       struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			} `json:"media"`
			SemanticAnnotationIds []interface{} `json:"semantic_annotation_ids"`
		}{
			TweetText: info,
			Reply: struct {
				InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
				ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
			}{
				InReplyToTweetId:    matches[1],
				ExcludeReplyUserIds: []interface{}{},
			},
			DarkRequest: false,
			Media: struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			}{
				MediaEntities:     []interface{}{},
				PossiblySensitive: false,
			},
			SemanticAnnotationIds: []interface{}{},
		},
		Features: struct {
			TweetypieUnmentionOptimizationEnabled                          bool `json:"tweetypie_unmention_optimization_enabled"`
			ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
			GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
			ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
			LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
			ResponsiveWebTwitterArticleTweetConsumptionEnabled             bool `json:"responsive_web_twitter_article_tweet_consumption_enabled"`
			TweetAwardsWebTippingEnabled                                   bool `json:"tweet_awards_web_tipping_enabled"`
			LongformNotetweetsRichTextReadEnabled                          bool `json:"longform_notetweets_rich_text_read_enabled"`
			LongformNotetweetsInlineMediaEnabled                           bool `json:"longform_notetweets_inline_media_enabled"`
			ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
			VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
			FreedomOfSpeechNotReachFetchEnabled                            bool `json:"freedom_of_speech_not_reach_fetch_enabled"`
			StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
			TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
			ResponsiveWebMediaDownloadVideoEnabled                         bool `json:"responsive_web_media_download_video_enabled"`
			ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
			ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
			ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
		}{
			TweetypieUnmentionOptimizationEnabled:                          true,
			ResponsiveWebEditTweetApiEnabled:                               true,
			GraphqlIsTranslatableRwebTweetIsTranslatableEnabled:            true,
			ViewCountsEverywhereApiEnabled:                                 true,
			LongformNotetweetsConsumptionEnabled:                           true,
			ResponsiveWebTwitterArticleTweetConsumptionEnabled:             false,
			TweetAwardsWebTippingEnabled:                                   false,
			LongformNotetweetsRichTextReadEnabled:                          true,
			LongformNotetweetsInlineMediaEnabled:                           true,
			ResponsiveWebGraphqlExcludeDirectiveEnabled:                    true,
			VerifiedPhoneLabelEnabled:                                      false,
			FreedomOfSpeechNotReachFetchEnabled:                            true,
			StandardizedNudgesMisinfo:                                      true,
			TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled: true,
			ResponsiveWebMediaDownloadVideoEnabled:                         false,
			ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled:      false,
			ResponsiveWebGraphqlTimelineNavigationEnabled:                  true,
			ResponsiveWebEnhanceCardsEnabled:                               false,
		},
		QueryId: "SoVnbfCycZ7fERGCwpZkYA",
	}

	//将请求的参数json化
	marshal, err := json.Marshal(data)
	if err != nil {
		return 500
	}

	//创建一个请求体
	request, err := http.NewRequest("POST", "https://twitter.com/i/api/graphql/PIZtQLRIYtSa9AtW_fI2Mw/CreateTweet", bytes.NewReader(marshal))
	if err != nil {
		return 500
	}

	//设置请求头部
	request.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
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
	request.Header.Set("origin", "https://twitter.com")
	request.Header.Set("content-type", "application/json")

	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})

	//发送请求
	response, err := client.Do(request)
	if err != nil {
		return 500
	}

	//判断状态码
	if response.Status != "200 OK" {
		return 500
	}
	return 200
}
