package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/middleware"
	"tutu-gin/http/web"
)

type UserRouter struct{}

func (u *UserRouter) Init(router *gin.Engine) {
	parseController := web.NewUser()
	v1 := router.Group("/v1")
	v1.Use(middleware.ErrorHandle(), middleware.StatusHandle())
	{
		v1.GET("/qrcode/get", parseController.Qrcode)
		v1.GET("/qrcode/test", parseController.Test)
		v1.GET("/qrcode/check", parseController.Check)
	}
}
