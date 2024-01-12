package youtube

import (
	"net/http"
	"strings"
	"tutu-gin/lib/extractors"

	"github.com/iawia002/lia/array"
	"github.com/kkdai/youtube/v2"
	"github.com/pkg/errors"

	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	e := New()
	extractors.Register("youtube", e)
	extractors.Register("youtu", e) // youtu.be
}

const referer = "https://www.youtube.com"

type extractor struct {
	client *youtube.Client
}

// New returns a youtube extractor.
func New() extractors.Extractor {
	return &extractor{
		client: &youtube.Client{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
				},
			},
		},
	}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {
	if !option.Playlist {
		video, err := e.client.GetVideo(url)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return []*extractors.Data{e.youtubeDownload(url, video)}, nil
	}

	playlist, err := e.client.GetPlaylist(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	needDownloadItems := utils.NeedDownloadList(option.Items, option.ItemStart, option.ItemEnd, len(playlist.Videos))
	extractedData := make([]*extractors.Data, len(needDownloadItems))
	wgp := utils.NewWaitGroupPool(option.ThreadNumber)
	dataIndex := 0
	for index, videoEntry := range playlist.Videos {
		if !array.ItemInArray(index+1, needDownloadItems) {
			continue
		}

		wgp.Add()
		go func(index int, entry *youtube.PlaylistEntry, extractedData []*extractors.Data) {
			defer wgp.Done()
			video, err := e.client.VideoFromPlaylistEntry(entry)
			if err != nil {
				return
			}
			extractedData[index] = e.youtubeDownload(url, video)
		}(dataIndex, videoEntry, extractedData)
		dataIndex++
	}
	wgp.Wait()
	return extractedData, nil
}

// youtubeDownload download function for single url
func (e *extractor) youtubeDownload(url string, video *youtube.Video) *extractors.Data {
	streams := make(map[string]*extractors.Stream, len(video.Formats))

	for i := range video.Formats {
		f := &video.Formats[i]

		if !strings.Contains(f.MimeType, "video/mp4") {
			continue
		}

		stream, ok := streams[f.Quality]
		var hasAudio bool
		if f.AudioChannels <= 0 {
			hasAudio = true
		}

		if ok && stream.NoAudio && f.AudioChannels > 0 {
			stream = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL: f.URL,
					},
				},
				Quality: f.QualityLabel,
				NeedMux: true,
				NoAudio: hasAudio,
			}
		}

		stream = &extractors.Stream{
			Parts: []*extractors.Part{
				{
					URL:  f.URL,
					Size: int64(f.Height),
				},
			},
			Quality: f.QualityLabel,
			NeedMux: true,
			NoAudio: hasAudio,
		}
		streams[f.Quality] = stream
	}

	var maxSize uint
	var cover string

	for _, thumbnail := range video.Thumbnails {
		if thumbnail.Height > maxSize {
			cover = thumbnail.URL
		}
	}

	return &extractors.Data{
		Site:    "YouTube youtube.com",
		Title:   video.Title,
		Type:    "video",
		Streams: streams,
		URL:     url,
		Cover:   cover,
	}
}

func (e *extractor) genPartByFormat(video *youtube.Video, f *youtube.Format) (*extractors.Part, error) {
	ext := getStreamExt(f.MimeType)
	url, err := e.client.GetStreamURL(video, f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	size := f.ContentLength
	if size == 0 {
		size, _ = request.Size(url, referer)
	}
	return &extractors.Part{
		URL:  url,
		Size: size,
		Ext:  ext,
	}, nil
}

func getVideoAudio(v *youtube.Video, mimeType string) (*youtube.Format, error) {
	audioFormats := v.Formats.Type(mimeType).Type("audio")
	if len(audioFormats) == 0 {
		return nil, errors.New("no audio format found after filtering")
	}
	audioFormats.Sort()
	return &audioFormats[0], nil
}

func getStreamExt(streamType string) string {
	// video/webm; codecs="vp8.0, vorbis" --> webm
	exts := utils.MatchOneOf(streamType, `(\w+)/(\w+);`)
	if exts == nil || len(exts) < 3 {
		return ""
	}
	return exts[2]
}
