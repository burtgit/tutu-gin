package parserMapper

import (
	"github.com/goccy/go-json"
	"io/ioutil"
	"path/filepath"
	"tutu-gin/core/exception"
)

type ParserJsonMapper struct {
	Code      string   `json:"code"`
	Name      string   `json:"name"`
	LogoImage string   `json:"logoImage"`
	Patterns  []string `json:"patterns"`
}

func (p *ParserJsonMapper) GetList() (data map[string]ParserJsonMapper, err error) {

	jsonPath, err := filepath.Abs("./parser/parserInfrastructure")
	if err != nil {
		return nil, exception.DOMAIN_NOT_FOUND
	}

	jsonData, err := ioutil.ReadFile(jsonPath + "/platform.json")

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, exception.JSON_PARSE_FAIL
	}

	return
}

func NewParserJsonMapper() *ParserJsonMapper {
	return &ParserJsonMapper{}
}
