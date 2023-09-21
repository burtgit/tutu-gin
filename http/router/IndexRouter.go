package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/web"
)

type IndexRouter struct{}

func (i IndexRouter) Init(router *gin.Engine) {
	indexController := web.NewIndex()
	router.GET("/", indexController.Index)
}
