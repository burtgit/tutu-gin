package parserApplicaition

import (
	"regexp"
	"strings"

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

func (p *ParserService) getParseURL(link string) (error, string) {
	oldLink := link
	link = strings.ReplaceAll(link, " ", "+@@@+")

	startIndex := strings.LastIndex(link, "http://")
	if startIndex == -1 {
		startIndex = strings.LastIndex(link, "https://")
	}
	if startIndex == -1 {
		return nil, oldLink
	}

	// 去掉前面的中文
	link = link[startIndex:]

	endIndex := strings.Index(link, "+@@@+")
	if endIndex != -1 {
		link = link[:endIndex]
	}

	endIndex = strings.Index(link, "，")
	if endIndex != -1 {
		link = link[:endIndex]
	}

	reg := regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z]{2,5}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	if len(link) < 16 || !reg.MatchString(link) {
		return errors.New("链接有误"), ""
	}

	return nil, link
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

	err, pageUrl = p.getParseURL(pageUrl)

	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_NOT_FOUND))
	}

	if platform.Code == "TENCENT_TV" {
		result, err = parserAdapter.NewGetLux().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
	} else if platform.Code == "ZHIHU" {
		result, err = parserAdapter.NewGetLux().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
	} else if platform.Code == "KUAISHOU" {
		result, err = parserAdapter.NewGetSpareKuaishou().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})
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

func (p *ParserService) Agent(pageUrl string) (result *parserDto.ParserResultDto, err error) {
	platform, err := p.platformRepository.GetByDomain(pageUrl)
	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_NOT_FOUND))
	}
	result, err = parserAdapter.NewGetLux().Fetch(&parserDto.GetSpareFetchDto{PageUrl: pageUrl, Platform: platform})

	return result, err
}

func NewParserService() *ParserService {
	return &ParserService{
		platformRepository: parserRepository.NewPlatformRepository(),
	}
}
