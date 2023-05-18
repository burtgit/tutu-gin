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

	getSpare := parserAdapter.NewGetSpare()
	result, err = getSpare.Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})

	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_PARSE_FAIL))
	}

	// 解析成功事件
	go event.EventHandler(&parseEvent.ParseSuccessEvent{
		ParserResult: result,
		Url:          pageUrl,
		UserId:       userId,
		Ip:           ip,
	})
	return
}

func NewParserService() *ParserService {
	return &ParserService{
		platformRepository: parserRepository.NewPlatformRepository(),
	}
}
