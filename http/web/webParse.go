package web

import (
	"net/http"
	"strconv"
	"time"

	"tutu-gin/record/recordDomain/values"

	"tutu-gin/core/global"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"

	"tutu-gin/core/api"
	"tutu-gin/core/exception"
	"tutu-gin/http/validator/webValidator"
	"tutu-gin/parser/parserApplicaition"
)

type WebParse struct{}

func (w *WebParse) Parse(c *gin.Context) {
	var requestData webValidator.WebParseValidator

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	day, _ := strconv.ParseInt(time.Now().Format("20060102"), 10, 64)
	getInt64 := c.GetInt64("userId")

	var apply values.Apply

	global.DB.Where("user_id = ?", getInt64).Where("platform = ?", "zhishuzhan").Where("date = ?", day).Get(&apply)

	if apply.Id <= 0 {
		apply.Platform = "zhishuzhan"
		apply.UserId = getInt64
		apply.Date = day
		apply.TotalTimes = 1
		_, _ = global.DB.Insert(&apply)
	} else {
		global.DB.ID(apply.Id).Incr("total_times").Cols("total_times").Update(&apply)
	}

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Parse(requestData.PageUrl, c.ClientIP(), getInt64)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}
	global.DB.ID(apply.Id).Incr("success_times").Cols("success_times").Update(&apply)
	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

func NewWebParse() *WebParse {
	return &WebParse{}
}
