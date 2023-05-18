package web

import (
	"encoding/json"
	"net/http"
	"net/url"

	"tutu-gin/http/validator/webValidator"

	"tutu-gin/core/api"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/valyala/fasthttp"
	"tutu-gin/core/exception"
	"tutu-gin/core/global"
	"tutu-gin/http/vo"
)

type User struct{}

func (u *User) Qrcode(c *gin.Context) {
	cacheKey := "paqiakerui_access_token"
	token, err := global.REDIS_CLIENT.Get(c, cacheKey).Result()
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	if len(token) <= 0 {
		code, body, err := fasthttp.Get(nil, "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxd6322dba9f40d541&secret=3e05992970e12dcaf1074264742c6e12")
		if err != nil {
			c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
			return
		}

		if code != 200 {
			c.Error(exception.ValidatorError(errors.Annotate(errors.New("授权失败"), exception.API_PARAMETER_CHECK_FAIL)))
			return
		}
		var accessTokenResponse vo.AccessTokenResponse
		err = json.Unmarshal(body, &accessTokenResponse)
		if err != nil || accessTokenResponse.Errcode != 0 {
			c.Error(exception.ValidatorError(errors.Annotate(errors.New("授权失败"), exception.API_PARAMETER_CHECK_FAIL)))
			return
		}

		global.REDIS_CLIENT.Set(c, cacheKey, accessTokenResponse.AccessToken, 6000)
		token = accessTokenResponse.AccessToken
	}

	err, qrcode := u.getCode(token)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(errors.New("生成二维码失败"), exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(map[string]string{
		"Url":    "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(qrcode.Ticket),
		"Ticket": qrcode.Ticket,
	}))
}

func (u *User) getCode(token string) (error, vo.QrcodeResponse) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token)

	jsonParams := map[string]any{
		"expire_seconds": 86400,
		"action_name":    "QR_SCENE",
		"action_info": map[string]any{
			"scene": map[string]any{
				"scene_str": "test111",
			},
		},
	}

	body, _ := json.Marshal(jsonParams)
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	var qrcodeResponse vo.QrcodeResponse

	if err := fasthttp.Do(req, resp); err != nil {
		return err, qrcodeResponse
	}

	err := json.Unmarshal(resp.Body(), &qrcodeResponse)
	if err != nil {
		return err, qrcodeResponse
	}

	return nil, qrcodeResponse
}

func (u *User) Check(c *gin.Context) {
	var requestData webValidator.UserQrcodeValidator

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	code, body, err := fasthttp.Get(nil, "https://jxapi.sucps.com/auth/ticketCheck?ticket="+requestData.Ticket)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	if code != 200 {
		c.Error(exception.ValidatorError(errors.Annotate(errors.New("授权失败"), exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	c.JSON(http.StatusOK, string(body))
}

func NewUser() *User {
	return &User{}
}
