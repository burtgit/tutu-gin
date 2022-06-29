package main

import (
	"github.com/gin-gonic/gin"
	"tutu-gin/core/initialize"
	"tutu-gin/http/router"
)

func main() {
	r := gin.Default()
	initialize.InitConfig()
	router.RouterInit(r)
	r.Run(":8080")
}
