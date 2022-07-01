package exception

import (
	"github.com/juju/errors"
	"log"
)

var domainErrMaps = map[string]int{
	"资源未找到": 404,
	"解析失败，请联系微信kaolajiexi2": 500,
	"json解析失败":              201,
	"数据库操作失败":               500,
	"请求失败":                  500,
}

type DomainException struct {
	Msg  string
	Code int
	Data interface{}
	Err  *errors.Err
}

func (d DomainException) Error() string {
	return d.Msg
}

func DomainError(err error) *DomainException {
	e, _ := err.(*errors.Err)
	log.Println(e.StackTrace())
	return &DomainException{
		Code: domainErrMaps[e.Message()],
		Msg:  e.Message(),
		Err:  e,
	}
}

var (
	DOMAIN_NOT_FOUND       = "资源未找到"
	DOMAIN_PARSE_FAIL      = "解析失败，请联系微信kaolajiexi2"
	DOMAIN_JSON_PARSE_FAIL = "json解析失败"
	DOMAIN_DB_ACTION_FAIL  = "数据库操作失败"
	DOMAIN_REQUEST_FAIL    = "请求失败"
)
