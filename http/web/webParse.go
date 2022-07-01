package web

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"net/http"
	"tutu-gin/core/api"
	"tutu-gin/core/exception"
	"tutu-gin/http/validator/webValidator"
	"tutu-gin/parser/parserApplicaition"
)

type WebParse struct {
}

func (w *WebParse) Parse(c *gin.Context) {
	var requestData webValidator.WebParseValidator

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARAMETER_CHECK_FAIL)))
		return
	}

	parserService := parserApplicaition.NewParserService()
	result, err := parserService.Parse(requestData.PageUrl, c.ClientIP())
	if err != nil {
		c.Error(exception.ValidatorError(errors.Annotate(err, exception.API_PARSER_FAIL)))
		return
	}

	c.JSON(http.StatusOK, api.ApiSuccessResponse(result))
}

func NewWebParse() *WebParse {
	return &WebParse{}
}
