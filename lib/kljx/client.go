package kljx

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	errors2 "github.com/juju/errors"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
	"tutu-gin/core/exception"
	"tutu-gin/lib/kljx/response"
)

type client[T response.ClientResponse] struct {
	host string
}

func (c *client[T]) Apply(router string, params map[string]string) (error, T) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(c.host + router)

	params["t"] = strconv.FormatInt(time.Now().Unix(), 10)

	v, ok := params["token"]
	if !ok {
		v = "lilililili"
	}

	vv, ok := params["pageUrl"]
	if ok {
		v = v + vv
	}

	params["s"] = fmt.Sprintf("%x", md5.Sum([]byte(params["t"]+v+"lihuanjie111111111")))

	b, _ := json.Marshal(params)
	req.SetBody(b)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	var result response.Response[T]

	if err := fasthttp.DoTimeout(req, resp, time.Minute*5); err != nil {
		return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.API_REQUEST_FAIL)), result.Data
	}

	err := json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.API_REQUEST_FAIL)), result.Data
	}

	if result.Status != 200 {
		if result.Status == 403 || result.Msg == "未登录" {
			return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.NOT_LOGIN)), result.Data
		} else if result.Status == 202 {
			return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.NEED_VIP)), result.Data
		} else if result.Status == 204 {
			return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.OVER_NEED_VIP)), result.Data
		}
		return exception.ValidatorError(errors2.Annotate(errors.New(result.Msg), exception.API_PARSER_FAIL)), result.Data
	}

	return nil, result.Data
}

func NewClient[T response.ClientResponse]() *client[T] {
	return &client[T]{
		//host: "http://localhost:3000",
		host: "https://jxapi.sucps.com/",
	}
}
