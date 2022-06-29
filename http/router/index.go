package router

import "github.com/gin-gonic/gin"

func RouterInit(engine *gin.Engine) {
	// 初始化解析路由
	parseRouter := &ParseRouter{}
	parseRouter.Init(engine)

	indexRouter := &IndexRouter{}
	indexRouter.Init(engine)

	// 初始化静态资源
	engine.Static("/static/", "./http/resource")

	// 初始化html模板目录
	engine.LoadHTMLGlob("http/resource/templates/default/*")
}
