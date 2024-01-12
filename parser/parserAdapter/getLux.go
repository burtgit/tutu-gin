package parserAdapter

import (
	"tutu-gin/lib/extractors"

	"tutu-gin/parser/parserApplicaition/parserDto"
)

type GetLux struct{}

func (s *GetLux) Fetch(dto *parserDto.GetSpareFetchDto) (result *parserDto.ParserResultDto, err error) {
	extract, err := extractors.Extract(dto.PageUrl, extractors.Options{})
	if err != nil {
		return nil, err
	}

	if len(extract) < 1 {
		return nil, err
	}

	item := extract[0]
	formats := make([]parserDto.ParseFormat, 0)

	if item.IsNotVideo {
		result = &parserDto.ParserResultDto{
			Title:     item.Title,
			IsVideo:   false,
			CoverUrls: item.Cover,
			Pics:      item.Image,
		}
	} else {
		var bigSize int64
		var bigUrl string

		for _, v := range item.Streams {

			if len(v.Parts) < 1 || v.Size <= 0 {
				continue
			}

			if v.Size > bigSize && !v.NoAudio {
				bigSize = v.Size
				bigUrl = v.Parts[0].URL
			}

			var separate int

			if v.NoAudio {
				separate = 1
			}

			formats = append(formats, parserDto.ParseFormat{
				QualityNote: v.Quality,
				Separate:    separate,
				Vext:        v.Ext,
				Video:       v.Parts[0].URL,
			})
		}

		if len(bigUrl) <= 0 && len(formats) > 0 {
			bigUrl = formats[0].Video
		}

		result = &parserDto.ParserResultDto{
			Title:     item.Title,
			VideoUrls: bigUrl,
			IsVideo:   true,
			Formats:   formats,
			CoverUrls: item.Cover,
		}
	}

	return
}

func NewGetLux() *GetLux {
	return &GetLux{}
}
