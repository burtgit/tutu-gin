package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/middleware"
	"tutu-gin/http/web"
)

type ApiRouter struct{}

func (p *ApiRouter) Init(router *gin.Engine) {
	parseController := web.NewWebParse()
	v1 := router.Group("/v1")
	v1.Use(middleware.ErrorHandle(), middleware.StatusHandle())
	{
		v1.POST("/api/parse", parseController.Api)
	}
}
