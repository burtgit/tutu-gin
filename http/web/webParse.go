package web

import (
	"net/http"

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

	getInt64 := c.GetInt64("userId")

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Parse(requestData.PageUrl, c.ClientIP(), getInt64)
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

func NewWebParse() *WebParse {
	return &WebParse{}
}
