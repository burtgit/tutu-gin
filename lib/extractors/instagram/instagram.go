package instagram

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/iawia002/lux/request"
	errors2 "github.com/juju/errors"
	"io"
	"net/http"
	netURL "net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"tutu-gin/utils"

	"github.com/pkg/errors"

	"tutu-gin/lib/extractors"
)

type instagram struct {
	GqlData `json:"gql_data"`
}

type GqlData struct {
	ShortcodeMedia struct {
		Typename       string `json:"__typename"`
		Id             string `json:"id"`
		Shortcode      string `json:"shortcode"`
		CommenterCount int    `json:"commenter_count"`
		Dimensions     struct {
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"dimensions"`
		DisplayUrl       string `json:"display_url"`
		VideoUrl         string `json:"video_url"`
		DisplayResources []struct {
			ConfigWidth  int    `json:"config_width"`
			ConfigHeight int    `json:"config_height"`
			Src          string `json:"src"`
		} `json:"display_resources"`
		IsVideo             bool        `json:"is_video"`
		MediaOverlayInfo    interface{} `json:"media_overlay_info"`
		SharingFrictionInfo struct {
			ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
			BloksAppUrl               interface{} `json:"bloks_app_url"`
		} `json:"sharing_friction_info"`
		Owner struct {
			Id               string `json:"id"`
			ProfilePicUrl    string `json:"profile_pic_url"`
			Username         string `json:"username"`
			FollowedByViewer bool   `json:"followed_by_viewer"`
			HasPublicStory   bool   `json:"has_public_story"`
			IsPrivate        bool   `json:"is_private"`
			IsUnpublished    bool   `json:"is_unpublished"`
			IsVerified       bool   `json:"is_verified"`
			EdgeFollowedBy   struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			EdgeOwnerToTimelineMedia struct {
				Count int `json:"count"`
				Edges []struct {
					Node struct {
						Id                 string `json:"id"`
						ThumbnailSrc       string `json:"thumbnail_src"`
						ThumbnailResources []struct {
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
							Src          string `json:"src"`
						} `json:"thumbnail_resources"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_media"`
			EdgeOwnerToTimelineVideoMedia struct {
				Edges []struct {
					Node struct {
						AccessibilityCaption interface{} `json:"accessibility_caption"`
						MediaOverlayInfo     interface{} `json:"media_overlay_info"`
						Permalink            string      `json:"permalink"`
						Shortcode            string      `json:"shortcode"`
						ThumbnailSrc         string      `json:"thumbnail_src"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_owner_to_timeline_video_media"`
		} `json:"owner"`
		TakenAtTimestamp   int `json:"taken_at_timestamp"`
		EdgeMediaToCaption struct {
			Edges []struct {
				Node struct {
					Text string `json:"text"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_media_to_caption"`
		EdgeMediaToSponsorUser struct {
			Edges []interface{} `json:"edges"`
		} `json:"edge_media_to_sponsor_user"`
		IsAffiliate               bool        `json:"is_affiliate"`
		IsPaidPartnership         bool        `json:"is_paid_partnership"`
		LikeAndViewCountsDisabled bool        `json:"like_and_view_counts_disabled"`
		MediaPreview              interface{} `json:"media_preview"`
		Location                  interface{} `json:"location"`
		EdgeMediaToComment        struct {
			Count int `json:"count"`
		} `json:"edge_media_to_comment"`
		EdgeLikedBy struct {
			Count int `json:"count"`
		} `json:"edge_liked_by"`
		CoauthorProducers     []interface{} `json:"coauthor_producers"`
		PinnedForUsers        []interface{} `json:"pinned_for_users"`
		EdgeSidecarToChildren struct {
			Edges []struct {
				Node struct {
					Typename       string `json:"__typename"`
					Id             string `json:"id"`
					Shortcode      string `json:"shortcode"`
					CommenterCount int    `json:"commenter_count"`
					Dimensions     struct {
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"dimensions"`
					DisplayUrl       string `json:"display_url"`
					DisplayResources []struct {
						ConfigWidth  int    `json:"config_width"`
						ConfigHeight int    `json:"config_height"`
						Src          string `json:"src"`
					} `json:"display_resources"`
					IsVideo                   bool        `json:"is_video"`
					Title                     interface{} `json:"title"`
					ProductType               string      `json:"product_type,omitempty"`
					VideoUrl                  string      `json:"video_url,omitempty"`
					VideoViewCount            int         `json:"video_view_count,omitempty"`
					ClipsMusicAttributionInfo interface{} `json:"clips_music_attribution_info"`
					MediaOverlayInfo          interface{} `json:"media_overlay_info"`
					SharingFrictionInfo       struct {
						ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
						BloksAppUrl               interface{} `json:"bloks_app_url"`
					} `json:"sharing_friction_info"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"edge_sidecar_to_children"`
	} `json:"shortcode_media"`
}

type insItem struct {
	Data struct {
		ShortcodeMedia struct {
			Typename   string `json:"__typename"`
			Id         string `json:"id"`
			Shortcode  string `json:"shortcode"`
			Dimensions struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			GatingInfo              interface{} `json:"gating_info"`
			FactCheckOverallRating  interface{} `json:"fact_check_overall_rating"`
			FactCheckInformation    interface{} `json:"fact_check_information"`
			SensitivityFrictionInfo interface{} `json:"sensitivity_friction_info"`
			SharingFrictionInfo     struct {
				ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
				BloksAppUrl               interface{} `json:"bloks_app_url"`
			} `json:"sharing_friction_info"`
			MediaOverlayInfo interface{} `json:"media_overlay_info"`
			MediaPreview     string      `json:"media_preview"`
			DisplayUrl       string      `json:"display_url"`
			DisplayResources []struct {
				Src          string `json:"src"`
				ConfigWidth  int    `json:"config_width"`
				ConfigHeight int    `json:"config_height"`
			} `json:"display_resources"`
			AccessibilityCaption interface{} `json:"accessibility_caption"`
			DashInfo             struct {
				IsDashEligible    bool        `json:"is_dash_eligible"`
				VideoDashManifest interface{} `json:"video_dash_manifest"`
				NumberOfQualities int         `json:"number_of_qualities"`
			} `json:"dash_info"`
			HasAudio              bool        `json:"has_audio"`
			VideoUrl              string      `json:"video_url"`
			VideoViewCount        int         `json:"video_view_count"`
			VideoPlayCount        int         `json:"video_play_count"`
			IsVideo               bool        `json:"is_video"`
			TrackingToken         string      `json:"tracking_token"`
			UpcomingEvent         interface{} `json:"upcoming_event"`
			EdgeMediaToTaggedUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_tagged_user"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						CreatedAt string `json:"created_at"`
						Text      string `json:"text"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			CanSeeInsightsAsBrand     bool `json:"can_see_insights_as_brand"`
			CaptionIsEdited           bool `json:"caption_is_edited"`
			HasRankedComments         bool `json:"has_ranked_comments"`
			LikeAndViewCountsDisabled bool `json:"like_and_view_counts_disabled"`
			EdgeMediaToParentComment  struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool        `json:"has_next_page"`
					EndCursor   interface{} `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						Id              string `json:"id"`
						Text            string `json:"text"`
						CreatedAt       int    `json:"created_at"`
						DidReportAsSpam bool   `json:"did_report_as_spam"`
						Owner           struct {
							Id            string `json:"id"`
							IsVerified    bool   `json:"is_verified"`
							ProfilePicUrl string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"owner"`
						ViewerHasLiked bool `json:"viewer_has_liked"`
						EdgeLikedBy    struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						IsRestrictedPending  bool `json:"is_restricted_pending"`
						EdgeThreadedComments struct {
							Count    int `json:"count"`
							PageInfo struct {
								HasNextPage bool        `json:"has_next_page"`
								EndCursor   interface{} `json:"end_cursor"`
							} `json:"page_info"`
							Edges []interface{} `json:"edges"`
						} `json:"edge_threaded_comments"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_parent_comment"`
			EdgeMediaToHoistedComment struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_hoisted_comment"`
			EdgeMediaPreviewComment struct {
				Count int `json:"count"`
				Edges []struct {
					Node struct {
						Id              string `json:"id"`
						Text            string `json:"text"`
						CreatedAt       int    `json:"created_at"`
						DidReportAsSpam bool   `json:"did_report_as_spam"`
						Owner           struct {
							Id            string `json:"id"`
							IsVerified    bool   `json:"is_verified"`
							ProfilePicUrl string `json:"profile_pic_url"`
							Username      string `json:"username"`
						} `json:"owner"`
						ViewerHasLiked bool `json:"viewer_has_liked"`
						EdgeLikedBy    struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						IsRestrictedPending bool `json:"is_restricted_pending"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_preview_comment"`
			CommentsDisabled            bool `json:"comments_disabled"`
			CommentingDisabledForViewer bool `json:"commenting_disabled_for_viewer"`
			TakenAtTimestamp            int  `json:"taken_at_timestamp"`
			EdgeMediaPreviewLike        struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_preview_like"`
			EdgeMediaToSponsorUser struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_media_to_sponsor_user"`
			IsAffiliate                bool        `json:"is_affiliate"`
			IsPaidPartnership          bool        `json:"is_paid_partnership"`
			Location                   interface{} `json:"location"`
			NftAssetInfo               interface{} `json:"nft_asset_info"`
			ViewerHasLiked             bool        `json:"viewer_has_liked"`
			ViewerHasSaved             bool        `json:"viewer_has_saved"`
			ViewerHasSavedToCollection bool        `json:"viewer_has_saved_to_collection"`
			ViewerInPhotoOfYou         bool        `json:"viewer_in_photo_of_you"`
			ViewerCanReshare           bool        `json:"viewer_can_reshare"`
			Owner                      struct {
				Id                        string      `json:"id"`
				IsVerified                bool        `json:"is_verified"`
				ProfilePicUrl             string      `json:"profile_pic_url"`
				Username                  string      `json:"username"`
				BlockedByViewer           bool        `json:"blocked_by_viewer"`
				RestrictedByViewer        interface{} `json:"restricted_by_viewer"`
				FollowedByViewer          bool        `json:"followed_by_viewer"`
				FullName                  string      `json:"full_name"`
				HasBlockedViewer          bool        `json:"has_blocked_viewer"`
				IsEmbedsDisabled          bool        `json:"is_embeds_disabled"`
				IsPrivate                 bool        `json:"is_private"`
				IsUnpublished             bool        `json:"is_unpublished"`
				RequestedByViewer         bool        `json:"requested_by_viewer"`
				PassTieringRecommendation bool        `json:"pass_tiering_recommendation"`
				EdgeOwnerToTimelineMedia  struct {
					Count int `json:"count"`
				} `json:"edge_owner_to_timeline_media"`
				EdgeFollowedBy struct {
					Count int `json:"count"`
				} `json:"edge_followed_by"`
			} `json:"owner"`
			IsAd                       bool `json:"is_ad"`
			EdgeWebMediaToRelatedMedia struct {
				Edges []interface{} `json:"edges"`
			} `json:"edge_web_media_to_related_media"`
			CoauthorProducers         []interface{} `json:"coauthor_producers"`
			PinnedForUsers            []interface{} `json:"pinned_for_users"`
			EncodingStatus            interface{}   `json:"encoding_status"`
			IsPublished               bool          `json:"is_published"`
			ProductType               string        `json:"product_type"`
			Title                     string        `json:"title"`
			VideoDuration             float64       `json:"video_duration"`
			ThumbnailSrc              string        `json:"thumbnail_src"`
			ClipsMusicAttributionInfo struct {
				ArtistName            string `json:"artist_name"`
				SongName              string `json:"song_name"`
				UsesOriginalAudio     bool   `json:"uses_original_audio"`
				ShouldMuteAudio       bool   `json:"should_mute_audio"`
				ShouldMuteAudioReason string `json:"should_mute_audio_reason"`
				AudioId               string `json:"audio_id"`
			} `json:"clips_music_attribution_info"`
			EdgeRelatedProfiles struct {
				Edges []struct {
					Node struct {
						Id             string `json:"id"`
						FullName       string `json:"full_name"`
						IsPrivate      bool   `json:"is_private"`
						IsVerified     bool   `json:"is_verified"`
						ProfilePicUrl  string `json:"profile_pic_url"`
						Username       string `json:"username"`
						EdgeFollowedBy struct {
							Count int `json:"count"`
						} `json:"edge_followed_by"`
						EdgeOwnerToTimelineMedia struct {
							Count int `json:"count"`
							Edges []struct {
								Node struct {
									Typename             string `json:"__typename"`
									Id                   string `json:"id"`
									Shortcode            string `json:"shortcode"`
									EdgeMediaPreviewLike struct {
										Count int `json:"count"`
									} `json:"edge_media_preview_like"`
									EdgeMediaPreviewComment struct {
										Count int `json:"count"`
									} `json:"edge_media_preview_comment"`
									ThumbnailSrc string `json:"thumbnail_src"`
									Owner        struct {
										Id       string `json:"id"`
										Username string `json:"username"`
									} `json:"owner"`
									GatingInfo          interface{} `json:"gating_info"`
									SharingFrictionInfo struct {
										ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
										BloksAppUrl               interface{} `json:"bloks_app_url"`
									} `json:"sharing_friction_info"`
									MediaOverlayInfo     interface{} `json:"media_overlay_info"`
									IsVideo              bool        `json:"is_video"`
									AccessibilityCaption *string     `json:"accessibility_caption"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_owner_to_timeline_media"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_related_profiles"`
			EdgeSidecarToChildren struct {
				Edges []struct {
					Node struct {
						Typename       string `json:"__typename"`
						Id             string `json:"id"`
						Shortcode      string `json:"shortcode"`
						CommenterCount int    `json:"commenter_count"`
						Dimensions     struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayUrl       string `json:"display_url"`
						DisplayResources []struct {
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
							Src          string `json:"src"`
						} `json:"display_resources"`
						IsVideo                   bool        `json:"is_video"`
						Title                     interface{} `json:"title"`
						ProductType               string      `json:"product_type,omitempty"`
						VideoUrl                  string      `json:"video_url,omitempty"`
						VideoViewCount            int         `json:"video_view_count,omitempty"`
						ClipsMusicAttributionInfo interface{} `json:"clips_music_attribution_info"`
						MediaOverlayInfo          interface{} `json:"media_overlay_info"`
						SharingFrictionInfo       struct {
							ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
							BloksAppUrl               interface{} `json:"bloks_app_url"`
						} `json:"sharing_friction_info"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_sidecar_to_children"`
		} `json:"shortcode_media"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
	Status string `json:"status"`
}

type extractor struct{}

// New returns a instagram extractor.
func New() extractors.Extractor {
	return &extractor{}
}

func getVideoId(url string) string {

	validURL := `(?P<url>https?://(?:www\.)?instagram\.com(?:/[^/]+)?/(?:p|tv|reel)/(?P<id>[^/?#&]+))`
	re := regexp.MustCompile(validURL)
	match := re.FindStringSubmatch(url)

	var videoId string

	for i, n := range re.SubexpNames() {
		if i != 0 && n != "" {
			if n == "id" {
				videoId = match[i]
			}
		}
	}

	return videoId
}

func convertStringToNumber(str string) int64 {

	table := map[rune]int64{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
		'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19,
		'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29,
		'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
		'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49,
		'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59,
		'8': 60, '9': 61, '-': 62, '_': 63,
	}

	var result, base int64
	base = int64(len(table))
	for _, char := range str {
		result = result*base + table[char]
	}
	return result
}

func getCsrfToken(videoId string) (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			Proxy:              http.ProxyFromEnvironment,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, "https://i.instagram.com/api/v1/web/get_ruling_for_content/?content_type=MEDIA&target_id="+strconv.FormatInt(convertStringToNumber(videoId), 10), nil)
	if err != nil {
		return "", fmt.Errorf("创建请求出错: %v", err)
	}
	//{'X-Ig-App-Id': '936619743392459', 'X-Asbd-Id': '198387', 'X-Ig-Www-Claim': '0', 'Origin': 'https://www.instagram.com', 'Accept': '*/*', 'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36'}

	req.Header.Set("X-Ig-App-Id", "936619743392459")
	req.Header.Set("X-Asbd-Id", "198387")
	req.Header.Set("X-Ig-Www-Claim", "0")
	req.Header.Set("Origin", "https://www.instagram.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")

	//if _, ok := headers["Referer"]; !ok {
	//	req.Header.Set("Referer", urlStr)
	//}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors2.Annotate(err, "创建请求出错")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b))
		return "", errors2.Annotate(errors.New("请求返回状态码有误2"), "请求返回状态码有误2")
	}

	// 获取cookie里面的csrftoken内容
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			return cookie.Value, nil
		}
	}

	return "", errors2.Annotate(errors.New("未找到csrftoken"), "未找到csrftoken")
}

// Extract is the main function to extract the data.
func (e *extractor) ExtractV2(url string, option extractors.Options) ([]*extractors.Data, error) {
	videoId := getVideoId(url)
	token, err := getCsrfToken(videoId)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy:              http.ProxyFromEnvironment,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, "https://www.instagram.com/graphql/query/?query_hash=9f8827793ef34641b2fb195d4d41151c&variables=%7B%22shortcode%22%3A%22"+videoId+"%22%2C%22child_comment_count%22%3A3%2C%22fetch_comment_count%22%3A40%2C%22parent_comment_count%22%3A24%2C%22has_threaded_comments%22%3Atrue%7D", nil)
	if err != nil {
		return nil, errors2.Annotate(err, "创建请求出错")
	}

	req.Header.Set("X-Ig-App-Id", "936619743392459")
	req.Header.Set("X-Asbd-Id", "198387")
	req.Header.Set("X-Ig-Www-Claim", "0")
	req.Header.Set("Origin", "https://www.instagram.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Csrftoken", token)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", url)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")

	//if _, ok := headers["Referer"]; !ok {
	//	req.Header.Set("Referer", urlStr)
	//}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors2.Annotate(err, "创建请求出错")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		fmt.Println(string(b))
		return nil, errors2.Annotate(errors.New("请求返回状态码有误1"), "请求返回状态码有误1")
	}

	b, _ := io.ReadAll(resp.Body)

	var detail insItem
	streams := make(map[string]*extractors.Stream)

	err = json.Unmarshal(b, &detail)
	if err != nil {
		return nil, errors2.Annotate(err, "json 解析失败")
	}

	var images []string
	var isNotVideo bool
	dataType := extractors.DataTypeVideo

	var title string

	if len(detail.Data.ShortcodeMedia.EdgeMediaToCaption.Edges) > 0 {
		title = detail.Data.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
	}

	cover := detail.Data.ShortcodeMedia.DisplayUrl

	if detail.Data.ShortcodeMedia.IsVideo {
		stream := &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  detail.Data.ShortcodeMedia.VideoUrl,
					Size: 850,
				},
			},
			Size:    850,
			Quality: "1080p",
		}
		streams["1080"] = stream
	} else {
		dataType = extractors.DataTypeImage
		isNotVideo = true
		for _, edge := range detail.Data.ShortcodeMedia.EdgeSidecarToChildren.Edges {
			images = append(images, edge.Node.DisplayUrl)
		}
	}

	return []*extractors.Data{
		{
			Site:       "Instagram instagram.com",
			Title:      title,
			Type:       dataType,
			Streams:    streams,
			URL:        url,
			Image:      images,
			Cover:      cover,
			IsNotVideo: isNotVideo,
		},
	}, nil

}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {
	// Instagram is forcing a login to access the page, so we use the embed page to bypass that.
	url = strings.Replace(url, `/pink/p`, `/p`, -1)
	u, err := netURL.Parse(url)
	if err != nil {
		return nil, errors2.Annotate(errors.New("url解析失败"), "url解析失败")
	}

	u.Path = path.Join(u.Path, "embed")

	html, err := request.Get(u.String(), url, map[string]string{
		"Cookie":         "ig_did=88184AC9-0D63-4462-AA47-D0BFF15B3572; datr=JOpNZQjtwKw9_sjZ2aX8zARw; csrftoken=PES3HGWe06P0LHWg1Shn3vUAjtHsLuWT; ig_nrcb=1; mid=ZU3qJgAEAAG0owZzr_p7Yf2HuI4M; dpr=3",
		"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Viewport-Width": "756",
		"Sec-Fetch-Mode": "navigate",
		"Sec-Fetch-Dest": "document",
		"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dataStrings := utils.MatchOneOf(html, `gql_data.*\}\}`)
	if dataStrings == nil || len(dataStrings) < 1 {
		// 检查是否为图片类型
		if strings.Contains(html, "在 Instagram 观看") || strings.Contains(html, "Watch on Instagram") {
			return nil, errors2.Annotate(errors.New("未匹配到数据"), "未匹配到数据")
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

		if err != nil {
			return nil, errors2.Annotate(errors.New("html解析失败"), "html解析失败")
		}

		imgs := make([]string, 0)

		doc.Find("img.EmbeddedMediaImage").Each(func(i int, selection *goquery.Selection) {
			srcset := selection.AttrOr("srcset", "")

			scs := strings.Split(srcset, ",")

			var maxSize, maxImg string

			for _, v := range scs {
				item := strings.Split(v, " ")
				if len(item) > 1 {
					if item[1] > maxSize {
						maxSize = item[1]
						maxImg = item[0]
					}
				}
			}

			imgs = append(imgs, maxImg)
		})

		if len(imgs) <= 0 {
			return nil, errors2.Annotate(errors.New("未能匹配到"), "未能匹配到")
		}
		html, err = request.Get(url, url, map[string]string{
			"Cookie":         "ig_did=88184AC9-0D63-4462-AA47-D0BFF15B3572; datr=JOpNZQjtwKw9_sjZ2aX8zARw; csrftoken=PES3HGWe06P0LHWg1Shn3vUAjtHsLuWT; ig_nrcb=1; mid=ZU3qJgAEAAG0owZzr_p7Yf2HuI4M; dpr=3",
			"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Viewport-Width": "756",
			"Sec-Fetch-Mode": "navigate",
			"Sec-Fetch-Dest": "document",
			"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		doc, err = goquery.NewDocumentFromReader(strings.NewReader(html))

		if err != nil {
			return nil, errors2.Annotate(errors.New("html解析失败"), "html解析失败")
		}
		var title string
		doc.Find("title").Each(func(i int, selection *goquery.Selection) {
			title = selection.Text()
		})

		return []*extractors.Data{
			{
				Site:       "Instagram instagram.com",
				Title:      title,
				Type:       extractors.DataTypeImage,
				Streams:    nil,
				URL:        url,
				Image:      imgs,
				Cover:      imgs[0],
				IsNotVideo: true,
			},
		}, nil
	}
	dataString := strings.Replace(dataStrings[0], `\"`, `"`, -1)
	dataString = strings.Replace(dataString, `\\\/`, `/`, -1)
	dataString = strings.Replace(dataString, `\\u`, `\u`, -1)
	dataString = strings.Replace(dataString, `\\n`, ``, -1)
	dataString = `{"` + dataString

	streams := make(map[string]*extractors.Stream)

	var data instagram
	if err = json.Unmarshal([]byte(dataString), &data); err != nil {
		return nil, errors2.Annotate(errors.New("json解析失败"), "json解析失败")
	}
	var images []string
	var isNotVideo bool
	dataType := extractors.DataTypeVideo

	var title string

	if len(data.GqlData.ShortcodeMedia.EdgeMediaToCaption.Edges) > 0 {
		title = data.GqlData.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
	}

	cover := data.GqlData.ShortcodeMedia.DisplayUrl

	if data.GqlData.ShortcodeMedia.IsVideo {
		stream := &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  data.GqlData.ShortcodeMedia.VideoUrl,
					Size: 850,
				},
			},
			Size:    850,
			Quality: "1080p",
		}
		streams["1080"] = stream
	} else {
		dataType = extractors.DataTypeImage
		isNotVideo = true
		for _, edge := range data.GqlData.ShortcodeMedia.EdgeSidecarToChildren.Edges {
			images = append(images, edge.Node.DisplayUrl)
		}
	}

	return []*extractors.Data{
		{
			Site:       "Instagram instagram.com",
			Title:      title,
			Type:       dataType,
			Streams:    streams,
			URL:        url,
			Image:      images,
			Cover:      cover,
			IsNotVideo: isNotVideo,
		},
	}, nil
}

func EncodeBaseN(videoId string) int {

	if len(videoId) > 11 {
		videoId = videoId[:11]
	}

	var mapping = map[rune]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13,
		'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25,
		'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31, 'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37,
		'm': 38, 'n': 39, 'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47, 'w': 48, 'x': 49,
		'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55, '4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61,
		'-': 62, '_': 63,
	}

	result, base := 0, len(mapping)
	for _, char := range videoId {
		result = result*base + mapping[char]
	}
	return result
}
