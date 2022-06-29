package api

type ApiSuccesss struct {
	Code int
	Msg  string
	Data interface{}
}

func ApiSuccessResponse(data interface{}) *ApiSuccesss {
	return &ApiSuccesss{
		Code: 200,
		Msg:  "请求成功",
		Data: data,
	}
}
