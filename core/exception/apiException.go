package exception

type ApiException struct {
	Code int
	Msg  string
	Data interface{}
}

func (a *ApiException) Error() string {
	return a.Msg
}

func ValidatorError(msg string) *ApiException {
	return &ApiException{
		Code: 201,
		Msg:  msg,
	}
}

var API_PARAMETER_CHECK_FAIL = ValidatorError("参数请求有误")
