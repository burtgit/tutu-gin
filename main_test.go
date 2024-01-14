package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type RequestParams struct {
	Query   map[string]string
	Headers map[string]string
}

func GetHTTPRes(Link string, params RequestParams) *fasthttp.Response {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书认证
	}

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	client := &fasthttp.Client{ReadBufferSize: 8192, TLSConfig: tlsConfig}

	req.Header.SetMethod("GET")
	for key, value := range params.Headers {
		req.Header.Set(key, value)
	}

	req.SetRequestURI(Link)
	for key, value := range params.Query {
		req.URI().QueryArgs().Add(key, value)
	}

	err := client.Do(req, res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func TestTwitter(t *testing.T) {
	url := "https://vxtwitter.com/tankman2002/status/1746023931679522906"
	if len(url) == 0 {
		errorMessage := "No URL specified"
		t.Fatal(errorMessage)
	}

	TweetID := (regexp.MustCompile((`.*(?:twitter|x).com/.+status/([A-Za-z0-9]+)`))).FindStringSubmatch(url)[1]
	csrfToken := strings.ReplaceAll((uuid.New()).String(), "-", "")
	Headers := map[string]string{
		"Authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"Cookie":                    fmt.Sprintf("auth_token=ee4ebd1070835b90a9b8016d1e6c6130ccc89637; ct0=%v; ", csrfToken),
		"x-twitter-active-user":     "yes",
		"x-twitter-auth-type":       "OAuth2Session",
		"x-twitter-client-language": "en",
		"x-csrf-token":              csrfToken,
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0",
	}
	variables := map[string]interface{}{
		"focalTweetId":                           TweetID,
		"referrer":                               "messages",
		"includePromotedContent":                 true,
		"withCommunity":                          true,
		"withQuickPromoteEligibilityTweetFields": true,
		"withBirdwatchNotes":                     true,
		"withVoice":                              true,
		"withV2Timeline":                         true,
	}
	features := map[string]interface{}{
		"rweb_lists_timeline_redesign_enabled":                                    true,
		"responsive_web_graphql_exclude_directive_enabled":                        true,
		"verified_phone_label_enabled":                                            false,
		"creator_subscriptions_tweet_preview_api_enabled":                         true,
		"responsive_web_graphql_timeline_navigation_enabled":                      true,
		"responsive_web_graphql_skip_user_profile_image_extensions_enabled":       false,
		"tweetypie_unmention_optimization_enabled":                                true,
		"responsive_web_edit_tweet_api_enabled":                                   true,
		"graphql_is_translatable_rweb_tweet_is_translatable_enabled":              false,
		"view_counts_everywhere_api_enabled":                                      true,
		"longform_notetweets_consumption_enabled":                                 true,
		"responsive_web_twitter_article_tweet_consumption_enabled":                false,
		"tweet_awards_web_tipping_enabled":                                        false,
		"freedom_of_speech_not_reach_fetch_enabled":                               true,
		"standardized_nudges_misinfo":                                             true,
		"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": true,
		"longform_notetweets_rich_text_read_enabled":                              true,
		"longform_notetweets_inline_media_enabled":                                true,
		"responsive_web_media_download_video_enabled":                             false,
		"responsive_web_enhance_cards_enabled":                                    false,
	}
	fieldtoggles := map[string]interface{}{
		"withAuxiliaryUserLabels":     false,
		"withArticleRichContentState": false,
	}

	variablesJson, _ := json.Marshal(variables)
	featuresJson, _ := json.Marshal(features)
	fieldTogglesJson, _ := json.Marshal(fieldtoggles)

	Query := map[string]string{
		"variables":    string(variablesJson),
		"features":     string(featuresJson),
		"fieldToggles": string(fieldTogglesJson),
	}

	body := GetHTTPRes("https://twitter.com/i/api/graphql/NmCeCgkVlsRGS1cAwqtgmw/TweetDetail", RequestParams{Query: Query, Headers: Headers}).Body()
	t.Log(body)
	//s := gjson.ParseBytes(body).String()
	//indexedMedia := &handler.IndexedMedia{}
	//var caption string
	//results := gjson.Get(s, fmt.Sprintf(`data.threaded_conversation_with_injections_v2.instructions.0.entries.#(entryId="tweet-%v").content.itemContent.tweet_results.result`, string(TweetID)))
	//if results.Get("__typename").String() == "TweetWithVisibilityResults" {
	//	results = results.Get("tweet")
	//}
	//caption = results.Get("legacy.full_text").String()
	//
	//medias := results.Get("legacy.extended_entities.media")
	//for _, media := range medias.Array() {
	//	var videoType string
	//	if slices.Contains([]string{"animated_gif", "video"}, media.Get("type").String()) {
	//		videoType = "video"
	//	}
	//
	//	if videoType != "video" {
	//		indexedMedia.Medias = append(indexedMedia.Medias, handler.Medias{
	//			Height: int(media.Get("original_info.height").Int()),
	//			Width:  int(media.Get("original_info.width").Int()),
	//			Source: media.Get("media_url_https").String(),
	//			Video:  false,
	//		})
	//	} else {
	//		indexedMedia.Medias = append(indexedMedia.Medias, handler.Medias{
	//			Height: int(media.Get("original_info.height").Int()),
	//			Width:  int(media.Get("original_info.width").Int()),
	//			Source: media.Get("video_info.variants.0.url").String(),
	//			Video:  true,
	//		})
	//	}
	//}
	//
	//ixt := handler.IndexedMedia{
	//	URL:     url,
	//	Medias:  indexedMedia.Medias,
	//	Caption: caption}
	//
	//jsonResponse, _ := json.Marshal(ixt)
	//err := cache.GetRedisClient().Set(context.Background(), TweetID, jsonResponse, 24*time.Hour*60).Err()
	//if err != nil {
	//	log.Println("Error setting cache:", err)
	//}
	//ctx.Response.Header.Add("Content-Type", "application/json")
	//json.NewEncoder(ctx).Encode(ixt)
}

func TestVersion(t *testing.T) {

	var NeedUpdate bool
	v1 := strings.Split("1.0.1", ".")
	v2 := strings.Split("1.1.0", ".")

	for i := 0; i < len(v1) || i < len(v2); i++ {
		num1 := 0
		num2 := 0

		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}

		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}

		if num1 > num2 {
			NeedUpdate = true
			break
		}
	}

	t.Log(NeedUpdate)
}
