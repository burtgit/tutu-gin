package web

import (
	"net/http"
	"strconv"
	"sync"
	"time"
	"tutu-gin/lib/kljx"
	"tutu-gin/lib/kljx/response"

	"tutu-gin/core/api"

	"github.com/juju/errors"

	"tutu-gin/core/exception"
	"tutu-gin/http/validator/webValidator"

	"github.com/gin-gonic/gin"
)

type Pay struct{}

func (i *Pay) Index(c *gin.Context) {

	var token string
	tokens, e := c.Request.Cookie("tokens")

	if e == nil {
		token = tokens.Value
	}

	var result response.Menu
	var err error
	var userDetail response.User

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		err, result = kljx.NewClient[response.Menu]().Apply(kljx.PayInfo, map[string]string{
			"token": token,
		})
	}()

	go func() {
		defer wg.Done()
		err, userDetail = kljx.NewClient[response.User]().Apply(kljx.UserInfo, map[string]string{
			"token": token,
		})
	}()

	wg.Wait()

	var errorTip string
	if err != nil {
		errorTip = err.Error()
	}

	hasType := 0
	mode := "single"
	if len(c.Query("hasType")) > 0 {
		hasType = 1
	}

	if len(c.Query("type")) > 0 {
		mode = c.Query("type")
	}

	if c.Query("res") == "2" {
		errorTip = "海外视频解析仅限会员使用，请先购买会员"
	} else if c.Query("res") == "1" {
		errorTip = "今日的免费次数已用完，请先购买会员"
	}

	c.HTML(http.StatusOK, "pay2.html", gin.H{
		"menuList":     result.MenuList,
		"messageData":  result.Message,
		"errorTip":     errorTip,
		"token":        userDetail.Token,
		"vip_times":    userDetail.Times,
		"batch_time":   userDetail.BatchTime,
		"vip_end_time": userDetail.EndTime,
		"message":      userDetail.Message,
		"user_id":      userDetail.Id,
		"domain":       "zhishuzhan.com",
		"mode":         mode,
		"has_type":     hasType,
		"vip_type":     1,
	})
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

	var token string
	tokens, err := c.Request.Cookie("tokens")

	if err == nil {
		token = tokens.Value
	}

	err, result := kljx.NewClient[response.PayApplyResult]().Apply(kljx.PayApply, map[string]string{
		"token":    token,
		"menu":     strconv.FormatInt(requestData.Menu, 10),
		"pay_type": strconv.FormatInt(requestData.PayType, 10),
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
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

	var token string
	tokens, err := c.Request.Cookie("tokens")

	if err == nil {
		token = tokens.Value
	}

	err, _ = kljx.NewClient[any]().Apply(kljx.PayCheck, map[string]string{
		"token":    token,
		"order_id": strconv.FormatInt(requestData.OrderId, 10),
		"pay_type": strconv.FormatInt(requestData.PayType, 10),
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse("支付成功"))
}

func NewPay() *Pay {
	return &Pay{}
}
