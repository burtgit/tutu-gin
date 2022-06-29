package global

import (
	"tutu-gin/core/config"
	"xorm.io/xorm"
)

var (
	DB             *xorm.Engine
	REDIS_CLIENT   interface{}
	SERVICE_CONFIG config.ServiceConfig
)
