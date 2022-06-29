package parserApplicaition

import (
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
func (p *ParserService) Parse(pageUrl string, ip string) (result *parserDto.ParserResultDto, err error) {

	platform, err := p.platformRepository.GetByDomain(pageUrl)

	if err != nil || platform == nil {
		return nil, exception.DOMAIN_NOT_FOUND
	}

	getSpare := parserAdapter.NewGetSpare()
	result, err = getSpare.Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})

	// 解析成功事件
	if err == nil {
		go event.EventHandler(&parseEvent.ParseSuccessEvent{
			ParserResult: result,
			Url:          pageUrl,
			UserId:       11,
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
