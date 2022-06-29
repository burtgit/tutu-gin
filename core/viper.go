package core

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"tutu-gin/core/global"
)

var configDebugFile = "config.debug.yaml"
var configReleaseFile = "config.release.yaml"

func ViperConfigLoad() {
	var configPath string

	if gin.Mode() == gin.ReleaseMode {
		configPath = configReleaseFile
	} else {
		configPath = configDebugFile
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err)
	}

	err = viper.Unmarshal(&global.SERVICE_CONFIG)
	if err != nil {
		log.Println(err)
	}
	log.Println(global.SERVICE_CONFIG)
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println(in.Name)
		err = viper.Unmarshal(&global.SERVICE_CONFIG)
		if err != nil {
			log.Println(err)
		}
	})
	viper.WatchConfig()
}
