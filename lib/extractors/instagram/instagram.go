package instagram

import (
	"encoding/json"
	"fmt"
	"github.com/iawia002/lux/request"
	errors2 "github.com/juju/errors"
	netURL "net/url"
	"path"
	"strings"
	"tutu-gin/utils"

	"github.com/pkg/errors"

	"tutu-gin/lib/extractors"
)

type instagram struct {
	GqlData struct {
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
	} `json:"gql_data"`
}

type extractor struct{}

// New returns a instagram extractor.
func New() extractors.Extractor {
	return &extractor{}
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
		fmt.Println(html)
		return nil, errors2.Annotate(errors.New("未匹配到数据"), "未匹配到数据")
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
	title := data.GqlData.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
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
