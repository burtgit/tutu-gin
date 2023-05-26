package geekbang

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/iawia002/lux/extractors"
	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	extractors.Register("geekbang", New())
}

type geekData struct {
	Code  int             `json:"code"`
	Error json.RawMessage `json:"error"`
	Data  struct {
		VideoID      string `json:"video_id"`
		Title        string `json:"article_sharetitle"`
		ColumnHadSub bool   `json:"column_had_sub"`
	} `json:"data"`
}

type videoPlayAuth struct {
	Code  int             `json:"code"`
	Error json.RawMessage `json:"error"`
	Data  struct {
		PlayAuth string `json:"play_auth"`
	} `json:"data"`
}

type playInfo struct {
	VideoBase struct {
		VideoID  string `json:"VideoId"`
		Title    string `json:"Title"`
		CoverURL string `josn:"CoverURL"`
	} `json:"VideoBase"`
	PlayInfoList struct {
		PlayInfo []struct {
			URL        string `json:"PlayURL"`
			Size       int64  `json:"Size"`
			Definition string `json:"Definition"`
		} `json:"PlayInfo"`
	} `json:"PlayInfoList"`
}

type geekURLInfo struct {
	URL  string
	Size int64
}

func geekM3u8(url string) ([]geekURLInfo, error) {
	var (
		data []geekURLInfo
		temp geekURLInfo
		size int64
		err  error
	)
	urls, err := utils.M3u8URLs(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, u := range urls {
		temp = geekURLInfo{
			URL:  u,
			Size: size,
		}
		data = append(data, temp)
	}
	return data, nil
}

type extractor struct{}

// New returns a geekbang extractor.
func New() extractors.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, _ extractors.Options) ([]*extractors.Data, error) {
	var err error
	matches := utils.MatchOneOf(url, `https?://time.geekbang.org/course/detail/(\d+)-(\d+)`)
	if matches == nil || len(matches) < 3 {
		return nil, errors.WithStack(extractors.ErrURLParseFailed)
	}

	// Get video information
	heanders := map[string]string{"Origin": "https://time.geekbang.org", "Content-Type": "application/json", "Referer": url}
	params := strings.NewReader(fmt.Sprintf(`{"id": %q}`, matches[2]))
	res, err := request.Request(http.MethodPost, "https://time.geekbang.org/serv/v1/article", params, heanders)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close() // nolint

	var data geekData
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, errors.WithStack(err)
	}

	if data.Code < 0 {
		return nil, errors.New(string(data.Error))
	}

	if data.Data.VideoID == "" && !data.Data.ColumnHadSub {
		return nil, errors.New("请先购买课程，或使用Cookie登录。")
	}

	// Get video license token information
	params = strings.NewReader("{\"source_type\":1,\"aid\":" + matches[2] + ",\"video_id\":\"" + data.Data.VideoID + "\"}")
	res, err = request.Request(http.MethodPost, "https://time.geekbang.org/serv/v3/source_auth/video_play_auth", params, heanders)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close() // nolint

	var playAuth videoPlayAuth
	if err = json.NewDecoder(res.Body).Decode(&playAuth); err != nil {
		return nil, errors.WithStack(err)
	}

	if playAuth.Code < 0 {
		return nil, errors.New(string(playAuth.Error))
	}

	// Get video playback information
	heanders = map[string]string{"Accept-Encoding": ""}
	res, err = request.Request(http.MethodGet, "http://ali.mantv.top/play/info?playAuth="+playAuth.Data.PlayAuth, nil, heanders)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close() // nolint

	var playInfo playInfo
	if err = json.NewDecoder(res.Body).Decode(&playInfo); err != nil {
		return nil, errors.WithStack(err)
	}

	title := data.Data.Title

	streams := make(map[string]*extractors.Stream, len(playInfo.PlayInfoList.PlayInfo))

	for _, media := range playInfo.PlayInfoList.PlayInfo {
		m3u8URLs, err := geekM3u8(media.URL)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		urls := make([]*extractors.Part, len(m3u8URLs))
		for index, u := range m3u8URLs {
			urls[index] = &extractors.Part{
				URL:  u.URL,
				Size: u.Size,
				Ext:  "ts",
			}
		}

		streams[media.Definition] = &extractors.Stream{
			Parts: urls,
			Size:  media.Size,
		}
	}

	return []*extractors.Data{
		{
			Site:    "极客时间 geekbang.org",
			Title:   title,
			Type:    extractors.DataTypeVideo,
			Streams: streams,
			URL:     url,
		},
	}, nil
}
