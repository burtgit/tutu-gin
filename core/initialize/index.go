package initialize

import "tutu-gin/core"

func InitConfig() {
	core.ViperConfigLoad() // 加载配置文件
	InitMysql()            // 初始化数据库
	InitRedis()
	InitParse()
}
