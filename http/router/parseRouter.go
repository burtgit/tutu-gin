package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/middleware"
	"tutu-gin/http/web"
)

type ParseRouter struct{}

func (p *ParseRouter) Init(router *gin.Engine) {
	parseController := web.NewWebParse()
	v1 := router.Group("/v1")
	v1.Use(middleware.ErrorHandle(), middleware.StatusHandle())
	{
		v1.POST("/parse/index", middleware.AuthHandle(), parseController.Parse)
		v1.POST("/parse/dana", parseController.ParseDaNa)
	}

	v2 := router.Group("/v2")
	v2.Use(middleware.ErrorHandle(), middleware.StatusHandle())
	{
		v2.POST("/parse/index", parseController.ParseV2)
	}
}
