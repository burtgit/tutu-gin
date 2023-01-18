package parserAdapter

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/juju/errors"
	"io"
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
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_REQUEST_FAIL))
	}
	defer resp.Body.Close()
	var getSpareResult GetSpareResult
	err = json.NewDecoder(resp.Body).Decode(&getSpareResult)
	if err != nil {
		ss, _ := io.ReadAll(resp.Body)
		fmt.Println(string(ss))
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_JSON_PARSE_FAIL))
	}
	log.Println(getSpareResult)
	if getSpareResult.Code != "0001" {
		return nil, exception.DomainError(errors.Annotate(errors.New(getSpareResult.Message), exception.DOMAIN_PARSE_FAIL))
	}

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
