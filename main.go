package main

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/core/initialize"
	"tutu-gin/http/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	initialize.InitConfig()
	router.RouterInit(r)
	r.Run(":5634")
}
