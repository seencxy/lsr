package galxe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 绑定twitter请求参数
type BindTwitterRequest struct {
	Variables struct {
		TweetText   string `json:"tweet_text"`
		DarkRequest bool   `json:"dark_request"`
		Media       struct {
			MediaEntities     []interface{} `json:"media_entities"`
			PossiblySensitive bool          `json:"possibly_sensitive"`
		} `json:"media"`
		SemanticAnnotationIDs []interface{} `json:"semantic_annotation_ids"`
	} `json:"variables"`
	Features map[string]bool `json:"features"`
	QueryID  string          `json:"queryId"`
}

// 绑定推特响应参数
type BindTwitterResponse struct {
	Data struct {
		CreateTweet struct {
			TweetResults struct {
				Result struct {
					RestId string `json:"rest_id"`
					Core   struct {
						UserResults struct {
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
									NeedsPhoneVerification  bool          `json:"needs_phone_verification"`
									NormalFollowersCount    int           `json:"normal_followers_count"`
									PinnedTweetIdsStr       []interface{} `json:"pinned_tweet_ids_str"`
									PossiblySensitive       bool          `json:"possibly_sensitive"`
									ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
									ProfileInterstitialType string        `json:"profile_interstitial_type"`
									ScreenName              string        `json:"screen_name"`
									StatusesCount           int           `json:"statuses_count"`
									TranslatorType          string        `json:"translator_type"`
									Verified                bool          `json:"verified"`
									WantRetweets            bool          `json:"want_retweets"`
									WithheldInCountries     []interface{} `json:"withheld_in_countries"`
								} `json:"legacy"`
							} `json:"result"`
						} `json:"user_results"`
					} `json:"core"`
					Card struct {
						RestId string `json:"rest_id"`
						Legacy struct {
							BindingValues []struct {
								Key   string `json:"key"`
								Value struct {
									ImageValue struct {
										Height int    `json:"height"`
										Width  int    `json:"width"`
										Url    string `json:"url"`
									} `json:"image_value,omitempty"`
									Type        string `json:"type"`
									StringValue string `json:"string_value,omitempty"`
									ScribeKey   string `json:"scribe_key,omitempty"`
									UserValue   struct {
										IdStr string        `json:"id_str"`
										Path  []interface{} `json:"path"`
									} `json:"user_value,omitempty"`
									ImageColorValue struct {
										Palette []struct {
											Rgb struct {
												Blue  int `json:"blue"`
												Green int `json:"green"`
												Red   int `json:"red"`
											} `json:"rgb"`
											Percentage float64 `json:"percentage"`
										} `json:"palette"`
									} `json:"image_color_value,omitempty"`
								} `json:"value"`
							} `json:"binding_values"`
							CardPlatform struct {
								Platform struct {
									Audience struct {
										Name string `json:"name"`
									} `json:"audience"`
									Device struct {
										Name    string `json:"name"`
										Version string `json:"version"`
									} `json:"device"`
								} `json:"platform"`
							} `json:"card_platform"`
							Name            string `json:"name"`
							Url             string `json:"url"`
							UserRefsResults []struct {
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
											Url struct {
												Urls []struct {
													DisplayUrl  string `json:"display_url"`
													ExpandedUrl string `json:"expanded_url"`
													Url         string `json:"url"`
													Indices     []int  `json:"indices"`
												} `json:"urls"`
											} `json:"url"`
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
										Url                     string        `json:"url"`
										Verified                bool          `json:"verified"`
										VerifiedType            string        `json:"verified_type"`
										WantRetweets            bool          `json:"want_retweets"`
										WithheldInCountries     []interface{} `json:"withheld_in_countries"`
									} `json:"legacy"`
									Professional struct {
										RestId           string        `json:"rest_id"`
										ProfessionalType string        `json:"professional_type"`
										Category         []interface{} `json:"category"`
									} `json:"professional"`
								} `json:"result"`
							} `json:"user_refs_results"`
						} `json:"legacy"`
					} `json:"card"`
					UnifiedCard struct {
						CardFetchState string `json:"card_fetch_state"`
					} `json:"unified_card"`
					EditControl struct {
						EditTweetIds       []string `json:"edit_tweet_ids"`
						EditableUntilMsecs string   `json:"editable_until_msecs"`
						IsEditEligible     bool     `json:"is_edit_eligible"`
						EditsRemaining     string   `json:"edits_remaining"`
					} `json:"edit_control"`
					IsTranslatable bool `json:"is_translatable"`
					Views          struct {
						State string `json:"state"`
					} `json:"views"`
					Source string `json:"source"`
					Legacy struct {
						BookmarkCount     int    `json:"bookmark_count"`
						Bookmarked        bool   `json:"bookmarked"`
						CreatedAt         string `json:"created_at"`
						ConversationIdStr string `json:"conversation_id_str"`
						DisplayTextRange  []int  `json:"display_text_range"`
						Entities          struct {
							UserMentions []struct {
								IdStr      string `json:"id_str"`
								Name       string `json:"name"`
								ScreenName string `json:"screen_name"`
								Indices    []int  `json:"indices"`
							} `json:"user_mentions"`
							Urls []struct {
								DisplayUrl  string `json:"display_url"`
								ExpandedUrl string `json:"expanded_url"`
								Url         string `json:"url"`
								Indices     []int  `json:"indices"`
							} `json:"urls"`
							Hashtags []struct {
								Indices []int  `json:"indices"`
								Text    string `json:"text"`
							} `json:"hashtags"`
							Symbols []interface{} `json:"symbols"`
						} `json:"entities"`
						FavoriteCount             int    `json:"favorite_count"`
						Favorited                 bool   `json:"favorited"`
						FullText                  string `json:"full_text"`
						IsQuoteStatus             bool   `json:"is_quote_status"`
						Lang                      string `json:"lang"`
						PossiblySensitive         bool   `json:"possibly_sensitive"`
						PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
						QuoteCount                int    `json:"quote_count"`
						ReplyCount                int    `json:"reply_count"`
						RetweetCount              int    `json:"retweet_count"`
						Retweeted                 bool   `json:"retweeted"`
						UserIdStr                 string `json:"user_id_str"`
						IdStr                     string `json:"id_str"`
					} `json:"legacy"`
					UnmentionInfo struct {
					} `json:"unmention_info"`
				} `json:"result"`
			} `json:"tweet_results"`
		} `json:"create_tweet"`
	} `json:"data"`
}

