package exception

import (
	"fmt"
	"github.com/juju/errors"
)

type ApiException struct {
	Code int
	Data interface{}
	Msg  string
	Err  *errors.Err
}

var apiErrMaps = map[string]int{
	API_PARSER_FAIL:          201,
	API_PAY_FAIL:             201,
	API_PARAMETER_CHECK_FAIL: 201,
	API_REQUEST_FAIL:         201,
	NEED_VIP:                 202,
	NOT_LOGIN:                403,
	OVER_NEED_VIP:            204,
}

func (a *ApiException) Error() string {
	return a.Msg
}

func ValidatorError(err error) *ApiException {
	e, _ := err.(*errors.Err)
	fmt.Println(e.StackTrace())
	return &ApiException{
		Code: apiErrMaps[e.Message()],
		Msg:  e.Message(),
		Err:  e,
	}
}

var (
	API_PARAMETER_CHECK_FAIL = "参数请求有误"
	API_PARSER_FAIL          = "解析失败，请联系微信：kaolajiexi2"
	API_REQUEST_FAIL         = "请求超时，请联系微信：kaolajiexi2"
	API_PAY_FAIL             = "支付失败，请联系微信：kaolajiexi2"
	NOT_LOGIN                = "未登录"
	NEED_VIP                 = "次数不够"
	OVER_NEED_VIP            = "海外视频需要开通vip"
)
