package twitter

import (
	"encoding/json"
	"fmt"
	errors2 "github.com/juju/errors"
	"regexp"
	"strconv"
	"strings"
	"tutu-gin/lib/extractors"

	"github.com/pkg/errors"

	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	extractors.Register("twitter", New())
}

type twitter struct {
	Track struct {
		URL string `json:"playbackUrl"`
	} `json:"track"`
	TweetID  string
	Username string
}

type extractor struct{}

// New returns a twitter extractor.
func New() extractors.Extractor {
	return &extractor{}
}

type twitterInfo struct {
	CommunityNote  interface{}   `json:"communityNote"`
	ConversationID string        `json:"conversationID"`
	Date           string        `json:"date"`
	DateEpoch      int           `json:"date_epoch"`
	Hashtags       []interface{} `json:"hashtags"`
	Likes          int           `json:"likes"`
	MediaURLs      []string      `json:"mediaURLs"`
	MediaExtended  []struct {
		AltText        interface{} `json:"altText"`
		DurationMillis int         `json:"duration_millis"`
		Size           struct {
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"size"`
		ThumbnailUrl string `json:"thumbnail_url"`
		Type         string `json:"type"`
		Url          string `json:"url"`
	} `json:"media_extended"`
	PossiblySensitive   bool        `json:"possibly_sensitive"`
	QrtURL              interface{} `json:"qrtURL"`
	Replies             int         `json:"replies"`
	Retweets            int         `json:"retweets"`
	Text                string      `json:"text"`
	TweetID             string      `json:"tweetID"`
	TweetURL            string      `json:"tweetURL"`
	UserName            string      `json:"user_name"`
	UserProfileImageUrl string      `json:"user_profile_image_url"`
	UserScreenName      string      `json:"user_screen_name"`
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {

	if strings.Contains(url, "twitter.com") {
		pattern := `/status/(\d+)`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(url)

		if len(matches) > 1 {
			url = "https://api.vxtwitter.com/Twitter/status/" + matches[1]
		} else {
			url = strings.Replace(url, "twitter.com", "api.vxtwitter.com", -1)
		}
	} else {
		url = strings.Replace(url, "x.com", "api.vxtwitter.com", -1)
	}

	b, err := request.Get(url, url, map[string]string{
		"User-Agent": "TelegramBot (like TwitterBot)",
	})
	//resp, err := request.Request(http.MethodGet, strings.Replace(url, "twitter.com", "api.vxtwitter.com", -1), nil, map[string]string{
	//	"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	//})
	//if err != nil {
	//	return nil, errors2.Annotate(errors.New("twitter获取失败"), "twitter获取失败")
	//}
	//defer resp.Body.Close() // nolint
	//b, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(b))
	var detail twitterInfo
	err = json.Unmarshal([]byte(b), &detail)

	if err != nil {
		return nil, errors2.Annotate(errors.New("twitter的json解析失败"), "twitter的json解析失败")
	}

	if detail.MediaExtended == nil || len(detail.MediaExtended) == 0 {
		return nil, errors2.Annotate(errors.New("视频或图片不存在"), "视频或图片不存在")
	}

	quality := strconv.Itoa(detail.MediaExtended[0].Size.Height)
	streams := make(map[string]*extractors.Stream)
	var images []string
	var isNotVideo bool

	if detail.MediaExtended[0].Type == "image" {
		isNotVideo = true
		images = append(images, detail.MediaExtended[0].Url)
	} else {
		stream := &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  detail.MediaExtended[0].Url,
					Size: 850,
				},
			},
			Size:    850,
			Quality: quality,
		}
		streams[quality] = stream
	}

	return []*extractors.Data{
		{
			Site:       "Twitter twitter.com",
			Title:      detail.Text,
			Type:       extractors.DataTypeImage,
			Streams:    streams,
			URL:        url,
			Image:      images,
			Cover:      detail.MediaExtended[0].ThumbnailUrl,
			IsNotVideo: isNotVideo,
		},
	}, nil
}

func download(data twitter, uri string) ([]*extractors.Data, error) {
	var (
		err  error
		size int64
	)
	streams := make(map[string]*extractors.Stream)
	switch {
	// if video file is m3u8 and ts
	case strings.Contains(data.Track.URL, ".m3u8"):
		m3u8urls, err := utils.M3u8URLs(data.Track.URL)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for index, m3u8 := range m3u8urls {
			var totalSize int64
			ts, err := utils.M3u8URLs(m3u8)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			urls := make([]*extractors.Part, 0, len(ts))
			for _, i := range ts {
				size, err := request.Size(i, uri)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				temp := &extractors.Part{
					URL:  i,
					Size: size,
					Ext:  "ts",
				}
				totalSize += size
				urls = append(urls, temp)
			}
			qualityString := utils.MatchOneOf(m3u8, `/(\d+x\d+)/`)[1]
			quality := strconv.Itoa(index + 1)
			streams[quality] = &extractors.Stream{
				Parts:   urls,
				Size:    totalSize,
				Quality: qualityString,
			}
		}

	// if video file is mp4
	case strings.Contains(data.Track.URL, ".mp4"):
		size, err = request.Size(data.Track.URL, uri)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		urlData := &extractors.Part{
			URL:  data.Track.URL,
			Size: size,
			Ext:  "mp4",
		}
		streams["default"] = &extractors.Stream{
			Parts: []*extractors.Part{urlData},
			Size:  size,
		}
	}

	return []*extractors.Data{
		{
			Site:    "Twitter twitter.com",
			Title:   fmt.Sprintf("%s %s", data.Username, data.TweetID),
			Type:    extractors.DataTypeVideo,
			Streams: streams,
			URL:     uri,
		},
	}, nil
}
