package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"tutu-gin/core/api"

	"github.com/valyala/fasthttp"

	"github.com/juju/errors"

	"tutu-gin/core/exception"
	"tutu-gin/http/validator/webValidator"

	"github.com/gin-gonic/gin"
)

type Pay struct{}

func (i *Pay) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "pay.html", nil)
}

func (i *Pay) Mobile(c *gin.Context) {
	c.HTML(http.StatusOK, "mobile_pay.html", nil)
}

func (i *Pay) Apply(c *gin.Context) {
	var requestData webValidator.PayApplyValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	type Result struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
		Data   struct {
			Money      string `json:"money"`
			OrderId    int64  `json:"order_id"`
			PayType    string `json:"pay_type"`
			NeedRemark bool   `json:"need_remark"`
			UserId     int64  `json:"user_id"`
			Qrcode     string `json:"qrcode"`
		} `json:"data"`
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURI("http://localhost:3000/app/pay/apply")

	token, _ := c.Cookie("tokens")

	req.SetBody([]byte(`{"menu":` + strconv.FormatInt(requestData.Vip, 10) + `,"pay_type":"` + strconv.FormatInt(requestData.PayMethod, 10) + `","token":"` + token + `"}`))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PAY_FAIL)))
		return
	}

	var result Result

	err := json.Unmarshal(resp.Body(), &result)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PAY_FAIL)))
		return
	}

	payUrl := "/v1/pay/wechat"

	if requestData.PayMethod == 1 {
		payUrl = "/v1/pay/alipay"
	}

	link := payUrl + "?user_id=" + strconv.FormatInt(result.Data.UserId, 10) + "&price=" + result.Data.Money + "&order=" + strconv.FormatInt(result.Data.OrderId, 10) + "&vip=" + strconv.FormatInt(requestData.Vip, 10)

	c.JSON(http.StatusOK, api.ApiSuccessResponse(link))
}

func (i *Pay) Wechat(c *gin.Context) {
	var requestData webValidator.PayDetailValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	c.HTML(http.StatusOK, "pay_wechat.html", gin.H{
		"user_id":     requestData.UserId,
		"order":       requestData.Order,
		"price":       requestData.Price,
		"vip":         requestData.Vip,
		"create_time": time.Now().Format("2006-01-02 15:04:05"),
	})
}

func (i *Pay) Alipay(c *gin.Context) {
	var requestData webValidator.PayDetailValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	c.HTML(http.StatusOK, "pay_alipay.html", gin.H{
		"user_id":     requestData.UserId,
		"order":       requestData.Order,
		"price":       requestData.Price,
		"vip":         requestData.Vip,
		"create_time": time.Now().Format("2006-01-02 15:04:05"),
	})
}

func (i *Pay) Check(c *gin.Context) {
	var requestData webValidator.PayCheckValidator

	if err := c.ShouldBind(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	type Result struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
		Data   any    `json:"data"`
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(http.MethodPost)
	req.SetRequestURI("http://localhost:3000/app/pay/check")

	token, _ := c.Cookie("tokens")

	req.SetBody([]byte(`{"order_id":` + strconv.FormatInt(requestData.OrderId, 10) + `,"pay_type":"` + strconv.FormatInt(requestData.PayType, 10) + `","token":"` + token + `"}`))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PAY_FAIL)))
		return
	}

	var result Result

	err := json.Unmarshal(resp.Body(), &result)
	if err != nil || result.Status != 200 {
		c.Error(exception.ValidatorError(errors.Annotate(errors.New(exception.API_PAY_FAIL), exception.API_PAY_FAIL)))
		return
	}
	c.JSON(http.StatusOK, api.ApiSuccessResponse("支付成功"))
}

func NewPay() *Pay {
	return &Pay{}
}
