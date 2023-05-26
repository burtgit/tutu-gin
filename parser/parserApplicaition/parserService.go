package parserApplicaition

import (
	"github.com/juju/errors"
	"tutu-gin/core/event"
	"tutu-gin/core/exception"
	"tutu-gin/parser/parserAdapter"
	"tutu-gin/parser/parserApplicaition/parserDto"
	"tutu-gin/parser/parserDoMain/parseEvent"
	"tutu-gin/parser/parserDoMain/parserRepository"
)

type ParserService struct {
	platformRepository *parserRepository.PlatformRepository
}

// Parse 解析操作
func (p *ParserService) Parse(pageUrl string, ip string, userId int64) (result *parserDto.ParserResultDto, err error) {
	platform, err := p.platformRepository.GetByDomain(pageUrl)
	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_NOT_FOUND))
	}

	if platform == nil {
		return nil, exception.DomainError(errors.Annotate(errors.New("platform not found"), exception.DOMAIN_NOT_FOUND))
	}

	if platform.Code == "TENCENT_TV" {
		result, err = parserAdapter.NewGetLux().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
	} else if platform.Code == "ZHIHU" {
		result, err = parserAdapter.NewGetLux().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
	} else {
		getSpare := parserAdapter.NewGetSpare()
		result, err = getSpare.Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
	}

	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_PARSE_FAIL))
	}

	if userId != 1 {
		// 解析成功事件
		go event.EventHandler(&parseEvent.ParseSuccessEvent{
			ParserResult: result,
			Url:          pageUrl,
			UserId:       userId,
			Ip:           ip,
		})
	}
	return
}

func NewParserService() *ParserService {
	return &ParserService{
		platformRepository: parserRepository.NewPlatformRepository(),
	}
}
