package router

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/http/web"
)

type IndexRouter struct{}

func (i IndexRouter) Init(router *gin.Engine) {
	indexController := web.NewIndex()
	router.GET("/", indexController.Index)
	router.GET("/xiaohongshu.html", indexController.Index)
	router.GET("/weibo.html", indexController.Index)
	router.GET("/xigua.html", indexController.Index)
	router.GET("/kuaishou.html", indexController.Index)
	router.GET("/bilibili.html", indexController.Index)
	router.GET("/instagram.html", indexController.Index)
	router.GET("/youtube.html", indexController.Index)
	router.GET("/tiktok.html", indexController.Index)
}