// 尝试绑定twitter
func BindTwitter(client http.Client, auth_token string, gid string, address string, sign string) (bool, error) {
	cookies, ct0, err := GetTwitterInfo(client, auth_token)
	if err != nil {
		return false, err
	}

	// 创建结构体实例并赋值
	data := BindTwitterRequest{
		Variables: struct {
			TweetText   string `json:"tweet_text"`
			DarkRequest bool   `json:"dark_request"`
			Media       struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			} `json:"media"`
			SemanticAnnotationIDs []interface{} `json:"semantic_annotation_ids"`
		}{
			TweetText:   fmt.Sprintf("Verifying my Twitter account for my #GalxeID gid:%s @Galxe \n\n galxe.com/galxeid ", gid),
			DarkRequest: false,
			Media: struct {
				MediaEntities     []interface{} `json:"media_entities"`
				PossiblySensitive bool          `json:"possibly_sensitive"`
			}{
				MediaEntities:     []interface{}{},
				PossiblySensitive: false,
			},
			SemanticAnnotationIDs: []interface{}{},
		},
		Features: map[string]bool{
			"tweetypie_unmention_optimization_enabled":                                true,
			"responsive_web_edit_tweet_api_enabled":                                   true,
			"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              true,
			"view_counts_everywhere_api_enabled":                                      true,
			"longform_notetweets_consumption_enabled":                                 true,
			"responsive_web_twitter_article_tweet_consumption_enabled":                false,
			"tweet_awards_web_tipping_enabled":                                        false,
			"longform_notetweets_rich_text_read_enabled":                              true,
			"longform_notetweets_inline_media_enabled":                                true,
			"responsive_web_graphql_exclude_directive_enabled":                        true,
			"verified_phone_label_enabled":                                            false,
			"freedom_of_speech_not_reach_fetch_enabled":                               true,
			"standardized_nudges_misinfo":                                             true,
			"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
			"responsive_web_media_download_video_enabled":                             false,
			"responsive_web_graphql_skip_user_profile_image_extensions_enabled":       false,
			"responsive_web_graphql_timeline_navigation_enabled":                      true,
			"responsive_web_enhance_cards_enabled":                                    false,
		},
		QueryID: "SoVnbfCycZ7fERGCwpZkYA",
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// 创建请求
	request, err := http.NewRequest("POST", "https://twitter.com/i/api/graphql/SoVnbfCycZ7fERGCwpZkYA/CreateTweet", bytes.NewReader(marshal))
	if err != nil {
		return false, err
	}

	for _, v := range cookies {
		request.AddCookie(v)
	}

	request.Header.Add("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	request.Header.Add("x-twitter-active-user", "yes")
	request.Header.Add("x-twitter-client-language", "en")
	request.Header.Add("x-twitter-auth-type", "OAuth2Session")
	request.Header.Add("x-csrf-token", ct0)
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})
	request.AddCookie(&http.Cookie{Name: "x-csrf-token", Value: ct0})

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	// 首先对请求状态坐下判断
	if response.StatusCode == 200 {
		all, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}

		var result BindTwitterResponse
		err = json.Unmarshal(all, &result)
		if err != nil {
			return false, err
		}
		twitter_url := fmt.Sprintf("https://twitter.com/%s/status/%s", result.Data.CreateTweet.TweetResults.Result.Core.UserResults.Result.Legacy.ScreenName, result.Data.CreateTweet.TweetResults.Result.RestId)
		account, err := VerifyTwitterAccount(client, address, twitter_url, sign)
		if err != nil || account == false {
			return false, err
		}

		return true, nil
	}

	return false, errors.New("绑定twitter失败: token需要验证")
}

func GetTwitterInfo(client http.Client, auth_token string) ([]*http.Cookie, string, error) {
	//首先对https://twitter.com/home发送一个请求获取ct0
	request, err := http.NewRequest("GET", "https://twitter.com/home", nil)
	if err != nil {
		return nil, "", err
	}

	request.AddCookie(&http.Cookie{Name: "auth_token", Value: auth_token})

	do, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}

	if do.StatusCode == 302 {
		return nil, "", errors.New("账号需要认证...")
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
