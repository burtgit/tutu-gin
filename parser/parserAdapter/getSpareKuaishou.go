package parserAdapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/juju/errors"

	"tutu-gin/core/exception"
	"tutu-gin/parser/parserApplicaition/parserDto"
)

type GetSpareKuaishou struct{}

func (s *GetSpareKuaishou) Fetch(dto *parserDto.GetSpareFetchDto) (result *parserDto.ParserResultDto, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURI("https://qrcode.layzz.cn/test/moreDyAnalyse")

	params := map[string]string{
		"link":  dto.PageUrl,
		"token": "cnbk#ucyitzkc-auther-hengtongtongxin",
	}

	b, _ := json.Marshal(params)
	fmt.Println(string(b))
	req.SetBody(b)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err = fasthttp.Do(req, resp); err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_JSON_PARSE_FAIL))
	}

	var getSpareResult GetSpareResult

	err = json.Unmarshal(resp.Body(), &getSpareResult)
	if err != nil {
		fmt.Println(string(resp.Body()))
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_JSON_PARSE_FAIL))
	}

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

func NewGetSpareKuaishou() *GetSpareKuaishou {
	return &GetSpareKuaishou{}
}
