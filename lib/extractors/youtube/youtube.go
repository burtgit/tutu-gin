package youtube

import (
	"fmt"
	"github.com/iawia002/lia/array"
	"net/http"
	"strconv"
	"strings"
	"tutu-gin/lib/extractors"
	"tutu-gin/lib/youtube"
	//"github.com/kkdai/youtube/v2"
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

type youtubeResponse struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Formats []struct {
		FormatId   string  `json:"format_id"`
		FormatNote string  `json:"format_note,omitempty"`
		Ext        string  `json:"ext"`
		Protocol   string  `json:"protocol"`
		Acodec     string  `json:"acodec,omitempty"`
		Vcodec     string  `json:"vcodec"`
		Url        string  `json:"url"`
		Width      int     `json:"width,omitempty"`
		Height     int     `json:"height,omitempty"`
		Fps        float64 `json:"fps,omitempty"`
		Rows       int     `json:"rows,omitempty"`
		Columns    int     `json:"columns,omitempty"`
		Fragments  []struct {
			Url      string  `json:"url"`
			Duration float64 `json:"duration"`
		} `json:"fragments,omitempty"`
		Resolution  string  `json:"resolution"`
		AspectRatio float64 `json:"aspect_ratio"`
		HttpHeaders struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		AudioExt           string      `json:"audio_ext"`
		VideoExt           string      `json:"video_ext"`
		Vbr                float64     `json:"vbr"`
		Abr                float64     `json:"abr"`
		Tbr                float64     `json:"tbr"`
		Format             string      `json:"format"`
		FormatIndex        interface{} `json:"format_index"`
		ManifestUrl        string      `json:"manifest_url,omitempty"`
		Language           interface{} `json:"language"`
		Preference         interface{} `json:"preference"`
		Quality            float64     `json:"quality,omitempty"`
		HasDrm             bool        `json:"has_drm,omitempty"`
		SourcePreference   int         `json:"source_preference,omitempty"`
		Asr                int         `json:"asr,omitempty"`
		Filesize           int         `json:"filesize,omitempty"`
		AudioChannels      int         `json:"audio_channels,omitempty"`
		LanguagePreference int         `json:"language_preference,omitempty"`
		DynamicRange       string      `json:"dynamic_range,omitempty"`
		Container          string      `json:"container,omitempty"`
		DownloaderOptions  struct {
			HttpChunkSize int `json:"http_chunk_size"`
		} `json:"downloader_options,omitempty"`
		FilesizeApprox int `json:"filesize_approx,omitempty"`
	} `json:"formats"`
	Thumbnails []struct {
		Url        string `json:"url"`
		Preference int    `json:"preference"`
		Id         string `json:"id"`
		Height     int    `json:"height,omitempty"`
		Width      int    `json:"width,omitempty"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"thumbnails"`
	Thumbnail         string      `json:"thumbnail"`
	Description       string      `json:"description"`
	ChannelId         string      `json:"channel_id"`
	ChannelUrl        string      `json:"channel_url"`
	Duration          int         `json:"duration"`
	ViewCount         int         `json:"view_count"`
	AverageRating     interface{} `json:"average_rating"`
	AgeLimit          int         `json:"age_limit"`
	WebpageUrl        string      `json:"webpage_url"`
	Categories        []string    `json:"categories"`
	Tags              []string    `json:"tags"`
	PlayableInEmbed   bool        `json:"playable_in_embed"`
	LiveStatus        string      `json:"live_status"`
	ReleaseTimestamp  interface{} `json:"release_timestamp"`
	FormatSortFields  []string    `json:"_format_sort_fields"`
	AutomaticCaptions struct {
	} `json:"automatic_captions"`
	Subtitles struct {
	} `json:"subtitles"`
	CommentCount int `json:"comment_count"`
	Chapters     []struct {
		StartTime float64 `json:"start_time"`
		Title     string  `json:"title"`
		EndTime   float64 `json:"end_time"`
	} `json:"chapters"`
	Heatmap []struct {
		StartTime float64 `json:"start_time"`
		EndTime   float64 `json:"end_time"`
		Value     float64 `json:"value"`
	} `json:"heatmap"`
	LikeCount            int         `json:"like_count"`
	Channel              string      `json:"channel"`
	ChannelFollowerCount int         `json:"channel_follower_count"`
	Uploader             string      `json:"uploader"`
	UploaderId           string      `json:"uploader_id"`
	UploaderUrl          string      `json:"uploader_url"`
	UploadDate           string      `json:"upload_date"`
	Availability         string      `json:"availability"`
	OriginalUrl          string      `json:"original_url"`
	WebpageUrlBasename   string      `json:"webpage_url_basename"`
	WebpageUrlDomain     string      `json:"webpage_url_domain"`
	Extractor            string      `json:"extractor"`
	ExtractorKey         string      `json:"extractor_key"`
	Playlist             interface{} `json:"playlist"`
	PlaylistIndex        interface{} `json:"playlist_index"`
	DisplayId            string      `json:"display_id"`
	Fulltitle            string      `json:"fulltitle"`
	DurationString       string      `json:"duration_string"`
	ReleaseYear          interface{} `json:"release_year"`
	IsLive               bool        `json:"is_live"`
	WasLive              bool        `json:"was_live"`
	RequestedSubtitles   interface{} `json:"requested_subtitles"`
	HasDrm               interface{} `json:"_has_drm"`
	Epoch                int         `json:"epoch"`
	RequestedFormats     []struct {
		FormatId         string      `json:"format_id"`
		FormatIndex      interface{} `json:"format_index"`
		Url              string      `json:"url"`
		ManifestUrl      string      `json:"manifest_url,omitempty"`
		Tbr              float64     `json:"tbr"`
		Ext              string      `json:"ext"`
		Fps              float64     `json:"fps"`
		Protocol         string      `json:"protocol"`
		Preference       interface{} `json:"preference"`
		Quality          float64     `json:"quality"`
		HasDrm           bool        `json:"has_drm"`
		Width            int         `json:"width"`
		Height           int         `json:"height"`
		Vcodec           string      `json:"vcodec"`
		Acodec           string      `json:"acodec"`
		DynamicRange     string      `json:"dynamic_range"`
		SourcePreference int         `json:"source_preference"`
		FormatNote       string      `json:"format_note"`
		Resolution       string      `json:"resolution"`
		AspectRatio      float64     `json:"aspect_ratio"`
		HttpHeaders      struct {
			UserAgent      string `json:"User-Agent"`
			Accept         string `json:"Accept"`
			AcceptLanguage string `json:"Accept-Language"`
			SecFetchMode   string `json:"Sec-Fetch-Mode"`
		} `json:"http_headers"`
		VideoExt           string      `json:"video_ext"`
		AudioExt           string      `json:"audio_ext"`
		Abr                float64     `json:"abr"`
		Vbr                float64     `json:"vbr"`
		Format             string      `json:"format"`
		Asr                int         `json:"asr,omitempty"`
		Filesize           int         `json:"filesize,omitempty"`
		AudioChannels      int         `json:"audio_channels,omitempty"`
		Language           interface{} `json:"language"`
		LanguagePreference int         `json:"language_preference,omitempty"`
		Container          string      `json:"container,omitempty"`
		DownloaderOptions  struct {
			HttpChunkSize int `json:"http_chunk_size"`
		} `json:"downloader_options,omitempty"`
	} `json:"requested_formats"`
	Format         string      `json:"format"`
	FormatId       string      `json:"format_id"`
	Ext            string      `json:"ext"`
	Protocol       string      `json:"protocol"`
	Language       interface{} `json:"language"`
	FormatNote     string      `json:"format_note"`
	FilesizeApprox int         `json:"filesize_approx"`
	Tbr            float64     `json:"tbr"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Resolution     string      `json:"resolution"`
	Fps            float64     `json:"fps"`
	DynamicRange   string      `json:"dynamic_range"`
	Vcodec         string      `json:"vcodec"`
	Vbr            float64     `json:"vbr"`
	StretchedRatio interface{} `json:"stretched_ratio"`
	AspectRatio    float64     `json:"aspect_ratio"`
	Acodec         string      `json:"acodec"`
	Abr            float64     `json:"abr"`
	Asr            int         `json:"asr"`
	AudioChannels  int         `json:"audio_channels"`
	Filename       string      `json:"_filename"`
	Filename1      string      `json:"filename"`
	Type           string      `json:"_type"`
	Version        struct {
		Version        string `json:"version"`
		CurrentGitHead string `json:"current_git_head"`
		ReleaseGitHead string `json:"release_git_head"`
		Repository     string `json:"repository"`
	} `json:"_version"`
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {
	if !option.Playlist {
		count := 3
		video, err := e.client.GetVideo(url)
		if err != nil {
			// 发现返回400，直接循环执行10次
			if strings.Contains(err.Error(), "status code: 400") {
				for i := 0; i < count; i++ {
					fmt.Println("油管的正在尝试....", i+1)
					video, err = e.client.GetVideo(url)
					if err == nil {
						break
					}
				}
			}
			if err != nil {
				return nil, err
			}
		}
		return []*extractors.Data{e.youtubeDownload(url, video)}, nil
		//res, err := yt_dl.Client(url)
		//
		//if err != nil {
		//	return nil, err
		//}
		//
		//var resultJson youtubeResponse
		//
		//err = json.Unmarshal([]byte(res), &resultJson)
		//if err == nil {
		//	return []*extractors.Data{e.youtubeDownloadV2(url, resultJson)}, nil
		//}
		//
		//return nil, errors2.Annotate(err, "解析body失败")
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
		f := video.Formats[i]
		quality := strconv.Itoa(f.Height)

		if !strings.Contains(f.MimeType, "video/mp4") && !strings.Contains(f.MimeType, "video/webm") {
			continue
		}

		stream, ok := streams[quality]
		var hasAudio bool
		if f.AudioChannels <= 0 {
			hasAudio = true
		}

		ext := "mp4"

		if strings.Contains(f.MimeType, "video/webm") {
			ext = "webm"
		}

		if ok {
			if stream.NoAudio && f.AudioChannels > 0 {
				if len(f.URL) <= 0 {
					f.URL, _ = e.client.GetStreamURL(video, &f)
				}
				stream = &extractors.Stream{
					Parts: []*extractors.Part{
						{
							URL:  f.URL,
							Size: int64(f.Height),
							Ext:  ext,
						},
					},
					Quality: f.QualityLabel,
					NeedMux: true,
					NoAudio: hasAudio,
				}
			}
		} else {
			if len(f.URL) <= 0 {
				f.URL, _ = e.client.GetStreamURL(video, &f)
			}
			stream = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL:  f.URL,
						Size: int64(f.Height),
						Ext:  ext,
					},
				},
				Quality: f.QualityLabel,
				NeedMux: true,
				NoAudio: hasAudio,
			}
		}

		streams[quality] = stream
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

// youtubeDownload download function for single url
func (e *extractor) youtubeDownloadV2(url string, video youtubeResponse) *extractors.Data {
	streams := make(map[string]*extractors.Stream, len(video.Formats))

	for i := range video.Formats {
		f := video.Formats[i]

		if f.Height <= 0 {
			continue
		}

		quality := strconv.Itoa(f.Height)

		if !strings.Contains(f.Protocol, "https") {
			continue
		}

		if !strings.Contains(f.Ext, "mp4") && !strings.Contains(f.Ext, "webm") {
			continue
		}

		stream, ok := streams[quality]
		var hasAudio bool
		if f.AudioChannels <= 0 {
			hasAudio = true
		}

		ext := "mp4"

		if strings.Contains(f.Ext, "webm") {
			ext = "webm"
		}

		if ok {
			if stream.NoAudio && f.AudioChannels > 0 {
				if len(f.Url) <= 0 {
					continue
				}
				stream = &extractors.Stream{
					Parts: []*extractors.Part{
						{
							URL:  f.Url,
							Size: int64(f.Height),
							Ext:  ext,
						},
					},
					Quality: f.FormatNote,
					NeedMux: true,
					NoAudio: hasAudio,
				}
			}
		} else {
			if len(f.Url) <= 0 {
				continue
			}
			stream = &extractors.Stream{
				Parts: []*extractors.Part{
					{
						URL:  f.Url,
						Size: int64(f.Height),
						Ext:  ext,
					},
				},
				Quality: f.FormatNote,
				NeedMux: true,
				NoAudio: hasAudio,
			}
		}

		streams[quality] = stream
	}

	var maxSize int
	var cover string

	for _, thumbnail := range video.Thumbnails {
		if thumbnail.Height > maxSize {
			cover = thumbnail.Url
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

func lazyLast(iterable []interface{}) (bool, interface{}) {
	iterator := iterable
	var prev interface{}
	if len(iterator) > 0 {
		prev = iterator[0]
		iterator = iterator[1:]
	} else {
		return false, nil
	}

	for _, item := range iterator {
		yield := []interface{}{false, prev}
		// Process yield value here (e.g., print, store, etc.)
		_ = yield
		prev = item
	}

	yield := []interface{}{true, prev}
	// Process yield value here (e.g., print, store, etc.)
	_ = yield

	return yield[0].(bool), yield[1]
}
