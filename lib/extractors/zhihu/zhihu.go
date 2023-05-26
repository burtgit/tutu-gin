package zhihu

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/iawia002/lux/request"

	"tutu-gin/utils"

	"tutu-gin/lib/extractors"

	"github.com/pkg/errors"
)

const (
	videoURL = "www.zhihu.com/zvideo"
	api      = "https://lens.zhihu.com/api/v4/videos/"
)

type extractor struct{}

func New() extractors.Extractor {
	return &extractor{}
}

func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {
	if !strings.Contains(url, videoURL) {
		return nil, errors.WithStack(extractors.ErrURLParseFailed)
	}

	html, err := request.Get(url, url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	videoID := utils.MatchOneOf(html, `"videoId":"(\d+)"`)
	titleMatch := utils.MatchOneOf(html, `<title.*?>(.*?)</title>`)

	if len(videoID) <= 1 {
		return nil, errors.New("zhihu video id extract failed")
	}

	title := "Unknown"
	if len(titleMatch) > 1 {
		title = titleMatch[1]
	}

	resp, err := request.GetByte(fmt.Sprintf("%s%s", api, videoID[1]), url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var data video
	if err = json.Unmarshal(resp, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	streams := make(map[string]*extractors.Stream)
	resolutions := map[string]resolution{
		"FHD": data.PlayList.FHD,
		"HD":  data.PlayList.HD,
		"SD":  data.PlayList.SD,
		"LD":  data.PlayList.LD,
	}

	if data.PlayListV2.SD.Size > 0 {
		resolutions["SD"] = data.PlayListV2.SD
	}

	if data.PlayListV2.HD.Size > 0 {
		resolutions["HD"] = data.PlayListV2.HD
	}

	if data.PlayListV2.LD.Size > 0 {
		resolutions["LD"] = data.PlayListV2.LD
	}

	if data.PlayListV2.FHD.Size > 0 {
		resolutions["FHD"] = data.PlayListV2.FHD
	}

	for k, v := range resolutions {
		stream := &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  v.PlayURL,
					Size: v.Size,
					Ext:  v.Format,
				},
			},
			Size:    v.Size,
			Quality: strconv.Itoa(v.Width),
		}
		streams[k] = stream
	}
	return []*extractors.Data{
		{
			Site:    "知乎 zhihu.com",
			Title:   title,
			Streams: streams,
			Cover:   data.CoverUrl,
			Type:    extractors.DataTypeVideo,
			URL:     url,
		},
	}, nil
}
