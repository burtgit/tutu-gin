package parserRepository

import (
	"github.com/juju/errors"
	"strings"
	"tutu-gin/core/exception"
	"tutu-gin/parser/parserDoMain"
	"tutu-gin/parser/parserInfrastructure/parserMapper"
)

type PlatformRepository struct {
	mapper *parserMapper.ParserJsonMapper
}

// GetByDomain 获取链接获取平台信息
func (p *PlatformRepository) GetByDomain(url string) (platform *parserDoMain.Platform, err error) {

	platform = &parserDoMain.Platform{
		Code:   "Default",
		Name:   "default",
		Domain: []string{},
	}

	platformList, err := p.getList()
	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_NOT_FOUND))
	}

	for _, v := range platformList {
		for _, domainItem := range v.Domain {
			if strings.Contains(url, domainItem) {
				platform = v
				return
			}
		}
	}

	return
}

// 获取平台信息
func (p *PlatformRepository) getList() (platformList []*parserDoMain.Platform, err error) {

	// 这边加入缓存
	data, err := p.mapper.GetList()
	if err != nil {
		return nil, exception.DomainError(errors.Annotate(err, exception.DOMAIN_NOT_FOUND))
	}

	for _, v := range data {
		platformList = append(platformList, parserDoMain.NewPlatform(v.Code, v.Name, v.Patterns))
	}

	return
}

func NewPlatformRepository() *PlatformRepository {
	return &PlatformRepository{
		mapper: parserMapper.NewParserJsonMapper(),
	}
}
