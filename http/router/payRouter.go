package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/middleware"
	"tutu-gin/http/web"
)

type PayRouter struct{}

func (p *PayRouter) Init(router *gin.Engine) {
	payController := web.NewPay()
	v1 := router.Group("/v1")
	v1.Use(middleware.ErrorHandle(), middleware.StatusHandle())
	{
		v1.GET("/pay/index", payController.Index)
		v1.GET("/pay/mobile", middleware.AuthHandle(), payController.Mobile)
		v1.POST("/pay/apply", payController.Apply)
		v1.GET("/pay/wechat", payController.Wechat)
		v1.GET("/pay/alipay", payController.Alipay)
		v1.POST("/pay/check", payController.Check)
	}
}
