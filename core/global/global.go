package global

import (
	"github.com/redis/go-redis/v9"
	"tutu-gin/core/config"
	"xorm.io/xorm"
)

var (
	DB             *xorm.Engine
	REDIS_CLIENT   *redis.Client
	SERVICE_CONFIG config.ServiceConfig
)
