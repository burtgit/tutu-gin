package tiktok

import (
	"encoding/json"
	errors2 "github.com/juju/errors"
	"regexp"
	"strings"
	"tutu-gin/lib/extractors"
	yt_dl "tutu-gin/lib/yt-dl"

	"github.com/pkg/errors"

	"github.com/iawia002/lux/request"
)

func init() {
	extractors.Register("tiktok", New())
}

type extractor struct{}

type tiktokContent struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Duration     int    `json:"duration"`
	Timestamp    int    `json:"timestamp"`
	Creator      string `json:"creator"`
	Uploader     string `json:"uploader"`
	UploaderId   string `json:"uploader_id"`
	ViewCount    int    `json:"view_count"`
	LikeCount    int    `json:"like_count"`
	RepostCount  int    `json:"repost_count"`
	CommentCount int    `json:"comment_count"`
	Track        string `json:"track"`
	Artist       string `json:"artist"`
	ChannelId    string `json:"channel_id"`
	UploaderUrl  string `json:"uploader_url"`
	Formats      []struct {
		Url          string  `json:"url"`
		Ext          string  `json:"ext"`
		Width        int     `json:"width"`
		Height       int     `json:"height"`
		Protocol     string  `json:"protocol"`
		Resolution   string  `json:"resolution"`
		DynamicRange string  `json:"dynamic_range"`
		AspectRatio  float64 `json:"aspect_ratio"`
		Cookies      string  `json:"cookies"`
		HttpHeaders  struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
			Referer        string `json:"Referer"`
		} `json:"http_headers"`
		VideoExt string      `json:"video_ext"`
		AudioExt string      `json:"audio_ext"`
		Vbr      interface{} `json:"vbr"`
		Abr      interface{} `json:"abr"`
		Tbr      interface{} `json:"tbr"`
		FormatId string      `json:"format_id"`
		Format   string      `json:"format"`
	} `json:"formats"`
	Thumbnails []struct {
		Url        string `json:"url"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Id         string `json:"id"`
		Resolution string `json:"resolution"`
	} `json:"thumbnails"`
	HttpHeaders struct {
		UserAgent      string `json:"User-Agent"`
		Accept         string `json:"Accept"`
		AcceptLanguage string `json:"Accept-Language"`
		SecFetchMode   string `json:"Sec-Fetch-Mode"`
		Referer        string `json:"Referer"`
	} `json:"http_headers"`
	WebpageUrl         string      `json:"webpage_url"`
	OriginalUrl        string      `json:"original_url"`
	WebpageUrlBasename string      `json:"webpage_url_basename"`
	WebpageUrlDomain   string      `json:"webpage_url_domain"`
	Extractor          string      `json:"extractor"`
	ExtractorKey       string      `json:"extractor_key"`
	Playlist           interface{} `json:"playlist"`
	PlaylistIndex      interface{} `json:"playlist_index"`
	Thumbnail          string      `json:"thumbnail"`
	DisplayId          string      `json:"display_id"`
	Fulltitle          string      `json:"fulltitle"`
	DurationString     string      `json:"duration_string"`
	UploadDate         string      `json:"upload_date"`
	ReleaseYear        interface{} `json:"release_year"`
	Artists            []string    `json:"artists"`
	Creators           []string    `json:"creators"`
	RequestedSubtitles interface{} `json:"requested_subtitles"`
	HasDrm             interface{} `json:"_has_drm"`
	Epoch              int         `json:"epoch"`
	FormatId           string      `json:"format_id"`
	Url                string      `json:"url"`
	Ext                string      `json:"ext"`
	Width              int         `json:"width"`
	Height             int         `json:"height"`
	Protocol           string      `json:"protocol"`
	Resolution         string      `json:"resolution"`
	DynamicRange       string      `json:"dynamic_range"`
	AspectRatio        float64     `json:"aspect_ratio"`
	Cookies            string      `json:"cookies"`
	VideoExt           string      `json:"video_ext"`
	AudioExt           string      `json:"audio_ext"`
	Vbr                interface{} `json:"vbr"`
	Abr                interface{} `json:"abr"`
	Tbr                interface{} `json:"tbr"`
	Format             string      `json:"format"`
	Filename           string      `json:"_filename"`
	Filename1          string      `json:"filename"`
	Type               string      `json:"_type"`
	Version            struct {
		Version        string `json:"version"`
		CurrentGitHead string `json:"current_git_head"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}

// New returns a tiktok extractor.
func New() extractors.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {

	// yt-dl执行出来的结果

	res, err := yt_dl.Client(url)
	if err != nil {
		return nil, err
	}

	var data tiktokContent

	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, errors2.Annotate(err, "json 解析失败")
	}

	if len(data.Format) <= 2 {
		return nil, errors2.Annotate(errors.New("视频存在不可播放"), "视频存在不可播放")
	}

	html, err := request.Get(url, url, map[string]string{
		// tiktok require a user agent
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:98.0) Gecko/20100101 Firefox/98.0",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	urlMatcherRegExp := regexp.MustCompile(`"downloadAddr":\s*"([^"]+)"`)

	downloadURLMatcher := urlMatcherRegExp.FindStringSubmatch(html)

	if len(downloadURLMatcher) == 0 {
		return nil, errors.WithStack(extractors.ErrURLParseFailed)
	}

	videoURL := strings.ReplaceAll(downloadURLMatcher[1], `\u002F`, "/")

	titleMatcherRegExp := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`)

	titleMatcher := titleMatcherRegExp.FindStringSubmatch(html)

	if len(titleMatcher) == 0 {
		return nil, errors.WithStack(extractors.ErrURLParseFailed)
	}

	title := titleMatcher[1]

	titleArr := strings.Split(title, "|")

	if len(titleArr) == 1 {
		title = titleArr[0]
	} else {
		title = strings.TrimSpace(strings.Join(titleArr[:len(titleArr)-1], "|"))
	}

	streams := make(map[string]*extractors.Stream)

	size, err := request.Size(videoURL, url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	urlData := &extractors.Part{
		URL:  videoURL,
		Size: size,
		Ext:  "mp4",
	}
	streams["default"] = &extractors.Stream{
		Parts: []*extractors.Part{urlData},
		Size:  size,
	}

	return []*extractors.Data{
		{
			Site:    "TikTok tiktok.com",
			Title:   title,
			Type:    extractors.DataTypeVideo,
			Streams: streams,
			URL:     url,
		},
	}, nil
}
