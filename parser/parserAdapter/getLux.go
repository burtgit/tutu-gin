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

	var bigSize int64
	var bigUrl string

	for _, v := range item.Streams {

		if len(v.Parts) < 1 || v.Size <= 0 {
			continue
		}

		if v.Size > bigSize {
			bigSize = v.Size
			bigUrl = v.Parts[0].URL
		}
		formats = append(formats, parserDto.ParseFormat{
			QualityNote: v.Quality,
			Separate:    0,
			Vext:        v.Ext,
			Video:       v.Parts[0].URL,
		})
	}

	result = &parserDto.ParserResultDto{
		Title:     item.Title,
		VideoUrls: bigUrl,
		IsVideo:   true,
		Formats:   formats,
		CoverUrls: item.Cover,
	}

	return
}

func NewGetLux() *GetLux {
	return &GetLux{}
}
