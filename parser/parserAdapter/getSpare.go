package parserAdapter

import (
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"net/url"
	"tutu-gin/core/exception"
	"tutu-gin/core/global"
	"tutu-gin/parser/parserApplicaition/parserDto"
)

type GetSpare struct {
}

type GetSpareResult struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    GetSpareResultData `json:"data"`
}

type GetSpareResultData struct {
	PlayAddr string   `json:"playAddr"`
	Cover    string   `json:"cover"`
	Desc     string   `json:"desc"`
	Type     int      `json:"type"`
	Pics     []string `json:"pics"`
	Music    string   `json:"music"`
	Size     string   `json:"size"`
}

func (s *GetSpare) Fetch(dto *parserDto.GetSpareFetchDto) (result *parserDto.ParserResultDto, err error) {

	requestBody := url.Values{}
	requestBody.Set("token", global.SERVICE_CONFIG.ParserConfig.GetSpare.Token)
	requestBody.Set("link", dto.PageUrl)

	resp, err := http.PostForm(global.SERVICE_CONFIG.ParserConfig.GetSpare.RequestUrl, requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	var getSpareResult GetSpareResult

	err = json.NewDecoder(resp.Body).Decode(&getSpareResult)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(getSpareResult)
	if getSpareResult.Code != "0001" {
		log.Println(err)
		return nil, exception.PARSE_FAIL
	}
	log.Println(getSpareResult)
	result = &parserDto.ParserResultDto{
		Title:     getSpareResult.Data.Desc,
		CoverUrls: getSpareResult.Data.Cover,
		Pics:      getSpareResult.Data.Pics,
		VideoUrls: getSpareResult.Data.PlayAddr,
		IsVideo:   true,
	}

	if getSpareResult.Data.Type == 2 {
		result.IsVideo = false
	}

	return

}

func NewGetSpare() *GetSpare {
	return &GetSpare{}
}
